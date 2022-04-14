from typing import Optional
from fastapi import APIRouter
from src.db.review_repository import Reviews

reviews = APIRouter()


@reviews.get("/")
def get(tag: Optional[int] = None, limit: int = 10, offset: int = 0):
    if tag:
        return Reviews().by_tag(tag)
    return Reviews().find_all(limit=limit, offset=offset)


@reviews.get("/{id_}")
def get_one(id_):
    return Reviews().by_id(id_)
