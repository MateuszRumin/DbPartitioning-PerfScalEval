from __future__ import annotations

from contextlib import contextmanager
from datetime import date, datetime, time, timedelta
from typing import Any, Iterator, Sequence

import bleach
import mysql.connector
from mysql.connector import pooling
import streamlit as st


st.set_page_config(
    page_title="Mini Forum",
    page_icon="💬",
    layout="wide",
)


ALLOWED_TAGS = {
    "a", "blockquote", "br", "code", "em", "h1", "h2", "h3", "h4",
    "li", "ol", "p", "pre", "strong", "ul",
}
ALLOWED_ATTRIBUTES = {
    "a": ["href", "title", "target", "rel"],
}
CONTENT_LICENSE = "CC BY-SA 4.0"
MAX_TAGS = 10


# -----------------------------------------------------------------------------
# Pomocnicze funkcje interfejsu
# -----------------------------------------------------------------------------


def show_flash() -> None:
    message = st.session_state.pop("flash", None)
    if message:
        st.success(message)


def set_flash(message: str) -> None:
    st.session_state["flash"] = message


def clean_html(value: str | None) -> str:
    if not value:
        return "_Brak treści._"

    return bleach.clean(
        value,
        tags=ALLOWED_TAGS,
        attributes=ALLOWED_ATTRIBUTES,
        protocols={"http", "https", "mailto"},
        strip=True,
    )


def parse_tags(raw_tags: str | None) -> list[str]:
    if not raw_tags:
        return []

    if "><" in raw_tags or raw_tags.startswith("<"):
        parts = raw_tags.replace("><", " ").replace("<", "").replace(">", "").split()
    else:
        parts = raw_tags.replace(",", " ").split()

    tags: list[str] = []
    for part in parts:
        tag = part.strip().lower().replace("<", "").replace(">", "")
        if tag and tag not in tags:
            tags.append(tag[:50])
    return tags[:MAX_TAGS]


def normalize_tags(raw_tags: str) -> str | None:
    tags = parse_tags(raw_tags)
    if not tags:
        return None
    return "".join(f"<{tag}>" for tag in tags)


def format_tags(raw_tags: str | None) -> str:
    return " ".join(f"`{tag}`" for tag in parse_tags(raw_tags))


def format_date(value: datetime | None) -> str:
    if value is None:
        return "brak daty"
    return value.strftime("%Y-%m-%d %H:%M")


def inclusive_date_bounds(date_from: date, date_to: date) -> tuple[datetime, datetime]:
    return (
        datetime.combine(date_from, time.min),
        datetime.combine(date_to + timedelta(days=1), time.min),
    )


# -----------------------------------------------------------------------------
# Połączenie i operacje bazodanowe
# -----------------------------------------------------------------------------


@st.cache_resource(show_spinner=False)
def get_pool() -> pooling.MySQLConnectionPool:
    try:
        cfg = st.secrets["mysql"]
    except (KeyError, FileNotFoundError) as exc:
        raise RuntimeError(
            "Brak konfiguracji [mysql] w pliku .streamlit/secrets.toml."
        ) from exc

    return pooling.MySQLConnectionPool(
        pool_name="streamlit_forum_pool",
        pool_size=int(cfg.get("pool_size", 5)),
        pool_reset_session=True,
        host=str(cfg.get("host", "127.0.0.1")),
        port=int(cfg.get("port", 3306)),
        user=str(cfg["user"]),
        password=str(cfg["password"]),
        database=str(cfg["database"]),
        charset="utf8mb4",
        collation="utf8mb4_0900_ai_ci",
        autocommit=False,
    )


@contextmanager
def db_cursor(*, dictionary: bool = True) -> Iterator[tuple[Any, Any]]:
    connection = get_pool().get_connection()
    cursor = None

    try:
        connection.ping(reconnect=True, attempts=2, delay=1)
        cursor = connection.cursor(dictionary=dictionary)
        yield connection, cursor
    finally:
        if cursor is not None:
            cursor.close()
        connection.close()


def fetch_all(sql: str, params: Sequence[Any] = ()) -> list[dict[str, Any]]:
    with db_cursor() as (_, cursor):
        cursor.execute(sql, tuple(params))
        return list(cursor.fetchall())


def fetch_one(sql: str, params: Sequence[Any] = ()) -> dict[str, Any] | None:
    with db_cursor() as (_, cursor):
        cursor.execute(sql, tuple(params))
        return cursor.fetchone()


def execute_transaction(statements: Sequence[tuple[str, Sequence[Any]]]) -> int:
    with db_cursor(dictionary=False) as (connection, cursor):
        try:
            lastrowid = 0
            for sql, params in statements:
                cursor.execute(sql, tuple(params))
                if cursor.lastrowid:
                    lastrowid = int(cursor.lastrowid)
            connection.commit()
            return lastrowid
        except Exception:
            connection.rollback()
            raise


