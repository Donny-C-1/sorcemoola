<script>
	let { size } = $props();
</script>

<div class="logo" style:--size={size}>
	<a href="/"><span class="logo_text">Sorce<span class="highlight">Moola</span></span></a>
</div>

<style>
	.logo {
		font-size: var(--size, 1.8rem);
		font-weight: 700;
		font-family: "Quicksand";
	}

	.logo_text {
		letter-spacing: -0.5px;
	}

	.highlight {
		color: var(--primary-color);
	}
</style>
