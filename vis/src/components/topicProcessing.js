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

    let albumNode = {
      id: uniqueAlbumId,
      rawId: albumId,
      label: `${album.name} - ${album.artists.map(a => a.name).join(', ')}`,
      artists: album.artists
    };
    if (topicId !== null) {
      albumNode["parent"] = `topic_${topicId}`;
    }
    nodes.push(albumNode);
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
        let sameAlbum = (newNode.rawId !== undefined &&
          existing.rawId !== undefined &&
          newNode.rawId === existing.rawId);
        let sameArtist = (newNode.artists !== undefined &&
          existing.artists !== undefined &&
          existing.artists.some(a => newNode.artists.map(n => n.id).includes(a.id)));
        if (sameAlbum || sameArtist) {
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
  });
  return {
    nodes: allNodes,
    edges: allEdges,
  };
}
