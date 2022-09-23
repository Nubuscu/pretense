<script>
  import Graph from "$lib/Graph.svelte";
  import { topics, processed } from "$lib/stores";

  export let data;

  $topics = data.topics;
  $processed = data.singles;
</script>

{#if data.multi !== null}
  <Graph input={data.multi} />
{/if}
