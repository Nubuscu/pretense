import json
from .db import BaseRepository, get_cursor
import logging
from src.models import Topic, TopicMeta
from src.db.review_repository import Reviews
from src.db.album_repository import Albums
from src.db.artist_repository import Artists

LOG = logging.getLogger(__name__)

SELECT_ALL_SQL = "SELECT id, name FROM topic LIMIT %(limit)s OFFSET %(offset)s"

SELECT_ONE_SQL = (
    "SELECT"
    " topic.id,"
    " topic.name,"
    " topic.meta,"
    " array_agg(rel_album_topic.album_id),"
    " array_agg(rel_artist_topic.artist_id),"
    " array_agg(rel_review_topic.review_id)"
    " FROM topic"
    " LEFT JOIN rel_album_topic ON rel_album_topic.topic_id = topic.id"
    " LEFT JOIN rel_artist_topic ON rel_artist_topic.topic_id = topic.id"
    " LEFT JOIN rel_review_topic ON rel_review_topic.topic_id = topic.id"
    " WHERE topic.id = %(topic_id)s"
    " GROUP BY 1, 2"
)


class Topics(BaseRepository):
    def find_all(self, limit=10, offset=0):
        with get_cursor() as cursor:
            cursor.execute(
                SELECT_ALL_SQL,
                {"limit": limit, "offset": offset},
            )
            results = cursor.fetchall()
        return [Topic(id=id_, name=name) for (id_, name) in results]

    def by_id(self, id_) -> Topic:
        with get_cursor() as cursor:
            cursor.execute(SELECT_ONE_SQL, {"topic_id": id_})
            result = cursor.fetchone()
            if not result:
                return None
            (topic_id, name, meta, album_ids, artist_ids, review_ids) = result
            if meta:
                try:
                    meta = json.loads(meta)
                except ValueError:
                    LOG.warning(
                        "Failed to load metadata of topic %s, %s", topic_id, meta
                    )
                    meta = None

            return Topic(
                id=topic_id,
                name=name,
                meta=TopicMeta(layout=meta),
                albums=[
                    Albums().by_id(album_id) for album_id in set(album_ids) if album_id
                ],
                artists=[
                    Artists().by_id(artist_id)
                    for artist_id in set(artist_ids)
                    if artist_id
                ],
                reviews=[
                    Reviews().by_id(review_id)
                    for review_id in set(review_ids)
                    if review_id
                ],
            )

    def by_tag(self, tag_id):
        return super().by_tag(tag_id)
