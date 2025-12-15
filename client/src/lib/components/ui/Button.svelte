<script>
	let {
		handler,
		primary = true,
		neutral = false,
		inverted = false,
		large = false,
		wide = false,
		disabled = false,
		href = undefined,
		type = 'button',
		children
	} = $props();
</script>

{#if href}
	<a {href} class:primary class:neutral class:inverted class:large {disabled}>{@render children?.()}</a>
{:else}
	<button onclick={handler} {type} class:primary class:neutral class:inverted class:large class:wide {disabled}
		>{@render children?.()}</button
	>
{/if}

<style>
	button,
	a {
		display: inline-flex;
		align-items: center;
		justify-content: center;
		font-family: 'Poppins', sans-serif;
		font-weight: 600;
		font-size: 1rem;
		padding: 12px 24px;
		border-radius: var(--radius-md);
		text-decoration: none;
		cursor: pointer;
		transition: all 0.2s ease-in-out;
		border: none;
		outline: none;
		position: relative;
		overflow: hidden;
		z-index: 1;
	}

	button::after,
	a::after {
		content: '';
		position: absolute;
		width: 100%;
		height: 100%;
		top: 0;
		left: -100%;
		background: linear-gradient(
			90deg,
			rgba(255, 255, 255, 0) 0%,
			rgba(255, 255, 255, 0.2) 50%,
			rgba(255, 255, 255, 0) 100%
		);
		transition: left 0.5s ease;
		z-index: -1;
	}

	button:hover::after,
	a:hover::after {
		left: 100%;
	}

	.primary {
		background-color: var(--primary-color);
		color: white;
		box-shadow: 0 4px 15px rgba(52, 162, 115, 0.3);
	}

	.primary.inverted {
		background-color: white;
		color: var(--primary-color);
		box-shadow: 0 4px 15px rgba(255, 255, 255, 0.4);
	}

	.primary:hover {
		background-color: var(--primary-dark);
		transform: translateY(-2px);
		box-shadow: 0 6px 20px rgba(52, 162, 115, 0.4);
	}

	.primary.inverted:hover {
		background-color: var(--light-gray, #f0f0f0);
		transform: translateY(-2px);
		box-shadow: 0 6px 20px rgba(255, 255, 255, 0.5);
	}

	.primary:active {
		transform: translateY(1px);
		box-shadow: 0 2px 10px rgba(52, 162, 115, 0.3);
	}

	/* Neutral button styles */
	.neutral {
		background-color: var(--neutral-color, #6c757d);
		color: white;
		box-shadow: 0 4px 15px rgba(108, 117, 125, 0.3);
	}

	.neutral.inverted {
		background-color: white;
		color: var(--neutral-color, #6c757d);
		box-shadow: 0 4px 15px rgba(255, 255, 255, 0.4);
	}

	.neutral:hover {
		background-color: var(--neutral-dark, #5a6268);
		transform: translateY(-2px);
		box-shadow: 0 6px 20px rgba(108, 117, 125, 0.4);
	}

	.neutral.inverted:hover {
		background-color: var(--light-gray, #f0f0f0);
		transform: translateY(-2px);
		box-shadow: 0 6px 20px rgba(255, 255, 255, 0.5);
	}

	.neutral:active {
		transform: translateY(1px);
		box-shadow: 0 2px 10px rgba(108, 117, 125, 0.3);
	}

	/* Outline button styles */
	button:not(.primary):not(.neutral),
	a:not(.primary):not(.neutral) {
		background-color: transparent;
		color: var(--primary-color);
		border: 2px solid var(--primary-color);
	}

	button:not(.primary):not(.neutral).inverted,
	a:not(.primary):not(.neutral).inverted {
		background-color: transparent;
		color: white;
		border: 2px solid white;
	}

	button:not(.primary):not(.neutral):hover,
	a:not(.primary):not(.neutral):hover {
		background-color: rgba(52, 162, 115, 0.1);
		transform: translateY(-2px);
	}

	button:not(.primary):not(.neutral):active,
	a:not(.primary):not(.neutral):active {
		transform: translateY(1px);
	}

	button:not(.primary):not(.neutral).inverted:hover,
	a:not(.primary):not(.neutral).inverted:hover {
		background-color: rgba(255, 255, 255, 0.1);
		transform: translateY(-2px);
	}

	/* Outlined neutral button */
	button:not(.primary).neutral,
	a:not(.primary).neutral {
		background-color: transparent;
		color: var(--text-dark, #6c757d);
		border: 2px solid var(--text-light, #6c757d);
	}

	button:not(.primary).neutral:hover,
	a:not(.primary).neutral:hover {
		background-color: rgba(108, 117, 125, 0.1);
		transform: translateY(-2px);
	}

	button:not(.primary).neutral.inverted,
	a:not(.primary).neutral.inverted {
		background-color: transparent;
		color: white;
		border: 2px solid white;
	}

	.large {
		font-size: 1.125rem;
		padding: 14px 28px;
	}

	.wide {
		width: 100%;
	}

	@media (max-width: 576px) {
		button,
		a {
			width: 100%;
		}

		.large {
			padding: 12px 24px;
			font-size: 1rem;
		}
	}
</style>