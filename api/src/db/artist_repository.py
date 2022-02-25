from .db import BaseRepository, get_cursor
import logging
from models import Artist

LOG = logging.getLogger(__name__)

SELECT_ALL_SQL = "SELECT id, name FROM artist LIMIT %(limit)s OFFSET %(offset)s"


class Artists(BaseRepository):
    def find_all(self, limit=10, offset=0):
        with get_cursor() as cursor:
            cursor.execute(
                SELECT_ALL_SQL,
                {"limit": limit, "offset": offset},
            )
            results = cursor.fetchall()
        out = []
        for (id_, name) in results:
            a = Artist(_id=id_, name=name)
            out.append(a)
        return out

    def by_id(self, id_):
        return super().by_id(id_)

    def by_tag(self, tag_id):
        return super().by_tag(tag_id)
