<script>
	import Graph from '$lib/Graph.svelte';
	import { selectedTopicId, openModal } from '$lib/stores.js';
	import { multiTopic } from '$lib/topicProcessing';
	import SvelteMarkdown from 'svelte-markdown';
	import CollapsibleSection from '$lib/CollapsibleSection.svelte';
	import Modal from '$lib/Modal.svelte';

	/* @type { import('./$houdini').PageData } */
	export let data;

	let reviews = [];
	let topicTitle = '';
	let staticText = `
A project in which I try to overanalyze my taste in music. I thought a typical blog felt too linear, so here we are.

This is as much for my own interest as it is anyone who might stumble across it.
It's a way to grapple with what I listen to, how my tastes have evolved over time, what discoveries were made when, etc.
There might be a few recommendations but this is mostly me writing to myself.

Click on one of the "topic" nodes to expand more text below.
	`;

	$: ({ AllTopics } = data);
	$: processed = multiTopic($AllTopics.data.topics.edges);

	$: if ($selectedTopicId !== null) {
		let selectedTopic = $AllTopics.data.topics.edges.find(
			(element) => element.node.id == $selectedTopicId
		).node;
		reviews = selectedTopic.reviewedBy;
		topicTitle = selectedTopic.name;
	}
</script>

<html lang="en">
	<div class="container">
		<div class="container-inner">
			<h1>Pretense</h1>
			<CollapsibleSection headerText="intro blurb">
				<SvelteMarkdown source={staticText} />
			</CollapsibleSection>
			<Graph input={processed} />
			<Modal bind:showModal={$openModal}>
				<h2 slot="header">{topicTitle}</h2>
				<div class="text-content">
					{#each reviews as { name, body }}
						<h3>{name}</h3>
						<SvelteMarkdown source={body} />
					{/each}
				</div>
			</Modal>
		</div>
	</div>
</html>

<style>
	html {
		color: white;
	}
	.container {
		height: 100%;
		width: 100%;
		display: flex;
		align-items: center;
		justify-content: center;
		flex-flow: column;
	}
	.container-inner {
		width: 100%;
		height: 100%;
	}
</style>
