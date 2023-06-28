<script lang="ts">
	import { onMount, createEventDispatcher } from 'svelte';
	import Board from './02-Molecules/Board.svelte';
	import { writable } from 'svelte/store';
	import { TARGET_BOARD_3, TARGET_BOARD_4, TARGET_BOARD_5 } from '../const';
	
	export let boardSize: number;
	export let search: string;
	export let heuristic: string;

	let solution_pres: number[][] = [];
	var TARGET_BOARD: number[][]
	if (boardSize === 3) {
		TARGET_BOARD = TARGET_BOARD_3
	} else if (boardSize === 4) {
		TARGET_BOARD = TARGET_BOARD_4
	} else {
		TARGET_BOARD = TARGET_BOARD_5
	};

	async function solvePuzzle(puzzle: number[][]) {
		var puzzleArr: number[] = []
		for (let i = 0; i < puzzle.length; i++) {
			for (let j = 0; j < puzzle[i].length; j++) {
				puzzleArr.push(puzzle[i][j]);
			}
		}
		const response = await fetch("http://localhost:3020/solve",{
			method:  'POST',
			headers: {
			'Content-Type': 'application/json'
			},
			body: JSON.stringify({
			puzzle: puzzleArr,
			search:   search,
			heuristic:    heuristic
			})
		})
		if (response.ok) {
  		return await response.json();
			
		} else {
			throw new Error("help");
		}
		
	};
	let solution: number[][][] = [];
	async function handleClick() {
		// Now set it to the real fetch promise 
		const result = await solvePuzzle(board);
		solution = result.states;
		solution_pres = solution[0]
	};
	const handleFullLeftClick = () => {
		solution_pres = solution[0];
	};
	const handleLeftClick = () => {
		let index = solution.indexOf(solution_pres);
		if (index > 0) {
			solution_pres = solution[index - 1];
		}
	};
	const handlePlayClick = () => {
		const interval = setInterval(function () {
			let index = solution.indexOf(solution_pres);
			if (index == solution.length - 1) {
				clearInterval(interval);
			} else {
				handleRightClick();
			}
		}, 50);
	};
	const handleRightClick = () => {
		let index = solution.indexOf(solution_pres);
		if (index < solution.length - 1) {
			solution_pres = solution[index + 1];
		}
	};
	const handleFullRightClick = () => {
		solution_pres = solution[solution.length - 1]
	};

	const findPos = (boardSize: number, value: number, puzzle: number[][]) => {
		for (let i = 0; i < puzzle.length; i++) {
			for (let j = 0; j < puzzle[i].length; j++) {
				if (puzzle[i][j] === value) {
					return i * boardSize + j;
				}
			}
		}
		return 0;
	};

	const generateRandomBoard = (boardSize: number) => {
		let tmp: number[][];
		tmp = TARGET_BOARD
		for (let i = 0; i < 100 + Math.floor(Math.random() * 100); i++) {
			var idx = findPos(boardSize, 0, tmp)
			var arr = []
			if (idx%boardSize > 0 ){
				arr.push(idx-1)
			}
			if (idx%boardSize < boardSize-1) {
				arr.push(idx+1)
			}
			if (Math.floor(idx/boardSize) > 0 && idx > boardSize) {
				arr.push(idx-boardSize)
			}
			if (Math.floor(idx/boardSize) < boardSize-1) {
				arr.push(idx+boardSize)
			}
			var swapIndex = arr[Math.floor(Math.random() * arr.length)]
			tmp[Math.floor(idx/boardSize)][idx%boardSize] = tmp[Math.floor(swapIndex/boardSize)][swapIndex%boardSize]
			tmp[Math.floor(swapIndex/boardSize)][swapIndex%boardSize] = 0
		}
		return tmp
	};

	let board: number[][] = generateRandomBoard(boardSize)

	let emptyRow: number;
	let emptyCol: number;
	let moveCount = 0;
	let solved = false;
	const dispatch = createEventDispatcher();

	const moveCallback = (row: number, col: number) => {
		solution_pres = []
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

	const isSolved = (board: number[][]) => {
		for (let i = 0; i < boardSize; i++) {
			for (let j = 0; j < boardSize; j++) {
				if (board[i][j] != TARGET_BOARD[i][j]) {
					return false;
				}
			}
		}
		return true;
	};

	let clickable=true;
	let not_clickable=false;
</script>

<div>
	{#if solved}
		<p>Le puzzle est résolu !</p>
	{:else}
		<p>Nombre de coups : {moveCount}</p>
	{/if}
	<Board {board} {moveCallback} {clickable} />
	<button on:click={handleClick} >Solve</button>

	{#if solution_pres.length !== 0}
		{#key solution_pres}
			<Board board={solution_pres} {moveCallback} clickable={not_clickable} />
		{/key}
		<button on:click={handleFullLeftClick} >FullLeft</button>
		<button on:click={handleLeftClick} >Left</button>
		<button on:click={handlePlayClick} >Play</button>
		<button on:click={handleRightClick} >Right</button>
		<button on:click={handleFullRightClick} >FullRight</button>
	{/if}
	</div>
