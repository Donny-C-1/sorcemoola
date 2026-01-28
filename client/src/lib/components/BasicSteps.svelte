<script>
	/**
	 * Props:
	 * formData: A bindable object containing the campaign data.
	 */
	let { formData = $bindable({ image: "" }) } = $props();

	// Local state for image preview (to show immediately upon upload)
	let imagePreview = $state(formData.image || null);

	const categories = [
		'Technology', 'Art', 'Comics', 'Design', 'Film', 'Food', 'Games', 'Music', 'Publishing'
	];

	const durationOptions = [30, 60];

	// Handle File Upload to generate a preview URL
	function handleImageUpload(event) {
		const input = event.target;
		if (input.files && input.files[0]) {
			const file = input.files[0];
			
			// In a real app, you might upload this to S3 here.
			// For now, we create a local object URL for immediate preview.
			imagePreview = URL.createObjectURL(file);
			
			// Update the parent data (assuming parent expects a file object or URL)
			formData.image = imagePreview; 
		}
	}
</script>

<div class="basics-container">
	
	<!-- LEFT COLUMN: Input Form -->
	<div class="form-section">
		
		<!-- 1. Identity Group -->
		<section class="group-container">
			<div class="form-group">
				<div class="label-row">
					<label for="campaign-title">Campaign Title</label>
					<span class="char-count" aria-live="polite">
						{formData.title?.length || 0} / 60
					</span>
				</div>
				<input 
					type="text" 
					id="campaign-title" 
					bind:value={formData.title} 
					maxlength="60" 
					placeholder="The Future of Coffee"
					aria-describedby="title-help"
				/>
				<p id="title-help" class="helper-text">What is the clear, distinct name of your project?</p>
			</div>

			<div class="form-group">
				<div class="label-row">
					<label for="campaign-tagline">Short Tagline</label>
					<span class="char-count" aria-live="polite">
						{formData.tagline?.length || 0} / 135
					</span>
				</div>
				<textarea 
					id="campaign-tagline" 
					bind:value={formData.tagline} 
					maxlength="135" 
					rows="3" 
					placeholder="A smart coffee maker that learns your brewing habits."
					aria-describedby="tagline-help"
				></textarea>
				<p id="tagline-help" class="helper-text">This appears on your project card and search results.</p>
			</div>
		</section>

		<!-- 2. Classification Group -->
		<section class="group-container grid-2">
			<div class="form-group">
				<label for="category-select">Category</label>
				<div class="select-wrapper">
					<select id="category-select" bind:value={formData.category}>
						<option value="" disabled selected>Select a category...</option>
						{#each categories as cat}
							<option value={cat}>{cat}</option>
						{/each}
					</select>
					<svg class="select-icon" viewBox="0 0 24 24" width="16" height="16" stroke="currentColor" stroke-width="2" fill="none"><polyline points="6 9 12 15 18 9"></polyline></svg>
				</div>
			</div>

			<div class="form-group">
				<label for="location-input">Location</label>
				<div class="input-icon-wrapper">
					<svg class="input-icon" viewBox="0 0 24 24" width="18" height="18" stroke="currentColor" stroke-width="2" fill="none"><path d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0 1 18 0z"></path><circle cx="12" cy="10" r="3"></circle></svg>
					<input 
						type="text" 
						id="location-input" 
						bind:value={formData.location} 
						placeholder="New York, NY"
					/>
				</div>
			</div>
		</section>

		<!-- 3. Visuals Group -->
		<section class="group-container">
			<div class="form-group">
				<label for="image-upload">Card Image</label>
				<div class="file-dropzone">
					<input 
						type="file" 
						id="image-upload" 
						accept="image/*" 
						onchange={handleImageUpload}
						aria-describedby="image-help"
					/>
					<div class="dropzone-content">
						<div class="upload-icon">
							<svg viewBox="0 0 24 24" width="32" height="32" stroke="currentColor" stroke-width="2" fill="none"><rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect><circle cx="8.5" cy="8.5" r="1.5"></circle><polyline points="21 15 16 10 5 21"></polyline></svg>
						</div>
						<span class="upload-text"><strong>Upload an image</strong> or drag and drop</span>
						<span class="upload-hint">JPG, PNG, GIF up to 10MB (16:9 ratio recommended)</span>
					</div>
				</div>
			</div>
		</section>

		<!-- 4. Funding Group -->
		<section class="group-container">
			<div class="form-group">
				<label for="goal-input">Funding Goal</label>
				<div class="input-icon-wrapper">
					<span class="currency-symbol">$</span>
					<input 
						type="number" 
						id="goal-input" 
						bind:value={formData.goal} 
						placeholder="10,000"
						min="1"
					/>
				</div>
				<p class="helper-text">If you don't reach this goal, you won't receive funds.</p>
			</div>

			<fieldset class="form-group">
				<legend class="label-legend">Campaign Duration</legend>
				<div class="duration-grid">
					{#each durationOptions as days}
						<label class="radio-card {formData.duration === days ? 'selected' : ''}">
							<input type="radio" name="duration" value={days} bind:group={formData.duration} />
							<span class="radio-label">{days} Days</span>
						</label>
					{/each}
					<!-- Custom Input Mockup -->
					<label class="radio-card custom-card">
						<span class="radio-label text-muted">Custom End Date</span>
						<!-- In a full implementation, clicking this would show a date picker -->
					</label>
				</div>
			</fieldset>
		</section>

	</div>

	<!-- RIGHT COLUMN: Preview -->
	<aside class="preview-section">
		<div class="sticky-wrapper">
			<h3 class="preview-heading">Card Preview</h3>
			
			<div class="card-preview">
				<div class="card-image-area">
					{#if imagePreview}
						<img src={imagePreview} alt="Campaign Preview" />
					{:else}
						<div class="image-placeholder">
							<span>No Image Selected</span>
						</div>
					{/if}
				</div>
				
				<div class="card-content">
					<span class="card-category">{formData.category || 'Category'}</span>
					<h4 class="card-title">{formData.title || 'Your Campaign Title'}</h4>
					<p class="card-desc">{formData.tagline || 'Your catchy tagline will appear here to summarize your project...'}</p>
					
					<div class="card-footer">
						<div class="progress-bar-bg">
							<div class="progress-bar-fill" style="width: 0%"></div>
						</div>
						<div class="card-stats">
							<span class="raised"><strong>$0</strong> raised</span>
							<span class="percent">0%</span>
						</div>
					</div>
				</div>
			</div>

			<div class="tip-box">
				<span class="tip-icon">💡</span>
				<p>This card is the first thing potential backers see. Use a high-quality image!</p>
			</div>
		</div>
	</aside>

</div>

<style>
	/* --- VARIABLES --- */
	.basics-container {
		--primary: #34a273;
		--primary-hover: #2d8f65;
		--primary-light: rgba(52, 162, 115, 0.1);
		--text-dark: #111827;
		--text-gray: #4b5563;
		--text-light: #9ca3af;
		--border: #e5e7eb;
		--bg-white: #ffffff;
		--bg-page: #f9fafb;
		--focus-ring: 0 0 0 3px rgba(52, 162, 115, 0.2);
		--radius: 8px;
		--radius-lg: 12px;
	}

	/* --- LAYOUT --- */
	.basics-container {
		display: grid;
		grid-template-columns: 1fr;
		gap: 3rem;
		width: 100%;
	}

	@media (min-width: 960px) {
		.basics-container {
			grid-template-columns: 2fr 1fr;
		}
	}

	/* --- FORM STYLES --- */
	.form-section {
		display: flex;
		flex-direction: column;
		gap: 2.5rem;
	}

	.group-container {
		display: flex;
		flex-direction: column;
		gap: 1.5rem;
	}

	.group-container.grid-2 {
		display: grid;
		grid-template-columns: 1fr;
		gap: 1.5rem;
	}

	@media (min-width: 600px) {
		.group-container.grid-2 {
			grid-template-columns: 1fr 1fr;
		}
	}

	.form-group {
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
	}

	/* Labels & Text */
	label, .label-legend {
		font-weight: 600;
		font-size: 0.95rem;
		color: var(--text-dark);
		display: block;
	}

	.label-row {
		display: flex;
		justify-content: space-between;
		align-items: baseline;
	}

	.char-count {
		font-size: 0.75rem;
		color: var(--text-light);
	}

	.helper-text {
		font-size: 0.85rem;
		color: var(--text-gray);
		margin: 0;
	}

	/* Inputs */
	input[type="text"],
	input[type="number"],
	textarea,
	select {
		width: 100%;
		padding: 0.75rem 1rem;
		border: 1px solid var(--border);
		border-radius: var(--radius);
		font-size: 1rem;
		background-color: var(--bg-white);
		color: var(--text-dark);
		transition: all 0.2s ease;
		font-family: inherit;
	}

	input:focus,
	textarea:focus,
	select:focus {
		outline: none;
		border-color: var(--primary);
		box-shadow: var(--focus-ring);
	}

	input::placeholder,
	textarea::placeholder {
		color: #d1d5db;
	}

	textarea {
		resize: vertical;
		min-height: 5rem;
	}

	/* Input with Icon Wrapper */
	.input-icon-wrapper {
		position: relative;
	}

	.input-icon-wrapper input {
		padding-left: 2.5rem; /* Space for icon */
	}

	.input-icon, .currency-symbol {
		position: absolute;
		left: 0.75rem;
		top: 50%;
		transform: translateY(-50%);
		color: var(--text-light);
		pointer-events: none;
	}

	.currency-symbol {
		font-weight: 700;
		color: var(--text-gray);
	}

	/* Select Customization */
	.select-wrapper {
		position: relative;
	}

	.select-wrapper select {
		appearance: none;
		cursor: pointer;
	}

	.select-icon {
		position: absolute;
		right: 1rem;
		top: 50%;
		transform: translateY(-50%);
		color: var(--text-light);
		pointer-events: none;
	}

	/* File Upload Dropzone */
	.file-dropzone {
		position: relative;
		border: 2px dashed var(--border);
		border-radius: var(--radius-lg);
		padding: 2rem;
		text-align: center;
		background-color: var(--bg-page);
		transition: all 0.2s;
		cursor: pointer;
	}

	.file-dropzone:hover, .file-dropzone:focus-within {
		border-color: var(--primary);
		background-color: var(--primary-light);
	}

	.file-dropzone input {
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		opacity: 0;
		cursor: pointer;
	}

	.dropzone-content {
		pointer-events: none;
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 0.5rem;
	}

	.upload-icon {
		color: var(--text-light);
		margin-bottom: 0.5rem;
	}

	.file-dropzone:hover .upload-icon {
		color: var(--primary);
	}

	.upload-text {
		font-size: 0.95rem;
		color: var(--text-dark);
	}

	.upload-hint {
		font-size: 0.75rem;
		color: var(--text-light);
	}

	/* Radio Grid (Duration) */
	.duration-grid {
		display: grid;
		grid-template-columns: repeat(3, 1fr);
		gap: 1rem;
	}

	.radio-card {
		border: 1px solid var(--border);
		border-radius: var(--radius);
		padding: 1rem;
		text-align: center;
		cursor: pointer;
		position: relative;
		transition: all 0.2s;
		background: var(--bg-white);
	}

	.radio-card input {
		position: absolute;
		opacity: 0;
		width: 0;
		height: 0;
	}

	.radio-card:hover {
		border-color: var(--primary);
	}

	.radio-card.selected {
		border-color: var(--primary);
		background-color: var(--primary-light);
		color: var(--primary);
		font-weight: bold;
		box-shadow: inset 0 0 0 1px var(--primary);
	}

	.custom-card {
		border-style: dashed;
		color: var(--text-light);
	}

	/* --- PREVIEW SECTION --- */
	.preview-section {
		display: none; /* Hidden on mobile by default */
	}

	@media (min-width: 960px) {
		.preview-section {
			display: block;
		}
	}

	.sticky-wrapper {
		position: sticky;
		top: 6rem;
	}

	.preview-heading {
		font-size: 0.75rem;
		text-transform: uppercase;
		letter-spacing: 0.05em;
		color: var(--text-light);
		margin-bottom: 1rem;
		font-weight: 700;
	}

	.card-preview {
		background: var(--bg-white);
		border: 1px solid var(--border);
		border-radius: var(--radius-lg);
		overflow: hidden;
		box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05);
	}

	.card-image-area {
		width: 100%;
		height: 200px;
		background-color: #f3f4f6;
		display: flex;
		align-items: center;
		justify-content: center;
		overflow: hidden;
	}

	.card-image-area img {
		width: 100%;
		height: 100%;
		object-fit: cover;
	}

	.image-placeholder {
		color: #d1d5db;
		font-size: 0.875rem;
		font-weight: 500;
	}

	.card-content {
		padding: 1.25rem;
	}

	.card-category {
		color: var(--primary);
		font-size: 0.75rem;
		font-weight: 700;
		text-transform: uppercase;
		display: block;
		margin-bottom: 0.5rem;
	}

	.card-title {
		margin: 0 0 0.5rem 0;
		font-size: 1.125rem;
		font-weight: 700;
		line-height: 1.4;
		color: var(--text-dark);
	}

	.card-desc {
		font-size: 0.875rem;
		color: var(--text-gray);
		line-height: 1.5;
		margin: 0 0 1.25rem 0;
		display: -webkit-box;
		-webkit-line-clamp: 2;
		-webkit-box-orient: vertical;
		overflow: hidden;
	}

	.card-footer {
		border-top: 1px solid #f3f4f6;
		padding-top: 1rem;
	}

	.progress-bar-bg {
		width: 100%;
		height: 6px;
		background-color: #f3f4f6;
		border-radius: 99px;
		margin-bottom: 0.75rem;
	}

	.progress-bar-fill {
		height: 100%;
		background-color: var(--primary);
		border-radius: 99px;
	}

	.card-stats {
		display: flex;
		justify-content: space-between;
		font-size: 0.875rem;
	}

	.card-stats strong {
		color: var(--text-dark);
	}

	.card-stats .percent {
		color: var(--text-gray);
	}

	/* Tip Box */
	.tip-box {
		margin-top: 1.5rem;
		background-color: #eff6ff;
		border: 1px solid #dbeafe;
		border-radius: var(--radius);
		padding: 1rem;
		display: flex;
		gap: 0.75rem;
	}

	.tip-icon {
		font-size: 1.25rem;
	}

	.tip-box p {
		margin: 0;
		font-size: 0.875rem;
		color: #1e40af;
		line-height: 1.5;
	}
</style>