def get_user(user_id: int) -> dict[str, Any] | None:
    return fetch_one(
        """
        SELECT id, display_name, reputation
        FROM users
        WHERE id = %s
        LIMIT 1
        """,
        (user_id,),
    )


def require_user(user_id: int) -> dict[str, Any]:
    user = get_user(user_id)
    if user is None:
        raise ValueError(f"Użytkownik o ID {user_id} nie istnieje.")
    return user


@st.cache_data(ttl=300, show_spinner=False)
def get_question_date_bounds() -> tuple[date, date]:
    row = fetch_one(
        """
        SELECT
            MIN(creation_date) AS min_date,
            MAX(creation_date) AS max_date
        FROM posts
        WHERE post_type_id = 1
        """
    )

    if not row or row["min_date"] is None or row["max_date"] is None:
        today = datetime.now().date()
        return today, today

    return row["min_date"].date(), row["max_date"].date()


def list_questions(
    *,
    limit: int,
    offset: int,
    search_text: str = "",
    search_scope: str = "Wszystko",
    tag: str = "",
    date_from: date | None = None,
    date_to: date | None = None,
    sort_order: str = "Najnowsze",
) -> tuple[list[dict[str, Any]], bool]:
    conditions = ["q.post_type_id = 1"]
    params: list[Any] = []

    normalized_search = search_text.strip()
    if normalized_search:
        pattern = f"%{normalized_search}%"
        if search_scope == "Tytuł":
            conditions.append("q.post_title LIKE %s")
            params.append(pattern)
        elif search_scope == "Treść":
            conditions.append("q.post_body LIKE %s")
            params.append(pattern)
        elif search_scope == "Tagi":
            conditions.append("q.tags LIKE %s")
            params.append(pattern)
        else:
            conditions.append(
                "(q.post_title LIKE %s OR q.post_body LIKE %s OR q.tags LIKE %s)"
            )
            params.extend([pattern, pattern, pattern])

    normalized_tag = parse_tags(tag)
    if normalized_tag:
        conditions.append("q.tags LIKE %s")
        params.append(f"%<{normalized_tag[0]}>%")

    if date_from is not None and date_to is not None:
        range_start, range_end = inclusive_date_bounds(date_from, date_to)
        conditions.append("q.creation_date >= %s AND q.creation_date < %s")
        params.extend([range_start, range_end])

    order_by = {
        "Najnowsze": "q.creation_date DESC, q.id DESC",
        "Najstarsze": "q.creation_date ASC, q.id ASC",
        "Najwyższy wynik": "q.score DESC, q.creation_date DESC",
        "Najwięcej odpowiedzi": "q.answer_count DESC, q.creation_date DESC",
        "Najwięcej wyświetleń": "q.view_count DESC, q.creation_date DESC",
    }.get(sort_order, "q.creation_date DESC, q.id DESC")

    params.extend([limit + 1, offset])
    rows = fetch_all(
        f"""
        SELECT
            q.id,
            q.creation_date,
            q.score,
            q.view_count,
            q.post_title,
            q.tags,
            q.answer_count,
            q.comment_count,
            q.owner_user_id,
            COALESCE(u.display_name, '[usunięty]') AS author_name
        FROM posts AS q
        LEFT JOIN users AS u
            ON u.id = q.owner_user_id
        WHERE {' AND '.join(conditions)}
        ORDER BY {order_by}
        LIMIT %s OFFSET %s
        """,
        params,
    )

    has_more = len(rows) > limit
    return rows[:limit], has_more


def get_question(question_id: int) -> dict[str, Any] | None:
    return fetch_one(
        """
        SELECT
            q.id,
            q.creation_date,
            q.last_edit_date,
            q.score,
            q.view_count,
            q.post_title,
            q.post_body,
            q.tags,
            q.answer_count,
            q.comment_count,
            q.accepted_answer_id,
            q.owner_user_id,
            COALESCE(u.display_name, '[usunięty]') AS author_name
        FROM posts AS q
        LEFT JOIN users AS u
            ON u.id = q.owner_user_id
        WHERE q.id = %s
          AND q.post_type_id = 1
        LIMIT 1
        """,
        (question_id,),
    )


def get_answers(question_id: int) -> list[dict[str, Any]]:
    return fetch_all(
        """
        SELECT
            a.id,
            a.creation_date,
            a.last_edit_date,
            a.score,
            a.post_body,
            a.comment_count,
            a.owner_user_id,
            COALESCE(u.display_name, '[usunięty]') AS author_name
        FROM posts AS a
        LEFT JOIN users AS u
            ON u.id = a.owner_user_id
        WHERE a.parent_id = %s
          AND a.post_type_id = 2
        ORDER BY a.score DESC, a.creation_date ASC
        """,
        (question_id,),
    )


