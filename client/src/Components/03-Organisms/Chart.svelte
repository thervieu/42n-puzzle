<script lang="ts">
	import { BarChart } from 'chartist';
	import { generateRandomBoard, solvePuzzle } from '../../Helpers/board';
	import { averageMoves, averageSC, averageTC, averageTime } from '../../Helpers/stats';
	import Button from '../01-Atoms/Button.svelte';
	import Chart from '../01-Atoms/Chart.svelte';
	import InputsRow from '../02-Molecules/InputsRow.svelte';
	import LABELS from '../../Consts/labels';

	let graph = 0;
	let graphed = 0;
	let loadingBenchmark = false;
	let data: any[] = [];
	let chart: any;

	async function DoBenchmark() {
		loadingBenchmark = true;
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
		let searches = ['uniform', 'greedy', 'both'];
		let heuristics = ['hamming', 'manhattan', 'euclidean'];
		for (let k = 0; k < 3; k++) {
			let p1 = generateRandomBoard();
			let p2 = generateRandomBoard();
			let p3 = generateRandomBoard();
			let p4 = generateRandomBoard();
			let p5 = generateRandomBoard();
			let p6 = generateRandomBoard();
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
					if (searches[i] === 'both' && heuristics[j] === 'hamming') {
						b_h.push(r1);
						b_h.push(r2);
						b_h.push(r3);
						b_h.push(r4);
						b_h.push(r5);
						b_h.push(r6);
					}
					if (searches[i] === 'both' && heuristics[j] === 'manhattan') {
						b_m.push(r1);
						b_m.push(r2);
						b_m.push(r3);
						b_m.push(r4);
						b_m.push(r5);
						b_m.push(r6);
					}
					if (searches[i] === 'both' && heuristics[j] === 'euclidean') {
						b_e.push(r1);
						b_e.push(r2);
						b_e.push(r3);
						b_e.push(r4);
						b_e.push(r5);
						b_e.push(r6);
					}
				}
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
				[
					averageTime(u_h),
					averageTime(u_m),
					averageTime(u_e),
					averageTime(g_h),
					averageTime(g_m),
					averageTime(g_e),
					averageTime(b_h),
					averageTime(b_m),
					averageTime(b_e)
				],
				[
					averageMoves(u_h),
					averageMoves(u_m),
					averageMoves(u_e),
					averageMoves(g_h),
					averageMoves(g_m),
					averageMoves(g_e),
					averageMoves(b_h),
					averageMoves(b_m),
					averageMoves(b_e)
				],
				[
					averageTC(u_h),
					averageTC(u_m),
					averageTC(u_e),
					averageTC(g_h),
					averageTC(g_m),
					averageTC(g_e),
					averageTC(b_h),
					averageTC(b_m),
					averageTC(b_e)
				],
				[
					averageSC(u_h),
					averageSC(u_m),
					averageSC(u_e),
					averageSC(g_h),
					averageSC(g_m),
					averageSC(g_e),
					averageSC(b_h),
					averageSC(b_m),
					averageSC(b_e)
				]
			],
			[
				[averageTime(u), averageTime(g), averageTime(b)],
				[averageMoves(u), averageMoves(g), averageMoves(b)],
				[averageTC(u), averageTC(g), averageTC(b)],
				[averageSC(u), averageSC(g), averageSC(b)]
			],
			[
				[averageTime(h), averageTime(m), averageTime(e)],
				[averageMoves(h), averageMoves(m), averageMoves(e)],
				[averageTC(h), averageTC(m), averageTC(e)],
				[averageSC(h), averageSC(m), averageSC(e)]
			]
		];
		loadingBenchmark = false;
	}

	export async function changeGraph(event: any) {
		graph = event.currentTarget.value;
		chart = new BarChart(
			'#bar-chart',
			{
				labels: LABELS[graph],
				series: data[graph][graphed]
			},
			{
				seriesBarDistance: 30,
				distributeSeries: true
			}
		);
	}
	async function changeGraphed(event: any) {
		graphed = event.currentTarget.value;
		chart = new BarChart(
			'#bar-chart',
			{
				labels: LABELS[graph],
				series: data[graph][graphed]
			},
			{
				seriesBarDistance: 30,
				distributeSeries: true
			}
		);
	}
</script>

<div>
	<Button action={DoBenchmark} text={'Benchmark'} />
	{#if data.length !== 0}
		<div>
			<p>Graph by:</p>
			<InputsRow
				name={'graph'}
				inputs={[
					{ label: 'all', value: 0 },
					{ label: 'search', value: 1 },
					{ label: 'heuristic', value: 2 }
				]}
				checkedValue={graph}
				handleChange={changeGraph}
			/>
		</div>
		<!-- Graph Button -->
		<div>
			<p>Graph Of:</p>
			<InputsRow
				name={'graphed'}
				inputs={[
					{ label: 'time', value: 0 },
					{ label: 'moves', value: 1 },
					{ label: 'time complexity', value: 2 },
					{ label: 'space complexity', value: 3 }
				]}
				checkedValue={graphed}
				handleChange={changeGraphed}
			/>
		</div>
		<Chart />
	{/if}
</div>
