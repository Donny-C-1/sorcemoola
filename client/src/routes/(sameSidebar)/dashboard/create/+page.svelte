<script>
	import { enhance } from "$app/forms";
	import Button from "$lib/components/ui/Button.svelte";
	import LoadingOverlay from "$lib/components/LoadingOverlay.svelte";
	import Snackbar from "$lib/components/ui/Snackbar.svelte";

	let { form } = $props();

	const categories = ["Animals", "Business", "Community", "Competitions", "Creative", "Education", "Emergencies", "Environment", "Events", "Faith", "Family", "Funerals & Memorials", "Medical", "Monthly Bills", "Newlyweds", "Other", "Sports", "Travel", "Ukraine Relief", "Volunteer", "Wishes"];

	// Steps
	const steps = [
		{ id: 1, name: "Basics" },
		{ id: 2, name: "Category" },
		{ id: 3, name: "Story" },
		{ id: 4, name: "Goal" },
		{ id: 5, name: "Banner" }
	];

	let currentStep = $state(1);
	let isSubmitting = $state(false);
	let isSnackbarVisible = $state(false);

	let imageFiles = $state(null);
	// Form data
	let campaignName = $state("");
	let description = $state("");
	let selectedCategory = $state("");
	let story = $state("");
	let fundGoal = $state("");
	let imageFile = $state(null);
	let imagePreview = $state(null);
	let isDragging = $state(false);

	// Validation errors
	let errors = $state({
		campaignName: "",
		description: "",
		category: "",
		story: "",
		fundGoal: ""
	});

	const MAX_STORY_LENGTH = 1000;
	const MAX_DESCRIPTION_LENGTH = 150;

	function nextStep() {
		if (validateCurrentStep()) {
			if (currentStep < steps.length) {
				currentStep++;
			}
		}
	}

	function prevStep() {
		if (currentStep > 1) {
			currentStep--;
			// Clear errors when going back
			clearErrors();
		}
	}

	function clearErrors() {
		errors = {
			campaignName: "",
			description: "",
			category: "",
			story: "",
			fundGoal: ""
		};
	}

	function validateCurrentStep() {
		clearErrors();

		switch (currentStep) {
			case 1: // Basics
				let isValid = true;
				if (!campaignName.trim()) {
					errors.campaignName = "Campaign name is required";
					isValid = false;
				}
				if (!description.trim()) {
					errors.description = "Campaign tagline is required";
					isValid = false;
				}
				return isValid;

			case 2: // Category
				if (!selectedCategory) {
					errors.category = "Please select a category for your campaign";
					return false;
				}
				return true;

			case 3: // Story
				if (!story.trim()) {
					errors.story = "Please tell your campaign story";
					return false;
				}
				if (story.trim().length < 50) {
					errors.story = "Your story should be at least 50 characters";
					return false;
				}
				return true;

			case 4: // Goal
				if (!fundGoal || parseFloat(fundGoal) < 1000) {
					errors.fundGoal = "Please enter a funding goal of at least ₦1,000";
					return false;
				}
				return true;

			case 5: // Banner (optional)
				return true;

			default:
				return true;
		}
	}

	// Image upload handlers
	function handleFileSelect(event) {
		const file = event.target.files?.[0];
		processFile(file);
	}

	function handleDrop(event) {
		event.preventDefault();
		isDragging = false;
		const file = event.dataTransfer?.files?.[0];
		processFile(file);
	}

	function processFile(file) {
		if (file && file.type.startsWith("image/")) {
			const dataTransfer = new DataTransfer();
			dataTransfer.items.add(file);
			imageFiles = dataTransfer.files;

			// imageFile = file;
			const reader = new FileReader();
			reader.onload = (e) => {
				imagePreview = e.target?.result;
			};
			reader.readAsDataURL(file);
		}
	}

	function removeImage() {
		imageFile = null;
		imagePreview = null;
	}

	// Calculate progress
	let progressPercentage = $derived((currentStep / steps.length) * 100);
</script>

