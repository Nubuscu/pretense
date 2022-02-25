from .db import BaseRepository, get_cursor
from models import Album, Artist
import logging

LOG = logging.getLogger(__name__)

SELECT_ALL_SQL = (
    "SELECT al.id, al.title, json_object_agg(ar.id, ar.name) FROM album al"
    " JOIN rel_album_artist rel ON rel.album_id = al.id"
    " JOIN artist ar on rel.artist_id = ar.id"
    " GROUP BY 1, 2"
    " LIMIT %(limit)s OFFSET %(offset)s"
)


class Albums(BaseRepository):
    def find_all(self, limit=10, offset=0):
        with get_cursor() as cursor:
            cursor.execute(
                SELECT_ALL_SQL,
                {"limit": limit, "offset": offset},
            )
            results = cursor.fetchall()
        out = []
        for (album_id, title, artists_json) in results:
            a = Album(
                _id=album_id,
                name=title,
                artists=[
                    Artist(_id=artist_id, name=artist_name)
                    for (artist_id, artist_name) in artists_json.items()
                ],
            )
            out.append(a)
        return out

    def by_id(self, id_):
        return super().by_id(id_)

    def by_tag(self, tag_id):
        return super().by_tag(tag_id)
