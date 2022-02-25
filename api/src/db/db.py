import psycopg2
import os
import contextlib
from abc import ABC, abstractmethod

_CONNECTION_DETAILS = {
    "database": os.environ["DB_DATABASE"],
    "host": os.environ["DB_HOST"],
    "port": os.environ["DB_PORT"],
    "user": os.environ["DB_USER"],
    "password": os.environ["DB_PASS"],
}


@contextlib.contextmanager
def get_cursor():
    with psycopg2.connect(**_CONNECTION_DETAILS) as conn:
        with conn.cursor() as cursor:
            yield cursor


class BaseRepository(ABC):
    @abstractmethod
    def by_id(self, id_):
        pass

    @abstractmethod
    def by_tag(self, tag_id):
        pass

    @abstractmethod
    def find_all(self, limit=10, offset=0):
        pass


if __name__ == "__main__":
    with get_cursor() as cursor:
        title = "Sonder"
        artist = "Tesseract"
        cursor.execute(
            """
        INSERT INTO album (title) VALUES (%(title)s) RETURNING id;
        """,
            {"title": title},
        )
        album_id = cursor.fetchone()[0]

        cursor.execute(
            """
        INSERT INTO artist (name) VALUES (%(name)s) RETURNING id;
        """,
            {"name": artist},
        )
        artist_id = cursor.fetchone()[0]

        cursor.execute(
            """
        INSERT INTO artist (name) VALUES (%(name)s) RETURNING id;
        """,
            {"name": "Jeff"},
        )
        artist2_id = cursor.fetchone()[0]

        cursor.execute(
            """
        INSERT INTO rel_album_artist (album_id, artist_id) VALUES (%(album_id)s, %(artist_id)s);
        """,
            {"album_id": album_id, "artist_id": artist_id},
        )
        cursor.execute(
            """
        INSERT INTO rel_album_artist (album_id, artist_id) VALUES (%(album_id)s, %(artist_id)s);
        """,
            {"album_id": album_id, "artist_id": artist2_id},
        )
