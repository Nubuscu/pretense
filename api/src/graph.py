from gremlin_python.process.anonymous_traversal import traversal
from gremlin_python.driver.driver_remote_connection import DriverRemoteConnection
from gremlin_python.process.traversal import T, Cardinality, WithOptions
from gremlin_python.process.graph_traversal import GraphTraversalSource, __, unfold
from typing import Dict
from uuid import uuid4


from src.models import Album, Artist, Topic, Review

single = Cardinality.single


class GraphRepository:
    g = None
    _id = T.id

    def __init__(self):
        self.g: GraphTraversalSource = traversal().withRemote(
            DriverRemoteConnection("ws://localhost:8182/gremlin", "g")
        )

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
            unfold(),
            __.V(review).add_e("reviews").to(subject),
        ).next()

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

    def get_album(self, title) -> Album:
        raw = self.g.V().has("album", "name", title).next()
        artists = self.g.V(raw).in_("wrote").to_list()
        return Album(
            id=raw.id,
            name=self.g.V(raw).values("name").next(),
            artists=[
                Artist(id=a.id, name=self.g.V(a).values("name").next())
                for a in artists
                if a
            ],
        )

    def get_topic(self, id_=None):
        if not id_:
            # no id given, get a big list of empty topics
            raw_topics = (
                self.g.V().has_label("topic").valueMap().with_(WithOptions.tokens)
            )
            return [
                Topic(id_=raw.get(T.id), name=raw.get("name", [None])[0])
                for raw in raw_topics
            ]
        else:
            parsed_topic, parsed_reviews, album_names = (
                self.g.V(id_)
                .union(
                    # re-select the topic
                    __.V(id_).valueMap().with_(WithOptions.tokens),
                    # get the review
                    __.in_("reviews").valueMap().with_(WithOptions.tokens).fold(),
                    # get all the included albums in a list
                    __.out("includes").dedup().values("name").fold(),
                )
                .toList()
            )
            return Topic(
                id_=parsed_topic[T.id],
                albums=[self.get_album(name) for name in album_names],
                reviews=[
                    Review(r[T.id], body=r.get("content")) for r in parsed_reviews
                ],
            )


class ContentWriter:
    def __init__(self, repo=None):
        self.repo = repo or GraphRepository()
        super().__init__()

    def write_topic(self, title, review_content, album_names):
        topic = self.repo.upsert_topic(title)
        self.repo.add_review(title, topic.id, review_content)
        albums = [self.repo.get_album(name) for name in album_names]
        for album in albums:
            album_v = self.repo.g.V(album.id_).next()
            self.repo.g.V(topic).add_e("includes").to(album_v).next()
