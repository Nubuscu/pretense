from functools import partial
import logging

from typing import List, Optional
from fastapi import Depends
import strawberry
from strawberry.types import Info

from dependency_injector.wiring import inject, Provide
from strawberry.fastapi import GraphQLRouter, BaseContext
from src.graph import GraphRepository

from src import models
from src.di import Container

LOG = logging.getLogger(__name__)
## base resolvers for the query root-level objects
def get_topics(root, info: Info, id: Optional[int] = None) -> List["Topic"]:
    with info.context.repo as ctx_repo:
        return ctx_repo.get_topic(id)


def get_albums(
    root, info: Info, id: Optional[int] = None, name: Optional[str] = None
) -> List["Album"]:
    with info.context.repo as ctx_repo:
        return ctx_repo.get_album(name=name, id_=id)


def get_artists(
    root, info: Info, id: Optional[int] = None, name: Optional[str] = None
) -> List["Artist"]:
    with info.context.repo as ctx_repo:
        return ctx_repo.get_artist(name=name, id_=id)


def get_related_for_single(root, info, root_type, return_type) -> list:
    with info.context.repo as ctx_repo:
        return ctx_repo.get_related_for(root.id_, root_type, return_type)


@strawberry.experimental.pydantic.type(model=models.Tag, all_fields=True)
class Tag:
    pass


@strawberry.experimental.pydantic.type(model=models.Review, all_fields=True)
class Review:
    pass


@strawberry.experimental.pydantic.type(model=models.Base)
class Base:
    id_: strawberry.auto
    # hiding tags for now
    # reviews need to be done individually, different query for each


@strawberry.experimental.pydantic.type(model=models.Artist)
class Artist(Base):
    name: str

    @strawberry.field()
    def albums(root, info) -> List["Album"]:
        return get_related_for_single(root, info, "artist", "album")


@strawberry.experimental.pydantic.type(model=models.Album)
class Album(Base):
    name: str

    @strawberry.field()
    def artists(root, info) -> List[Artist]:
        return get_related_for_single(root, info, "album", "artist")

    @strawberry.field()
    def topics(root, info) -> List["Topic"]:
        return get_related_for_single(root, info, "album", "topic")


@strawberry.experimental.pydantic.type(model=models.Topic)
class Topic(Base):
    name: str

    @strawberry.field()
    def albums(root, info) -> List[Album]:
        return get_related_for_single(root, info, "topic", "album")

    # hiding artists in in topic - everything's at the album level for now
    @strawberry.field()
    def reviews(root, info) -> List[Review]:
        return get_related_for_single(root, info, "topic", "review")


@strawberry.type
class Query:
    """Base query type. Fields here are the root of any graphql requests"""

    topics: List[Topic] = strawberry.field(resolver=get_topics)
    albums: List[Album] = strawberry.field(resolver=get_albums)
    artists: List[Artist] = strawberry.field(resolver=get_artists)


class Context(BaseContext):
    """Context object to make data-access layers available in th request context"""

    def __init__(self, repo: GraphRepository):
        self.repo = repo


@inject
async def get_topics_context(
    repo: GraphRepository = Depends(Provide[Container.graphql_graph_repo]),
) -> Context:
    """Extra processing done in the request context."""
    return Context(repo=repo)


# register things so FastAPI can serve it
schema = strawberry.Schema(Query)
router = GraphQLRouter(schema, context_getter=get_topics_context)