def get_comments(post_ids: Sequence[int]) -> dict[int, list[dict[str, Any]]]:
    if not post_ids:
        return {}

    placeholders = ", ".join(["%s"] * len(post_ids))
    rows = fetch_all(
        f"""
        SELECT
            c.id,
            c.post_id,
            c.comment_text,
            c.creation_date,
            c.score,
            c.user_id,
            COALESCE(u.display_name, '[usunięty]') AS author_name
        FROM comments AS c
        LEFT JOIN users AS u
            ON u.id = c.user_id
        WHERE c.post_id IN ({placeholders})
        ORDER BY c.creation_date ASC, c.id ASC
        """,
        tuple(post_ids),
    )

    grouped: dict[int, list[dict[str, Any]]] = {}
    for row in rows:
        grouped.setdefault(int(row["post_id"]), []).append(row)
    return grouped


def increment_view_count(question_id: int) -> None:
    execute_transaction(
        [
            (
                """
                UPDATE posts
                SET view_count = COALESCE(view_count, 0) + 1
                WHERE id = %s
                  AND post_type_id = 1
                """,
                (question_id,),
            )
        ]
    )


def create_question(
    *,
    user_id: int,
    title: str,
    body: str,
    tags: str,
) -> int:
    require_user(user_id)

    normalized_tags = normalize_tags(tags)
    if normalized_tags is None:
        raise ValueError("Pytanie musi zawierać co najmniej jeden tag.")

    return execute_transaction(
        [
            (
                """
                INSERT INTO posts (
                    post_type_id,
                    creation_date,
                    score,
                    view_count,
                    post_body,
                    owner_user_id,
                    last_activity_date,
                    post_title,
                    tags,
                    answer_count,
                    comment_count,
                    content_license
                )
                VALUES (
                    1,
                    NOW(),
                    0,
                    0,
                    %s,
                    %s,
                    NOW(),
                    %s,
                    %s,
                    0,
                    0,
                    %s
                )
                """,
                (body, user_id, title, normalized_tags, CONTENT_LICENSE),
            )
        ]
    )


def create_answer(*, question_id: int, user_id: int, body: str) -> int:
    require_user(user_id)

    with db_cursor(dictionary=False) as (connection, cursor):
        try:
            cursor.execute(
                """
                SELECT id
                FROM posts
                WHERE id = %s
                  AND post_type_id = 1
                LIMIT 1
                FOR UPDATE
                """,
                (question_id,),
            )
            if cursor.fetchone() is None:
                raise ValueError("Wybrany temat nie istnieje.")

            cursor.execute(
                """
                INSERT INTO posts (
                    post_type_id,
                    parent_id,
                    creation_date,
                    score,
                    post_body,
                    owner_user_id,
                    last_activity_date,
                    comment_count,
                    content_license
                )
                VALUES (2, %s, NOW(), 0, %s, %s, NOW(), 0, %s)
                """,
                (question_id, body, user_id, CONTENT_LICENSE),
            )
            answer_id = int(cursor.lastrowid)

            cursor.execute(
                """
                UPDATE posts
                SET
                    answer_count = COALESCE(answer_count, 0) + 1,
                    last_activity_date = NOW()
                WHERE id = %s
                  AND post_type_id = 1
                """,
                (question_id,),
            )
            connection.commit()
            return answer_id
        except Exception:
            connection.rollback()
            raise


def update_answer(*, answer_id: int, user_id: int, body: str) -> None:
    require_user(user_id)

    with db_cursor(dictionary=False) as (connection, cursor):
        try:
            cursor.execute(
                """
                UPDATE posts
                SET
                    post_body = %s,
                    last_edit_date = NOW(),
                    last_activity_date = NOW()
                WHERE id = %s
                  AND post_type_id = 2
                  AND owner_user_id = %s
                """,
                (body, answer_id, user_id),
            )
            if cursor.rowcount != 1:
                raise ValueError("Nie znaleziono odpowiedzi lub nie jesteś jej autorem.")
            connection.commit()
        except Exception:
            connection.rollback()
            raise


def delete_answer(*, answer_id: int, user_id: int) -> None:
    require_user(user_id)

    with db_cursor(dictionary=False) as (connection, cursor):
        try:
            cursor.execute(
                """
                SELECT parent_id
                FROM posts
                WHERE id = %s
                  AND post_type_id = 2
                  AND owner_user_id = %s
                LIMIT 1
                FOR UPDATE
                """,
                (answer_id, user_id),
            )
            row = cursor.fetchone()
            if row is None:
                raise ValueError("Nie znaleziono odpowiedzi lub nie jesteś jej autorem.")

            question_id = int(row[0])
            cursor.execute("DELETE FROM comments WHERE post_id = %s", (answer_id,))
            cursor.execute(
                """
                DELETE FROM posts
                WHERE id = %s
                  AND post_type_id = 2
                  AND owner_user_id = %s
                """,
                (answer_id, user_id),
            )
            if cursor.rowcount != 1:
                raise ValueError("Nie udało się usunąć odpowiedzi.")

            cursor.execute(
                """
                UPDATE posts
                SET
                    answer_count = GREATEST(COALESCE(answer_count, 0) - 1, 0),
                    accepted_answer_id = CASE
                        WHEN accepted_answer_id = %s THEN NULL
                        ELSE accepted_answer_id
                    END,
                    last_activity_date = NOW()
                WHERE id = %s
                  AND post_type_id = 1
                """,
                (answer_id, question_id),
            )
            connection.commit()
        except Exception:
            connection.rollback()
            raise


