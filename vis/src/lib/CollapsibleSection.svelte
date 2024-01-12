<script>
	// stolen from https://svelte.dev/repl/a5f4d395b15a44d48a6b2239ef705fc4?version=3.35.0
	import GraphStyles from '$lib/styles.js'; 
    export let headerText;

	let expanded = false;
</script>

<div class="collapsible">
	<h3>
		<button aria-expanded={expanded} on:click={() => (expanded = !expanded)}
			>{headerText}
		</button>
	</h3>

	<div class="contents" hidden={!expanded}>
		<slot />
	</div>
</div>

<style>
	.collapsible {
		border-bottom: 1px solid #797979
	}

	h3 {
		margin: 0;
	}

	button {
		background-color: #2c2c2c;
		color:  #ffffff;
		display: flex;
		justify-content: space-between;
		width: 100%;
		border: none;
		margin: 0;
		padding: 1em 0.5em;
	}

	button[aria-expanded='true'] {
		border-bottom: 1px solid #797979;
	}
</style>
