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
## RESOLVERS ##
def get_topics(root, info: Info, id: Optional[int] = None) -> List["Topic"]:
    with info.context.repo as ctx_repo:
        return ctx_repo.get_topic(id)


def get_albums(
    root, info: Info, id: Optional[int] = None, name: Optional[str] = None
) -> List["Album"]:
    with info.context.repo as ctx_repo:
        return ctx_repo.get_album(name=name, id_=id)


def get_albums_for_topic(root: "Topic", info: Info) -> List["Album"]:
    with info.context.repo as ctx_repo:
        return ctx_repo.get_albums_for_topic(root.id_)


def get_artists(
    root, info: Info, id: Optional[int] = None, name: Optional[str] = None
) -> List["Artist"]:
    with info.context.repo as ctx_repo:
        return ctx_repo.get_artist(name=name, id_=id)


def get_artists_for_album(root: "Album", info: Info) -> List["Artist"]:
    with info.context.repo as ctx_repo:
        return ctx_repo.get_artists_for_album(root.id_)


def get_reviews(root, info: Info, id: Optional[int] = None) -> List["Review"]:
    with info.context.repo as ctx_repo:
        return ctx_repo.get_review(root.id_)


def get_reviews_for_topic(root: "Topic", info: Info) -> List["Review"]:
    with info.context.repo as ctx_repo:
        return ctx_repo.get_reviews_for_topic(root.id_)


# copy all the pydantic models so Strawberry knows how to make them
# can overwrite fields here or exclude them if need be.
# TODO: pull the nested re-querying up to this layer
# - get_topic just grabs a topic
# - add a resolver for all albums that might be related to that topic
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


@strawberry.experimental.pydantic.type(model=models.Artist, all_fields=True)
class Artist(Base):
    pass


@strawberry.experimental.pydantic.type(model=models.Album)
class Album(Base):
    name: str
    artists: List[Artist] = strawberry.field(resolver=get_artists_for_album)


@strawberry.experimental.pydantic.type(model=models.Topic)
class Topic(Base):
    name: str
    albums: List[Album] = strawberry.field(resolver=get_albums_for_topic)
    # hiding artists in in topic - everything's at the album level for now
    reviews: List[Review] = strawberry.field(resolver=get_reviews_for_topic)


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
    """Extra processing done in the request context.

    Consider caching expensive calls.
    """
    return Context(repo=repo)


# register things so FastAPI can serve it
schema = strawberry.Schema(Query)
router = GraphQLRouter(schema, context_getter=get_topics_context)
