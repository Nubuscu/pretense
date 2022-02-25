from .db import BaseRepository, get_cursor
import logging
from models import Review

SELECT_ALL_SQL = "SELECT id, title, body FROM review LIMIT %(limit)s OFFSET %(offset)s"


class Reviews(BaseRepository):
    def find_all(self, limit=10, offset=0):
        with get_cursor() as cursor:
            cursor.execute(
                SELECT_ALL_SQL,
                {"limit": limit, "offset": offset},
            )
            results = cursor.fetchall()
        return [Review(id=id_, name=name, body=body) for (id_, name, body) in results]

    def by_id(self, id_):
        return super().by_id(id_)

    def by_tag(self, tag_id):
        return super().by_tag(tag_id)
