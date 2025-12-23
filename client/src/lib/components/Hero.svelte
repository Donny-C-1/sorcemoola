<script>
	import Button from './ui/Button.svelte';
</script>

<section class="hero-section">
	<div class="overlay"></div>
	<div class="parallax-bg"></div>
	<div class="hero-content container">
		<div class="hero-text">
			<h1>Fund Your Dreams, Support Others</h1>
			<p>
				Whether it’s an urgent campaign or a collaborative project, SorceMoola helps you achieve your funding goal, securely.
			</p>
			<div class="hero-buttons">
				<Button primary={true} large={true} href="/dashboard/create">Start a Campaign</Button>
				<Button primary={false} large={true} href="/explore">Donate Now</Button>
			</div>
		</div>

		<div class="hero-image">
			<img
				src="images/hero-image.webp"
				alt="People collaborating on projects"
				loading="eager"
			/>
		</div>
	</div>
</section>

<style>
	.hero-section {
		position: relative;
		height: 90vh;
    padding: 0;
		min-height: 38rem;
		display: flex;
		align-items: center;
		overflow: hidden;
		background-color: var(--primary-bg);
	}

	.overlay {
		position: absolute;
    inset: 0;
		background: linear-gradient(135deg, rgba(52, 162, 115, 0.1) 0%, rgba(52, 162, 115, 0.2) 100%);
		z-index: 1;
	}

	.parallax-bg {
		position: absolute;
		inset: 0;
		background-image: url('/images/dots-pattern.svg');
		opacity: 0.4;
		z-index: 0;
		transition: transform 0.5s ease-out;
	}

	.hero-section:hover .parallax-bg {
		transform: translateY(-20px);
	}

	.hero-content {
		position: relative;
		z-index: 2;
		display: flex;
		justify-content: space-between;
		align-items: center;
	}

	.hero-text {
		flex: 1;
		max-width: 37.5rem;
		padding-right: var(--spacing-xl);
	}

	.hero-text h1 {
		color: var(--text-dark);
		margin-bottom: var(--spacing-md);
		font-size: 3.5rem;
		font-weight: 700;
		line-height: 1.2;
	}

	.hero-text p {
		font-size: 1.25rem;
		color: var(--text-medium);
		margin-bottom: var(--spacing-lg);
	}

	.hero-buttons {
		display: flex;
		gap: var(--spacing-md);
		flex-wrap: wrap;
	}

	.hero-image {
		flex: 1;
		max-width: 32rem;
		display: flex;
		justify-content: flex-end;
		animation: float 6s ease-in-out infinite;
	}

	.hero-image img {
		width: 100%;
		height: auto;
	}

	@keyframes float {
		0% {
			transform: translateY(0px);
		}
		50% {
			transform: translateY(-20px);
		}
		100% {
			transform: translateY(0px);
		}
	}

	@media (max-width: 992px) {
    .hero-section {
			height: auto;
			padding-top: var(--spacing-md);
			padding-bottom: var(--spacing-xl);
		}

		.hero-content {
			flex-direction: column-reverse;
			text-align: center;
		}

		.hero-text {
			max-width: 100%;
			padding-right: 0;
			margin-top: var(--spacing-xl);
		}

		.hero-buttons {
			justify-content: center;
		}

		.hero-image {
			max-width: 400px;
			margin-top: var(--spacing-xl);
			justify-content: center;
		}
	}

	@media (max-width: 576px) {
		.hero-text h1 {
			font-size: 2.5rem;
		}

		.hero-image {
			max-width: 280px;
		}

		.hero-buttons {
			flex-direction: column;
			gap: var(--spacing-sm);
		}
	}
</style>
