from typing import List
from uuid import UUID

from fastapi import APIRouter, Depends
from src.models import Album
from src.graph import GraphRepository

router = APIRouter()


@router.get("/")
def get(repo: GraphRepository = Depends(GraphRepository)) -> List[Album]:
    return repo.get_album()


@router.get("/{id_}")
def get_one(id_: int, repo: GraphRepository = Depends(GraphRepository)) -> Album:
    return repo.get_album(id_)
