import asyncio
import os
from typing import List, Optional, Union
from gremlin_python.process.anonymous_traversal import traversal
from gremlin_python.driver.aiohttp.transport import AiohttpTransport
from gremlin_python.driver.driver_remote_connection import DriverRemoteConnection
from gremlin_python.process.traversal import T, Cardinality, WithOptions
from gremlin_python.process.graph_traversal import (
    GraphTraversalSource,
    __,
    unfold,
    value_map,
    in_,
    out,
)
from src.models import Album, Artist, Topic, Review

single = Cardinality.single


def connection_factory(
    host=os.environ.get("DB_HOST"),
    port=os.environ.get("DB_PORT"),
    username=os.environ.get("DB_USER"),
    password=os.environ.get("DB_PASS"),
    call_from_event_loop=None,
) -> DriverRemoteConnection:
    conn_string = f"ws://{host}:{port}/gremlin"
    conn = DriverRemoteConnection(
        conn_string,
        "g",
        username=username,
        password=password,
        transport_factory=lambda: AiohttpTransport(
            call_from_event_loop=call_from_event_loop
        ),
    )
    return conn


class GraphRepository:
    def __init__(self, conn: DriverRemoteConnection | None = None):
        self.provided_connection = bool(conn)
        self.g: GraphTraversalSource = traversal().with_remote(conn)
        self.conn = conn

    def __enter__(self):
        if not self.provided_connection:
            self.conn = connection_factory()
            self.g: GraphTraversalSource = traversal().with_remote(self.conn)
        return self

    def __exit__(self, *_args):
        if not self.provided_connection:
            self.conn.close()

    def upsert_artist(self, name):
        # unique by name
        return (
            self.g.V()
            .has_label("artist")
            .has("name", name)
            .fold()
            .coalesce(unfold(), __.add_v("artist").property("name", name))
            .next()
        )

    def upsert_album(self, name, artist_names):
        # unique by name
        album = (
            self.g.V()
            .has_label("album")
            .has("name", name)
            .fold()
            .coalesce(unfold(), __.add_v("album").property("name", name))
            .next()
        )

        for artist in artist_names:
            artist = self.upsert_artist(artist)
            self.g.V(artist).add_e("wrote").to(album).iterate()
        return self.g.V(album).next()

    def add_review(self, name, subject_v_id, content):
        # upsert review vertex by name
        review = (
            self.g.V()
            .has_label("review")
            .has("name", name)
            .fold()
            .coalesce(unfold(), __.add_v("review").property("name", name))
            .next()
        )
        # always update review content
        self.g.V(review).property(single, "content", content).next()

        # upsert edge to subject
        subject = self.g.V(subject_v_id).next()
        self.g.V(review).out("reviews").fold().coalesce(
            unfold(), __.V(review).add_e("reviews").to(subject)
        ).iterate()

    def upsert_topic(self, name):
        # unique by name
        return (
            self.g.V()
            .has_label("topic")
            .has("name", name)
            .fold()
            .coalesce(unfold(), __.addV("topic").property("name", name))
            .next()
        )

    def get_album(self, name=None, id_=None) -> list[Album]:

        tvsl = self.g.V().has_label("album")
        if name:
            tvsl = tvsl.has("name", name)
        elif id_:
            tvsl = tvsl.has_id(id_)

        raw_list = (
            tvsl.project("album").by(value_map().with_(WithOptions.tokens)).toList()
        )
        retval = []
        for raw in raw_list:
            raw_album = raw["album"]
            retval.append(
                Album(
                    id=raw_album[T.id],
                    name=raw_album["name"][0],
                )
            )
        return retval

    def get_albums_for_topic(self, topic_id) -> list[Album]:
        raw_list = (
            self.g.V(topic_id)
            .out("includes")
            .dedup()
            .project("album")
            .by(value_map().with_(WithOptions.tokens))
            .toList()
        )
        retval = []
        for raw in raw_list:
            raw_album = raw["album"]
            retval.append(
                Album(
                    id=raw_album[T.id],
                    name=raw_album["name"][0],
                )
            )
        return retval

    def get_artist(self, name=None, id_=None) -> list[Artist]:
        tvsl = self.g.V().has_label("artist")
        if name:
            tvsl = tvsl.has("name", name)
        elif id_:
            tvsl = tvsl.has_id(id_)

        raw_list = (
            tvsl.project("artist").by(value_map().with_(WithOptions.tokens)).toList()
        )
        retval = []
        for raw in raw_list:
            raw_artist = raw["artist"]
            retval.append(
                Artist(
                    id=raw_artist[T.id],
                    name=raw_artist["name"][0],
                )
            )
        return retval

    def get_artists_for_album(self, album_id) -> list[Artist]:
        raw_list = (
            self.g.V(album_id)
            .in_("wrote")
            .dedup()
            .project("artist")
            .by(value_map().with_(WithOptions.tokens))
            .toList()
        )
        retval = []
        for raw in raw_list:
            raw_artist = raw["artist"]
            retval.append(
                Artist(
                    id=raw_artist[T.id],
                    name=raw_artist["name"][0],
                )
            )
        return retval

    def get_topic(self, id_: int = None) -> list[Topic]:
        tvsl = self.g.V().has_label("topic")
        if id_:
            tvsl = tvsl.has_id(id_)

        raw_list = (
            tvsl.project("topic").by(value_map().with_(WithOptions.tokens)).toList()
        )
        retval = []
        for raw in raw_list:
            raw_topic = raw["topic"]
            names = raw_topic["name"]
            retval.append(
                Topic(
                    id=raw_topic[T.id],
                    name=names[0] if names else None,
                )
            )
        return retval

    def get_review(self, id_: int = None) -> list[Review]:
        tvsl = self.g.V().has_label("review")
        if id_:
            tvsl = tvsl.has_id(id_)

        raw_list = (
            tvsl.project("review").by(value_map().with_(WithOptions.tokens)).toList()
        )
        retval = []
        for raw in raw_list:
            raw_review = raw["review"]
            titles = raw_review.get("title")
            retval.append(
                Review(
                    id=raw_review[T.id],
                    name=titles[0] if titles else None,
                    body=raw_review["content"][0],
                )
            )
        return retval

    def get_reviews_for_topic(self, topic_id: int) -> list[Review]:
        raw_list = (
            self.g.V(topic_id)
            .in_("reviews")
            .dedup()
            .project("review")
            .by(value_map().with_(WithOptions.tokens))
            .toList()
        )
        retval = []
        for raw in raw_list:
            raw_review = raw["review"]
            titles = raw_review.get("title")
            retval.append(
                Review(
                    id=raw_review[T.id],
                    name=titles[0] if titles else None,
                    body=raw_review["content"][0],
                )
            )
        return retval

    def get_unrelated_albums(self) -> list[Album]:
        """Find all the albums that aren't associated with a Topic (yet)"""
        ids = self.g.V().has_label("album").not_(in_("reviews")).id_().toList()
        return [self.get_album(id_=id_) for id_ in ids]


class ContentWriter:
    def __init__(self, repo=None):
        self.repo = repo or GraphRepository()

    def write_topic(self, title, review_content, album_names):
        with self.repo as repo:
            topic = repo.upsert_topic(title)
            repo.add_review(title, topic.id, review_content)
            albums = [repo.get_album(name=name)[0] for name in album_names]
            for album in albums:
                album_v = repo.g.V(album.id_).next()
                repo.g.V(topic).add_e("includes").to(album_v).iterate()
