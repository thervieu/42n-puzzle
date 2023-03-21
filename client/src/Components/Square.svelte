<script>
	export let value;
	export let index;

	function isCorrect() {
		return value == index + 1;
	}
	// class={isCorrect() ? 'correct' : 'incorrect'}
</script>

<div class={isCorrect() ? 'correct' : 'incorrect'}>{value}</div>

<style>
	.ctn {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 50px;
		width: 50px;
		border: 1px solid black;
		border-radius: 5px;
		margin: 2px;
	}
	.correct {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 50px;
		width: 50px;
		border: 1px solid black;
		border-radius: 5px;
		margin: 2px;
		background-color: #333333;
		color: #00ff00;
	}
	.incorrect {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 50px;
		width: 50px;
		border: 1px solid black;
		border-radius: 5px;
		margin: 2px;
		background-color: #333333;
		color: #ff0000;
	}
</style>
