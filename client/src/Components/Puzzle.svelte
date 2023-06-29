<script lang="ts">
	import { onMount, createEventDispatcher } from 'svelte';
	import Board from './02-Molecules/Board.svelte';
	import { writable } from 'svelte/store';
	import { TARGET_BOARD_3, TARGET_BOARD_4, TARGET_BOARD_5 } from '../const';
	import { averageMoves, averageSC, averageTC, averageTime } from '../Helpers/stats';
	import { Bar } from 'svelte-chartjs';
	import { browser } from '$app/environment';


	import {
		Chart,
		Title,
		Tooltip,
		Legend,
		BarElement,
		CategoryScale,
		LinearScale,
	} from 'chart.js';

	Chart.register(
		Title,
		Tooltip,
		Legend,
		BarElement,
		CategoryScale,
		LinearScale
	);
	let dataChart = {
		labels: [""],
		datasets: [
			{
			label: '',
			data: [""],
			},
		],
	};
	const labels: string[][] = [
		["uniform and hamming", "uniform and manhattan", "uniform and euclidean",
		"greedy and hamming", "greedy and manhattan", "greedy and euclidean",
		"both and hamming", "both and manhattan", "both and euclidean"],
		["uniform", "greedy", "both"],
		["hamming", "manhattan", "euclidian"]
	];
	const titles = ["Graph of avg time for 30 puzzles", "Graph of avg moves for 30 puzzles",
		"Graph of avg time complexity for 30 puzzles", "Graph of avg space complexity for 30 puzzles"];

	export let boardSize: number;
	export let search: string;
	export let heuristic: string;

	let solution_pres: number[][] = [];
	var TARGET_BOARD: number[][];
	if (boardSize === 3) {
		TARGET_BOARD = TARGET_BOARD_3
	} else if (boardSize === 4) {
		TARGET_BOARD = TARGET_BOARD_4
	} else {
		TARGET_BOARD = TARGET_BOARD_5
	};

	async function solvePuzzle(puzzle: number[][], search_: string, heuristic_: string) {
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
			search: search_,
			heuristic: heuristic_
			})
		})
		if (response.ok) {
  		return await response.json();
			
		} else {
			throw new Error("help");
		}
	};
	let solution: number[][][] = [];
	let time: string;
	let moves: number;
	let time_complexity: number;
	let space_complexity: number;
	async function handleClick() {
		// Now set it to the real fetch promise 
		const result = await solvePuzzle(board, search, heuristic);
		solution = result.states;
		solution_pres = solution[0];
		time = result.time.substring(12, result.time.length);
		moves = result.moves;
		time_complexity = result.time_complexity;
		space_complexity = result.space_complexity;
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
		solution = []
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

	let data: any[] = [];
	async function DoBenchmark() {
		let u_h, u_m, u_e;
		let g_h, g_m, g_e;
		let b_h, b_m, b_e;
		u_h = [];
		u_m = [];
		u_e = [];
		g_h = [];
		g_m = [];
		g_e = [];
		b_h = [];
		b_m = [];
		b_e = [];
		let searches = ["uniform", "greedy", "both"];
		let heuristics = ["hamming", "manhattan", "euclidean"];
		for (let i in searches) {
			for (let j in heuristics) {
				for (let k = 0; k < 5; k++) {
					let p1 = generateRandomBoard(3);
					let p2 = generateRandomBoard(3);
					let p3 = generateRandomBoard(3);
					let p4 = generateRandomBoard(3);
					let p5 = generateRandomBoard(3);
					let p6 = generateRandomBoard(3);
					const [r1, r2, r3, r4, r5, r6] = await Promise.all([solvePuzzle(p1, searches[i], heuristics[j]),
						solvePuzzle(p2, searches[i], heuristics[j]),
						solvePuzzle(p3, searches[i], heuristics[j]),
						solvePuzzle(p4, searches[i], heuristics[j]),
						solvePuzzle(p5, searches[i], heuristics[j]),
						solvePuzzle(p6, searches[i], heuristics[j])]);
					r1.time = r1.time.substring(12, r1.time.length);
					r2.time = r2.time.substring(12, r2.time.length);
					r3.time = r3.time.substring(12, r3.time.length);
					r4.time = r4.time.substring(12, r4.time.length);
					r5.time = r5.time.substring(12, r5.time.length);
					r6.time = r6.time.substring(12, r6.time.length);
					if (searches[i] === "uniform" && heuristics[j] === "hamming") {
						u_h.push(r1); u_h.push(r2); u_h.push(r3); u_h.push(r4); u_h.push(r5); u_h.push(r6);
					}
					if (searches[i] === "uniform" && heuristics[j] === "manhattan") {
						u_m.push(r1); u_m.push(r2); u_m.push(r3); u_m.push(r4); u_m.push(r5); u_m.push(r6);
					}
					if (searches[i] === "uniform" && heuristics[j] === "euclidean") {
						u_e.push(r1); u_e.push(r2); u_e.push(r3); u_e.push(r4); u_e.push(r5); u_e.push(r6);
					}
					if (searches[i] === "greedy" && heuristics[j] === "hamming") {
						g_h.push(r1); g_h.push(r2); g_h.push(r3); g_h.push(r4); g_h.push(r5); g_h.push(r6);
					}
					if (searches[i] === "greedy" && heuristics[j] === "manhattan") {
						g_m.push(r1); g_m.push(r2); g_m.push(r3); g_m.push(r4); g_m.push(r5); g_m.push(r6);
					}
					if (searches[i] === "greedy" && heuristics[j] === "euclidean") {
						g_e.push(r1); g_e.push(r2); g_e.push(r3); g_e.push(r4); g_e.push(r5); g_e.push(r6);
					}
					if (searches[i] === "both" && heuristics[j] === "hamming") {
						b_h.push(r1); b_h.push(r2); b_h.push(r3); b_h.push(r4); b_h.push(r5); b_h.push(r6);
					}
					if (searches[i] === "both" && heuristics[j] === "manhattan") {
						b_m.push(r1); b_m.push(r2); b_m.push(r3); b_m.push(r4); b_m.push(r5); b_m.push(r6);
					}
					if (searches[i] === "both" && heuristics[j] === "euclidean") {
						b_e.push(r1); b_e.push(r2); b_e.push(r3); b_e.push(r4); b_e.push(r5); b_e.push(r6);
					}
				}
				console.log(i, j)
			}
		}
		let u = u_h.concat(u_m).concat(u_e);
		let g = g_h.concat(g_m).concat(g_e);
		let b = b_h.concat(b_m).concat(b_e);
		let h = b_h.concat(u_h).concat(g_h);
		let m = b_m.concat(u_m).concat(g_m);
		let e = b_e.concat(u_e).concat(g_e);
		data = [
			[
				[averageTime(u_h), averageTime(u_m), averageTime(u_e),
				averageTime(g_h), averageTime(g_m), averageTime(g_e),
				averageTime(b_h), averageTime(b_m), averageTime(b_e)],
				[averageMoves(u_h), averageMoves(u_m), averageMoves(u_e),
				averageMoves(g_h), averageMoves(g_m), averageMoves(g_e),
				averageMoves(b_h), averageMoves(b_m), averageMoves(b_e)],
				[averageTC(u_h), averageTC(u_m), averageTC(u_e),
				averageTC(g_h), averageTC(g_m), averageTC(g_e),
				averageTC(b_h), averageTC(b_m), averageTC(b_e)],
				[averageSC(u_h), averageSC(u_m), averageSC(u_e),
				averageSC(g_h), averageSC(g_m), averageSC(g_e),
				averageSC(b_h), averageSC(b_m), averageSC(b_e)]
			],
			[
				[averageTime(u), averageTime(g), averageTime(b)],
				[averageMoves(u), averageMoves(g), averageMoves(b)],
				[averageTC(u), averageTC(g), averageTC(b)],
				[averageSC(u), averageSC(g), averageSC(b)],
			],
			[
				[averageTime(h), averageTime(m), averageTime(e)],
				[averageMoves(h), averageMoves(m), averageMoves(e)],
				[averageTC(h), averageTC(m), averageTC(e)],
				[averageSC(h), averageSC(m), averageSC(e)],
			],
		];
		dataChart = {
			labels: labels[graph],
			datasets: [
				{
					label: titles[graphed],
					data: data[graph][graphed]
				},
			],
		};
		console.log(data);
		console.log(dataChart);
	};
	let clickable=true;
	let not_clickable=false;
	let graph = 0;
	let graphed = 0;

	const changeGraph = (event: any) => {
		graph = event.currentTarget.value;
		dataChart = {
			labels: labels[graph],
			datasets: [
				{
					label: titles[graphed],
					data: data[graph][graphed]
				},
			],
		};
		console.log(dataChart)
	};
	const changeGraphed = (event: any) => {
		graphed = event.currentTarget.value;
		dataChart = {
			labels: labels[graph],
			datasets: [
				{
					label: titles[graphed],
					data: data[graph][graphed]
				},
			],
		};
		console.log(dataChart)
	};
	// Bar chart
	// var chart;
	// if (browser) {
	// 	chart = new Chart(document.getElementById("bar-chart"), {
	// 		type: 'bar',
	// 		data: {
	// 		labels: labels[graph],
	// 		datasets: [
	// 				{
	// 					label: titles[graphed],
	// 					data: data[graph][graphed],
	// 				}
	// 			]
	// 		},
	// 		options: {
	// 		legend: { display: false },
	// 		title: {
	// 			display: true,
	// 			text: 'Predicted world population (millions) in 2050'
	// 		}
	// 		}
	// 	});
	// }
</script>

<div>
	{#if solved}
		<p>Le puzzle est résolu !</p>
	{:else}
		<p>Nombre de coups : {moveCount}</p>
	{/if}
	<Board {board} {moveCallback} {clickable} />
	<button on:click={handleClick} >Solve</button>

	{#if solution.length !== 0}
		<p>Move {solution.indexOf(solution_pres)}</p>
		{#key solution_pres}
			<Board board={solution_pres} {moveCallback} clickable={not_clickable} />
		{/key}
		<button on:click={handleFullLeftClick} >FullLeft</button>
		<button on:click={handleLeftClick} >Left</button>
		<button on:click={handlePlayClick} >Play</button>
		<button on:click={handleRightClick} >Right</button>
		<button on:click={handleFullRightClick} >FullRight</button>
		<div>
			<p>Took {time}</p>
			<p>Solved in {moves} moves</p>
			<p>Time complexity {time_complexity}</p>
			<p>Space complexity {space_complexity}</p>
		</div>
	{/if}
	<div>
		<button on:click={DoBenchmark} >Benchmark</button>
		{#if data.length !== 0}
			<div>
				<p>Graph by:</p>
				<label>
					<input checked={graph === 0} on:change={changeGraph} type="radio" name="graph" value={0}  />
					all
				</label>
				<label>
					<input checked={graph === 1} on:change={changeGraph} type="radio" name="graph" value={1} />
					search
				</label>
				<label>
					<input checked={graph === 2} on:change={changeGraph} type="radio" name="graph" value={2} />
					heuristic
				</label>
			</div>
			<div>
				<p>Graph Of:</p>
				<label>
					<input checked={graphed === 0} on:change={changeGraphed} type="radio" name="graphed" value={0}  />
					time
				</label>
				<label>
					<input checked={graphed === 1} on:change={changeGraphed} type="radio" name="graphed" value={1} />
					moves
				</label>
				<label>
					<input checked={graphed === 2} on:change={changeGraphed} type="radio" name="graphed" value={2} />
					time complexity
				</label>
				<label>
					<input checked={graphed === 3} on:change={changeGraphed} type="radio" name="graphed" value={3} />
					space complexity
				</label>
			</div>
			{#key dataChart}
				<canvas id="bar-chart" width="400" height="250"></canvas>
			{/key}
		{/if}
	</div>
</div>
