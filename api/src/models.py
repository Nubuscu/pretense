from typing import List, NewType, Optional
from uuid import UUID

from pydantic import BaseModel, Field


class Tag(BaseModel):
    """A generic string tag to search other objects by."""

    id_: int = Field(None, alias="id")
    key: str


class Review(BaseModel):
    """Long(er) form text about another object."""

    id_: int = Field(None, alias="id")
    title: Optional[str]
    body: str


class Base(BaseModel):
    """Base model from which reviewable/taggable objects are created."""

    id_: int = Field(None, alias="id")
    tags: Optional[List[Tag]]
    reviews: Optional[List[Review]]


class Artist(Base):
    """Represents an artist/musician."""

    name: str


class Album(Base):
    """Represents an album."""

    name: str
    artists: Optional[List[Artist]]


class Topic(Base):
    """A broad(er) thing that can be talked to.

    e.g.
    - music encountered while learning bass
    - how the parents' music morphed into my tastes
    - the uni years

    expecting one review per topic, but not strictly.
    """

    name: str
    albums: Optional[List[Album]]
    artists: Optional[List[Artist]]


class Node(BaseModel):
    id_: str = Field(None, alias="id")
    label: str


class Edge(BaseModel):
    id_: str = Field(None, alias="id")
    label: Optional[str]
    source: str  # node id
    target: str  # node id


class GraphDTO(BaseModel):
    nodes: List[Node] = list
    edges: List[Edge] = list
