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
    with repo as ctx_repo:
        return ctx_repo.get_album()


@router.get("/unrelated")
@inject
def get_unrelated(
    repo: GraphRepository = Depends(Provide[Container.graph_repo]),
) -> List[Album]:
    with repo as ctx_repo:
        return ctx_repo.get_unrelated_albums()


@router.get("/{id_}")
@inject
def get_one(
    id_: int, repo: GraphRepository = Depends(Provide[Container.graph_repo])
) -> Album:
    with repo as ctx_repo:
        return (res := ctx_repo.get_album(id_))[0] if res else None
