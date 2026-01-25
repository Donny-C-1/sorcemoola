<script>
	import ProgressBar from "./ProgressBar.svelte";

	let { project } = $props();
</script>

<div class="project-card">
	<div class="project-image">
		<img src={project.imageUrl} alt={project.title} loading="lazy" />
		<div class="project-category">{project.category}</div>
	</div>

	<div class="project-content">
		<h3 class="project-title">{project.title}</h3>
		<p class="project-description">{project.description}</p>

		<div class="project-creator">
			<img src={project.creatorAvatar} alt={project.creatorName} loading="lazy" />
			<span>by {project.creatorName}</span>
		</div>

		<div class="project-progress">
			<ProgressBar value={project.fundedPercentage} />
			<div class="project-stats">
				<div class="stat">
					<span class="value">${project.fundedAmount.toLocaleString()}</span>
					<span class="label">raised of ${project.goalAmount.toLocaleString()}</span>
				</div>
				<div class="stat">
					<span class="value">{project.daysLeft}</span>
					<span class="label">days left</span>
				</div>
			</div>
		</div>

		<a href={`/campaigns/${project.slug}`} class="project-link">View Campaign</a>
	</div>
</div>

<style>
	.project-card {
		background-color: white;
		border-radius: var(--radius-lg);
		overflow: hidden;
		box-shadow: var(--shadow-md);
		transition: all 0.3s ease;
		height: 100%;
		display: flex;
		flex-direction: column;
	}

	.project-image {
		position: relative;
		height: 200px;
		overflow: hidden;
	}

	.project-image img {
		width: 100%;
		height: 100%;
		object-fit: cover;
		transition: transform 0.5s ease;
	}

	.project-card:hover .project-image img {
		transform: scale(1.05);
	}

	.project-category {
		position: absolute;
		top: var(--spacing-sm);
		right: var(--spacing-sm);
		background-color: var(--primary-color);
		color: white;
		padding: 4px 12px;
		border-radius: var(--radius-sm);
		font-size: 0.875rem;
		font-weight: 500;
	}

	.project-content {
		padding: var(--spacing-md);
		display: flex;
		flex-direction: column;
		flex-grow: 1;
	}

	.project-title {
		margin-top: 0;
		margin-bottom: var(--spacing-sm);
		font-size: 1.25rem;
		line-height: 1.4;
	}

	.project-description {
		color: var(--text-medium);
		font-size: 0.9rem;
		margin-bottom: var(--spacing-md);
		flex-grow: 1;
	}

	.project-creator {
		display: flex;
		align-items: center;
		margin-bottom: var(--spacing-md);
	}

	.project-creator img {
		width: 30px;
		height: 30px;
		border-radius: 50%;
		margin-right: var(--spacing-xs);
	}

	.project-creator span {
		font-size: 0.875rem;
		color: var(--text-medium);
	}

	.project-progress {
		margin-bottom: var(--spacing-md);
	}

	.project-stats {
		display: flex;
		justify-content: space-between;
		margin-top: var(--spacing-xs);
	}

	.stat {
		display: flex;
		flex-direction: column;
	}

	.stat .value {
		font-weight: 600;
		color: var(--text-dark);
	}

	.stat .label {
		font-size: 0.75rem;
		color: var(--text-medium);
	}

	.project-link {
		display: inline-block;
		text-align: center;
		padding: var(--spacing-sm);
		background-color: var(--primary-bg);
		color: var(--primary-color);
		font-weight: 600;
		border-radius: var(--radius-md);
		text-decoration: none;
		transition: all 0.2s ease;
	}

	.project-link:hover {
		background-color: var(--primary-color);
		color: white;
	}
</style>
