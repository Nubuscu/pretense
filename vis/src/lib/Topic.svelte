<script>
  import Graph from "./Graph.svelte";
  import GraphNode from "./GraphNode.svelte";
  import GraphEdge from "./GraphEdge.svelte";
  import SvelteMarkdown from "svelte-markdown";
  import { singleTopic } from "./topicProcessing.js";

  export let topic_id;

  let root = `${import.meta.env.VITE_BACKEND_HOST}:${
    import.meta.env.VITE_BACKEND_PORT
  }`;
  let nodes = [];
  let edges = [];
  let content = {};
  $: fetch(`${root}/v1/topics/${topic_id}`)
    .then((response) => response.json())
    .then((data) => {
      let res = singleTopic(data);
      nodes = res.nodes;
      edges = res.edges;
      content = res.content;
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
