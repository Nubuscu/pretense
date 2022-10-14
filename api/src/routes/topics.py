from typing import List

from fastapi import APIRouter, Depends
from dependency_injector.wiring import inject, Provide

from src.di import Container
from src.models import Topic
from src.graph import GraphRepository

router = APIRouter()


@router.get("/")
@inject
def get(repo: GraphRepository = Depends(Provide[Container.graph_repo])) -> List[Topic]:
    with repo as ctx_repo:
        return ctx_repo.get_topic()


@router.get("/{id_}")
@inject
def get_one(
    id_: int, repo: GraphRepository = Depends(Provide[Container.graph_repo])
) -> Topic:
    with repo as ctx_repo:
        res = ctx_repo.get_topic(id_)
        if res:
            return res[0]
        return None
