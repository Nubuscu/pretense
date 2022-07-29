from typing import List

from fastapi import APIRouter, Depends
from src.models import Topic
from src.graph import GraphRepository

router = APIRouter()


@router.get("/")
def get(repo: GraphRepository = Depends(GraphRepository)) -> List[Topic]:
    return repo.get_topic()


@router.get("/{id_}")
def get_one(id_: int, repo: GraphRepository = Depends(GraphRepository)) -> Topic:
    return repo.get_topic(id_)
