<script lang="ts">
	export let action: () => any;
	export let text: string;
	let isLoading = false;

	async function handleClick() {
		isLoading = true;

		try {
			await action();
			// Le code asynchrone à exécuter après l'appel à handleClick
		} catch (error) {
			console.error(error);
		} finally {
			isLoading = false;
		}
	}
</script>

{#if isLoading}
	<p>Chargement...</p>
{:else}
	<button on:click={handleClick}>{text}</button>
{/if}
