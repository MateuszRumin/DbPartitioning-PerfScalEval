from datetime import datetime, time, timedelta
from time import perf_counter

import mysql.connector
import streamlit as st

st.set_page_config(page_title="Forum", layout="wide")


def sql(query, params=(), one=False):
    cfg = st.secrets["mysql"]
    connection = mysql.connector.connect(
        host=cfg.get("host", "localhost"),
        port=cfg.get("port", 3306),
        user=cfg["user"],
        password=cfg["password"],
        database=cfg["database"],
    )
    cursor = connection.cursor(dictionary=True)
    try:
        cursor.execute(query, params)
        return cursor.fetchone() if one else cursor.fetchall()
    finally:
        cursor.close()
        connection.close()


def tags(value):
    if not value:
        return ""
    values = value.replace("><", " ").replace("<", "").replace(">", "").split()
    return " ".join(f"`{value}`" for value in values)


def show_comments(post_id):
    comments = sql(
        """
        SELECT comment_text, creation_date, score
        FROM comments
        WHERE post_id = %s
        ORDER BY creation_date
        """,
        (post_id,),
    )
    with st.expander(f"Komentarze ({len(comments)})", expanded=False):
        if not comments:
            st.caption("Brak komentarzy.")
        for comment in comments:
            st.caption(
                f"{comment['creation_date']:%Y-%m-%d %H:%M} · "
                f"wynik: {comment['score'] or 0}"
            )
            st.write(comment["comment_text"] or "")


def show_history(post_id):
    history = sql(
        """
        SELECT post_history_type_id, creation_date, post_text
        FROM post_history
        WHERE post_id = %s
        ORDER BY creation_date
        """,
        (post_id,),
    )
    with st.expander(f"Historia posta ({len(history)})", expanded=False):
        if not history:
            st.caption("Brak historii posta.")
        for entry in history:
            st.caption(
                f"{entry['creation_date']:%Y-%m-%d %H:%M} · "
                f"typ: {entry['post_history_type_id']}"
            )
            if entry["post_text"]:
                st.write(entry["post_text"])
            st.divider()


def show_question(question_id):
    load_started = perf_counter()
    load_time = st.empty()

    if st.button("← Wróć"):
        del st.session_state["question_id"]
        st.rerun()

    question = sql(
        """
        SELECT id, post_title, post_body, creation_date,
               score, view_count, tags, answer_count, accepted_answer_id
        FROM posts
        WHERE id = %s AND post_type_id = 1
        """,
        (question_id,),
        one=True,
    )
    if not question:
        st.error("Nie znaleziono pytania.")
        return

    st.title(question["post_title"] or f"Pytanie {question_id}")
    st.markdown(tags(question["tags"]))
    st.caption(
        f"{question['creation_date']:%Y-%m-%d %H:%M} · "
        f"wynik: {question['score'] or 0} · "
        f"wyświetlenia: {question['view_count'] or 0}"
    )
    st.markdown(question["post_body"] or "Brak treści", unsafe_allow_html=True)
    show_comments(question_id)
    show_history(question_id)

    accepted_answer_id = question["accepted_answer_id"]
    answers = sql(
        """
        SELECT id, parent_id, post_type_id, post_body, creation_date, score
        FROM posts
        WHERE parent_id = %s
           OR id = %s
        ORDER BY (id = %s) DESC, score DESC, creation_date
        """,
        (
            question_id,
            accepted_answer_id,
            accepted_answer_id,
        ),
    )

    st.header(f"Odpowiedzi ({len(answers)})")
    if not answers:
        if (question["answer_count"] or 0) > 0:
            st.warning(
                "Pytanie ma answer_count > 0, ale w tabeli posts nie znaleziono "
                "rekordów z parent_id równym id pytania."
            )
        else:
            st.caption("Brak odpowiedzi.")

    for answer in answers:
        with st.container(border=True):
            is_accepted = answer["id"] == accepted_answer_id
            accepted = " · zaakceptowana" if is_accepted else ""
            st.caption(
                f"Odpowiedź #{answer['id']} · "
                f"{answer['creation_date']:%Y-%m-%d %H:%M} · "
                f"wynik: {answer['score'] or 0}{accepted}"
            )
            st.markdown(
                answer["post_body"] or "Brak treści odpowiedzi",
                unsafe_allow_html=True,
            )
            show_comments(answer["id"])
            show_history(answer["id"])

    load_time.caption(
        f"Czas pobrania i przygotowania danych: "
        f"{perf_counter() - load_started:.3f} s"
    )


