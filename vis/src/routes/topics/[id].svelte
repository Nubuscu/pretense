<script>
  import Topic from "../../components/Topic.svelte";
  import Navigation from "../../components/Navigation.svelte";
  import { Container, Row, Col, Styles } from "sveltestrap";
  export let id;
  export let topics;
</script>

<html lang="en">
  <Container class="main">
    <Row class="fullheight">
      <Col xs="2">
        <Navigation {topics} />
      </Col>
      <Col xs="10">
        <Topic topic_id={id} />
      </Col>
    </Row>
  </Container>
</html>
