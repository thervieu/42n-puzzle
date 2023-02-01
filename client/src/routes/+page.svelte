<script>
	// import Range from '../Components/Range.svelte';
	import Puzzle from '../Components/Puzzle.svelte';
	let selected = '20';

	function onChange(event) {
		selected = event.currentTarget.value;
	}
</script>

<main class="globalLayout">
	<h1>N-Puzzle</h1>
	<div class="section">
		<h3>Size</h3>
		<!-- <Range /> -->
	</div>
	<div class="section">
		<h3>Search</h3>
		<label>
			<input checked={selected === 1} on:change={onChange} type="radio" name="amount" value="1" /> 10
		</label>
		<label>
			<input checked={selected === 2} on:change={onChange} type="radio" name="amount" value="2" /> 10
		</label>
	</div>
	<div class="section">
		<h3>Heuristic</h3>
		<label>
			<input checked={selected === 3} on:change={onChange} type="radio" name="amount" value="3" /> 10
		</label>
		<label>
			<input checked={selected === 4} on:change={onChange} type="radio" name="amount" value="4" /> 10
		</label>
		<label>
			<input checked={selected === 5} on:change={onChange} type="radio" name="amount" value="5" /> 10
		</label>
	</div>
	<Puzzle />
	<button>Solve</button>
</main>

<style>
	.globalLayout {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}
	.section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		width: 100%;
	}
</style>