def show_search():
    load_started = perf_counter()

    st.title("Wyszukiwanie pytań")
    load_time = st.empty()

    bounds = sql(
        """
        SELECT MIN(creation_date) AS first_date,
               MAX(creation_date) AS last_date
        FROM posts
        WHERE post_type_id = 1
        """,
        one=True,
    )
    first_date = bounds["first_date"].date()
    last_date = bounds["last_date"].date()

    
    st.session_state.setdefault("saved_title", "")
    st.session_state.setdefault("saved_tag", "")
    st.session_state.setdefault("saved_start_date", first_date)
    st.session_state.setdefault("saved_end_date", last_date)
    st.session_state.setdefault("saved_limit", 50)

    widget_defaults = {
        "title_widget": st.session_state["saved_title"],
        "tag_widget": st.session_state["saved_tag"],
        "start_date_widget": st.session_state["saved_start_date"],
        "end_date_widget": st.session_state["saved_end_date"],
        "limit_widget": st.session_state["saved_limit"],
    }
    for key, value in widget_defaults.items():
        if key not in st.session_state:
            st.session_state[key] = value

    title = st.text_input("Tytuł zawiera", key="title_widget")
    tag = st.text_input(
        "Tag",
        placeholder="np. mysql",
        key="tag_widget",
    )
    start_date = st.date_input(
        "Data 1",
        min_value=first_date,
        max_value=last_date,
        key="start_date_widget",
    )
    end_date = st.date_input(
        "Data 2",
        min_value=first_date,
        max_value=last_date,
        key="end_date_widget",
    )
    limit = st.number_input(
        "Limit",
        min_value=1,
        max_value=500,
        key="limit_widget",
    )

    st.session_state["saved_title"] = title
    st.session_state["saved_tag"] = tag
    st.session_state["saved_start_date"] = start_date
    st.session_state["saved_end_date"] = end_date
    st.session_state["saved_limit"] = int(limit)

    if start_date > end_date:
        st.info("Data początkowa nie może być późniejsza niż końcowa.")
        return

    start = datetime.combine(start_date, time.min)
    end = datetime.combine(end_date + timedelta(days=1), time.min)

    conditions = [
        "post_type_id = 1",
        "creation_date >= %s",
        "creation_date < %s",
    ]
    params = [start, end]

    if title.strip():
        conditions.append("post_title LIKE %s")
        params.append(f"%{title.strip()}%")

    clean_tag = tag.strip().lower().replace("<", "").replace(">", "")
    if clean_tag:
        conditions.append("tags LIKE %s")
        params.append(f"%<{clean_tag}>%")

    params.append(int(limit))
    questions = sql(
        f"""
        SELECT id, post_title, creation_date,
               score, answer_count, tags
        FROM posts
        WHERE {' AND '.join(conditions)}
        ORDER BY creation_date DESC
        LIMIT %s
        """,
        tuple(params),
    )

    st.write(f"Znaleziono: {len(questions)}")
    for question in questions:
        st.subheader(question["post_title"] or f"Pytanie {question['id']}")
        st.markdown(tags(question["tags"]))
        st.caption(
            f"{question['creation_date']:%Y-%m-%d %H:%M} · "
            f"wynik: {question['score'] or 0} · "
            f"odpowiedzi: {question['answer_count'] or 0}"
        )
        if st.button("Otwórz", key=question["id"]):
            st.session_state["question_id"] = question["id"]
            st.rerun()
        st.divider()

    load_time.caption(
        f"Czas pobrania i przygotowania danych: "
        f"{perf_counter() - load_started:.3f} s"
    )


try:
    if "question_id" in st.session_state:
        show_question(st.session_state["question_id"])
    else:
        show_search()
except (mysql.connector.Error, KeyError, TypeError) as error:
    st.error(f"Błąd: {error}")