def set_accepted_answer(
    *,
    question_id: int,
    answer_id: int | None,
    user_id: int,
) -> None:
    require_user(user_id)

    with db_cursor(dictionary=False) as (connection, cursor):
        try:
            cursor.execute(
                """
                SELECT id
                FROM posts
                WHERE id = %s
                  AND post_type_id = 1
                  AND owner_user_id = %s
                LIMIT 1
                FOR UPDATE
                """,
                (question_id, user_id),
            )
            if cursor.fetchone() is None:
                raise ValueError("Tylko autor pytania może zaakceptować odpowiedź.")

            if answer_id is not None:
                cursor.execute(
                    """
                    SELECT id
                    FROM posts
                    WHERE id = %s
                      AND parent_id = %s
                      AND post_type_id = 2
                    LIMIT 1
                    """,
                    (answer_id, question_id),
                )
                if cursor.fetchone() is None:
                    raise ValueError("Odpowiedź nie należy do tego pytania.")

            cursor.execute(
                """
                UPDATE posts
                SET
                    accepted_answer_id = %s,
                    last_activity_date = NOW()
                WHERE id = %s
                  AND post_type_id = 1
                  AND owner_user_id = %s
                """,
                (answer_id, question_id, user_id),
            )
            connection.commit()
        except Exception:
            connection.rollback()
            raise


def create_comment(*, post_id: int, user_id: int, body: str) -> int:
    require_user(user_id)

    with db_cursor(dictionary=False) as (connection, cursor):
        try:
            cursor.execute(
                """
                SELECT id
                FROM posts
                WHERE id = %s
                LIMIT 1
                FOR UPDATE
                """,
                (post_id,),
            )
            if cursor.fetchone() is None:
                raise ValueError("Post, do którego dodawany jest komentarz, nie istnieje.")

            cursor.execute(
                """
                INSERT INTO comments (
                    post_id,
                    score,
                    comment_text,
                    creation_date,
                    user_id,
                    content_license
                )
                VALUES (%s, 0, %s, NOW(), %s, %s)
                """,
                (post_id, body, user_id, CONTENT_LICENSE),
            )
            comment_id = int(cursor.lastrowid)

            cursor.execute(
                """
                UPDATE posts
                SET
                    comment_count = COALESCE(comment_count, 0) + 1,
                    last_activity_date = NOW()
                WHERE id = %s
                """,
                (post_id,),
            )
            connection.commit()
            return comment_id
        except Exception:
            connection.rollback()
            raise


def update_comment(
    *,
    comment_id: int,
    post_id: int,
    user_id: int,
    body: str,
) -> None:
    require_user(user_id)

    with db_cursor(dictionary=False) as (connection, cursor):
        try:
            cursor.execute(
                """
                UPDATE comments
                SET comment_text = %s
                WHERE id = %s
                  AND post_id = %s
                  AND user_id = %s
                """,
                (body, comment_id, post_id, user_id),
            )
            if cursor.rowcount != 1:
                raise ValueError("Nie znaleziono komentarza lub nie jesteś jego autorem.")
            connection.commit()
        except Exception:
            connection.rollback()
            raise


def delete_comment(*, comment_id: int, post_id: int, user_id: int) -> None:
    require_user(user_id)

    with db_cursor(dictionary=False) as (connection, cursor):
        try:
            cursor.execute(
                """
                DELETE FROM comments
                WHERE id = %s
                  AND post_id = %s
                  AND user_id = %s
                """,
                (comment_id, post_id, user_id),
            )
            if cursor.rowcount != 1:
                raise ValueError("Nie znaleziono komentarza lub nie jesteś jego autorem.")

            cursor.execute(
                """
                UPDATE posts
                SET
                    comment_count = GREATEST(COALESCE(comment_count, 0) - 1, 0),
                    last_activity_date = NOW()
                WHERE id = %s
                """,
                (post_id,),
            )
            connection.commit()
        except Exception:
            connection.rollback()
            raise


# -----------------------------------------------------------------------------
# Renderowanie komentarzy, odpowiedzi i tematów
# -----------------------------------------------------------------------------


