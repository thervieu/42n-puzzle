<script lang="ts">
	import { createEventDispatcher, onMount } from 'svelte';
	import Subtitle from '../Components/01-Atoms/Subtitle.svelte';
	import Title from '../Components/01-Atoms/Title.svelte';
	import InputsRow from '../Components/02-Molecules/InputsRow.svelte';
	import Solution from '../Components/03-Organisms/Solution.svelte';
	import Puzzle from '../Components/03-Organisms/Puzzle.svelte';
	import Chart from '../Components/03-Organisms/Chart.svelte';

	import { generateRandomBoard, isSolved, isValidMove, solvePuzzle } from '../Helpers/board';
	import RACCOON from '../Consts/imgs';

	let search = 'uniform';
	let heuristic = 'hamming';
	let size = 3;

	let imagePaths: string[][] = [];
	let imagePath: string = RACCOON;
	let board: number[][] = generateRandomBoard();
	let isImageLoaded = false;
	let solution: number[][][] = [];
	let emptyRow: number;
	let emptyCol: number;
	let moveCount = 0;
	let solved = false;
	const dispatch = createEventDispatcher();

	function onChangeSearch(event: { currentTarget: { value: string } }) {
		search = event.currentTarget.value;
	}
	function onChangeHeuristic(event: { currentTarget: { value: string } }) {
		heuristic = event.currentTarget.value;
	}

	onMount(() => {
		const canvas = document.createElement('canvas');
		let ctx = canvas.getContext('2d');

		const image = new Image();

		image.src = imagePath;
		image.onload = () => {
			const cellWidth = image.width / board.length;
			const cellHeight = image.height / board.length;

			for (let i = 0; i < board.length; i++) {
				imagePaths[i] = [];
				for (let j = 0; j < board[i].length; j++) {
					ctx?.clearRect(0, 0, canvas.width, canvas.height);
					ctx?.drawImage(
						image,
						j * cellWidth,
						i * cellHeight,
						cellWidth,
						cellHeight,
						0,
						0,
						canvas.width,
						canvas.height
					);
					imagePaths[i][j] = canvas.toDataURL();
				}
			}
			isImageLoaded = true;
		};
	});

	const moveCallback = (row: number, col: number) => {
		solution = [];
		// Vérifier si le mouvement est valide
		if (isValidMove(row, col, emptyRow, emptyCol)) {
			// Déplacer la case
			const temp = board[row][col];
			board[row][col] = board[emptyRow][emptyCol];
			board[emptyRow][emptyCol] = temp;
			emptyRow = row;
			emptyCol = col;
			moveCount++;
			dispatch('move', moveCount);
			solved = isSolved(board);
		}
	};
</script>

<main class="globalLayout">
	<Title value="N-Puzzle" />
	<div class="section">
		<Subtitle value="Search" />
		<InputsRow
			name={'search'}
			inputs={[
				{ label: 'uniform cost', value: 'uniform' },
				{ label: 'greedy', value: 'greedy' },
				{ label: 'both', value: 'both' }
			]}
			checkedValue={search}
			handleChange={onChangeSearch}
		/>
	</div>
	<div class="section">
		<Subtitle value="Heuristic" />
		<InputsRow
			name={'Heuristic'}
			inputs={[
				{ label: 'hamming', value: 0 },
				{ label: 'manhattan', value: 1 },
				{ label: 'euclidean', value: 2 }
			]}
			checkedValue={heuristic}
			handleChange={onChangeHeuristic}
		/>
	</div>
	<!-- PUZZLE -->
	{#key size}
		<Puzzle />
	{/key}
	<!-- SOLVER -->
	<Solution {solution} {moveCallback} {isImageLoaded} {search} {heuristic} {imagePaths} />
	<!-- CHART -->
	<Chart />
</main>

<style>
	.globalLayout {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;

		background-image: url('https://cdn.create.vista.com/api/media/medium/170344426/stock-photo-wooden-background?token=');
		flex-grow: 1;
		min-height: 100vh;
		padding: 20px;
	}
	.section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		width: 100%;
	}
</style>
