<!-- src/components/Categories.svelte -->
<script>
	import CategoryCard from "./ui/CategoryCard.svelte";
	import { categories } from "$lib/data/categories";
</script>

<section id="categories" class="categories-section">
	<div class="container">
		<div class="section-title animate-on-scroll">
			<h2>Explore Categories</h2>
			<p>Discover projects in various categories that match your interests and passions.</p>
		</div>

		<div class="categories-grid animate-on-scroll">
			{#each categories as category, index}
				<CategoryCard {category} style="transition-delay: {index * 100}ms" />
			{/each}
		</div>
	</div>
</section>

<style>
	.categories-section {
		background-color: white;
	}

	.categories-grid {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
		gap: var(--spacing-md);
	}

	@media (max-width: 576px) {
		.categories-grid {
			grid-template-columns: repeat(2, 1fr);
			gap: var(--spacing-sm);
		}
	}
</style>
