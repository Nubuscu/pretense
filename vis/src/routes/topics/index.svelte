<script>
  import Topic from "../../components/Topic.svelte";

  let root = `${import.meta.env.VITE_BACKEND_HOST}:${
    import.meta.env.VITE_BACKEND_PORT
  }`;
  let topics = [];
  $: fetch(`${root}/v1/topics`)
    .then((response) => response.json())
    .then((data) => {
      topics = data.map((val) => val.id);
    })
    .catch((err) => console.error(err.code));
</script>

<html lang="en">
  <h1>Pretense</h1>
  {#each topics as topicId}
    <Topic topic_id={topicId} />
  {/each}
</html>
