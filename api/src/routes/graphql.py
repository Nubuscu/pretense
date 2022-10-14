from typing import List, Optional
from fastapi import Depends
import strawberry
from strawberry.types import Info

from dependency_injector.wiring import inject, Provide
from strawberry.fastapi import GraphQLRouter, BaseContext
from src.graph import GraphRepository

from src import models
from src.di import Container

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


@strawberry.experimental.pydantic.type(model=models.Base, all_fields=True)
class Base:
    pass


@strawberry.experimental.pydantic.type(model=models.Artist, all_fields=True)
class Artist:
    pass


@strawberry.experimental.pydantic.type(model=models.Album, all_fields=True)
class Album:
    pass


@strawberry.experimental.pydantic.type(model=models.Topic, all_fields=True)
class Topic:
    pass


@strawberry.type
class Query:
    @strawberry.field
    def topics(self, info: Info, id: Optional[int] = None) -> List[Topic]:
        with info.context.repo as ctx_repo:
            return ctx_repo.get_topic(id)

    @strawberry.field
    def albums(
        self, info: Info, id: Optional[int] = None, title: Optional[str] = None
    ) -> List[Album]:
        with info.context.repo as ctx_repo:
            return ctx_repo.get_album(title=title, id_=id)


class Context(BaseContext):
    def __init__(self, repo: GraphRepository):
        self.repo = repo


@inject
async def get_topics_context(
    repo: GraphRepository = Depends(Provide[Container.graphql_graph_repo]),
) -> Context:
    """Extra processing done in the request context.

    Consider caching expensive calls.
    """
    with repo as ctx_repo:
        return Context(repo=ctx_repo)


schema = strawberry.Schema(Query)


router = GraphQLRouter(schema, context_getter=get_topics_context)
