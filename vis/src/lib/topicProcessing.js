/*
 * A set of functions to turn reasonable backend json into unreasonable node/edge graph data.
 */

/**
 * Format a topic as a central node linked to all the albums.
 * @param {object} topicData json blob from the backend detailing a topic
 * @param {number} topicId id of the topic (if the topic should be displayed)
 * @returns nodes, edges, and (review) text content of the topic
 */
export function singleTopic(topicData, topicId = null) {
  let nodes = [];
  let edges = [];

  let topicNodeId;
  if (topicId !== null) {
    topicNodeId = `topic_${topicId}`
    nodes.push({
      id: topicNodeId,
      label: topicData.name,
      isTopic: true,
      cluster: topicNodeId,
    });
  }

  const content = topicData.reviews[0];
  topicData.albums.forEach((album) => {
    let albumNodeId = `album_${album.id}`;

    let albumNode = {
      id: albumNodeId,
      label: `${album.name} - ${album.artists.map(a => a.name).join(', ')}`,
      artists: album.artists,
      cluster: topicNodeId,  // for cytoscape-cise
    };
    if (topicId !== null) {
      edges.push({
        id: `${topicNodeId}_${albumNodeId}`,
        source: albumNodeId,
        target: topicNodeId
      })
    }
    nodes.push(albumNode);
  });
  return {
    nodes: nodes,
    edges: edges,
    content: content,
  };
}

/**
 * process multiple topics into nodes and edges.
 * @param {array[object]} topicsData multiple topics in a list
 * @param {function} method a function to process a single topic
 * @returns nodes and edges
 */
export function multiTopic(topicsData, method = singleTopic) {
  let allNodes = [];
  let allEdges = [];
  topicsData.forEach((topicData) => {
    let {
      nodes: newNodes,
      edges: newEdges,
      content: _content,
    } = method(topicData, topicData.id);
    newNodes.forEach(newNode => {
      // just add topics, no dups or artists to process
      if (newNode.isTopic) {
        allNodes.push(newNode);
        return;
      }
      let existingNode = allNodes.find(existing => existing.id === newNode.id);
      // new node isn't really new, just update edge references
      if (existingNode !== undefined) {
        newEdges.forEach(edge => {
          if (edge.source === newNode.id) {
            edge.source = existingNode.id
          } else if (edge.target === newNode.id) {
            edge.target = existingNode.id
          }
        })
        return
      }
      // look for other albums by the same artist and link them
      let candidateArtistIds = newNode.artists.map(ar => ar.id);
      allNodes.filter(n => !n.isTopic).forEach(ex => {
        ex.artists.forEach(artist => {
          if (candidateArtistIds.includes(artist.id)) {
            newEdges.push({
              id: `${ex.id}_${newNode.id}`,
              source: ex.id,
              target: newNode.id,
            });
          }
        })
      });
      allNodes.push(newNode);
    })
    // add all the remaining edges, if source/dest both exist
    newEdges.forEach(newEdge => {
      let sourceExists = allNodes.map(n => n.id).includes(newEdge.source);
      let targetExists = allNodes.map(n => n.id).includes(newEdge.target);
      let sourceIsNotTarget = newEdge.source !== newEdge.target
      if (sourceExists && targetExists && sourceIsNotTarget) {
        allEdges.push(newEdge)
      }
    });
  })
  return {
    nodes: allNodes,
    edges: allEdges,
  };
}
