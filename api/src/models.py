from typing import List, NewType, Optional
from uuid import uuid4

from pydantic import BaseModel

UUID = NewType("UUID", uuid4)


class Tag(BaseModel):
    """A generic string tag to search other objects by."""

    _id: UUID
    key: str


class Review(BaseModel):
    """Long(er) form text about another object."""

    _id: UUID
    title: Optional[str]
    body: str


class Base(BaseModel):
    """Base model from which reviewable/taggable objects are created."""

    _id: UUID
    tags: Optional[List[Tag]]
    reviews: Optional[List[Review]]


class Artist(Base):
    """Represents an artist/musician."""

    name: str


class Album(Base):
    """Represents an album."""

    name: str
    artists: Optional[List[Artist]]
