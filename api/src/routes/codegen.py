from typing import Any, Callable, Dict, List
from uuid import UUID
from pymongo.collection import Collection
from fastapi import APIRouter, Depends
from src.models import Base
from typing import TypeVar, Generic


def router_for_collection(get_collection_fn: Callable, type_: Base) -> APIRouter:
    router = APIRouter()

    @router.get("/")
    def get(collection: Collection = Depends(get_collection_fn)) -> List[Base]:
        return [type_.from_orm(t) for t in collection.find()]

    @router.get("/{id_}")
    def get_one(id_: UUID, collection: Collection = Depends(get_collection_fn)) -> Base:
        resp = collection.find_one({"id": id_})
        return type_.from_orm(resp) if resp else None

    return router
