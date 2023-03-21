<script>
	import Range from '../Components/01-Atoms/Range.svelte';
	import Subtitle from '../Components/01-Atoms/Subtitle.svelte';
	import Title from '../Components/01-Atoms/Title.svelte';
	import Puzzle from '../Components/Puzzle.svelte';
	let selected = '20';

	function onChange(event) {
		selected = event.currentTarget.value;
	}
</script>

<main class="globalLayout">
	<Title value="N-Puzzle" />
	<div class="section">
		<Subtitle value="Size" />
		<Range value={3} />
	</div>
	<div class="section">
		<Subtitle value="Search" />
		<div>
			<label>
				<input checked={selected === 1} on:change={onChange} type="radio" name="amount" value="1" />
				10
			</label>
			<label>
				<input checked={selected === 2} on:change={onChange} type="radio" name="amount" value="2" />
				10
			</label>
		</div>
	</div>
	<div class="section">
		<Subtitle value="Heuristic" />
		<div>
			<label>
				<input checked={selected === 3} on:change={onChange} type="radio" name="amount" value="3" />
				10
			</label>
			<label>
				<input checked={selected === 4} on:change={onChange} type="radio" name="amount" value="4" />
				10
			</label>
			<label>
				<input checked={selected === 5} on:change={onChange} type="radio" name="amount" value="5" />
				10
			</label>
		</div>
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
