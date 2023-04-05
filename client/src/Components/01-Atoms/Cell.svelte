<script lang="ts">
	export let i: number;
	export let j: number;
	export let cell: number;
	export let moveCallback: (row: number, col: number) => void;
</script>

<div
	class="cell {cell === 0 ? 'empty' : ''}"
	on:click={() => moveCallback(i, j)}
	role="presentation"
>
	{cell === 0 ? '' : cell}
</div>

<style>
	.cell {
		display: flex;
		justify-content: center;
		align-items: center;
		width: 50px;
		height: 50px;
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
</style>
