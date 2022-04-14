from typing import Optional
from src.models import Edge, GraphDTO, Node
from fastapi import APIRouter
from src.db.album_repository import Albums

albums = APIRouter()


@albums.get("/")
def get(tag: Optional[int] = None, limit: int = 10, offset: int = 0):
    if tag:
        return Albums().by_tag(tag)
    return Albums().find_all(limit=limit, offset=offset)


@albums.get("/{id_}")
def get_one(id_):
    return Albums().by_id(id_)


@albums.get("/graph/{album_id}")
def get_graph_from_node(album_id: int) -> GraphDTO:
    album = Albums().by_id(album_id)
    dto = GraphDTO()
    if album:
        _nodes, _edges = _to_nodes_and_edges(album)
        dto.nodes += _nodes
        dto.edges += _edges
    return dto


@albums.get("/graph/all")
def get_all_graph(limit: int = 10, offset: int = 0):
    albums = Albums().find_all(limit=limit, offset=offset)
    dto = GraphDTO()
    for album in albums:
        _nodes, _edges = _to_nodes_and_edges(album)
        dto.nodes += _nodes
        dto.edges += _edges

    return dto


def _to_nodes_and_edges(album):
    nodes = []
    edges = []
    album_node_id = f"album-{album.id_}"
    for artist in album.artists:
        artist_node_id = f"artist-{artist.id_}"
        if not any(node.id_ == artist_node_id for node in nodes):
            nodes.append(Node(id=artist_node_id, label=artist.name))
        edges.append(
            Edge(
                id=f"{album_node_id}-{artist_node_id}",
                source=artist_node_id,
                target=album_node_id,
            )
        )
    return nodes, edges