def render_comments(
    *,
    post_id: int,
    comments: list[dict[str, Any]],
    current_user_id: int,
) -> None:
    with st.expander(f"Komentarze ({len(comments)})"):
        if comments:
            for comment in comments:
                comment_id = int(comment["id"])
                is_owner = comment["user_id"] == current_user_id

                st.markdown(
                    f"**{comment['author_name']}** · "
                    f"{format_date(comment['creation_date'])} · "
                    f"wynik: {comment['score'] or 0}"
                )
                st.write(comment["comment_text"] or "")

                if is_owner:
                    action_col1, action_col2, _ = st.columns([1, 1, 5])
                    with action_col1:
                        if st.button(
                            "Edytuj",
                            key=f"edit_comment_button_{comment_id}",
                            use_container_width=True,
                        ):
                            st.session_state["editing_comment_id"] = comment_id
                            st.session_state.pop("deleting_comment_id", None)
                            st.rerun()
                    with action_col2:
                        if st.button(
                            "Usuń",
                            key=f"delete_comment_button_{comment_id}",
                            use_container_width=True,
                        ):
                            st.session_state["deleting_comment_id"] = comment_id
                            st.session_state.pop("editing_comment_id", None)
                            st.rerun()

                    if st.session_state.get("editing_comment_id") == comment_id:
                        with st.form(f"edit_comment_form_{comment_id}"):
                            edited_body = st.text_area(
                                "Treść komentarza",
                                value=comment["comment_text"] or "",
                                max_chars=4000,
                            )
                            save_col, cancel_col = st.columns(2)
                            save = save_col.form_submit_button("Zapisz", use_container_width=True)
                            cancel = cancel_col.form_submit_button("Anuluj", use_container_width=True)

                        if cancel:
                            st.session_state.pop("editing_comment_id", None)
                            st.rerun()
                        if save:
                            if not edited_body.strip():
                                st.warning("Komentarz nie może być pusty.")
                            else:
                                try:
                                    update_comment(
                                        comment_id=comment_id,
                                        post_id=post_id,
                                        user_id=current_user_id,
                                        body=edited_body.strip(),
                                    )
                                    st.session_state.pop("editing_comment_id", None)
                                    set_flash("Komentarz został zaktualizowany.")
                                    st.rerun()
                                except (mysql.connector.Error, ValueError, RuntimeError) as exc:
                                    st.error(str(exc))

                    if st.session_state.get("deleting_comment_id") == comment_id:
                        with st.form(f"delete_comment_form_{comment_id}"):
                            st.warning("Usunięcie komentarza jest nieodwracalne.")
                            confirmed = st.checkbox("Potwierdzam usunięcie komentarza")
                            delete_col, cancel_col = st.columns(2)
                            delete_submitted = delete_col.form_submit_button(
                                "Usuń komentarz",
                                use_container_width=True,
                            )
                            cancel = cancel_col.form_submit_button(
                                "Anuluj",
                                use_container_width=True,
                            )

                        if cancel:
                            st.session_state.pop("deleting_comment_id", None)
                            st.rerun()
                        if delete_submitted:
                            if not confirmed:
                                st.warning("Zaznacz potwierdzenie usunięcia.")
                            else:
                                try:
                                    delete_comment(
                                        comment_id=comment_id,
                                        post_id=post_id,
                                        user_id=current_user_id,
                                    )
                                    st.session_state.pop("deleting_comment_id", None)
                                    set_flash("Komentarz został usunięty.")
                                    st.rerun()
                                except (mysql.connector.Error, ValueError, RuntimeError) as exc:
                                    st.error(str(exc))

                st.divider()
        else:
            st.caption("Brak komentarzy.")

        with st.form(f"comment_form_{post_id}", clear_on_submit=True):
            comment_body = st.text_area(
                "Dodaj komentarz",
                max_chars=4000,
                key=f"comment_body_{post_id}",
            )
            submitted = st.form_submit_button("Dodaj komentarz")

        if submitted:
            if not comment_body.strip():
                st.warning("Komentarz nie może być pusty.")
            else:
                try:
                    create_comment(
                        post_id=post_id,
                        user_id=current_user_id,
                        body=comment_body.strip(),
                    )
                    set_flash("Komentarz został dodany.")
                    st.rerun()
                except (mysql.connector.Error, ValueError, RuntimeError) as exc:
                    st.error(str(exc))


