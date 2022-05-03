<script>
  import Graph from "./Graph.svelte";
  import GraphNode from "./GraphNode.svelte";
  import GraphEdge from "./GraphEdge.svelte";
  import SvelteMarkdown from "svelte-markdown";

  export let topic_id;

  let root = `${import.meta.env.VITE_BACKEND_HOST}:${
    import.meta.env.VITE_BACKEND_PORT
  }`;
  let topic_node_id = `topic_${topic_id}`;
  let nodes = [
    {
      id: topic_node_id,
    },
  ];
  let edges = [];
  let content = {};
  $: fetch(`${root}/v1/topics/${topic_id}`)
    .then((response) => response.json())
    .then((data) => {
      content = data.reviews[0];
      data.albums.forEach((album) => {
        let album_id = `album_${album.id}`;
        nodes = [
          ...nodes,
          {
            id: album_id,
            label: album.name,
            parent: topic_node_id,
          },
        ];
        album.artists.forEach((artist) => {
          let artist_id = `artist_${artist.id}`;
          nodes = [
            ...nodes,
            {
              id: artist_id,
              label: artist.name,
              parent: topic_node_id,
            },
          ];
          edges = [
            ...edges,
            {
              id: `${artist_id}_${album_id}`,
              source: artist_id,
              target: album_id,
            },
          ];
        });
      });
    })
    .catch((err) => console.error(err.code));
</script>

<div>
  <div class="content">
    <h2>{content.title}</h2>
    <SvelteMarkdown source={content.body} />
  </div>
  <Graph>
    {#each nodes as node}
      <GraphNode {node} />
    {/each}

    {#each edges as edge}
      <GraphEdge {edge} />
    {/each}
  </Graph>
</div>

<style>
  .content {
    padding: 15px;
  }
</style>
