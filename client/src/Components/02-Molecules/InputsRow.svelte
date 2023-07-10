<script lang="ts">
	export let name: string;
	export let inputs: string[];
	export let checkedValue: string;
	export let handleChange: (e: any) => void;
</script>

<div>
	{#each inputs as input}
		<label>
			<input
				type="radio"
				name={name}
				value={input}
				checked={checkedValue === input}
				on:change={handleChange}
			/>
			{input}
		</label>
	{/each}
</div>