def render_question_list() -> None:
    st.title("Mini Forum")
    st.caption("Wyszukiwanie pytań z tabeli `posts`, z uwzględnieniem treści i tagów.")

    try:
        min_date, max_date = get_question_date_bounds()
    except (mysql.connector.Error, RuntimeError):
        min_date = date(2008, 1, 1)
        max_date = datetime.now().date()

    with st.container(border=True):
        st.subheader("Wyszukiwarka")
        search_col, scope_col = st.columns([3, 1])
        search_text = search_col.text_input(
            "Szukana fraza",
            value=st.session_state.get("question_search_text", ""),
            placeholder="Tytuł, treść albo tag",
        )
        search_scope = scope_col.selectbox(
            "Zakres wyszukiwania",
            ["Wszystko", "Tytuł", "Treść", "Tagi"],
            index=["Wszystko", "Tytuł", "Treść", "Tagi"].index(
                st.session_state.get("question_search_scope", "Wszystko")
            ),
        )

        tag_col, date_col = st.columns([1, 2])
        tag_filter = tag_col.text_input(
            "Dokładny tag",
            value=st.session_state.get("question_tag_filter", ""),
            placeholder="np. mysql",
            help="Filtruje zapis Stack Overflow w formacie <tag>.",
        )
        selected_dates = date_col.date_input(
            "Zakres dat utworzenia",
            value=(
                st.session_state.get("question_date_from", min_date),
                st.session_state.get("question_date_to", max_date),
            ),
            min_value=min_date,
            max_value=max_date,
            format="YYYY-MM-DD",
        )

        sort_col, limit_col = st.columns(2)
        sort_order = sort_col.selectbox(
            "Sortowanie",
            [
                "Najnowsze",
                "Najstarsze",
                "Najwyższy wynik",
                "Najwięcej odpowiedzi",
                "Najwięcej wyświetleń",
            ],
        )
        page_size = limit_col.select_slider(
            "Liczba tematów na stronę",
            options=[10, 20, 30, 50, 100],
            value=int(st.session_state.get("question_page_size", 30)),
        )

    if isinstance(selected_dates, tuple) and len(selected_dates) == 2:
        date_from, date_to = selected_dates
    else:
        date_from = date_to = selected_dates  # type: ignore[assignment]

    current_filter = (
        search_text,
        search_scope,
        tag_filter,
        date_from,
        date_to,
        sort_order,
        page_size,
    )
    if st.session_state.get("question_filter_signature") != current_filter:
        st.session_state["question_filter_signature"] = current_filter
        st.session_state["question_page"] = 0

    st.session_state["question_search_text"] = search_text
    st.session_state["question_search_scope"] = search_scope
    st.session_state["question_tag_filter"] = tag_filter
    st.session_state["question_date_from"] = date_from
    st.session_state["question_date_to"] = date_to
    st.session_state["question_page_size"] = page_size

    page = int(st.session_state.get("question_page", 0))
    offset = page * page_size

    try:
        questions, has_more = list_questions(
            limit=page_size,
            offset=offset,
            search_text=search_text,
            search_scope=search_scope,
            tag=tag_filter,
            date_from=date_from,
            date_to=date_to,
            sort_order=sort_order,
        )
    except (mysql.connector.Error, RuntimeError) as exc:
        st.error(f"Nie udało się pobrać tematów: {exc}")
        return

    st.caption(f"Strona {page + 1} · rekordy od {offset + 1}")

    if not questions:
        st.info("Brak pytań spełniających podane kryteria.")
    else:
        for question in questions:
            with st.container(border=True):
                col_main, col_meta = st.columns([5, 1])

                with col_main:
                    title = question["post_title"] or f"Pytanie #{question['id']}"
                    st.subheader(title)
                    if question["tags"]:
                        st.markdown(format_tags(question["tags"]))
                    st.caption(
                        f"Autor: {question['author_name']} · "
                        f"{format_date(question['creation_date'])} · "
                        f"komentarze: {question['comment_count'] or 0}"
                    )

                with col_meta:
                    st.metric("Odpowiedzi", question["answer_count"] or 0)
                    st.caption(
                        f"Wynik: {question['score'] or 0}  \n"
                        f"Wyświetlenia: {question['view_count'] or 0}"
                    )
                    if st.button(
                        "Otwórz",
                        key=f"open_question_{question['id']}",
                        use_container_width=True,
                    ):
                        question_id = int(question["id"])
                        st.session_state["selected_question_id"] = question_id
                        viewed = st.session_state.setdefault("viewed_question_ids", set())
                        if question_id not in viewed:
                            try:
                                increment_view_count(question_id)
                                viewed.add(question_id)
                            except (mysql.connector.Error, RuntimeError):
                                pass
                        st.rerun()

    previous_col, page_col, next_col = st.columns([1, 3, 1])
    with previous_col:
        if st.button(
            "← Poprzednia",
            disabled=page == 0,
            use_container_width=True,
        ):
            st.session_state["question_page"] = max(page - 1, 0)
            st.rerun()
    with page_col:
        st.markdown(
            f"<div style='text-align:center;padding-top:0.5rem'>Strona {page + 1}</div>",
            unsafe_allow_html=True,
        )
    with next_col:
        if st.button(
            "Następna →",
            disabled=not has_more,
            use_container_width=True,
        ):
            st.session_state["question_page"] = page + 1
            st.rerun()


