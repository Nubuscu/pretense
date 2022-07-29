from typing import List

from fastapi import APIRouter, Depends
from dependency_injector.wiring import inject, Provide
from src.models import Album
from src.graph import GraphRepository

from src.di import Container

router = APIRouter()


@router.get("/")
@inject
def get(repo: GraphRepository = Depends(Provide[Container.graph_repo])) -> List[Album]:
    return repo.get_album()


@router.get("/{id_}")
def get_one(
    id_: int, repo: GraphRepository = Depends(Provide[Container.graph_repo])
) -> Album:
    return repo.get_album(id_)
