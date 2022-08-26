<script>
  import Graph from "./Graph.svelte";
  import SvelteMarkdown from "svelte-markdown";
  import { singleTopic } from "./topicProcessing.js";

  export let topic_id;
  let root = `${import.meta.env.VITE_BACKEND_HOST}:${
    import.meta.env.VITE_BACKEND_PORT
  }`;
  let graphInput = {
    nodes: [],
    edges: []
  };
  let content = {};
  $: fetch(`${root}/v1/topics/${topic_id}`)
    .then((response) => response.json())
    .then((data) => {
      let res = singleTopic(data);
      graphInput = {
        nodes: res.nodes,
        edges: res.edges,
      };
      content = res.content;
    })
    .catch((err) => console.error(err.code));
</script>

<div>
  <div class="content">
    <h2>{content.title}</h2>
    <SvelteMarkdown source={content.body} />
  </div>
  {#if graphInput !== {}}
    <Graph input={graphInput} />
  {/if}
</div>

<style>
  .content {
    padding: 15px;
  }
</style>
