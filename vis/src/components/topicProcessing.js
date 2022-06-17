/*
 * A set of functions to turn reasonable backend json into unreasonable node/edge graph data.
 */

/**
 *
 * @param {*} topicData json blob from the backend detailing a topic
 * @returns nodes, edges, and (review) text content of the topic
 */
export function singleTopic(topicData, topicId = null) {
  let nodes = [];
  let edges = [];

  if (topicId !== null) {
    nodes.push({ id: `topic_${topicId}`, label: topicData.name });
  }

  const content = topicData.reviews[0];
  topicData.albums.forEach((album) => {
    let albumId = `album_${album.id}`;
    let uniqueAlbumId = `${albumId}_${topicId}`;
    let album_node = {
      id: uniqueAlbumId,
      rawId: albumId,
      label: album.name,
    };
    if (topicId !== null) {
      album_node["parent"] = `topic_${topicId}`;
    }
    nodes.push(album_node);
    album.artists.forEach((artist) => {
      let artistId = `artist_${artist.id}`;
      let uniqueArtistId = `${artistId}_${topicId}`;
      let artistNode = {
        id: uniqueArtistId,
        rawId: artistId,
        label: artist.name,
      };
      if (topicId !== null) {
        artistNode["parent"] = `topic_${topicId}`;
      }
      let edge = {
        id: `${artistId}_${albumId}_${topicId}`,
        source: uniqueArtistId,
        target: uniqueAlbumId,
      };
      edges.push(edge);
      nodes.push(artistNode);
    });
  });
  return {
    nodes: nodes,
    edges: edges,
    content: content,
  };
}

export function multiTopic(topicsData) {
  let allNodes = [];
  let allEdges = [];
  topicsData.forEach((topicData) => {
    let {
      nodes: newNodes,
      edges: newEdges,
      content: _content,
    } = singleTopic(topicData, topicData.id);

    allNodes.forEach((existing) => {
      newNodes.forEach((newNode) => {
        if (
          newNode.rawId !== undefined &&
          existing.rawId !== undefined &&
          newNode.rawId === existing.rawId
        ) {
          newEdges.push({
            id: `${existing.id}_${newNode.id}`,
            source: existing.id,
            target: newNode.id,
          });
        }
      });
    });
    allNodes = allNodes.concat(newNodes);
    allEdges = allEdges.concat(newEdges);
    // iterate through nodes
    // flatmap into buckets of nodes with same album id
    // create edges for those
    // flatmap again for artist ids
    // create edges for those
  });
  // TODO process each topic individually
  //  check for duplicate albums, artists
  //  link between the dups
  return {
    nodes: allNodes,
    edges: allEdges,
  };
}
