<script>
  import Topic from "$lib/Topic.svelte";
  import Navigation from "$lib/Navigation.svelte";
  import { Container, Row, Col, Styles } from "sveltestrap";
  export let data;
  export let topics;
</script>

<html lang="en">
  <Container class="main">
    <Row class="fullheight">
      <Col xs="2">
        <Navigation {topics} />
      </Col>
      <Col xs="10">
        <Topic topic_id={data.id} />
      </Col>
    </Row>
  </Container>
</html>
