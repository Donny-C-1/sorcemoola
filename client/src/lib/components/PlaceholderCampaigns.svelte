<script>
	let { count = 3, layout = "grid" } = $props();

	// Generate random widths for shimmer elements
	const randomWidth = () => {
		const widths = ["60%", "70%", "80%", "90%"];
		return widths[Math.floor(Math.random() * widths.length)];
	};

	// Create array of placeholder items
	const placeholders = Array.from({ length: count }, (_, i) => ({
		id: `placeholder-${i}`,
		titleWidth: randomWidth(),
		descriptionLines: Math.floor(Math.random() * 2) + 2, // 2-3 lines
		progressValue: Math.floor(Math.random() * 85) + 10 // 10-95%
	}));
</script>

<div class="placeholder-container" class:list-layout={layout === "list"}>
	{#each placeholders as placeholder}
		<div class="placeholder-card">
			<div class="placeholder-image shimmer"></div>

			<div class="placeholder-content">
				<div class="placeholder-title shimmer" style="width: {placeholder.titleWidth};"></div>

				{#each Array(placeholder.descriptionLines) as _, i}
					<div class="placeholder-text shimmer" style="width: {randomWidth()};"></div>
				{/each}

				<div class="placeholder-stats">
					<div class="placeholder-progress-container">
						<div class="placeholder-progress shimmer" style="width: {placeholder.progressValue}%;"></div>
					</div>
					<div class="placeholder-amounts">
						<div class="placeholder-amount shimmer"></div>
						<div class="placeholder-amount shimmer"></div>
					</div>
				</div>

				<div class="placeholder-footer">
					<div class="placeholder-avatar shimmer"></div>
					<div class="placeholder-creator shimmer"></div>
				</div>
			</div>
		</div>
	{/each}
</div>

<style>
	.placeholder-container {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
		gap: 24px;
		width: 100%;
	}

	.list-layout {
		grid-template-columns: 1fr;
	}

	.placeholder-card {
		display: flex;
		flex-direction: column;
		border-radius: 12px;
		overflow: hidden;
		background: rgba(255, 255, 255, 0.9);
		box-shadow: 0 2px 12px rgba(0, 0, 0, 0.05);
		height: 100%;
	}

	.list-layout .placeholder-card {
		flex-direction: row;
		height: 180px;
	}

	.list-layout .placeholder-image {
		width: 280px;
		height: 100%;
		flex-shrink: 0;
	}

	.list-layout .placeholder-content {
		flex: 1;
	}

	.placeholder-image {
		aspect-ratio: 16/9;
		background: #f0f0f0;
		width: 100%;
	}

	.placeholder-content {
		padding: 20px;
		display: flex;
		flex-direction: column;
		gap: 16px;
	}

	.placeholder-title {
		height: 24px;
		border-radius: 4px;
		background: #f0f0f0;
	}

	.placeholder-text {
		height: 16px;
		border-radius: 4px;
		background: #f0f0f0;
	}

	.placeholder-stats {
		margin-top: auto;
	}

	.placeholder-progress-container {
		height: 12px;
		background: #f0f0f0;
		border-radius: 6px;
		overflow: hidden;
		margin-bottom: 10px;
	}

	.placeholder-progress {
		height: 100%;
		background: #e0e0e0;
		border-radius: 6px;
	}

	.placeholder-amounts {
		display: flex;
		justify-content: space-between;
		margin-bottom: 16px;
	}

	.placeholder-amount {
		height: 18px;
		width: 80px;
		border-radius: 4px;
		background: #f0f0f0;
	}

	.placeholder-footer {
		display: flex;
		align-items: center;
		gap: 12px;
	}

	.placeholder-avatar {
		width: 32px;
		height: 32px;
		border-radius: 50%;
		background: #f0f0f0;
	}

	.placeholder-creator {
		height: 16px;
		width: 120px;
		border-radius: 4px;
		background: #f0f0f0;
	}

	/* Shimmer animation */
	.shimmer {
		position: relative;
		overflow: hidden;
	}

	.shimmer::after {
		content: "";
		position: absolute;
		top: 0;
		right: 0;
		bottom: 0;
		left: 0;
		background: linear-gradient(90deg, rgba(255, 255, 255, 0) 0%, rgba(60, 185, 131, 0.2) 50%, rgba(255, 255, 255, 0) 100%);
		animation: shimmer 2s infinite;
		transform: translateX(-100%);
	}

	@keyframes shimmer {
		100% {
			transform: translateX(100%);
		}
	}

	/* Responsive adjustments */
	@media (max-width: 768px) {
		.placeholder-container:not(.list-layout) {
			grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
		}

		.list-layout .placeholder-card {
			flex-direction: column;
			height: auto;
		}

		.list-layout .placeholder-image {
			width: 100%;
			aspect-ratio: 16/9;
		}
	}
</style>
