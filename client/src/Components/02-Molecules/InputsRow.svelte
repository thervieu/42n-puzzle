<script lang="ts">
	export let inputs: { label: string; value: number }[];
	export let checkedValue: number;
	export let handleChange: (e: any) => void;
</script>

<div>
	{#each inputs as input}
		<label>
			<input
				type="radio"
				name="graphed"
				value={input.value}
				checked={checkedValue === input.value}
				on:change={handleChange}
			/>
			{input.label}
		</label>
	{/each}
</div>
