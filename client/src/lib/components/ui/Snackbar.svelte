<script>
	import { fade, fly } from "svelte/transition";

	let { children, visible = $bindable(false), type = "info", duration = 3000 } = $props();
    let timeout;

    $effect(() => {
        if (visible) {
            clearTimeout(timeout);
            timeout = setTimeout(() => visible = false, duration);
        }
        return () => clearTimeout(timeout);
    })
</script>

{#if visible}
	<div class="snackbar {type}" in:fly={{ y: 30, duration: 500 }} out:fade>{@render children()}</div>
{/if}

<style>
	.snackbar {
		position: fixed;
		bottom: 2rem;
		left: 50%;
		translate: -50%;
		z-index: 1;
		min-width: 16rem;
		text-align: center;
		border-radius: var(--radius-sm);
		padding: var(--spacing-sm);
		background-color: var(--text-dark);
		color: var(--text-white);
        box-shadow: var(--shadow-md);
	}

    .info {
        background-color: var(--primary-color);
    }

    .error {
        background-color: var(--error-color);
    }
</style>
