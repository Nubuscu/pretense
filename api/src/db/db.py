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
