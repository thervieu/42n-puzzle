<script lang="ts">
	import { onMount, createEventDispatcher } from 'svelte';
	import Button from '../01-Atoms/Button.svelte';
	import Board from '../02-Molecules/Board.svelte';
	import { generateRandomBoard, isSolved, isValidMove } from '../../Helpers/board';
	import RACCOON from '../../Consts/imgs';

	let solution: number[][][] = [];

	let board: number[][] = generateRandomBoard();

	async function makeRandomState() {
		// Now set it to the real fetch promise
		board = generateRandomBoard();
		solution = [];
	}

	let emptyRow: number;
	let emptyCol: number;
	let moveCount = 0;
	let solved = false;
	const dispatch = createEventDispatcher();

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

	onMount(() => {
		// Trouver la position de la case vide
		for (let i = 0; i < 3; i++) {
			for (let j = 0; j < 3; j++) {
				if (board[i][j] === 0) {
					emptyRow = i;
					emptyCol = j;
					break;
				}
			}
		}
	});

	let clickable = true;

	let imagePath: string = RACCOON;

	let imagePaths: string[][] = [];
	let isImageLoaded = false;

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
</script>

<div class="board-section">
	<img src={imagePath} alt="test" width={'240px'} />
	<br />
	<Button action={makeRandomState} text={'Randomize raccoon'} />
	{#key board}
		<Board
			{board}
			{moveCallback}
			{clickable}
			imagePaths={isImageLoaded ? imagePaths.flat() : null}
		/>
	{/key}
</div>

<style lang="scss">
	.board-section {
		text-align: center;
	}
	:global(.ct-series-a .ct-bar) {
		stroke: green;
		stroke-width: 20px;
	}
	:global(.ct-series-b .ct-bar) {
		stroke: rgb(0, 2, 128);
		stroke-width: 20px;
	}
	:global(.ct-series-c .ct-bar) {
		stroke: rgb(196, 216, 21);
		stroke-width: 20px;
	}
	:global(.ct-series-d .ct-bar) {
		stroke: rgb(128, 0, 75);
		stroke-width: 20px;
	}
	:global(.ct-series-e .ct-bar) {
		stroke: rgb(128, 58, 0);
		stroke-width: 20px;
	}
	:global(.ct-series-f .ct-bar) {
		stroke: rgb(89, 10, 92);
		stroke-width: 20px;
	}
	:global(.ct-series-g .ct-bar) {
		stroke: rgb(29, 184, 204);
		stroke-width: 20px;
	}
	:global(.ct-series-h .ct-bar) {
		stroke: rgb(81, 172, 39);
		stroke-width: 20px;
	}
	:global(.ct-series-i .ct-bar) {
		stroke: rgb(165, 3, 3);
		stroke-width: 20px;
	}

	:global(.ct-label) {
		color: white;
	}
</style>
