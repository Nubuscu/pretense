from src.db.db import get_cursor

INSERT_TOPIC_SQL = "INSERT INTO topic (name) VALUES (%(name)s)"
SELECT_TOPIC_SQL = "SELECT id FROM topic WHERE name = %(name)s"
INSERT_REVIEW_SQL = "INSERT INTO review (title, body) VALUES (%(title)s, %(body)s)"
SELECT_REVIEW_SQL = "SELECT id FROM review WHERE title = %(title)s"
INSERT_REVIEW_TOPIC_REL_SQL = "INSERT INTO rel_review_topic (review_id, topic_id) VALUES (%(review_id)s, %(topic_id)s)"

INSERT_ALBUM_RELATIONS_SQL = """
WITH cte_albums AS (
    SELECT id, %(topic_id)s FROM album
    WHERE title IN (
        'Sound Awake',
        'One',
        'It Hates You',
        'Koloss',
        'Laugh Tracks',
        'I Let It in and It Took Everything',
        'The Horrifying Truth'
        )
)
INSERT INTO rel_album_topic (album_id, topic_id) SELECT * FROM cte_albums
"""

body = """
Hi there! This is the first "review" on here, so I'll use it as an intro rant:
This is a side project to help me map out my music tastes and discoveries - where they came from, what lead to what, etc.
I'll probably write it like a series of reviews, which - as a nameless voice on the internet - can sound quite pretentious.
Some grounding opinions:
- there is no universally good/bad music. If you enjoy it, it's good (for you)
- being exact about subgenres is often silly but sometimes helpful
- I typically listen to some form of metal these days
- my roots are mostly in niche NZ and/or Christian rock, which progressed into Christian metalcore, then into other metalcore/deathcore
- some rap, some funk, some electronic thrown in somewhere along the line
- djent is a real genre, but only if you want it to be

Some things I seem to keep recommending to people are somewhat attached here. I'll try to leave words on each of them individually.
"""

with get_cursor() as cursor:
    cursor.execute(INSERT_TOPIC_SQL, {"name": "intro"})
    cursor.execute(SELECT_TOPIC_SQL, {"name": "intro"})
    topic_id = cursor.fetchone()

    cursor.execute(INSERT_REVIEW_SQL, {"title": "Gotta start somewhere", "body": body})
    cursor.execute(SELECT_REVIEW_SQL, {"title": "Gotta start somewhere"})
    review_id = cursor.fetchone()

    cursor.execute(
        INSERT_REVIEW_TOPIC_REL_SQL, {"review_id": review_id, "topic_id": topic_id}
    )

    cursor.execute(INSERT_ALBUM_RELATIONS_SQL, {"topic_id": topic_id})
