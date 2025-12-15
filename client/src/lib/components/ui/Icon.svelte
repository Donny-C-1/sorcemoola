<script>
	let { icon } = $props();
</script>

<span><i>{@html icon}</i></span>

<style>
	span {
		display: inline-grid;
		place-items: center;
		padding-inline: 0.75rem;
		color: var(--text-dark);
	}
	i {
		display: inline-grid;
		place-items: center;
		vertical-align: middle;
		width: 1.5rem;
		padding: 0.2rem;
	}
</style>
