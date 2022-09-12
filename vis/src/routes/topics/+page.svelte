<script>
  import Topic from "$lib/Topic.svelte";
  import { topics } from "$lib/stores";
</script>

<html lang="en">
  <h1>Pretense</h1>
  {#each topics as topicId}
    <Topic topic_id={topicId} />
  {/each}
</html>
