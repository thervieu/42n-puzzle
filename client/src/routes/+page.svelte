
<script>
	import Range from "../Components/Range.svelte"	
   
   	let selected = '20';
	
	function onChange(event) {
		selected = event.currentTarget.value;
	}
</script>

<h1>N-Puzzle</h1>
<h3>Size</h3>
<Range />
<h3>Search</h3>
<label>
	<input checked={selected===1} on:change={onChange} type="radio" name="amount" value="1" /> 10
</label>
<label>
	<input checked={selected===2} on:change={onChange} type="radio" name="amount" value="2" /> 10
</label>
<h3>Heuristic</h3>
<label>
	<input checked={selected===3} on:change={onChange} type="radio" name="amount" value="3" /> 10
</label>
<label>
	<input checked={selected===4} on:change={onChange} type="radio" name="amount" value="4" /> 10
</label>
<label>
	<input checked={selected===5} on:change={onChange} type="radio" name="amount" value="5" /> 10
</label>
<button>Solve</button>