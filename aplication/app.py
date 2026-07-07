from __future__ import annotations

from contextlib import contextmanager
from datetime import datetime
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


def normalize_tags(raw_tags: str) -> str | None:
    tags = []
    for part in raw_tags.replace(",", " ").split():
        tag = part.strip().strip("<>").lower()
        if tag and tag not in tags:
            tags.append(tag)

    if not tags:
        return None

    return "".join(f"<{tag}>" for tag in tags[:10])


def format_tags(raw_tags: str | None) -> str:
    if not raw_tags:
        return ""

    tags = raw_tags.replace("><", " ").replace("<", "").replace(">", "").split()
    return " ".join(f"`{tag}`" for tag in tags)


def format_date(value: datetime | None) -> str:
    if value is None:
        return "brak daty"
    return value.strftime("%Y-%m-%d %H:%M")


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


def list_questions(limit: int = 30) -> list[dict[str, Any]]:
    return fetch_all(
        """
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
        WHERE q.post_type_id = 1
        ORDER BY q.id DESC
        LIMIT %s
        """,
        (limit,),
    )


def get_question(question_id: int) -> dict[str, Any] | None:
    return fetch_one(
        """
        SELECT
            q.id,
            q.creation_date,
            q.score,
            q.view_count,
            q.post_title,
            q.post_body,
            q.tags,
            q.answer_count,
            q.comment_count,
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


def create_question(
    *,
    user_id: int,
    title: str,
    body: str,
    tags: str,
) -> int:
    require_user(user_id)

    normalized_tags = normalize_tags(tags)
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

    if get_question(question_id) is None:
        raise ValueError("Wybrany temat nie istnieje.")

    return execute_transaction(
        [
            (
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
                VALUES (
                    2,
                    %s,
                    NOW(),
                    0,
                    %s,
                    %s,
                    NOW(),
                    0,
                    %s
                )
                """,
                (question_id, body, user_id, CONTENT_LICENSE),
            ),
            (
                """
                UPDATE posts
                SET
                    answer_count = COALESCE(answer_count, 0) + 1,
                    last_activity_date = NOW()
                WHERE id = %s
                  AND post_type_id = 1
                """,
                (question_id,),
            ),
        ]
    )


def create_comment(*, post_id: int, user_id: int, body: str) -> int:
    require_user(user_id)

    return execute_transaction(
        [
            (
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
            ),
            (
                """
                UPDATE posts
                SET
                    comment_count = COALESCE(comment_count, 0) + 1,
                    last_activity_date = NOW()
                WHERE id = %s
                """,
                (post_id,),
            ),
        ]
    )


def render_comments(
    *,
    post_id: int,
    comments: list[dict[str, Any]],
    current_user_id: int,
) -> None:
    with st.expander(f"Komentarze ({len(comments)})"):
        if comments:
            for comment in comments:
                st.markdown(
                    f"**{comment['author_name']}** · "
                    f"{format_date(comment['creation_date'])} · "
                    f"wynik: {comment['score'] or 0}"
                )
                st.write(comment["comment_text"] or "")
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
    st.caption("Najnowsze pytania z tabeli `posts`.")

    limit = st.select_slider(
        "Liczba tematów",
        options=[10, 20, 30, 50, 100],
        value=30,
    )

    try:
        questions = list_questions(limit)
    except (mysql.connector.Error, RuntimeError) as exc:
        st.error(f"Nie udało się pobrać tematów: {exc}")
        return

    if not questions:
        st.info("Brak pytań w bazie.")
        return

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
                    f"{format_date(question['creation_date'])}"
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
                    st.session_state["selected_question_id"] = int(question["id"])
                    st.rerun()


def render_thread(question_id: int, current_user_id: int) -> None:
    if st.button("← Wróć do tematów"):
        st.session_state.pop("selected_question_id", None)
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

    st.caption(
        f"Autor: {question['author_name']} · "
        f"{format_date(question['creation_date'])} · "
        f"wynik: {question['score'] or 0}"
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
        with st.container(border=True):
            st.caption(
                f"Autor: {answer['author_name']} · "
                f"{format_date(answer['creation_date'])} · "
                f"wynik: {answer['score'] or 0}"
            )
            st.markdown(clean_html(answer["post_body"]), unsafe_allow_html=True)
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
        if not body.strip():
            st.warning("Treść odpowiedzi nie może być pusta.")
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
        "To jest wersja demonstracyjna. Nie implementuje logowania ani autoryzacji."
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
