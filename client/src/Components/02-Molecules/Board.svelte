<script lang="ts">
	import Cell from '../01-Atoms/Cell.svelte';
	export let board: number[][];
	export let moveCallback: (row: number, col: number) => void;
</script>

<div class="board">
	{#each board as row, i}
		{#each row as cell, j}
			<Cell {i} {j} {moveCallback} {cell} />
			<!-- <div
				class="cell {cell === 0 ? 'empty' : ''}"
				on:click={() => moveCallback(i, j)}
				role="presentation"
			>
				{cell === 0 ? '' : cell}
			</div> -->
		{/each}
	{/each}
</div>

<style>
	.board {
		display: grid;
		grid-template-columns: repeat(3, 1fr);
		grid-template-rows: repeat(3, 1fr);
	}
</style>
