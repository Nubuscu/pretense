from .db import BaseRepository, get_cursor
import logging
from models import Topic

LOG = logging.getLogger(__name__)

SELECT_ALL_SQL = "SELECT id, name FROM topic LIMIT %(limit)s OFFSET %(offset)s"


class Topics(BaseRepository):
    def find_all(self, limit=10, offset=0):
        with get_cursor() as cursor:
            cursor.execute(
                SELECT_ALL_SQL,
                {"limit": limit, "offset": offset},
            )
            results = cursor.fetchall()
        return [Topic(id=id_, name=name) for (id_, name) in results]

    def by_id(self, id_):
        return super().by_id(id_)

    def by_tag(self, tag_id):
        return super().by_tag(tag_id)
