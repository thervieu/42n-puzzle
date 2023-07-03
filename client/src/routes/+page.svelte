<script>
	import Range from '../Components/01-Atoms/Range.svelte';
	import Subtitle from '../Components/01-Atoms/Subtitle.svelte';
	import Title from '../Components/01-Atoms/Title.svelte';
	import Puzzle from '../Components/Puzzle.svelte';
	let search = "uniform";
	let heuristic = "hamming";
	let size = 3;
	function onChangeSearch(event) {
		search = event.currentTarget.value;
	}
	function onChangeHeuristic(event) {
		heuristic = event.currentTarget.value;
	}
</script>

<main class="globalLayout">
	<Title value="N-Puzzle" />
	<div class="section">
		<Subtitle value="Search" />
		<div>
			<label>
				<input checked={search === "uniform"} on:change={onChangeSearch} type="radio" name="search" value={"uniform"} />
				uniform cost
			</label>
			<label>
				<input checked={search === "greedy"} on:change={onChangeSearch} type="radio" name="search" value={"greedy"} />
				greedy
			</label>
			<label>
				<input checked={search === "both"} on:change={onChangeSearch} type="radio" name="search" value={"both"} />
				both
			</label>
		</div>
	</div>
	<div class="section">
		<Subtitle value="Heuristic" />
		<div>
			<label>
				<input checked={heuristic === "hamming"} on:change={onChangeHeuristic} type="radio" name="heuristic" value={"hamming"}  />
				hamming
			</label>
			<label>
				<input checked={heuristic === "manhattan"} on:change={onChangeHeuristic} type="radio" name="heuristic" value={"manhattan"} />
				manhattan
			</label>
			<label>
				<input checked={heuristic === "euclidean"} on:change={onChangeHeuristic} type="radio" name="heuristic" value={"euclidean"} />
				euclidean
			</label>
		</div>
	</div>
	{#key size}
		<Puzzle search={search} heuristic={heuristic}/>
	{/key}
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
