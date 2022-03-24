from .db import BaseRepository, get_cursor
from models import Album, Artist
from typing import List
import logging

LOG = logging.getLogger(__name__)

SELECT_ALL_SQL = (
    "SELECT al.id, al.title, json_object_agg(ar.id, ar.name) FROM album al"
    " JOIN rel_album_artist rel ON rel.album_id = al.id"
    " JOIN artist ar on rel.artist_id = ar.id"
    " GROUP BY 1, 2"
    " LIMIT %(limit)s OFFSET %(offset)s"
)

SELECT_BY_ID_SQL = (
    "SELECT al.id, al.title, json_object_agg(ar.id, ar.name) FROM album al"
    " JOIN rel_album_artist rel ON rel.album_id = al.id"
    " JOIN artist ar on rel.artist_id = ar.id"
    " WHERE al.id = %(id)s"
    " GROUP BY 1, 2"
    " LIMIT 1"
)


class Albums(BaseRepository):
    def find_all(self, limit=10, offset=0) -> List[Album]:
        with get_cursor() as cursor:
            cursor.execute(
                SELECT_ALL_SQL,
                {"limit": limit, "offset": offset},
            )
            results = cursor.fetchall()
        return [
            Album(
                id=album_id,
                name=title,
                artists=[
                    Artist(id=artist_id, name=artist_name)
                    for (artist_id, artist_name) in artists_json.items()
                ],
            )
            for (album_id, title, artists_json) in results
        ]

    def by_id(self, id_) -> Album:
        with get_cursor() as cursor:
            cursor.execute(SELECT_BY_ID_SQL, {"id": id_})
            result = cursor.fetchone()
            if not result:
                return None
            album_id, title, artists_json = result
            return Album(
                id=album_id,
                name=title,
                artists=[
                    Artist(id=artist_id, name=artist_name)
                    for (artist_id, artist_name) in artists_json.items()
                ],
            )

    def by_tag(self, tag_id) -> List[Album]:
        return super().by_tag(tag_id)
