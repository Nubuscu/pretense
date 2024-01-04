<script>
	import Graph from '$lib/Graph.svelte';
	import { multiTopic } from '$lib/topicProcessing';

	/* @type { import('./$houdini').PageData } */
    export let data

	$: ({ AllTopics } = data)

	$: processed = multiTopic($AllTopics.data.topics.edges)

</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>

<Graph input={processed} />
