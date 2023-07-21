<script lang="ts">
	import { BarChart } from 'chartist';
	import InputsRow from './InputsRow.svelte';
	import { generateRandomBoard, solvePuzzle } from '../../Helpers/board'
	import { averageMoves, averageSC, averageTC, averageTime } from '../../Helpers/stats';

	const labels = [
		[
			'uniform hamming',
			'uniform manhattan',
			'uniform euclidean',
			'greedy hamming',
			'greedy manhattan',
			'greedy euclidean'
		],
		['uniform', 'greedy'],
		['hamming', 'manhattan', 'euclidian']
	];
	let graph = "all";
	let graphed = "time";
    async function changeGraph(event) {
		graph = event.currentTarget.value;
		let graphs = ['all', 'search', 'heuristic'];
		let grapheds = ['time', 'moves', 'time complexity', 'space complexity'];
		chart = new BarChart(
			'#bar-chart',
			{
				labels: labels[graphs.indexOf(graph)],
				series: data[graphs.indexOf(graph)][grapheds.indexOf(graphed)]
			},
			{
				seriesBarDistance: 30,
				distributeSeries: true
			}
		);
	}
	async function changeGraphed(event) {
		graphed = event.currentTarget.value;
		let graphs = ['all', 'search', 'heuristic'];
		let grapheds = ['time', 'moves', 'time complexity', 'space complexity'];
		chart = new BarChart(
			'#bar-chart',
			{
				labels: labels[graphs.indexOf(graph)],
				series: data[graphs.indexOf(graph)][grapheds.indexOf(graphed)]
			},
			{
				seriesBarDistance: 30,
				distributeSeries: true
			}
		);
	};
    var chart: any;
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
		let searches = ['uniform', 'greedy'];
		let heuristics = ['hamming', 'manhattan', 'euclidean'];
		for (let k = 0; k < 3; k++) {
			let p1 = generateRandomBoard(3);
			let p2 = generateRandomBoard(3);
			let p3 = generateRandomBoard(3);
			let p4 = generateRandomBoard(3);
			let p5 = generateRandomBoard(3);
			let p6 = generateRandomBoard(3);
			for (let i in searches) {
				for (let j in heuristics) {
					const [r1, r2, r3, r4, r5, r6] = await Promise.all([
						solvePuzzle(p1, searches[i], heuristics[j]),
						solvePuzzle(p2, searches[i], heuristics[j]),
						solvePuzzle(p3, searches[i], heuristics[j]),
						solvePuzzle(p4, searches[i], heuristics[j]),
						solvePuzzle(p5, searches[i], heuristics[j]),
						solvePuzzle(p6, searches[i], heuristics[j])
					]);
					r1.time = r1.time.substring(12, r1.time.length);
					r2.time = r2.time.substring(12, r2.time.length);
					r3.time = r3.time.substring(12, r3.time.length);
					r4.time = r4.time.substring(12, r4.time.length);
					r5.time = r5.time.substring(12, r5.time.length);
					r6.time = r6.time.substring(12, r6.time.length);
					if (searches[i] === 'uniform' && heuristics[j] === 'hamming') {
						u_h.push(r1);
						u_h.push(r2);
						u_h.push(r3);
						u_h.push(r4);
						u_h.push(r5);
						u_h.push(r6);
					}
					if (searches[i] === 'uniform' && heuristics[j] === 'manhattan') {
						u_m.push(r1);
						u_m.push(r2);
						u_m.push(r3);
						u_m.push(r4);
						u_m.push(r5);
						u_m.push(r6);
					}
					if (searches[i] === 'uniform' && heuristics[j] === 'euclidean') {
						u_e.push(r1);
						u_e.push(r2);
						u_e.push(r3);
						u_e.push(r4);
						u_e.push(r5);
						u_e.push(r6);
					}
					if (searches[i] === 'greedy' && heuristics[j] === 'hamming') {
						g_h.push(r1);
						g_h.push(r2);
						g_h.push(r3);
						g_h.push(r4);
						g_h.push(r5);
						g_h.push(r6);
					}
					if (searches[i] === 'greedy' && heuristics[j] === 'manhattan') {
						g_m.push(r1);
						g_m.push(r2);
						g_m.push(r3);
						g_m.push(r4);
						g_m.push(r5);
						g_m.push(r6);
					}
					if (searches[i] === 'greedy' && heuristics[j] === 'euclidean') {
						g_e.push(r1);
						g_e.push(r2);
						g_e.push(r3);
						g_e.push(r4);
						g_e.push(r5);
						g_e.push(r6);
					}
				}
			}
		}
		let u = u_h.concat(u_m).concat(u_e);
		let g = g_h.concat(g_m).concat(g_e);
		let h = u_h.concat(g_h);
		let m = u_m.concat(g_m);
		let e = u_e.concat(g_e);
		data = [
			[
				[
					averageTime(u_h),
					averageTime(u_m),
					averageTime(u_e),
					averageTime(g_h),
					averageTime(g_m),
					averageTime(g_e),
				],
				[
					averageMoves(u_h),
					averageMoves(u_m),
					averageMoves(u_e),
					averageMoves(g_h),
					averageMoves(g_m),
					averageMoves(g_e),
				],
				[
					averageTC(u_h),
					averageTC(u_m),
					averageTC(u_e),
					averageTC(g_h),
					averageTC(g_m),
					averageTC(g_e),
				],
				[
					averageSC(u_h),
					averageSC(u_m),
					averageSC(u_e),
					averageSC(g_h),
					averageSC(g_m),
					averageSC(g_e),
				]
			],
			[
				[averageTime(u), averageTime(g)],
				[averageMoves(u), averageMoves(g)],
				[averageTC(u), averageTC(g)],
				[averageSC(u), averageSC(g)]
			],
			[
				[averageTime(h), averageTime(m), averageTime(e)],
				[averageMoves(h), averageMoves(m), averageMoves(e)],
				[averageTC(h), averageTC(m), averageTC(e)],
				[averageSC(h), averageSC(m), averageSC(e)]
			]
		];
	}
</script>


<div>
    <button on:click={DoBenchmark}>Benchmark</button>
    {#if data.length !== 0}
        <div>
            <p>Graph by:</p>
            <InputsRow
                name={"graph"}
                inputs={['all', 'search', 'heuristic']}
                checkedValue={graph}
                handleChange={changeGraph}
            />
        </div>
        <!-- Graph Button -->
        <div>
            <p>Graph Of:</p>
            <InputsRow
                name={"graphed"}
                inputs={['time', 'moves', 'time complexity', 'space complexity']}
                checkedValue={graphed}
                handleChange={changeGraphed}
            />
        </div>
        <!-- Chart -->
        <div id="bar-container">
            <div id="bar-chart" />
        </div>
    {/if}
</div>


<style lang="scss">
	#bar-chart {
		background-color: rgb(55, 71, 79);
		width: 1000px;
		height: 500px;
		font-family: Lato, Helvetica-Neue, monospace;
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
