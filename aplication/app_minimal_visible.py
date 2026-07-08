from datetime import datetime, time, timedelta

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
    with st.expander(f"Komentarze ({len(comments)})", expanded=bool(comments)):
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
    with st.expander(f"Historia posta ({len(history)})", expanded=bool(history)):
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
    if st.button("← Wróć"):
        del st.session_state["question_id"]
        st.rerun()

    question = sql(
        """
        SELECT id, post_title, post_body, creation_date,
               score, view_count, tags
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

    answers = sql(
        """
        SELECT id, post_body, creation_date, score
        FROM posts
        WHERE parent_id = %s AND post_type_id = 2
        ORDER BY score DESC, creation_date
        """,
        (question_id,),
    )

    st.header(f"Odpowiedzi ({len(answers)})")
    if not answers:
        st.caption("Brak odpowiedzi.")

    for answer in answers:
        with st.container(border=True):
            st.caption(
                f"{answer['creation_date']:%Y-%m-%d %H:%M} · "
                f"wynik: {answer['score'] or 0}"
            )
            st.markdown(answer["post_body"] or "Brak treści", unsafe_allow_html=True)
            show_comments(answer["id"])
            show_history(answer["id"])


def show_search():
    st.title("Wyszukiwanie pytań")

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

    title = st.text_input("Tytuł zawiera")
    tag = st.text_input("Tag", placeholder="np. mysql")
    start_date = st.date_input(
        "Zakres dat",
        value=(first_date),
        min_value=first_date,
        max_value=last_date,
    )
    end_date = st.date_input(
        "Zakres dat (drugi)",
        value= (last_date),
        min_value=first_date,
        max_value=last_date,
    )
    limit = st.number_input("Limit", min_value=1, max_value=500, value=50)

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


try:
    if "question_id" in st.session_state:
        show_question(st.session_state["question_id"])
    else:
        show_search()
except (mysql.connector.Error, KeyError, TypeError) as error:
    st.error(f"Błąd: {error}")
