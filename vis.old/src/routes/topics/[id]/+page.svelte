<script>
  import Graph from "$lib/Graph.svelte";
  import SvelteMarkdown from "svelte-markdown";
  export let data;
  let topicTitle = data.title ?? "";
  let reviews = data.reviews ?? "";
</script>

<html lang="en">
  <div>
    <div class="content">
      <h2>{topicTitle}</h2>
      {#each reviews as { title, body }}
        <h3>{title}</h3>
        <SvelteMarkdown source={body} />
      {/each}
    </div>
    <Graph input={data.graphInput} />
  </div>
</html>

<style>
  .content {
    padding: 15px;
  }
</style>