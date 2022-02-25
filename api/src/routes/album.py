from fastapi import APIRouter
from db.album_repository import Albums

albums = APIRouter()


@albums.get("/")
def get(limit: int = 10, offset: int = 0):
    return Albums().find_all(limit=limit, offset=offset)
