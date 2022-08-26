<script>
  import Graph from "$lib/Graph.svelte";
  import { Container, Row, Col } from "sveltestrap";
  import Navigation from "$lib/Navigation.svelte";
  import { topics } from "$lib/stores";

  export let data;

  topics.set(data.topics);
  let processed = data.processed;
</script>

<html lang="en">
  <Container class="main">
    <Row class="fullheight">
      <Col xs="2">
        <Navigation />
      </Col>
      <Col xs="10">
        {#if processed !== null}
          <Graph input={processed} />
        {/if}
      </Col>
    </Row>
  </Container>
</html>
