import os
from typing import List, Optional, Union
from gremlin_python.process.anonymous_traversal import traversal
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
) -> DriverRemoteConnection:
    conn_string = f"ws://{host}:{port}/gremlin"
    conn = DriverRemoteConnection(
        conn_string, "g", username=username, password=password
    )
    return conn


class GraphRepository:
    def __init__(self, conn: Optional[DriverRemoteConnection] = None):
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

    def get_album(self, title=None, id_=None) -> List[Album]:

        tvsl = self.g.V().has_label("album")
        if title:
            tvsl = tvsl.has("name", title)
        elif id_:
            tvsl = tvsl.has_id(id_)

        raw_list = (
            tvsl.project("album", "artists")
            .by(value_map().with_(WithOptions.tokens))
            .by(__.in_("wrote").value_map().with_(WithOptions.tokens).fold())
            .toList()
        )
        retval = []
        for raw in raw_list:
            raw_album = raw["album"]
            raw_artists = raw["artists"]
            retval.append(
                Album(
                    id=raw_album[T.id],
                    name=raw_album["name"][0],
                    artists=[
                        Artist(id=a[T.id], name=a["name"][0]) for a in raw_artists
                    ],
                )
            )
        assert retval, f"Album(s) not found: {title or id_}"
        if len(retval) == 1:
            return retval[0]
        return retval

    def get_topic(self, id_: int = None) -> Topic:
        """Get a topic from the backend.

        If id_ is not specified, return hollow Topics for all ids

        Args:
            id_ (int, optional): id to search by. Defaults to None.

        Returns:
            Topic: requested topic
        """
        if not id_:
            # no id given, get a big list of id-only, empty topics
            raw_topics = (
                self.g.V().has_label("topic").valueMap().with_(WithOptions.tokens)
            )
            return [
                Topic(id=raw.get(T.id), name=raw.get("name", [None])[0])
                for raw in raw_topics
            ]
        res = (
            self.g.V(id_)
            .project("topic", "reviews", "album_names")
            .by(value_map().with_(WithOptions.tokens))
            .by(__.in_("reviews").valueMap().with_(WithOptions.tokens).fold())
            .by(__.out("includes").dedup().values("name").fold())
            .next()
        )
        topic_id = res["topic"][T.id]
        return Topic(
            id=topic_id,
            name=res["topic"]["name"][0],
            reviews=[
                Review(id=r[T.id], body=r.get("content", [""])[0])
                for r in res["reviews"]
            ],
            albums=[self.get_album(title=name) for name in res["album_names"]],
        )

    def get_unrelated_albums(self) -> List[Album]:
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
            albums = [repo.get_album(title=name) for name in album_names]
            for album in albums:
                album_v = repo.g.V(album.id_).next()
                repo.g.V(topic).add_e("includes").to(album_v).iterate()
