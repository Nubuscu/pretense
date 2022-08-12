<script>
  import { onMount, setContext } from "svelte";
  import cytoscape from "cytoscape";
  import elk from "cytoscape-elk";
  import GraphStyles from "./GraphStyles.js";
  setContext("graphSharedState", {
    getCyInstance: () => cyInstance,
  });
  let refElement = null;
  let cyInstance = null;
  onMount(() => {
    cytoscape.use(elk);
    cyInstance = cytoscape({
      container: refElement,
      style: GraphStyles,
    });
    cyInstance.on("add", () => {
      cyInstance
        .makeLayout({
          name: "elk",
          animate: true,
          nodeDimensionsIncludeLabels: true,
          elk: {
            algorithm: "disco",
            // any of the options here: https://www.eclipse.org/elk/reference.html
            // are available, just drop the `org.eclipse` prefix.
            componentLayoutAlgorithm: "stress",
          },
        })
        .run();
    });
    cyInstance.on("tap", "node[isTopic]", (event) => {
      let data = event.target.data();
      // id ~= topic_N
      let topic_id = data.id.split("_").reverse()[0];
      window.location.href = `/topics/${topic_id}`;
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
