<script>
	import { onMount, setContext } from 'svelte';
	import cytoscape from 'cytoscape';
	import elk from 'cytoscape-elk';
	import cise from 'cytoscape-cise';
	import { colours, graphStyles } from '$lib/styles.js';
	import { selectedTopicId, openModal } from '$lib/stores.js';
	setContext('graphSharedState', {
		getCyInstance: () => cyInstance
	});
	let refElement = null;
	let cyInstance = null;
	export let input = null;
	$: if (input !== null && cyInstance !== null) {
		// batch process to stop it trying to reprocess the layout for every new node
		cyInstance.batch(() => {
			input.nodes.forEach((node) => {
				cyInstance.add({
					group: 'nodes',
					id: node.id,
					data: { ...node }
				});
			});
			input.edges.forEach((edge) => {
				cyInstance.add({
					group: 'edges',
					id: edge.id,
					data: { ...edge }
				});
			});
			// apply layout after everything's been added
			cyInstance
				.makeLayout({
					name: 'cise',
					clusters: (node) => node.cluster,
					allowNodesInsideCircle: true
				})
				.run();
			cyInstance
				.makeLayout({
					// use elk/disco afterwards to pack disconnected graph parts closer together
					name: 'elk',
					elk: {
						algorithm: 'disco'
					}
				})
				.run();
		});
	}
	onMount(() => {
		cytoscape.use(cise);
		cytoscape.use(elk);
		cyInstance = cytoscape({
			container: refElement,
			style: graphStyles,
			wheelSensitivity: 0.15
		});
		cyInstance.on('add', () => {});
		cyInstance.on('tap', 'node[isTopic]', (event) => {
			let data = event.target.data();
			// id ~= topic_N
			let topicId = data.id.split('_').reverse()[0];
			selectedTopicId.set(topicId);
			openModal.set(true);
		});
		cyInstance.on('mouseover', 'node', (event) => {
			let edges = event.target.connectedEdges();
			edges.forEach((e) => {
				e.style({ 'line-color': colours.edgeSelected });
			});
		});
		cyInstance.on('mouseout', 'node', (event) => {
			let edges = event.target.connectedEdges();
			edges.forEach((e) => {
				e.style({ 'line-color': colours.edgeDefault });
			});
		});
	});
</script>

<div class="graph" bind:this={refElement}>
	{#if cyInstance}
		<slot />
	{/if}
</div>

<style>
	.graph {
		min-height: 85vh;
		display: flex;
		min-width: 80vw;
		background-color: #393939;
	}
</style>
