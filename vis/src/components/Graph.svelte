<script>
  import { onMount, setContext } from "svelte";
  import cytoscape from "cytoscape";
  import cola from "cytoscape-cola";
  import cise from "cytoscape-cise";
  import GraphStyles from "./GraphStyles.js";
  setContext("graphSharedState", {
    getCyInstance: () => cyInstance,
  });
  let refElement = null;
  let cyInstance = null;
  onMount(() => {
    cytoscape.use(cise);
    cyInstance = cytoscape({
      container: refElement,
      style: GraphStyles,
    });
    cyInstance.on("add", () => {
      cyInstance
        .makeLayout({
          animate: true,
          name: "cise",
          clusters: (node) => node.cluster,
        })
        .run();
    });
  });
</script>

<div class="graph" bind:this={refElement}>
  {#if cyInstance}
    <slot />
  {/if}
</div>

<style>
  .graph {
    min-height: 100vh;
    background-color: white;
  }
</style>
