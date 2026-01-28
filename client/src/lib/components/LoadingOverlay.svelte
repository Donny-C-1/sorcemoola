<script>
	let { show = false, text = "" } = $props();
</script>

{#if show}
	<div class="loading-overlay">
		<div class="loading-content">
			{#if text}
				<p class="loading-text">{text}</p>
			{/if}
			<div class="pulsing-dots">
				<span class="dot"></span>
				<span class="dot"></span>
				<span class="dot"></span>
			</div>
		</div>
	</div>
{/if}

<style>
	.loading-overlay {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		background-color: rgba(52, 162, 115, 0.15);
		backdrop-filter: blur(8px);
		-webkit-backdrop-filter: blur(8px);
		z-index: 9999;
		display: flex;
		align-items: center;
		justify-content: center;
		pointer-events: all;
		cursor: not-allowed;
	}

	.loading-content {
		display: flex;
		flex-direction: row;
		align-items: center;
		gap: var(--spacing-md);
		padding: var(--spacing-xl);
		/* background-color: rgba(255, 255, 255, 0.95); */
		/* border-radius: var(--radius-lg); */
		/* box-shadow: var(--shadow-lg); */
	}

	.pulsing-dots {
		display: flex;
		align-items: center;
		gap: var(--spacing-xs);
	}

	.dot {
		width: 12px;
		height: 12px;
		background-color: var(--primary-color);
		border-radius: 50%;
		animation: pulse 1.4s ease-in-out infinite;
	}

	.dot:nth-child(1) {
		animation-delay: 0s;
	}

	.dot:nth-child(2) {
		animation-delay: 0.2s;
	}

	.dot:nth-child(3) {
		animation-delay: 0.4s;
	}

	@keyframes pulse {
		0%,
		100% {
			transform: scale(1);
			opacity: 1;
		}
		50% {
			transform: scale(1.5);
			opacity: 0.5;
		}
	}

	.loading-text {
		margin: 0;
		font-size: 1.125rem;
		font-weight: 600;
		color: var(--text-dark);
		text-align: center;
	}

	/* Customizable overlay color - can be changed via CSS variable */
	.loading-overlay {
		/* Default uses primary color with opacity */
		/* To customize, set --overlay-color in parent component or globally */
		background-color: var(--overlay-color, rgba(52, 162, 115, 0.15));
	}
</style>
