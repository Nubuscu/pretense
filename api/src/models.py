from uuid import uuid4

from pydantic import BaseModel, Optional


class Tag(BaseModel):
    _id: uuid4
    key: str


class Review(BaseModel):
    _id: uuid4
    title: Optional[str]
    body: str


class Base(BaseModel):
    _id: uuid4
    tags: Optional[set[Tag]]
    reviews: Optional[set[Review]]


class Artist(Base):
    name: str


class Album(Base):
    name: str
    artists: Optional[list[Artist]]
