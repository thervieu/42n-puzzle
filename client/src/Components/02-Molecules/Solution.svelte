<script lang="ts">
	import Board from './Board.svelte';
	export let time: string;
	export let moves: number;
	export let time_complexity: number;
	export let space_complexity: number;
	export let solution: number[][][];
	export let solution_pres: number[][];
	export let moveCallback: (row: number, col: number) => void;
	export let imagePaths: string[][] | null;
	export let isImageLoaded: boolean;

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
</script>

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
