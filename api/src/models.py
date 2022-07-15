from typing import Any, Collection, List, NewType, Optional
from uuid import UUID

from pydantic import BaseModel, Field

from src.db.db import AlbumMapper, DefaultMapper


class Tag(BaseModel):
    """A generic string tag to search other objects by."""

    id_: UUID = Field(alias="id")
    key: str


class Review(BaseModel):
    """Long(er) form text about another object."""

    id_: UUID = Field(alias="id")
    title: Optional[str]
    body: str


class Base(BaseModel):
    """Base model from which reviewable/taggable objects are created."""

    id_: UUID = Field(alias="id")
    tags: Optional[List[Tag]]
    reviews: Optional[List[Review]]

    class Config:
        orm_mode = True
        getter_dict = DefaultMapper


class Artist(Base):
    """Represents an artist/musician."""

    name: str


class Album(Base):
    """Represents an album."""

    name: str
    artists: Optional[List[Artist]]

    class Config(Base.Config):
        getter_dict = AlbumMapper


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
