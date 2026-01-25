<script>
	let { testimonial } = $props();
</script>

<div class="testimonial">
	<div class="testimonial_content">
		<div class="quote_mark">"</div>
		<p class="testimonial_text">{testimonial.quote}</p>
	</div>
	<div class="testimonial_author">
		<img src={testimonial.avatar} alt={testimonial.name} loading="lazy" />
		<div class="author_details">
			<h3 class="author_name">{testimonial.name}</h3>
			<p class="author_role">{testimonial.role}</p>
			<p class="project_name">{testimonial.project}</p>
		</div>
	</div>
</div>

<style>
	.testimonial {
		display: flex;
		flex-direction: column;
		position: relative;
		padding: var(--spacing-xl);
		background-color: white;
		border-radius: var(--radius-lg);
		box-shadow: var(--shadow-md);
		height: 100%;
		min-height: 25rem;
	}

	.testimonial_content {
		flex-grow: 1;
	}

	.quote_mark {
		font-size: 5rem;
		line-height: 1;
		position: absolute;
		top: var(--spacing-md);
		left: var(--spacing-md);
		color: rgba(52, 162, 115, 0.1);
		font-family: serif;
	}

	.testimonial_text {
		font-size: 1.25rem;
		line-height: 1.6;
		margin-top: var(--spacing-md);
	}

	.testimonial_author {
		display: flex;
		align-items: center;
		margin-top: var(--spacing-lg);
	}

	.testimonial_author img {
		width: 3.75rem;
		height: 3.75rem;
		border-radius: 50%;
		object-fit: cover;
		margin-right: var(--spacing-md);
	}

	.author_name {
		font-weight: 600;
		font-size: 1.1rem;
	}

	.author_role {
		font-size: 0.875rem;
		color: var(--primary-color);
		margin-top: 0.2rem;
	}

	.project_name {
		color: var(--text-medium);
		font-size: 0.875rem;
		margin-top: 0.2rem;
	}

	@media (max-width: 768px) {
		.testimonial {
			padding: var(--spacing-md);
		}

		.quote_mark {
			left: var(--spacing-sm);
			top: var(--spacing-sm);
		}

		.testimonial_text {
			font-size: 1.1rem;
		}
	}
</style>
