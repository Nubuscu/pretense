<script>
  import { Container, Row, Col } from "sveltestrap";
  import Navigation from "../lib/Navigation.svelte";
</script>

<html lang="en">
  <Container class="main">
    <Row class="fullheight">
      <Col xs="2">
        <Navigation />
      </Col>
      <Col xs="10">
        <slot />
      </Col>
    </Row>
  </Container>
</html>
