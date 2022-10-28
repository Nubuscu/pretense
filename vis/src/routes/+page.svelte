<script>
  import Graph from "$lib/Graph.svelte";

  export let data;
</script>

{#if data.multi !== null}
  <Graph input={data.multi} />
{/if}
