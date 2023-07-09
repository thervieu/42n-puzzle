<script lang="ts">
	export let name: string;
	export let inputs: { label: string; value: number | string }[];
	export let checkedValue: number | string;
	export let handleChange: (e: any) => void;
</script>

<div>
	{#each inputs as input}
		<label>
			<input
				type="radio"
				{name}
				value={input.value}
				checked={checkedValue === input.value}
				on:change={handleChange}
			/>
			{input.label}
		</label>
	{/each}
</div>
