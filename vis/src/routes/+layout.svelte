<script>
  import { Container, Row, Col } from "sveltestrap";
  import Navigation from "../lib/Navigation.svelte";
  export let data;
</script>

<html lang="en">
  <Container class="main">
    <Row class="fullheight">
      <Col xs="2">
        <Navigation topics={data.topics} />
      </Col>
      <Col xs="10">
        <slot />
      </Col>
    </Row>
  </Container>
</html>
