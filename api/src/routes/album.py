from typing import Optional
from models import Edge, GraphDTO, Node
from fastapi import APIRouter
from db.album_repository import Albums

albums = APIRouter()


@albums.get("/")
def get(tag: Optional[int] = None, limit: int = 10, offset: int = 0):
    if tag:
        return Albums().by_tag(tag)
    return Albums().find_all(limit=limit, offset=offset)


@albums.get("/{id_}")
def get_one(id_):
    return Albums().by_id(id_)


# @albums.get("/graph/{album_id}")
# def get_graph_from_node(album_id: int) -> GraphDTO:
#     album = Albums().by_id(album_id)
#     if not album:
#         return GraphDTO()
#     nodes = []
#     edges = []
#     album_node_id = f"album-{album.id_}"
#     nodes.append(Node(id=album_node_id, label=album.name))
#     for artist in album.artists:
#         artist_node_id = f"artist-{artist.id_}"
#         nodes.append(Node(id=artist_node_id, label=artist.name))
#         edges.append(
#             Edge(
#                 id=f"{album_node_id}-{artist_node_id}",
#                 source=artist_node_id,
#                 target=album_node_id,
#             )
#         )
#     return GraphDTO(nodes=nodes, edges=edges)

@albums.get("/graph/all")
def get_all_graph(limit: int = 10, offset: int = 0):
    albums = Albums().find_all(limit=limit, offset=offset)
    nodes = []
    edges = []
    for album in albums:
        album_node_id = f"album-{album.id_}"
        nodes.append(Node(id=album_node_id, label=album.name))
        for artist in album.artists:
            artist_node_id = f"artist-{artist.id_}"
            # TODO make Node hashable, use a set.
            if not any(node.id_ == artist_node_id for node in nodes):
                nodes.append(Node(id=artist_node_id, label=artist.name))
            edges.append(
                Edge(
                    id=f"{album_node_id}-{artist_node_id}",
                    source=artist_node_id,
                    target=album_node_id,
                )
            )
    return GraphDTO(nodes=nodes, edges=edges)
