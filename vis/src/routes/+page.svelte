<script>
  import Graph from "$lib/Graph.svelte";
  import { Container, Row, Col } from "sveltestrap";
  import Navigation from "$lib/Navigation.svelte";
  import { topics, processed } from "$lib/stores";

  export let data;

  topics.set(data.topics);
  processed.set(data.singles);
</script>

<html lang="en">
  <Container class="main">
    <Row class="fullheight">
      <Col xs="2">
        <Navigation />
      </Col>
      <Col xs="10">
        {#if data.multi !== null}
          <Graph input={data.multi} />
        {/if}
      </Col>
    </Row>
  </Container>
</html>
