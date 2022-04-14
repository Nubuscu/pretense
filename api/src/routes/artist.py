from typing import Optional
from fastapi import APIRouter
from src.db.artist_repository import Artists

artists = APIRouter()


@artists.get("/")
def get(tag: Optional[int] = None, limit: int = 10, offset: int = 0):
    if tag:
        return Artists().by_tag(tag)
    return Artists().find_all(limit=limit, offset=offset)


@artists.get("/{id_}")
def get_one(id_):
    return Artists().by_id(id_)
