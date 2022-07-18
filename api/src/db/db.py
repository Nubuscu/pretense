import logging
import os
from typing import Any, List
from uuid import uuid4, UUID
from datetime import datetime
from fastapi import Depends
from pymongo.mongo_client import MongoClient
from pymongo.database import Database
from pymongo.collection import Collection
from pymongo.errors import ConnectionFailure
from pydantic.utils import GetterDict
from cachetools import cached, TTLCache

LOG = logging.getLogger(__name__)

_CONNECTION_DETAILS = {
    "host": os.environ["DB_HOST"],
    "port": os.environ["DB_PORT"],
    "user": os.environ["DB_USER"],
    "pass": os.environ["DB_PASS"],
}
DB_NAME = os.environ["DB_DATABASE"]
URI = "mongodb://{user}:{pass}@{host}:{port}".format(**_CONNECTION_DETAILS)
CACHE_DURATION = 30000


# caching the db connection
# since this is read-only and (aside my annoyance while developing)
# it's probably safe to cache some of the results too
# especially for the mappers which will likely hit the same things over
@cached(TTLCache(maxsize=1024, ttl=CACHE_DURATION))
def get_db() -> Database:
    # TODO reconsider technology or approach
    #  - gremlindb (or comparable) for graph-based storage
    #  - mongo-like nosql recommends denormalizing data, unlike what I've done here
    #  - could probably aggregate better on query
    try:
        client = MongoClient(URI, uuidRepresentation="standard")
        return client[DB_NAME]
    except ConnectionFailure:
        LOG.error("Failed to connect to db; check credentials.")
        raise


def get_topics() -> Collection:
    return get_db()["topics"]


def get_albums() -> Collection:
    return get_db()["albums"]


def get_artists() -> Collection:
    return get_db()["artists"]


def get_reviews() -> Collection:
    return get_db()["reviews"]


def get_tags() -> Collection:
    return get_db()["tags"]


def safe_insert(collection: Collection, unique_by_query: dict, document: dict) -> dict:
    """Maintain uniqueness by a query value, only inserting if an entry does not exist.

    Returns found/created document."""
    document.update(**{"created_at": datetime.utcnow(), "id": uuid4()})
    ret_val = collection.find_one(unique_by_query)
    if not ret_val:
        collection.insert_one(document)
        ret_val = collection.find_one(unique_by_query)
    return ret_val


def _ensure_uuid(val, default):
    if val is None:
        return default
    if isinstance(val, UUID):
        return val
    try:
        return UUID(val)
    except:
        return default


def all_in_collection_by_ids(collection: Collection, ids: List[UUID]):
    return [collection.find_one({"id": id_}) for id_ in ids]


class DefaultMapper(GetterDict):
    """Running pydantic models in ORM mode means I need an ORM.

    This is it.
    flow:
    get something from the db by id
    do type_.from_orm(resp)
    this (or subclass) mapper translates resp into a pydantic model
    before it gets validated.
    """

    def get(self, key: str, default: Any, **_kwargs) -> Any:
        # TODO handle tags
        if self._obj is None:
            return default
        if key == "id":
            return _ensure_uuid(self._obj.get(key), default)
        if key == "reviews":
            return all_in_collection_by_ids(get_reviews(), self._obj["reviews"])
        if key == "tags":
            return all_in_collection_by_ids(get_reviews(), self._obj["tags"])
        return self._obj.get(key, default)


class AlbumMapper(DefaultMapper):
    def get(self, key: str, default: Any, **_kwargs) -> Any:
        if key == "artists":
            return all_in_collection_by_ids(get_artists(), self._obj["artists"])
        return super().get(key, default, **_kwargs)


class TopicMapper(DefaultMapper):
    def get(self, key: str, default: Any, **_kwargs) -> Any:
        if key == "artists":
            return all_in_collection_by_ids(get_artists(), self._obj["artists"])
        if key == "albums":
            return all_in_collection_by_ids(get_albums(), self._obj["albums"])
        return super().get(key, default, **_kwargs)
