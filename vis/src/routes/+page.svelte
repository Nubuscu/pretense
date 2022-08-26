<script>
  import Graph from "$lib/Graph.svelte";
  import GraphNode from "$lib/GraphNode.svelte";
  import GraphEdge from "$lib/GraphEdge.svelte";
  import { multiTopic } from "$lib/topicProcessing";
  import { onMount } from "svelte";
  import { Container, Row, Col } from "sveltestrap";
  import Navigation from "$lib/Navigation.svelte";

  let root = `${import.meta.env.VITE_BACKEND_HOST}:${
    import.meta.env.VITE_BACKEND_PORT
  }`;
  let processed = null;
  let topics = [];

  async function fetchTopic(id) {
    let resp = await fetch(`${root}/v1/topics/${id}`);
    return await resp.json();
  }
  async function fetchTopicIds() {
    return await fetch(`${root}/v1/topics`)
      .then((response) => response.json())
      .then((data) => {
        return data.map((val) => val.id);
      })
      .catch((err) => {
        console.error("it went bad", err.code);
        return [];
      });
  }
  onMount(async () => {
    fetchTopicIds()
      .then(async (ids) => {
        topics = await Promise.all(ids.map(async (id) => await fetchTopic(id)));
        processed = multiTopic(topics);
      })
      .catch((err) => {
        console.error("it went bad", err.code);
      });
  });
</script>

<html lang="en">
  <Container class="main">
    <Row class="fullheight">
      <Col xs="2">
        <Navigation {topics} />
      </Col>
      <Col xs="10">
        {#if processed !== null}
          <Graph>
            {#each processed.nodes as node}
              <GraphNode {node} />
            {/each}
            {#each processed.edges as edge}
              <GraphEdge {edge} />
            {/each}
          </Graph>
        {/if}
      </Col>
    </Row>
  </Container>
</html>
