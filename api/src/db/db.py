import logging
import os
from typing import Any
from uuid import uuid4, UUID
from datetime import datetime
from fastapi import Depends
from pymongo.mongo_client import MongoClient
from pymongo.database import Database
from pymongo.collection import Collection
from pymongo.errors import ConnectionFailure
from pydantic.utils import GetterDict

LOG = logging.getLogger(__name__)

_CONNECTION_DETAILS = {
    "host": os.environ["DB_HOST"],
    "port": os.environ["DB_PORT"],
    "user": os.environ["DB_USER"],
    "pass": os.environ["DB_PASS"],
}
DB_NAME = os.environ["DB_DATABASE"]
URI = "mongodb://{user}:{pass}@{host}:{port}".format(**_CONNECTION_DETAILS)


def get_db() -> Database:
    try:
        client = MongoClient(URI, uuidRepresentation="standard")
        return client[DB_NAME]
    except ConnectionFailure:
        LOG.error("Failed to connect to db; check credentials.")


def get_topics(db: Database = Depends(get_db)) -> Collection:
    return db["topics"]


def get_albums(db: Database = Depends(get_db)) -> Collection:
    return db["albums"]


def get_artists(db: Database = Depends(get_db)) -> Collection:
    return db["artists"]


def safe_insert(collection: Collection, unique_by_query: dict, document: dict):
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


class DefaultMapper(GetterDict):
    def get(self, key: str, default: Any, **_kwargs) -> Any:
        if self._obj is None:
            return default
        if key == "id":
            return _ensure_uuid(self._obj.get(key), default)
        return self._obj.get(key, default)


class AlbumMapper(DefaultMapper):
    def get(self, key: str, default: Any, **_kwargs) -> Any:
        if key == "artists":
            # TODO can't use Depends too far away from the api.
            # need to find another way to dedupe the client
            # is there anything in app.state?
            collection = get_artists(get_db())
            return [
                collection.find_one({"id": artist_id})
                for artist_id in self._obj["artists"]
            ]
        return super().get(key, default, **_kwargs)
