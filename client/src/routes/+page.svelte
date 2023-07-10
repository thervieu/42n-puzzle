<script lang="ts">
	import Range from '../Components/01-Atoms/Range.svelte';
	import Subtitle from '../Components/01-Atoms/Subtitle.svelte';
	import Title from '../Components/01-Atoms/Title.svelte';
	import InputsRow from '../Components/02-Molecules/InputsRow.svelte';
	import Puzzle from '../Components/Puzzle.svelte';
	let search = "uniform";
	let heuristic = "hamming";
	let size = 3;
	function onChangeSearch(event: any) {
		search = event.currentTarget.value;
	}
	function onChangeHeuristic(event: any) {
		heuristic = event.currentTarget.value;
	}
</script>

<main class="globalLayout">
	<Title value="N-Puzzle" />
	<div class="section">
		<Subtitle value="Size" />
		<Range bind:value={size} /> 
	</div>
	<div class="section">
		<Subtitle value="Search" />
		<InputsRow name={"search"}
			inputs={["uniform", "greedy", "both"]}
			checkedValue={"uniform"}
			handleChange={onChangeSearch} />
	</div>
	<div class="section">
		<Subtitle value="Heuristic" />
		<InputsRow name={"heuristic"}
			inputs={["hamming", "manhattan", "euclidean"]}
			checkedValue={"hamming"}
			handleChange={onChangeHeuristic} />
	</div>
	{#key size}
		<Puzzle size={size} search={search} heuristic={heuristic} />
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
