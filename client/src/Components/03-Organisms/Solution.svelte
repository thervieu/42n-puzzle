<script lang="ts">
	import RACCOON from '../../Consts/imgs';
	import { generateRandomBoard, solvePuzzle } from '../../Helpers/board';
	import Board from '../02-Molecules/Board.svelte';
	export let solution: number[][][];
	export let moveCallback: (row: number, col: number) => void;
	export let imagePaths: string[][] | null;
	export let isImageLoaded: boolean;

	export let search: string;
	export let heuristic: string;

	let board: number[][] = generateRandomBoard();

	let time: string;
	let moves: string;
	let time_complexity: number;
	let space_complexity: number;
	let solution_pres: number[][] = [];

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
		}, 100);
	};
	const handleRightClick = () => {
		let index = solution.indexOf(solution_pres);
		if (index < solution.length - 1) {
			solution_pres = solution[index + 1];
		}
	};

	const handleFullRightClick = () => {
		solution_pres = solution[solution.length - 1];
	};

	async function handleClick() {
		// Now set it to the real fetch promise
		const result = await solvePuzzle(board, search, heuristic);
		solution = result.states;
		solution_pres = solution[0];
		time = result.time.substring(12, result.time.length);
		moves = result.moves;
		time_complexity = result.time_complexity;
		space_complexity = result.space_complexity;
	}
</script>

<button on:click={handleClick}>Solve</button>
{#if solution.length !== 0}
	<p>Move {solution.indexOf(solution_pres)}</p>
	{#key solution_pres}
		<Board
			board={solution_pres}
			{moveCallback}
			clickable={false}
			imagePaths={isImageLoaded ? imagePaths.flat() : null}
		/>
	{/key}
	<button on:click={handleFullLeftClick}>FullLeft</button>
	<button on:click={handleLeftClick}>Left</button>
	<button on:click={handlePlayClick}>Play</button>
	<button on:click={handleRightClick}>Right</button>
	<button on:click={handleFullRightClick}>FullRight</button>
	<div>
		<p>Took {time}</p>
		<p>Solved in {moves} moves</p>
		<p>Time complexity {time_complexity}</p>
		<p>Space complexity {space_complexity}</p>
	</div>
{/if}