def render_answer_management(
    *,
    question: dict[str, Any],
    answer: dict[str, Any],
    current_user_id: int,
) -> None:
    answer_id = int(answer["id"])
    is_answer_owner = answer["owner_user_id"] == current_user_id
    is_question_owner = question["owner_user_id"] == current_user_id
    accepted_answer_id = question.get("accepted_answer_id")
    is_accepted = accepted_answer_id is not None and int(accepted_answer_id) == answer_id

    action_columns = st.columns([1, 1, 1, 4])

    if is_question_owner:
        with action_columns[0]:
            label = "Cofnij akceptację" if is_accepted else "Akceptuj"
            if st.button(
                label,
                key=f"accept_answer_{answer_id}",
                use_container_width=True,
            ):
                try:
                    set_accepted_answer(
                        question_id=int(question["id"]),
                        answer_id=None if is_accepted else answer_id,
                        user_id=current_user_id,
                    )
                    set_flash(
                        "Cofnięto akceptację odpowiedzi."
                        if is_accepted
                        else "Odpowiedź została zaakceptowana."
                    )
                    st.rerun()
                except (mysql.connector.Error, ValueError, RuntimeError) as exc:
                    st.error(str(exc))

    if is_answer_owner:
        with action_columns[1]:
            if st.button(
                "Edytuj",
                key=f"edit_answer_button_{answer_id}",
                use_container_width=True,
            ):
                st.session_state["editing_answer_id"] = answer_id
                st.session_state.pop("deleting_answer_id", None)
                st.rerun()
        with action_columns[2]:
            if st.button(
                "Usuń",
                key=f"delete_answer_button_{answer_id}",
                use_container_width=True,
            ):
                st.session_state["deleting_answer_id"] = answer_id
                st.session_state.pop("editing_answer_id", None)
                st.rerun()

    if st.session_state.get("editing_answer_id") == answer_id:
        with st.form(f"edit_answer_form_{answer_id}"):
            edited_body = st.text_area(
                "Treść odpowiedzi",
                value=answer["post_body"] or "",
                height=180,
                max_chars=10000,
            )
            save_col, cancel_col = st.columns(2)
            save = save_col.form_submit_button("Zapisz", use_container_width=True)
            cancel = cancel_col.form_submit_button("Anuluj", use_container_width=True)

        if cancel:
            st.session_state.pop("editing_answer_id", None)
            st.rerun()
        if save:
            if len(edited_body.strip()) < 10:
                st.warning("Treść odpowiedzi musi mieć co najmniej 10 znaków.")
            else:
                try:
                    update_answer(
                        answer_id=answer_id,
                        user_id=current_user_id,
                        body=edited_body.strip(),
                    )
                    st.session_state.pop("editing_answer_id", None)
                    set_flash("Odpowiedź została zaktualizowana.")
                    st.rerun()
                except (mysql.connector.Error, ValueError, RuntimeError) as exc:
                    st.error(str(exc))

    if st.session_state.get("deleting_answer_id") == answer_id:
        with st.form(f"delete_answer_form_{answer_id}"):
            st.warning("Usunięte zostaną również komentarze przypisane do odpowiedzi.")
            confirmed = st.checkbox("Potwierdzam usunięcie odpowiedzi")
            delete_col, cancel_col = st.columns(2)
            delete_submitted = delete_col.form_submit_button(
                "Usuń odpowiedź",
                use_container_width=True,
            )
            cancel = cancel_col.form_submit_button("Anuluj", use_container_width=True)

        if cancel:
            st.session_state.pop("deleting_answer_id", None)
            st.rerun()
        if delete_submitted:
            if not confirmed:
                st.warning("Zaznacz potwierdzenie usunięcia.")
            else:
                try:
                    delete_answer(answer_id=answer_id, user_id=current_user_id)
                    st.session_state.pop("deleting_answer_id", None)
                    set_flash("Odpowiedź została usunięta.")
                    st.rerun()
                except (mysql.connector.Error, ValueError, RuntimeError) as exc:
                    st.error(str(exc))


def render_thread(question_id: int, current_user_id: int) -> None:
    if st.button("← Wróć do tematów"):
        st.session_state.pop("selected_question_id", None)
        st.session_state.pop("editing_answer_id", None)
        st.session_state.pop("deleting_answer_id", None)
        st.session_state.pop("editing_comment_id", None)
        st.session_state.pop("deleting_comment_id", None)
        st.rerun()

    try:
        question = get_question(question_id)
        if question is None:
            st.error("Temat nie istnieje.")
            return

        answers = get_answers(question_id)
        post_ids = [question_id, *[int(answer["id"]) for answer in answers]]
        comments_by_post = get_comments(post_ids)
    except (mysql.connector.Error, RuntimeError) as exc:
        st.error(f"Nie udało się pobrać wątku: {exc}")
        return

    st.title(question["post_title"] or f"Pytanie #{question_id}")
    if question["tags"]:
        st.markdown(format_tags(question["tags"]))

    edit_info = ""
    if question.get("last_edit_date"):
        edit_info = f" · edytowano: {format_date(question['last_edit_date'])}"

    st.caption(
        f"Autor: {question['author_name']} · "
        f"{format_date(question['creation_date'])}{edit_info} · "
        f"wynik: {question['score'] or 0} · "
        f"wyświetlenia: {question['view_count'] or 0}"
    )
    st.markdown(clean_html(question["post_body"]), unsafe_allow_html=True)
    render_comments(
        post_id=question_id,
        comments=comments_by_post.get(question_id, []),
        current_user_id=current_user_id,
    )

    st.divider()
    st.header(f"Odpowiedzi ({len(answers)})")

    if not answers:
        st.info("Ten temat nie ma jeszcze odpowiedzi.")

    for answer in answers:
        answer_id = int(answer["id"])
        accepted_answer_id = question.get("accepted_answer_id")
        is_accepted = accepted_answer_id is not None and int(accepted_answer_id) == answer_id

        with st.container(border=True):
            if is_accepted:
                st.success("Zaakceptowana odpowiedź")

            edit_info = ""
            if answer.get("last_edit_date"):
                edit_info = f" · edytowano: {format_date(answer['last_edit_date'])}"

            st.caption(
                f"Autor: {answer['author_name']} · "
                f"{format_date(answer['creation_date'])}{edit_info} · "
                f"wynik: {answer['score'] or 0}"
            )
            st.markdown(clean_html(answer["post_body"]), unsafe_allow_html=True)

            render_answer_management(
                question=question,
                answer=answer,
                current_user_id=current_user_id,
            )
            render_comments(
                post_id=answer_id,
                comments=comments_by_post.get(answer_id, []),
                current_user_id=current_user_id,
            )

    st.divider()
    st.subheader("Dodaj odpowiedź")

    with st.form("answer_form", clear_on_submit=True):
        body = st.text_area(
            "Treść odpowiedzi",
            height=180,
            max_chars=10000,
        )
        submitted = st.form_submit_button("Opublikuj odpowiedź")

    if submitted:
        if len(body.strip()) < 10:
            st.warning("Treść odpowiedzi musi mieć co najmniej 10 znaków.")
        else:
            try:
                create_answer(
                    question_id=question_id,
                    user_id=current_user_id,
                    body=body.strip(),
                )
                set_flash("Odpowiedź została dodana.")
                st.rerun()
            except (mysql.connector.Error, ValueError, RuntimeError) as exc:
                st.error(str(exc))


