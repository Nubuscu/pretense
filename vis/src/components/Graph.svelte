<script>
  import { onMount, setContext } from "svelte";
  import cytoscape from "cytoscape";
  // a layout library thing. consider https://github.com/cytoscape/cytoscape.js-cola for topics
  import dagre from "cytoscape-dagre";
  import cola from "cytoscape-cola";
  import cxtmenu from "cytoscape-cxtmenu";
  import GraphStyles from "./GraphStyles.js";
  setContext("graphSharedState", {
    getCyInstance: () => cyInstance,
  });
  let refElement = null;
  let cyInstance = null;
  onMount(() => {
    cytoscape.use(dagre);
    cytoscape.use(cxtmenu);
    cytoscape.use(cola);
    cyInstance = cytoscape({
      container: refElement,
      style: GraphStyles,
    });
    cyInstance.on("add", () => {
      cyInstance
        .makeLayout({
          name: "dagre",
          rankDir: "TB",
          nodeSep: 150,
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
    min-height: 500px;
    left: 12.5%;
    width: 75%;
  }
</style>
