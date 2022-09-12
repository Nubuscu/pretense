<script>
  import Graph from "./Graph.svelte";
  import SvelteMarkdown from "svelte-markdown";
  import { processed } from "$lib/stores";

  export let topic_id;
  let topic_p = $processed[topic_id];
  let graphInput = {
    nodes: topic_p?.nodes,
    edges: topic_p?.edges,
  };
  let content = topic_p?.content;
</script>

<div>
  <div class="content">
    <h2>{content?.title}</h2>
    <SvelteMarkdown source={content?.body} />
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
