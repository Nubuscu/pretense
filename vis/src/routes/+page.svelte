<script>
	import Graph from '$lib/Graph.svelte';
	import { selectedTopicId } from '$lib/stores.js';
	import { multiTopic } from '$lib/topicProcessing';
	import SvelteMarkdown from 'svelte-markdown';

	/* @type { import('./$houdini').PageData } */
	export let data;

	let reviews = [];
	let topicTitle = '';
	$: ({ AllTopics } = data);
	$: processed = multiTopic($AllTopics.data.topics.edges);

	$: if ($selectedTopicId !== null) {
		let selectedTopic = $AllTopics.data.topics.edges.find(
			(element) => element.node.id == $selectedTopicId
		).node;
		reviews = selectedTopic.reviewedBy;
		topicTitle = selectedTopic.name
		console.log(reviews);
	}
</script>

<html lang="en">
	<h1>Pretense</h1>

	<div class="container">
		<Graph input={processed} />
		<div class="text-content">
		<h2>{topicTitle}</h2>
		{#each reviews as {name, body}}
		<h3>{name}</h3>
		<SvelteMarkdown source={body} />
		{/each}
		</div>
	</div>
</html>

<style>
	html {
		color: white;
	}
	.container {
		height: 100%;
		display: flex;
		align-items: center;
		justify-content: center;
		flex-flow: column;
	}
	.text-content {
		max-width: 80%;
	}
</style>
