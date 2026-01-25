<!-- src/components/FeaturedProjects.svelte -->
<script>
	import { onMount } from "svelte";
	import Button from "./ui/Button.svelte";
	import ProjectCard from "./ui/ProjectCard.svelte";
	import { featuredProjects } from "$lib/data/featured-projects";
	import { Splide, SplideSlide } from "@splidejs/svelte-splide";
	import "@splidejs/svelte-splide/css";

	let options = {
		perPage: 3,
		perMove: 1,
		type: "loop",
		gap: "1rem",
		pagination: false,
		breakpoints: {
			992: {
				perPage: 2
			},
			680: {
				perPage: 1,
				padding: "10%",
				arrows: false
			}
		}
	};
</script>

<section id="featured-projects" class="featured-projects">
	<div class="container">
		<div class="section-title animate-on-scroll">
			<h2>Featured Campaigns</h2>
			<p>Find and fund critical needs and community dreams from verified individuals, turning collective action into real, life-changing support.</p>
		</div>

		<div class="carousel-container animate-on-scroll">
			<Splide {options}>
				{#each featuredProjects as project, index}
					<SplideSlide>
						<ProjectCard {project} />
					</SplideSlide>
				{/each}
			</Splide>
		</div>

		<div class="view-all-container animate-on-scroll">
			<Button href="/explore" primary={false}>View All Campaigns</Button>
		</div>
	</div>
</section>

<style>
	.featured-projects {
		background-color: var(--primary-bg);
	}

	.carousel-container {
		position: relative;
		margin-bottom: var(--spacing-lg);
	}

	:global(.splide__track) {
		padding: var(--spacing-sm) 0;
	}
	:global(.splide__arrow) {
		background: #fff;
		box-shadow: var(--shadow-md);
		opacity: 1;
		transition: all 0.2s ease;
		width: 2.5rem;
		height: 2.5rem;
	}

	:global(.splide__arrow:hover) {
		background-color: var(--primary-color);
		box-shadow: var(--shadow-lg);
	}

	:global(.splide__arrow svg) {
		fill: var(--primary-color);
	}

	:global(.splide__arrow:hover svg) {
		fill: white;
	}

	.view-all-container {
		text-align: center;
		margin-top: var(--spacing-xl);
	}
</style>