def render_new_question(current_user_id: int) -> None:
    st.title("Nowy temat")

    with st.form("new_question_form", clear_on_submit=False):
        title = st.text_input(
            "Tytuł",
            max_chars=500,
        )
        tags = st.text_input(
            "Tagi",
            placeholder="mysql streamlit partycjonowanie",
            max_chars=500,
            help=f"Wymagany jest co najmniej jeden tag; maksymalnie {MAX_TAGS} tagów.",
        )
        body = st.text_area(
            "Treść",
            height=280,
            max_chars=10000,
        )
        submitted = st.form_submit_button("Opublikuj temat")

    if not submitted:
        return

    errors = []
    if len(title.strip()) < 5:
        errors.append("Tytuł musi mieć co najmniej 5 znaków.")
    if len(body.strip()) < 10:
        errors.append("Treść musi mieć co najmniej 10 znaków.")
    if not parse_tags(tags):
        errors.append("Podaj co najmniej jeden tag.")

    if errors:
        for error in errors:
            st.warning(error)
        return

    try:
        question_id = create_question(
            user_id=current_user_id,
            title=title.strip(),
            body=body.strip(),
            tags=tags,
        )
        st.session_state["selected_question_id"] = question_id
        st.session_state["page"] = "Tematy"
        set_flash("Temat został opublikowany.")
        st.rerun()
    except (mysql.connector.Error, ValueError, RuntimeError) as exc:
        st.error(str(exc))


def render_sidebar() -> tuple[str, int]:
    st.sidebar.title("Forum")

    current_page = st.session_state.get("page", "Tematy")
    page = st.sidebar.radio(
        "Nawigacja",
        ["Tematy", "Nowy temat"],
        index=0 if current_page == "Tematy" else 1,
    )
    st.session_state["page"] = page

    st.sidebar.divider()
    st.sidebar.subheader("Aktywny użytkownik")
    user_id = int(
        st.sidebar.number_input(
            "ID użytkownika",
            min_value=1,
            value=int(st.session_state.get("current_user_id", 1)),
            step=1,
        )
    )
    st.session_state["current_user_id"] = user_id

    try:
        user = get_user(user_id)
        if user:
            st.sidebar.success(
                f"{user['display_name'] or '[bez nazwy]'}\n\n"
                f"Reputacja: {user['reputation'] or 0}"
            )
        else:
            st.sidebar.warning("Nie znaleziono użytkownika.")
    except (mysql.connector.Error, RuntimeError) as exc:
        st.sidebar.error(f"Błąd połączenia: {exc}")

    st.sidebar.caption(
        "To jest wersja demonstracyjna. Autoryzacja operacji opiera się na wybranym ID użytkownika."
    )
    return page, user_id


def main() -> None:
    show_flash()

    try:
        page, current_user_id = render_sidebar()
    except RuntimeError as exc:
        st.error(str(exc))
        st.code(
            """
[mysql]
host = "127.0.0.1"
port = 3306
user = "forum"
password = "haslo"
database = "testdbp"
pool_size = 5
            """.strip(),
            language="toml",
        )
        st.stop()

    selected_question_id = st.session_state.get("selected_question_id")

    if page == "Nowy temat":
        render_new_question(current_user_id)
    elif selected_question_id is not None:
        render_thread(int(selected_question_id), current_user_id)
    else:
        render_question_list()


if __name__ == "__main__":
    main()
