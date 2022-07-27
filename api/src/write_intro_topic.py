from graph import get_cursor

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

title = "Starting somewhere"

body = """
I like music and I didn't think the internet had enough faceless opinions on subjective matter, so here we are.
This project is intended for me to grapple with what I listen to, how my tastes evolved over time, and what discoveries were made.
It's also an excuse for me to play with some new(ish, to me) technologies.

If I ever publish this, hi Mum!

Some grounding opinions and factoids:
- being exact about subgenres is often silly but sometimes helpful
- there is no universally good/bad music. If you enjoy it, it's good, even if your opinion is in the minority
- I typically listen to some form of metal these days
- my roots are mostly in niche NZ and/or Christian rock, which progressed into Christian metalcore, then into other metalcore/deathcore
- some mixture of rap, funk, and electronic thrown in along the line
- djent is a real genre, but only if you want it to be

Some things I seem to keep recommending to people are attached here.
I'll try to leave words on each of them individually, pending some work.
"""

with get_cursor() as cursor:
    cursor.execute(INSERT_TOPIC_SQL, {"name": "intro"})
    cursor.execute(SELECT_TOPIC_SQL, {"name": "intro"})
    topic_id = cursor.fetchone()

    cursor.execute(INSERT_REVIEW_SQL, {"title": title, "body": body})
    cursor.execute(SELECT_REVIEW_SQL, {"title": title})
    review_id = cursor.fetchone()

    cursor.execute(
        INSERT_REVIEW_TOPIC_REL_SQL, {"review_id": review_id, "topic_id": topic_id}
    )

    cursor.execute(INSERT_ALBUM_RELATIONS_SQL, {"topic_id": topic_id})
