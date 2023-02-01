<script>
	import Square from './Square.svelte';
	var tiles = [2, 3, 1, 4, 5, 6, 7, 8, 0];
</script>

<div>
	<ul>
		{#each tiles as _, i}
			<li><Square value={tiles[i]} /></li>
		{/each}
	</ul>
</div>

<style>
	ul {
		list-style-type: none;
	}
</style>
