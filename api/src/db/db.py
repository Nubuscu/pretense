import psycopg2
import os
import contextlib

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


if __name__ == "__main__":
    with get_cursor() as cursor:
        cursor.execute("SELECT * FROM album")
        print(cursor.fetchall())
