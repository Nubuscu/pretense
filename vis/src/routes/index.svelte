<script>
  import Graph from "../components/Graph.svelte";
  import GraphNode from "../components/GraphNode.svelte";
  import GraphEdge from "../components/GraphEdge.svelte";

  async function getGraph() {
    let root = "http://localhost:5000"; // TODO sensible value
    let res = await fetch(`${root}/v1/albums/graph/all?limit=100`);
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
