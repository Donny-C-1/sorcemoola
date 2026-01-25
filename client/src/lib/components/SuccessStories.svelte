<script>
	import { onMount } from "svelte";
	import { testimonials } from "$lib/data/testimonials";
	import { Splide, SplideSlide } from "@splidejs/svelte-splide";
	import "@splidejs/svelte-splide/css";
	import Testimonial from "./ui/Testimonial.svelte";

	let options = {
		perPage: 1,
		type: "loop",
		gap: "2rem",
		autoplay: true,
		padding: "1rem",
		breakpoints: {
			680: {
				arrows: false
			}
		}
	};
</script>

<section id="success-stories" class="success-stories">
	<div class="container">
		<div class="section-title animate-on-scroll">
			<h2>Success Stories</h2>
			<p>Hear from creators and backers who have found success through SorceMoola.</p>
		</div>

		<div class="testimonial-slider animate-on-scroll">
			<Splide {options}>
				{#each testimonials as testimonial, index}
					<SplideSlide>
						<Testimonial {testimonial} />
					</SplideSlide>
				{/each}
			</Splide>
		</div>
	</div>
</section>

<style>
	.success-stories {
		background-color: var(--primary-bg);
		position: relative;
		overflow: hidden;
	}

	.success-stories::before {
		content: "";
		position: absolute;
		top: -30px;
		right: -30px;
		width: 200px;
		height: 200px;
		background-color: rgba(52, 162, 115, 0.1);
		border-radius: 50%;
		z-index: 0;
	}

	.testimonial-slider {
		position: relative;
		margin: 0 auto;
		max-width: 900px;
		overflow: hidden;
	}

	:global(.splide__track) {
		padding: var(--spacing-xl) 0;
	}

	:global(.splide__pagination) {
		gap: 0.25rem;
		bottom: 0rem;
		z-index: 999;
	}

	:global(.splide__pagination__page) {
		background-color: var(--text-light);
		width: 0.75rem;
		height: 0.75rem;
		opacity: 1;
	}

	:global(.splide__pagination__page.is-active) {
		background-color: var(--primary-color);
	}
</style>
