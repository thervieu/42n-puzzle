<script lang="ts">
	import Cell from '../01-Atoms/Cell.svelte';
	import { isSolved } from '../../Helpers/board';
	export let board: number[][];
	export let moveCallback: (row: number, col: number) => void;
	export let clickable: boolean;
	export let imagePaths: string[] | null;
</script>

<div class="board" style:--size={board.length}>
	{#each board as row, i}
		{#each row as cell, j}
			{#if imagePaths !== null}
				<Cell
					size={board.length}
					{i}
					{j}
					{moveCallback}
					{cell}
					{clickable}
					image={imagePaths}
					isSolved={isSolved(board)}
				/>
			{:else}
				<Cell
					size={board.length}
					{i}
					{j}
					{moveCallback}
					{cell}
					{clickable}
					image={null}
					isSolved={isSolved(board)}
				/>
			{/if}
		{/each}
	{/each}
</div>

<style>
	.board {
		display: grid;
		grid-template-columns: repeat(var(--size), 1fr);
		grid-template-rows: repeat(var(--size), 1fr);
		width: 240px;
		margin: auto;
	}
</style>
