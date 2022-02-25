from typing import Optional
from fastapi import APIRouter
from db.album_repository import Albums

albums = APIRouter()


@albums.get("/")
def get(tag: Optional[int] = None, limit: int = 10, offset: int = 0):
    if tag:
        return Albums().by_tag(tag)
    return Albums().find_all(limit=limit, offset=offset)


@albums.get("/{id_}")
def get_one(id_):
    return Albums().by_id(id_)
