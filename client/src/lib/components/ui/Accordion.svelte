<script>
	let { children } = $props();
</script>

<div class="accordion">
	{@render children?.()}
</div>

<style>
	.accordion {
		border-radius: var(--radius-lg);
		background-color: white;
		box-shadow: var(--shadow-md);
		overflow: hidden;
	}
</style>
