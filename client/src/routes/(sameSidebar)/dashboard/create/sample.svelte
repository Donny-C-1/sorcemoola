<script>
	import { fade, fly } from "svelte/transition";
	import { cubicOut } from "svelte/easing";
	import BasicSteps from "$lib/components/BasicSteps.svelte";
	import StoryStep from "$lib/components/StoryStep.svelte";
	import RewardsStep from "$lib/components/RewardsStep.svelte";
	import PaymentStep from "$lib/components/PaymentStep.svelte";
	import ReviewStep from "$lib/components/ReviewStep.svelte";
	import Button from "$lib/components/ui/Button.svelte";
	import { SaveAll } from "lucide-svelte";

	let currentStep = $state(0);
	let isSaving = $state(false);

	const steps = [
		{ id: "basics", label: "Basics", description: "Name & Category" },
		{ id: "story", label: "Story", description: "Pitch & Media" },
		{ id: "payment", label: "Payment", description: "Bank & Verification" },
		{ id: "review", label: "Review", description: "Final Check" }
	];

	let progressPercentage = $derived((currentStep / (steps.length - 1)) * 100);
	let currentStepInfo = $derived(steps[currentStep]);

	function nextStep() {
		if (currentStep < steps.length - 1) {
			currentStep += 1;
			window.scrollTo({ top: 0, behavior: "smooth" });
		}
	}

	function prevStep() {
		if (currentStep > 0) {
			currentStep -= 1;
			window.scrollTo({ top: 0, behavior: "smooth" });
		}
	}

	function handleSaveExit() {
		isSaving = true;
		setTimeout(() => {
			isSaving = false;
			alert("Draft saved! Redirecting to dashboard...");
		}, 1000);
	}

	// --- Assets ---
	const icons = {
		check: `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"/></svg>`,
		chevronRight: `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m9 18 6-6-6-6"/></svg>`,
		chevronLeft: `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m15 18-6-6 6-6"/></svg>`,
		x: `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>`,
		logo: `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 2v8"/><path d="m4.93 10.93 1.41 1.41"/><path d="M2 18h2"/><path d="M20 18h2"/><path d="m19.07 10.93-1.41 1.41"/><path d="M22 22H2"/><path d="m8 22 4-10 4 10"/></svg>`
	};
</script>

