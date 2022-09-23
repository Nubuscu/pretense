<script>
  import Topic from "$lib/Topic.svelte";

  export let data;
</script>

<html lang="en">
  <Topic topic_id={data.id} />
</html>
