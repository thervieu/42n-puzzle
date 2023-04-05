<script lang="ts">
	import { onMount, createEventDispatcher } from 'svelte';
	import Board from './02-Molecules/Board.svelte';
	import { writable } from 'svelte/store';
	import { BOARD_SIZE, TARGET_BOARD } from '../const';

	const boardSize = 3;
	let board: number[][] = [
		[1, 2, 3],
		[4, 5, 6],
		[7, 0, 8]
	];

	let emptyRow: number;
	let emptyCol: number;
	let moveCount = 0;
	let solved = false;
	const dispatch = createEventDispatcher();

	const moveCallback = (row: number, col: number) => {
		// Vérifier si le mouvement est valide
		if (isValidMove(row, col)) {
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

	const isValidMove = (row: number, col: number) => {
		// Vérifier si la case cliquée est adjacente à la case vide
		if (Math.abs(row - emptyRow) + Math.abs(col - emptyCol) !== 1) {
			return false;
		}

		// Vérifier si la case cliquée est à une distance de 1 de la case vide
		if (row < 0 || row >= boardSize || col < 0 || col >= boardSize) {
			return false;
		}

		// Mouvement valide
		return true;
	};

	onMount(() => {
		// Trouver la position de la case vide
		for (let i = 0; i < boardSize; i++) {
			for (let j = 0; j < boardSize; j++) {
				if (board[i][j] === 0) {
					emptyRow = i;
					emptyCol = j;
					break;
				}
			}
		}
	});

	const isSolved = (board) => {
		for (let i = 0; i < BOARD_SIZE; i++) {
			for (let j = 0; j < BOARD_SIZE; j++) {
				if (board[i][j] != TARGET_BOARD[i][j]) {
					return false;
				}
			}
		}
		return true;
	};
</script>

<div>
	{#if solved}
		<p>Le puzzle est résolu !</p>
	{:else}
		<p>Nombre de coups : {moveCount}</p>
	{/if}
	<Board {board} {moveCallback} />
</div>
