from typing import Optional
from fastapi import APIRouter
from src.models import GraphDTO
from src.db.topic_repository import Topics

topics = APIRouter()


@topics.get("/")
def get(tag: Optional[int] = None, limit: int = 10, offset: int = 0):
    if tag:
        return Topics().by_tag(tag)
    return Topics().find_all(limit=limit, offset=offset)


@topics.get("/{id_}")
def get_one(id_: int):
    return Topics().by_id(id_)
