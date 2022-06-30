<script>
  import { onMount, setContext } from "svelte";
  import cytoscape from "cytoscape";
  // import dagre from "cytoscape-dagre";
  // import cxtmenu from "cytoscape-cxtmenu";
  import cola from "cytoscape-cola";
  import GraphStyles from "./GraphStyles.js";
  setContext("graphSharedState", {
    getCyInstance: () => cyInstance,
  });
  let refElement = null;
  let cyInstance = null;
  onMount(() => {
    // cytoscape.use(dagre);
    cytoscape.use(cola);
    cyInstance = cytoscape({
      container: refElement,
      style: GraphStyles,
    });
    cyInstance.on("add", () => {
      cyInstance
        .makeLayout({
          animate: true,
          name: "cola",
          avoidOverlap: true,
          handleDisconnected: true,
          nodeDimensionsIncludeLabels: true,
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
    left: 12.5%;
    width: 75%;
    background-color: white;
  }
</style>