<main>
	<LoadingOverlay show={isSubmitting} text="Creating your campaign..." />

	<div class="wizard-container">
		<!-- Progress Bar -->
		<div class="progress-section">
			<div class="step-indicators">
				{#each steps as step}
					<div class="step-indicator" class:active={currentStep === step.id} class:completed={currentStep > step.id}>
						<div class="step-circle">
							{#if currentStep > step.id}
								<svg width="14" height="14" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
									<path d="M20 6L9 17L4 12" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round" />
								</svg>
							{:else}
								{step.id}
							{/if}
						</div>
						<span class="step-name">{step.name}</span>
					</div>
					<div class="step-line" class:active={currentStep > step.id}></div>
				{/each}
			</div>
		</div>

		<!-- Form Container -->
		<div class="form-container">
			{#if form?.error}
				<div class="error-alert">
					<svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
						<path d="M10 0C4.48 0 0 4.48 0 10C0 15.52 4.48 20 10 20C15.52 20 20 15.52 20 10C20 4.48 15.52 0 10 0ZM11 15H9V13H11V15ZM11 11H9V5H11V11Z" fill="currentColor" />
					</svg>
					<div>
						<strong>{form.error.label}</strong>
						<p>{form.error.message}</p>
					</div>
				</div>
			{/if}

			<form
				action="?/createCampaign"
				method="post"
				enctype="multipart/form-data"
				use:enhance={({ cancel }) => {
					// When submitting (on step 5), validate all previous steps too
					if (currentStep === 5) {
						// Check all steps are valid
						const allValid = campaignName.trim() && description.trim() && selectedCategory && story.trim() && story.trim().length >= 50 && fundGoal && parseFloat(fundGoal) >= 1000;

						if (!allValid) {
							cancel();
							alert("Please complete all required fields");
							return;
						}
					}

					isSubmitting = true;
					return async ({ update }) => {
						isSubmitting = false;
						await update();
					};
				}}
			>
				<!-- Hidden inputs to carry data through all steps -->
				<input type="hidden" name="campaign_name" value={campaignName} />
				<input type="hidden" name="description" value={description} />
				<input type="hidden" name="category" value={selectedCategory} />
				<input type="hidden" name="campaign_story" value={story} />
				<input type="hidden" name="fund_goal" value={fundGoal} />

				<!-- Step 1: Campaign Basics -->
				{#if currentStep === 1}
					<div class="step-content">
						<div class="step-header">
							<h1>Let's start with the basics</h1>
							<p class="step-description">Give your campaign a memorable name and tagline</p>
						</div>

						<div class="form-group">
							<label for="campaign_name">Campaign Name *</label>
							<input type="text" id="campaign_name" bind:value={campaignName} placeholder="e.g., Help Sarah Fight Cancer" class:error={errors.campaignName} />
							{#if errors.campaignName}
								<span class="error-message">{errors.campaignName}</span>
							{/if}
						</div>

						<div class="form-group">
							<label for="description">Campaign Tagline *</label>
							<input type="text" id="description" bind:value={description} placeholder="A brief, compelling summary" maxlength={MAX_DESCRIPTION_LENGTH} class:error={errors.description} />
							<div class="field-footer">
								{#if errors.description}
									<span class="error-message">{errors.description}</span>
								{:else}
									<span class="helper-text">This will be the first thing people see</span>
								{/if}
								<span class="char-counter" class:warning={description.length > MAX_DESCRIPTION_LENGTH * 0.9}>
									{description.length}/{MAX_DESCRIPTION_LENGTH}
								</span>
							</div>
						</div>
					</div>
				{/if}

				<!-- Step 2: Category Selection -->
				{#if currentStep === 2}
					<div class="step-content">
						<div class="step-header">
							<h1>What best describes your campaign?</h1>
							<p class="step-description">Choose the category that fits your fundraising goal</p>
						</div>

						<div class="categories-grid">
							{#each categories as category}
								<button
									type="button"
									class="category-pill"
									class:selected={selectedCategory === category}
									onclick={() => {
										selectedCategory = category;
										errors.category = "";
									}}
								>
									{category}
								</button>
							{/each}
						</div>

						{#if errors.category}
							<p class="error-message centered">{errors.category}</p>
						{/if}
					</div>
				{/if}

				<!-- Step 3: Campaign Story -->
				{#if currentStep === 3}
					<div class="step-content">
						<div class="step-header">
							<h1>Tell your story</h1>
							<p class="step-description">Share why this campaign matters and how it will help</p>
						</div>

						<div class="form-group">
							<label for="campaign_story">Your Story *</label>
							<textarea
								id="campaign_story"
								bind:value={story}
								rows="8"
								placeholder="Share your story...

• Why does this campaign matter?
• Who will it help?
• How will the funds be used?
• What impact will it make?"
								maxlength={MAX_STORY_LENGTH}
								class:error={errors.story}
							></textarea>
							<div class="field-footer">
								{#if errors.story}
									<span class="error-message">{errors.story}</span>
								{:else}
									<span class="helper-text">Be authentic and specific</span>
								{/if}
								<span class="char-counter" class:warning={story.length > MAX_STORY_LENGTH * 0.9}>
									{story.length}/{MAX_STORY_LENGTH}
								</span>
							</div>
						</div>
					</div>
				{/if}

				<!-- Step 4: Funding Goal -->
				{#if currentStep === 4}
					<div class="step-content">
						<div class="step-header">
							<h1>Set your funding goal</h1>
							<p class="step-description">How much do you need to raise?</p>
						</div>

						<div class="form-group">
							<label for="fund_goal">Funding Goal *</label>
							<div class="input-wrapper">
								<span class="currency-symbol">₦</span>
								<input type="number" id="fund_goal" bind:value={fundGoal} placeholder="50000" min="1000" step="100" class:error={errors.fundGoal} />
							</div>
							{#if errors.fundGoal}
								<span class="error-message">{errors.fundGoal}</span>
							{:else}
								<span class="helper-text">Minimum goal: ₦1,000. Set a realistic target.</span>
							{/if}
						</div>

						<div class="tips-box">
							<p class="tips-title">💡 Tips for setting your goal:</p>
							<ul>
								<li>Include all necessary costs</li>
								<li>Add a buffer for transaction fees</li>
								<li>Be transparent about fund usage</li>
							</ul>
						</div>
					</div>
				{/if}

				<!-- Step 5: Banner Image -->
				{#if currentStep === 5}
					<div class="step-content">
						<div class="step-header">
							<h1>Add a banner image</h1>
							<p class="step-description">A great image can significantly increase support (Optional)</p>
						</div>

						<div
							class="image-upload"
							class:has-image={imagePreview}
							class:dragging={isDragging}
							ondragover={(e) => {
								e.preventDefault();
								isDragging = true;
							}}
							ondragleave={() => (isDragging = false)}
							ondrop={handleDrop}
							role="button"
							tabindex="0"
						>
							{#if imagePreview}
								<div class="image-preview">
									<img src={imagePreview} alt="Campaign banner preview" />
									<button type="button" class="remove-image" aria-label="remove-image" onclick={removeImage}>
										<svg width="18" height="18" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
											<path d="M18 6L6 18M6 6L18 18" stroke="currentColor" stroke-width="2" stroke-linecap="round" />
										</svg>
									</button>
								</div>
							{:else}
								<div class="upload-placeholder">
									<svg width="48" height="48" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
										<path d="M21 15V19C21 19.5304 20.7893 20.0391 20.4142 20.4142C20.0391 20.7893 19.5304 21 19 21H5C4.46957 21 3.96086 20.7893 3.58579 20.4142C3.21071 20.0391 3 19.5304 3 19V15" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
										<path d="M17 8L12 3L7 8" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
										<path d="M12 3V15" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
									</svg>
									<p class="upload-text">Drag and drop, or <label for="image-input" class="browse-link">browse</label></p>
									<p class="upload-hint">PNG, JPG up to 10MB</p>
								</div>
							{/if}
							<input type="file" name="campaign_banner" bind:files={imageFiles} id="image-input" accept="image/*" onchange={handleFileSelect} hidden />
						</div>

						<div class="tips-box">
							<p class="tips-title">📸 Image tips:</p>
							<ul>
								<li>Use high-quality, clear photos</li>
								<li>Show faces if possible - they create connection</li>
								<li>Landscape orientation works best</li>
							</ul>
						</div>
					</div>
				{/if}

				<!-- Navigation Buttons -->
				<div class="wizard-nav">
					{#if currentStep > 1}
						<button type="button" class="btn-secondary" onclick={prevStep}>
							<svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
								<path d="M19 12H5M5 12L12 19M5 12L12 5" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
							</svg>
							Back
						</button>
					{:else}
						<div></div>
					{/if}

					{#if currentStep < steps.length}
						<button type="button" class="btn-primary" onclick={nextStep}>
							Continue
							<svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
								<path d="M5 12H19M19 12L12 5M19 12L12 19" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
							</svg>
						</button>
					{:else}
						<Button type="submit">
							<svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
								<path d="M22 2L11 13" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
								<path d="M22 2L15 22L11 13L2 9L22 2Z" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
							</svg>
							Launch Campaign
						</Button>
					{/if}
				</div>
			</form>
		</div>
	</div>

	<Snackbar bind:visible={isSnackbarVisible} type={"error"}>{form?.error?.message || "Cannot connect to server"}</Snackbar>
</main>

<style>
	main {
		min-height: 100vh;
		background: linear-gradient(135deg, var(--primary-bg) 0%, #ffffff 100%);
		padding: var(--spacing-md) 0;
	}

	.wizard-container {
		max-width: 700px;
		margin: 0 auto;
		padding: 0 var(--spacing-md);
	}

	/* Progress Section */
	.progress-section {
		margin-bottom: var(--spacing-lg);
	}

	.progress-bar {
		height: 6px;
		background: #e0e0e0;
		border-radius: 100px;
		overflow: hidden;
		margin-bottom: var(--spacing-md);
	}

	.progress-fill {
		height: 100%;
		background: linear-gradient(90deg, var(--primary-color), var(--primary-light));
		transition: width 0.4s ease;
		border-radius: 100px;
	}

	.step-indicators {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: var(--spacing-xs);
	}

	.step-indicator {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 4px;
		flex: 1;
	}
	.step-circle {
		width: 32px;
		height: 32px;
		border-radius: 50%;
		background: white;
		border: 2px solid #e0e0e0;
		display: flex;
		align-items: center;
		justify-content: center;
		font-weight: 600;
		font-size: 0.875rem;
		color: var(--text-light);
		transition: all 0.3s ease;
	}

	.step-indicator.active .step-circle {
		background: var(--primary-color);
		border-color: var(--primary-color);
		color: white;
		transform: scale(1.1);
	}

	.step-indicator.completed .step-circle {
		background: var(--primary-color);
		border-color: var(--primary-color);
		color: white;
	}

	.step-name {
		font-size: 0.7rem;
		color: var(--text-light);
		font-weight: 500;
		text-align: center;
	}

	.step-indicator.active .step-name {
		color: var(--primary-color);
		font-weight: 600;
	}

	.step-indicator.completed .step-name {
		color: var(--text-medium);
	}

	.step-line {
		width: 60px;
		height: 2px;
		background: var(--primary-bg);
	}
	.step-line.active {
		background: var(--primary-color);
	}

	/* Form Container */
	.form-container {
		background: white;
		border-radius: var(--radius-lg);
		padding: var(--spacing-lg);
		box-shadow: var(--shadow-sm);
	}

	.error-alert {
		display: flex;
		align-items: flex-start;
		gap: var(--spacing-sm);
		padding: var(--spacing-sm);
		margin-bottom: var(--spacing-md);
		background-color: #fee;
		border: 2px solid var(--error-color);
		border-radius: var(--radius-md);
		color: var(--error-color);
		font-size: 0.875rem;
	}

	.error-alert svg {
		flex-shrink: 0;
		margin-top: 2px;
	}

	.error-alert strong {
		display: block;
		margin-bottom: 2px;
		font-weight: 600;
	}

	.error-alert p {
		margin: 0;
	}

	/* Step Content */
	.step-content {
		animation: fadeIn 0.3s ease;
	}

	@keyframes fadeIn {
		from {
			opacity: 0;
			transform: translateY(10px);
		}
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}

	.step-header {
		margin-bottom: var(--spacing-lg);
	}

	.step-header h1 {
		margin: 0 0 var(--spacing-xs) 0;
		font-size: 1.5rem;
		color: var(--text-dark);
	}

	.step-description {
		font-size: 1rem;
		color: var(--text-medium);
		margin: 0;
	}

	/* Form Groups */
	.form-group {
		margin-bottom: var(--spacing-md);
	}

	label {
		display: block;
		font-weight: 600;
		font-size: 0.9rem;
		color: var(--text-dark);
		margin-bottom: 6px;
	}

	input,
	textarea {
		width: 100%;
		padding: 10px 12px;
		font-size: 1rem;
		font-family: inherit;
		color: var(--text-dark);
		background-color: #fff;
		border: 2px solid #e0e0e0;
		border-radius: var(--radius-md);
		transition: all 0.2s ease;
	}

	input:focus,
	textarea:focus {
		outline: none;
		border-color: var(--primary-color);
		box-shadow: 0 0 0 3px rgba(52, 162, 115, 0.1);
	}

	input.error,
	textarea.error {
		border-color: var(--error-color);
	}

	input.error:focus,
	textarea.error:focus {
		box-shadow: 0 0 0 3px rgba(229, 57, 53, 0.1);
	}

	input::placeholder,
	textarea::placeholder {
		color: var(--text-light);
	}

	textarea {
		resize: vertical;
		min-height: 160px;
		line-height: 1.5;
	}

	.input-wrapper {
		position: relative;
	}

	.currency-symbol {
		position: absolute;
		left: 12px;
		top: 50%;
		transform: translateY(-50%);
		font-weight: 700;
		color: var(--text-medium);
		font-size: 1rem;
		pointer-events: none;
	}

	.input-wrapper input {
		padding-left: 2.25rem;
	}

	.field-footer {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-top: 6px;
		gap: var(--spacing-xs);
	}

	.helper-text {
		font-size: 0.8rem;
		color: var(--text-medium);
	}

	.error-message {
		font-size: 0.8rem;
		color: var(--error-color);
		font-weight: 500;
	}

	.error-message.centered {
		text-align: center;
		margin-top: var(--spacing-sm);
		display: block;
	}

	.char-counter {
		font-size: 0.8rem;
		color: var(--text-light);
		white-space: nowrap;
	}

	.char-counter.warning {
		color: var(--secondary-dark);
		font-weight: 600;
	}

	/* Category Pills */
	.categories-grid {
		display: flex;
		flex-wrap: wrap;
		gap: 8px;
	}

	.category-pill {
		padding: 8px 16px;
		border: 2px solid #e0e0e0;
		border-radius: 50px;
		background: white;
		color: var(--text-dark);
		font-size: 0.9rem;
		font-weight: 500;
		cursor: pointer;
		transition: all 0.2s ease;
		font-family: inherit;
	}

	.category-pill:hover {
		border-color: var(--primary-color);
		background: var(--primary-bg);
		transform: translateY(-1px);
	}

	.category-pill.selected {
		border-color: var(--primary-color);
		background: var(--primary-color);
		color: white;
		font-weight: 600;
	}

	/* Image Upload */
	.image-upload {
		border: 2px dashed #d0d0d0;
		border-radius: var(--radius-md);
		padding: var(--spacing-lg);
		text-align: center;
		background: #fafafa;
		transition: all 0.3s ease;
		cursor: pointer;
		min-height: 200px;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.image-upload:hover {
		border-color: var(--primary-color);
		background: var(--primary-bg);
	}

	.image-upload.dragging {
		border-color: var(--primary-color);
		background: var(--primary-bg);
		transform: scale(1.01);
	}

	.image-upload.has-image {
		padding: 0;
		border: none;
		background: transparent;
		min-height: auto;
	}

	.upload-placeholder svg {
		color: var(--primary-color);
		margin-bottom: var(--spacing-sm);
	}

	.upload-text {
		font-size: 1rem;
		color: var(--text-dark);
		margin: var(--spacing-xs) 0 4px 0;
	}

	.browse-link {
		color: var(--primary-color);
		font-weight: 600;
		cursor: pointer;
		text-decoration: underline;
	}

	.browse-link:hover {
		color: var(--primary-dark);
	}

	.upload-hint {
		font-size: 0.8rem;
		color: var(--text-light);
		margin: 0;
	}

	.image-preview {
		position: relative;
		width: 100%;
		border-radius: var(--radius-md);
		overflow: hidden;
	}

	.image-preview img {
		width: 100%;
		height: auto;
		display: block;
		max-height: 300px;
		object-fit: cover;
	}

	.remove-image {
		position: absolute;
		top: 8px;
		right: 8px;
		background: rgba(0, 0, 0, 0.7);
		color: white;
		border: none;
		padding: 6px;
		border-radius: 50%;
		cursor: pointer;
		display: flex;
		align-items: center;
		justify-content: center;
		transition: all 0.2s ease;
	}

	.remove-image:hover {
		background: var(--error-color);
		transform: scale(1.05);
	}

	/* Tips Box */
	.tips-box {
		background: var(--primary-bg);
		padding: var(--spacing-sm);
		border-radius: var(--radius-md);
		margin-top: var(--spacing-md);
	}

	.tips-title {
		margin: 0 0 6px 0;
		font-size: 0.875rem;
		font-weight: 600;
		color: var(--text-dark);
	}

	.tips-box ul {
		margin: 0;
		padding-left: var(--spacing-md);
		color: var(--text-medium);
	}

	.tips-box li {
		margin-bottom: 4px;
		font-size: 0.85rem;
		line-height: 1.4;
	}

	.tips-box li:last-child {
		margin-bottom: 0;
	}

	/* Wizard Navigation */
	.wizard-nav {
		display: flex;
		justify-content: space-between;
		align-items: center;
		gap: var(--spacing-sm);
		margin-top: var(--spacing-lg);
		padding-top: var(--spacing-md);
		border-top: 1px solid #e0e0e0;
	}

	.btn-secondary,
	.btn-primary {
		display: inline-flex;
		align-items: center;
		justify-content: center;
		gap: 6px;
		padding: 10px 20px;
		font-size: 0.95rem;
		font-weight: 600;
		border-radius: var(--radius-md);
		cursor: pointer;
		transition: all 0.2s ease;
		font-family: inherit;
		border: none;
	}

	.btn-secondary {
		background: white;
		color: var(--text-dark);
		border: 2px solid #e0e0e0;
	}

	.btn-secondary:hover {
		border-color: var(--text-medium);
		background: #f5f5f5;
	}

	.btn-primary {
		background: var(--primary-color);
		color: white;
		border: 2px solid var(--primary-color);
	}

	.btn-primary:hover {
		background: var(--primary-dark);
		border-color: var(--primary-dark);
		transform: translateY(-1px);
		box-shadow: var(--shadow-sm);
	}

	.wizard-nav :global(button[type="submit"]) {
		display: inline-flex;
		align-items: center;
		gap: 6px;
	}

	/* Responsive */
	@media (max-width: 768px) {
		main {
			padding: var(--spacing-sm) 0;
		}

		.wizard-container {
			padding: 0 var(--spacing-sm);
		}

		.form-container {
			padding: var(--spacing-md);
		}

		.step-header h1 {
			font-size: 1.25rem;
		}

		.step-description {
			font-size: 0.9rem;
		}

		.step-circle {
			width: 28px;
			height: 28px;
			font-size: 0.8rem;
		}

		.step-name {
			font-size: 0.65rem;
		}

		.category-pill {
			font-size: 0.85rem;
			padding: 6px 12px;
		}
	}
</style>
