from .db import BaseRepository, get_cursor
import logging
from models import Review

SELECT_ALL_SQL = "SELECT id, title, body FROM review LIMIT %(limit)s OFFSET %(offset)s"
SELECT_ONE_SQL = "SELECT id, title, body FROM review WHERE id = %(id)s LIMIT 1"


class Reviews(BaseRepository):
    def find_all(self, limit=10, offset=0):
        with get_cursor() as cursor:
            cursor.execute(SELECT_ALL_SQL, {"limit": limit, "offset": offset})
            results = cursor.fetchall()
        return [Review(id=id_, name=name, body=body) for (id_, name, body) in results]

    def by_id(self, id_):
        with get_cursor() as cursor:
            cursor.execute(SELECT_ONE_SQL, {"id": id_})
            result = cursor.fetchone()
            review_id, title, body = result
        return Review(id=review_id, title=title, body=body)

    def by_tag(self, tag_id):
        return super().by_tag(tag_id)
