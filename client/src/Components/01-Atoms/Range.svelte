<script lang="ts">
	export let min: number = 3;
	export let max: number = 5;
	export let value: number = min;
	export 

	function handleChange(event: Event) {
		if (event.target) {
			const target = event.target as HTMLInputElement;
			value = Number(target.value);
		}
	}
</script>

<input type="range" {min} {max} bind:value on:input={handleChange} />

<p>Value: {value}</p>
