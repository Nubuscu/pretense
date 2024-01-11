<script>
  import { onMount, setContext } from "svelte";
  import cytoscape from "cytoscape";
  import elk from "cytoscape-elk";
  import GraphStyles from "$lib/GraphStyles.js";
  import { selectedTopicId } from "$lib/stores.js"
  setContext("graphSharedState", {
    getCyInstance: () => cyInstance,
  });
  let refElement = null;
  let cyInstance = null;
  export let input = null;
  $: if (input !== null && cyInstance !== null) {
    // batch process to stop it trying to process too much
    cyInstance.batch(() => {
      input.nodes.forEach((node) => {
        cyInstance.add({
          group: "nodes",
          id: node.id,
          data: { ...node },
        });
      });
      input.edges.forEach((edge) => {
        cyInstance.add({
          group: "edges",
          id: edge.id,
          data: { ...edge },
        });
      });
      // apply layout after everything's been added
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
  }
  onMount(() => {
    cytoscape.use(elk);
    cyInstance = cytoscape({
      container: refElement,
      style: GraphStyles,
    });
    cyInstance.on("add", () => {});
    cyInstance.on("tap", "node[isTopic]", (event) => {
      let data = event.target.data();
      // id ~= topic_N
      let topicId = data.id.split("_").reverse()[0];
      selectedTopicId.set(topicId)
      
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
    min-height: 75vh;
    display: flex;
    min-width: 80vw;
    background-color: black;
  }
</style>