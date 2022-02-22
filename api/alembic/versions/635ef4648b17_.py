"""Initial table generation

Revision ID: 635ef4648b17
Revises: 
Create Date: 2022-02-19 12:13:25.275452

"""
from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision = '635ef4648b17'
down_revision = None
branch_labels = None
depends_on = None

SQL = """
-- Your SQL goes here
CREATE TABLE IF NOT EXISTS album (
    id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL
);
CREATE TABLE IF NOT EXISTS artist (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);
CREATE TABLE IF NOT EXISTS review (
    id SERIAL PRIMARY KEY,
    title VARCHAR,
    body VARCHAR NOT NULL
);
CREATE TABLE IF NOT EXISTS topic (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);

-- relation tables. rel_a_b in alphabetic order.
-- another option would be one big mostly-null table but I fear for the scalability
-- speed-wise, at least. This is a lot of tables too.

-- album_thing
CREATE TABLE IF NOT EXISTS rel_album_artist (
    id SERIAL PRIMARY KEY,
    album_id INTEGER REFERENCES album,
    artist_id INTEGER REFERENCES artist
);
CREATE TABLE IF NOT EXISTS rel_album_review (
    id SERIAL PRIMARY KEY,
    album_id INTEGER REFERENCES album,
    review_id INTEGER REFERENCES review
);
CREATE TABLE IF NOT EXISTS rel_album_topic (
    id SERIAL PRIMARY KEY,
    album_id INTEGER REFERENCES album,
    topic_id INTEGER REFERENCES topic
);
-- artist_thing
CREATE TABLE IF NOT EXISTS rel_artist_review (
    id SERIAL PRIMARY KEY,
    artist_id INTEGER REFERENCES artist,
    review_id INTEGER REFERENCES review
);
CREATE TABLE IF NOT EXISTS rel_artist_topic (
    id SERIAL PRIMARY KEY,
    artist_id INTEGER REFERENCES artist,
    topic_id INTEGER REFERENCES topic
);
-- review_thing
CREATE TABLE IF NOT EXISTS rel_review_topic (
    id SERIAL PRIMARY KEY,
    review_id INTEGER REFERENCES review,
    topic_id INTEGER REFERENCES topic
);
"""

def upgrade():
    op.execute(SQL)


def downgrade():
    pass
