from typing import List, Optional

from pydantic import BaseModel, Field


class Tag(BaseModel):
    """A generic string tag to search other objects by."""

    id_: int = Field(alias="id")
    key: str


class Review(BaseModel):
    """Long(er) form text about another object."""

    id_: int = Field(alias="id")
    title: str | None
    body: str


class Base(BaseModel):
    """Base model from which reviewable/taggable objects are created."""

    id_: int = Field(alias="id")
    tags: list[Tag] | None
    reviews: list[Review] | None


class Artist(Base):
    """Represents an artist/musician."""

    name: str


class Album(Base):
    """Represents an album."""

    name: str
    artists: list[Artist] | None


class Topic(Base):
    """A broad(er) thing that can be talked to.

    e.g.
    - music encountered while learning bass
    - how the parents' music morphed into my tastes
    - the uni years

    expecting one review per topic, but not strictly.
    """

    name: str
    albums: list[Album] | None
    artists: list[Artist] | None