<div class="wizard-layout">
	<header class="wizard-header">
		<div class="header-content container">
			<Button primary={false} handler={handleSaveExit}>
				<SaveAll size={18} />
				<span class="save_text">Save & Exit</span>
			</Button>

			<div class="stepper-container">
				<div class="mobile-stepper">
					<span class="step-count">Step {currentStep + 1} of {steps.length}</span>
					<p class="step-label-mobile">{currentStepInfo.label}</p>
				</div>

				<div class="desktop-stepper">
					<div class="progress-track"></div>
					<div class="progress-fill" style="width: {progressPercentage}%"></div>

					{#each steps as step, index}
						{@const isActive = index === currentStep}
						{@const isCompleted = index < currentStep}

						<div class="step-node">
							<div class="step-circle {isActive ? 'active' : ''} {isCompleted ? 'completed' : ''}">
								{#if isCompleted}
									{@html icons.check}
								{:else}
									<span>{index + 1}</span>
								{/if}
							</div>

							<div class="step-label" class:visible={isActive}>
								{step.label}
							</div>
						</div>
					{/each}
				</div>
			</div>
		</div>
	</header>

	<!-- 2. MAIN CONTENT AREA -->
	<main class="wizard-main">
		{#key currentStep}
			<div in:fly={{ y: 20, duration: 400, delay: 100, easing: cubicOut }} out:fade={{ duration: 200 }} class="step-content">
				<div class="step-header">
					<h1>{currentStepInfo.label}</h1>
					<p>{currentStepInfo.description}</p>
				</div>

				<!-- Placeholder Content Box -->
				<div class="form-placeholder">
					<!-- {#if currentStep === 0}
						<BasicSteps />
					{:else if currentStep === 1}
						<StoryStep />
					{:else if currentStep === 2}
						<RewardsStep />
					{:else if currentStep === 3}
						<PaymentStep />
					{:else if currentStep === 4}
						<ReviewStep />
					{/if} -->
					<!-- <div class="placeholder-icon">{currentStep + 1}</div>
					<p class="placeholder-title">Form content for <strong>{currentStepInfo.label}</strong> goes here.</p>
					<p class="placeholder-subtitle">Includes input fields, file uploaders, and validation logic specific to this stage.</p> -->
				</div>
			</div>
		{/key}
	</main>

	<!-- 3. FOOTER NAVIGATION -->
	<footer class="wizard-footer">
		<div class="footer-content">
			<button onclick={prevStep} disabled={currentStep === 0} class="btn-secondary">
				{@html icons.chevronLeft}
				<span>Back</span>
			</button>

			<div class="footer-actions">
				<span class="save-status">
					{isSaving ? "Saving..." : "Changes saved automatically"}
				</span>

				<button onclick={nextStep} class="btn-primary">
					{#if currentStep === steps.length - 1}
						<span>Submit for Review</span>
						{@html icons.check}
					{:else}
						<span>Next Step</span>
						{@html icons.chevronRight}
					{/if}
				</button>
			</div>
		</div>
	</footer>
</div>

<style>
	.wizard-layout {
		--primary: #34a273;
		--primary-hover: #2d8f65;
		--text-dark: #1f2937;
		--text-gray: #6b7280;
		--text-light: #9ca3af;
		--bg-light: #f9fafb;
		--border: #e5e7eb;
		--white: #ffffff;

		min-height: 100vh;
		display: flex;
		flex-direction: column;
		background-color: var(--primary-bg);
	}

	.wizard-header {
		background-color: var(--text-white);
		box-shadow: var(--shadow-sm);
	}

	.header-content {
		display: flex;
		align-items: center;
		flex-direction: column;
		justify-content: space-between;
		gap: var(--spacing-sm);
		padding: var(--spacing-sm);
	}

	.save_text {
		margin-left: var(--spacing-xs);
	}

	.stepper-container {
		flex-grow: 1;
		max-width: 42rem;
		padding-block: var(--spacing-xs);
	}

	.mobile-stepper {
		text-align: center;
	}

	.step-count {
		font-size: 0.75rem;
		font-weight: 700;
		text-transform: uppercase;
		letter-spacing: 0.05em;
		color: var(--text-light);
	}

	.step-label-mobile {
		font-size: 0.875rem;
		font-weight: 700;
		color: var(--text-dark);
		margin: 0;
	}

	.desktop-stepper {
		display: none;
		position: relative;
		align-items: center;
		justify-content: space-between;
		margin-inline: var(--spacing-sm);
	}

	.progress-track,
	.progress-fill {
		position: absolute;
		left: 0;
		top: 50%;
		transform: translateY(-50%);
		height: 4px;
		border-radius: 99px;
		z-index: 0;
	}

	.progress-track {
		width: 100%;
		background-color: var(--border);
	}

	.progress-fill {
		background-color: var(--primary);
		transition: width 0.5s cubic-bezier(0, 0, 0.2, 1);
	}

	.step-node {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 0.5rem;
		position: relative;
		z-index: 1;
	}

	.step-circle {
		width: 2rem;
		height: 2rem;
		border-radius: 50%;
		background-color: var(--white);
		border: 2px solid #d1d5db;
		display: flex;
		align-items: center;
		justify-content: center;
		font-size: 0.75rem;
		font-weight: 700;
		color: #d1d5db;
		transition: all 0.3s;
	}

	.step-circle.active {
		border-color: var(--primary);
		color: var(--primary);
		box-shadow: 0 0 0 4px rgba(52, 162, 115, 0.2);
		transform: scale(1.1);
	}

	.step-circle.completed {
		background-color: var(--primary);
		border-color: var(--primary);
		color: var(--white);
	}

	.step-label {
		position: absolute;
		top: 2.5rem;
		width: 8rem;
		text-align: center;
		font-size: 0.75rem;
		font-weight: 700;
		text-transform: uppercase;
		opacity: 0.5;
		transition: opacity 0.3s;
		pointer-events: none;
	}

	.step-label.visible {
		opacity: 1;
	}

	/* --- Main Content Styling --- */
	.wizard-main {
		flex: 1;
		width: 100%;
		max-width: 48rem;
		margin: 0 auto;
		padding: 3rem 1rem;
		display: flex;
		flex-direction: column;
		position: relative;
	}

	.step-content {
		flex: 1;
	}

	.step-header {
		text-align: center;
		margin-bottom: 2.5rem;
	}

	.step-header h1 {
		font-size: 2.25rem;
		font-weight: 800;
		color: var(--text-dark);
		margin: 0 0 0.75rem 0;
	}

	.step-header p {
		font-size: 1.125rem;
		color: var(--text-gray);
		margin: 0;
	}

	/* Form Placeholder */
	.form-placeholder {
		background-color: var(--white);
		border-radius: 1rem;
		border: 2px dashed var(--border);
		padding: 2.5rem;
		min-height: 400px;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		color: var(--text-light);
		box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
	}

	/* --- Footer Styling --- */
	.wizard-footer {
		background-color: rgba(255, 255, 255, 0.9);
		backdrop-filter: blur(12px);
		-webkit-backdrop-filter: blur(12px);
		border-top: 1px solid var(--border);
		position: sticky;
		bottom: 0;
		z-index: 40;
		padding: 1rem;
	}

	.footer-content {
		width: 100%;
		max-width: 1152px;
		margin: 0 auto;
		display: flex;
		align-items: center;
		justify-content: space-between;
	}

	@media (min-width: 640px) {
		.footer-content {
			padding: 0 2rem;
		}
	}

	/* Buttons */
	.btn-secondary {
		display: flex;
		align-items: center;
		gap: 0.5rem;
		padding: 0.75rem 1.5rem;
		border-radius: 0.5rem;
		font-weight: 600;
		border: none;
		background: none;
		color: var(--text-gray);
		cursor: pointer;
		transition: all 0.2s;
	}

	.btn-secondary:hover:not(:disabled) {
		background-color: #f3f4f6;
		color: var(--text-dark);
	}

	.btn-secondary:disabled {
		color: #d1d5db;
		cursor: not-allowed;
	}

	.footer-actions {
		display: flex;
		align-items: center;
		gap: 1rem;
	}

	.save-status {
		display: none;
		font-size: 0.75rem;
		font-weight: 500;
		color: var(--text-light);
	}

	@media (min-width: 640px) {
		.save-status {
			display: block;
		}
	}

	.btn-primary {
		display: flex;
		align-items: center;
		gap: 0.5rem;
		padding: 0.75rem 2rem;
		border-radius: 0.5rem;
		font-weight: 700;
		border: none;
		background-color: var(--primary);
		color: var(--white);
		cursor: pointer;
		transition: all 0.2s;
		box-shadow: 0 4px 6px -1px rgba(52, 162, 115, 0.3);
	}

	.btn-primary:hover {
		background-color: var(--primary-hover);
		box-shadow: 0 4px 6px -1px rgba(52, 162, 115, 0.4);
		transform: translateY(-1px);
	}

	.btn-primary:active {
		transform: scale(0.98);
	}

	@media screen and (min-width: 40rem) {
		.header-content {
			gap: 3rem;
			flex-direction: row;
		}

		.stepper-container {
			padding-block: var(--spacing-sm) var(--spacing-lg);
		}

		.mobile-stepper {
			display: none;
		}

		.desktop-stepper {
			display: flex;
		}
	}

	
	@media screen and (min-width: 48rem) {
		.wizard-header {
			position: sticky;
			top: 0;
			z-index: 1;
		}
	}
</style>
