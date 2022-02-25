from fastapi import APIRouter
from db.artist_repository import Artists

artists = APIRouter()


@artists.get("/")
def get(limit: int = 10, offset: int = 0):
    return Artists().find_all(limit=limit, offset=offset)
