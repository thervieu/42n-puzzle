<script lang="ts">
	import Cell from '../01-Atoms/Cell.svelte';
	export let board: number[][];
	export let moveCallback: (row: number, col: number) => void;
	export let clickable: boolean;
</script>

<div class="board" style:--size={board.length}>
	{#each board as row, i}
		{#each row as cell, j}
			<Cell {i} {j} {moveCallback} {cell} {clickable} />
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
		grid-template-columns: repeat(var(--size), 1fr);
		grid-template-rows: repeat(var(--size), 1fr);
	}
</style>
