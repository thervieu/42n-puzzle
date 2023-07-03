<script lang="ts">
	export let size: number;
	export let i: number;
	export let j: number;
	export let cell: number;
	export let moveCallback: (row: number, col: number) => void;
	export let clickable: boolean;
	export let image: string[] | undefined | null;
	const rev_board = [4, 0, 1, 2, 5, 8, 7, 6, 3];
</script>

<div
	class="cell {cell === 1 ? 'empty' : ''}"
	on:click={() => {
		clickable ? moveCallback(i, j) : 0;
	}}
	role="presentation"
>
	{#if image && cell !== 1}
		<img src={image[rev_board[cell]]} alt="cell {i}-{j}" height="80px" />
	{:else}
		{cell === 1 ? '' : cell}
	{/if}
</div>

<style>
	.cell {
		display: flex;
		justify-content: center;
		align-items: center;
		width: 80px;
		height: 80px;
		font-size: 20px;
		background-color: #e6e6e6;
		border: 1px solid #ccc;
		box-sizing: border-box;
		cursor: pointer;
		transition: background-color 0.3s ease;
		color: brown;
	}

	.cell:hover {
		background-color: #d9d9d9;
	}

	.cell.empty {
		background-color: lightgray;
		cursor: default;
	}

	img {
		max-width: 100%;
		max-height: 100%;
	}
</style>
