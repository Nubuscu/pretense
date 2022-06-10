/*
 * A set of functions to turn reasonable backend json into unreasonable node/edge graph data.
 */

/**
 *
 * @param {*} topicData json blob from the backend detailing a topic
 * @returns nodes, edges, and (review) text content of the topic
 */
export function singleTopic(topicData, topic_id = null) {
  let nodes = [];
  let edges = [];

  if (topic_id !== null) {
    nodes.push({ id: `topic_${topic_id}`, label: topicData.name });
  }

  const content = topicData.reviews[0];
  topicData.albums.forEach((album) => {
    let album_id = `album_${album.id}`;
    let unique_album_id = `${album_id}_${topic_id}`;
    let album_node = {
      id: unique_album_id,
      raw_id: album_id,
      label: album.name,
    };
    if (topic_id !== null) {
      album_node["parent"] = `topic_${topic_id}`;
    }
    nodes.push(album_node);
    album.artists.forEach((artist) => {
      let artist_id = `artist_${artist.id}`;
      let unique_artist_id = `${artist_id}_${topic_id}`;
      let artist_node = {
        id: unique_artist_id,
        raw_id: artist_id,
        label: artist.name,
      };
      if (topic_id !== null) {
        artist_node["parent"] = `topic_${topic_id}`;
      }
      let edge = {
        id: `${artist_id}_${album_id}_${topic_id}`,
        source: unique_artist_id,
        target: unique_album_id,
      };
      edges.push(edge);
      nodes.push(artist_node);
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
          newNode.raw_id !== undefined &&
          existing.raw_id !== undefined &&
          newNode.raw_id === existing.raw_id
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
