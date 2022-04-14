<script>
  import Graph from "../components/Graph.svelte";
  import GraphNode from "../components/GraphNode.svelte";
  import GraphEdge from "../components/GraphEdge.svelte";

  let root = `${import.meta.env.VITE_BACKEND_HOST}:${
    import.meta.env.VITE_BACKEND_PORT
  }`;
  async function getGraph() {
    let res = await fetch(`${root}/v1/topics/4`);
    if (res.ok) {
      return await res.json();
    } else {
      throw new Error(await res.text());
    }
  }
</script>

<h1>Pretense</h1>
{#await getGraph()}
  waiting...
{:then graphRes}
  <Graph>
    {#each graphRes.nodes as node}
      <GraphNode {node} />
    {/each}

    {#each graphRes.edges as edge}
      <GraphEdge {edge} />
    {/each}
  </Graph>
{:catch err}
  {err.message}
{/await}
