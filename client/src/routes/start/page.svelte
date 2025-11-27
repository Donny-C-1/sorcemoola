<script>

	let currentStep = 0;
	
	// Initial Form State
	let formData = {
		title: '',
		category: '',
		tagline: '',
		story: '',
		reason: '',
		goal: 5000,
		duration: 30,
		type: 'fixed', // 'fixed' or 'flexible'
		imageName: '' // Simulated image file name
	};

	// Steps Configuration
	const steps = [
		{ id: 0, title: 'The Basics' },
		{ id: 1, title: 'The Story' },
		{ id: 2, title: 'Funding' },
		{ id: 3, title: 'Visuals' },
		{ id: 4, title: 'Review' }
	];

	// Derived State for Progress
	$: progressPercentage = ((currentStep + 1) / steps.length) * 100;

	// Categories for Step 1
	const categories = [
		{ id: 'creative', label: 'Creative & Art', icon: '🎨' },
		{ id: 'social', label: 'Social Cause', icon: '❤️' },
		{ id: 'medical', label: 'Medical & Healing', icon: '🩺' },
		{ id: 'business', label: 'Business & Tech', icon: '🚀' },
		{ id: 'other', label: 'Other Dreams', icon: '✨' }
	];

	// Navigation Logic
	function nextStep() {
		if (!validateStep(currentStep)) {
			alert('Please fill in all required fields for this step.');
			return;
		}
		if (currentStep < steps.length - 1) {
			currentStep++;
		}
	}

	function prevStep() {
		if (currentStep > 0) {
			currentStep--;
		}
	}

	function validateStep(step) {
		switch (step) {
			case 0: return formData.title && formData.category && formData.tagline;
			case 1: return formData.story && formData.reason;
			case 2: return formData.goal > 0 && formData.duration > 0;
			case 3: return true; // Optional for demo or simple check
			default: return true;
		}
	}

	function handleLaunch() {
		alert('Campaign Launched! (Simulated)');
		// API call would go here
	}

	function selectCategory(catId) {
		formData.category = catId;
	}

	function handleFileChange(e) {
		formData.imageName = e.target.files[0]?.name || '';
	}

	// Formatting helpers
	const formatCurrency = (val) => new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(val);
</script>

<div class="page-container">
	<!-- Header / Progress -->
	<header class="wizard-header">
		<div class="logo">Sourcemoola</div>
		<div class="progress-track">
			<div class="progress-fill" style="width: {progressPercentage}%"></div>
		</div>
		<div class="steps-indicator">
			Step {currentStep + 1} of {steps.length}: <strong>{steps[currentStep].title}</strong>
		</div>
	</header>

	<main class="wizard-card">
		
		<!-- Step 1: The Basics -->
		{#if currentStep === 0}
			<div class="step-content fade-in">
				<h2>Let's start with the basics</h2>
				<p class="subtitle">Give your campaign a name and a home.</p>

				<div class="form-group">
					<label for="title">Campaign Title</label>
					<input 
						type="text" 
						id="title" 
						bind:value={formData.title} 
						placeholder="e.g., The Community Garden Project" 
					/>
				</div>

				<div class="form-group">
					<label>Choose a Category</label>
					<div class="category-grid">
						{#each categories as cat}
							<button 
								class="cat-card {formData.category === cat.id ? 'selected' : ''}"
								on:click={() => selectCategory(cat.id)}
								type="button"
							>
								<span class="cat-icon">{cat.icon}</span>
								<span class="cat-label">{cat.label}</span>
							</button>
						{/each}
					</div>
				</div>

				<div class="form-group">
					<label for="tagline">Short Tagline</label>
					<input 
						type="text" 
						id="tagline" 
						bind:value={formData.tagline} 
						placeholder="Write a catchy hook..." 
					/>
					<small>This will appear on the campaign card.</small>
				</div>
			</div>
		{/if}

		<!-- Step 2: The Story -->
		{#if currentStep === 1}
			<div class="step-content fade-in">
				<h2>Tell your story</h2>
				<p class="subtitle">Connect with backers by explaining your mission.</p>

				<div class="form-group">
					<label for="story">Campaign Description</label>
					<textarea 
						id="story" 
						bind:value={formData.story} 
						rows="6" 
						placeholder="What are you building or solving? be descriptive!"
					></textarea>
				</div>

				<div class="form-group">
					<label for="reason">Why does this matter?</label>
					<textarea 
						id="reason" 
						bind:value={formData.reason} 
						rows="3" 
						placeholder="Explain the impact of this project..."
					></textarea>
				</div>
			</div>
		{/if}

		<!-- Step 3: Funding -->
		{#if currentStep === 2}
			<div class="step-content fade-in">
				<h2>Funding & Duration</h2>
				<p class="subtitle">Set your targets realistically.</p>

				<div class="row-group">
					<div class="form-group">
						<label for="goal">Funding Goal ($)</label>
						<input 
							type="number" 
							id="goal" 
							bind:value={formData.goal} 
							min="1"
						/>
					</div>
					<div class="form-group">
						<label for="duration">Duration (Days)</label>
						<input 
							type="number" 
							id="duration" 
							bind:value={formData.duration} 
							min="1" 
							max="90"
						/>
					</div>
				</div>

				<div class="form-group">
					<label>Campaign Type</label>
					<div class="radio-group">
						<label class="radio-card {formData.type === 'fixed' ? 'selected' : ''}">
							<input type="radio" bind:group={formData.type} value="fixed" />
							<div class="radio-info">
								<strong>Fixed Goal</strong>
								<span>All-or-nothing. You only get funds if the goal is met.</span>
							</div>
						</label>
						<label class="radio-card {formData.type === 'flexible' ? 'selected' : ''}">
							<input type="radio" bind:group={formData.type} value="flexible" />
							<div class="radio-info">
								<strong>Flexible Goal</strong>
								<span>Keep what you raise, even if you miss the goal.</span>
							</div>
						</label>
					</div>
				</div>
			</div>
		{/if}

		<!-- Step 4: Visuals -->
		{#if currentStep === 3}
			<div class="step-content fade-in">
				<h2>Visuals</h2>
				<p class="subtitle">A great cover image increases funding by 50%.</p>

				<div class="dropzone">
					<div class="dropzone-content">
						<span class="icon">🖼️</span>
						<p><strong>Drop your cover image here</strong></p>
						<p class="sm">or click to browse (PNG, JPG)</p>
					</div>
					<!-- Simulated input -->
					<input 
						type="file" 
						accept="image/*" 
						on:change={handleFileChange}
					/>
				</div>

				{#if formData.imageName}
					<div class="file-preview">
						Selected: <strong>{formData.imageName}</strong>
					</div>
				{/if}
			</div>
		{/if}

		<!-- Step 5: Review -->
		{#if currentStep === 4}
			<div class="step-content fade-in">
				<h2>Ready to Launch?</h2>
				<p class="subtitle">Review your campaign details before going live.</p>

				<div class="preview-card">
					<div class="preview-header">
						<div class="preview-tag">Category: {formData.category}</div>
						<h3>{formData.title || 'Untitled Campaign'}</h3>
						<p class="preview-tagline">{formData.tagline || 'No tagline set'}</p>
					</div>
					<div class="preview-stats">
						<div class="stat">
							<span>Goal</span>
							<strong>{formatCurrency(formData.goal)}</strong>
						</div>
						<div class="stat">
							<span>Type</span>
							<strong style="text-transform: capitalize">{formData.type}</strong>
						</div>
						<div class="stat">
							<span>Duration</span>
							<strong>{formData.duration} Days</strong>
						</div>
					</div>
					<div class="preview-body">
						<h4>About</h4>
						<p>{formData.story || 'No story provided.'}</p>
					</div>
				</div>
			</div>
		{/if}

		<!-- Footer Actions -->
		<footer class="wizard-footer">
			<button 
				class="btn-secondary" 
				on:click={prevStep} 
				disabled={currentStep === 0}
			>
				Back
			</button>

			{#if currentStep === steps.length - 1}
				<button class="btn-primary launch-btn" on:click={handleLaunch}>
					Launch Campaign 🚀
				</button>
			{:else}
				<button class="btn-primary" on:click={nextStep}>
					Next Step
				</button>
			{/if}
		</footer>

	</main>
</div>

<style>
	/* --- Theme Variables --- */
	:root {
		--primary: #34a273;
		--primary-dark: #2a825c;
		--primary-light: #eafaf3;
		
		--text-main: #1a202c;
		--text-muted: #718096;
		
		--bg-body: #f4f6f8;
		--bg-card: #ffffff;
		
		--border: #e2e8f0;
		--radius: 12px;
		--shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
	}

	/* --- Layout --- */
	:global(body) {
		margin: 0;
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
		background-color: var(--bg-body);
		color: var(--text-main);
	}

	.page-container {
		display: flex;
		flex-direction: column;
		align-items: center;
		min-height: 100vh;
		padding: 2rem 1rem;
	}

	/* --- Header & Progress --- */
	.wizard-header {
		width: 100%;
		max-width: 700px;
		margin-bottom: 2rem;
		text-align: center;
	}

	.logo {
		font-size: 1.5rem;
		font-weight: 800;
		color: var(--primary);
		margin-bottom: 1.5rem;
		text-transform: uppercase;
		letter-spacing: 1px;
	}

	.progress-track {
		background: #e2e8f0;
		height: 8px;
		border-radius: 10px;
		overflow: hidden;
		margin-bottom: 1rem;
	}

	.progress-fill {
		background: var(--primary);
		height: 100%;
		transition: width 0.4s ease;
	}

	.steps-indicator {
		font-size: 0.9rem;
		color: var(--text-muted);
	}

	/* --- Main Card --- */
	.wizard-card {
		background: var(--bg-card);
		width: 100%;
		max-width: 700px;
		padding: 2.5rem;
		border-radius: var(--radius);
		box-shadow: var(--shadow);
		display: flex;
		flex-direction: column;
		min-height: 500px; /* Prevent jarring height jumps */
	}

	.step-content {
		flex: 1;
		display: flex;
		flex-direction: column;
		gap: 1.5rem;
	}

	/* --- Typography --- */
	h2 {
		margin: 0;
		font-size: 1.75rem;
		color: var(--text-main);
	}

	.subtitle {
		margin: 0 0 1rem 0;
		color: var(--text-muted);
		font-size: 1.05rem;
	}

	label {
		display: block;
		font-weight: 600;
		margin-bottom: 0.5rem;
		color: var(--text-main);
	}

	small {
		display: block;
		margin-top: 0.25rem;
		color: var(--text-muted);
	}

	/* --- Inputs --- */
	input[type="text"],
	input[type="number"],
	textarea {
		width: 100%;
		padding: 0.75rem 1rem;
		border: 2px solid var(--border);
		border-radius: 8px;
		font-size: 1rem;
		transition: border-color 0.2s;
		font-family: inherit;
		box-sizing: border-box;
	}

	input:focus,
	textarea:focus {
		outline: none;
		border-color: var(--primary);
	}

	/* --- Category Grid (Step 1) --- */
	.category-grid {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
		gap: 1rem;
	}

	.cat-card {
		background: #fff;
		border: 2px solid var(--border);
		border-radius: 8px;
		padding: 1rem;
		cursor: pointer;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		gap: 0.5rem;
		transition: all 0.2s;
	}

	.cat-card:hover {
		border-color: var(--primary);
		background-color: var(--primary-light);
	}

	.cat-card.selected {
		border-color: var(--primary);
		background-color: var(--primary-light);
		color: var(--primary-dark);
		box-shadow: 0 0 0 2px var(--primary-light);
	}

	.cat-icon { font-size: 1.5rem; }
	.cat-label { font-size: 0.9rem; font-weight: 600; text-align: center; }

	/* --- Row Groups (Step 3) --- */
	.row-group {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 1.5rem;
	}

	/* --- Radio Cards (Step 3) --- */
	.radio-group {
		display: flex;
		flex-direction: column;
		gap: 1rem;
	}

	.radio-card {
		display: flex;
		align-items: flex-start;
		gap: 1rem;
		padding: 1rem;
		border: 2px solid var(--border);
		border-radius: 8px;
		cursor: pointer;
		transition: all 0.2s;
	}

	.radio-card:hover { border-color: var(--primary); }

	.radio-card.selected {
		border-color: var(--primary);
		background-color: var(--primary-light);
	}

	.radio-card input { margin-top: 4px; }

	.radio-info strong { display: block; color: var(--text-main); }
	.radio-info span { color: var(--text-muted); font-size: 0.9rem; font-weight: normal; }

	/* --- Dropzone (Step 4) --- */
	.dropzone {
		border: 2px dashed var(--border);
		border-radius: 12px;
		padding: 3rem;
		text-align: center;
		position: relative;
		cursor: pointer;
		background: #fafbfc;
		transition: all 0.2s;
	}

	.dropzone:hover {
		border-color: var(--primary);
		background: var(--primary-light);
	}

	.dropzone input {
		position: absolute;
		top: 0; left: 0; width: 100%; height: 100%;
		opacity: 0;
		cursor: pointer;
	}

	.dropzone-content .icon { font-size: 3rem; display: block; margin-bottom: 1rem; }
	.sm { font-size: 0.85rem; color: var(--text-muted); }

	.file-preview {
		margin-top: 1rem;
		padding: 0.5rem;
		background: var(--primary-light);
		color: var(--primary-dark);
		border-radius: 6px;
		text-align: center;
	}

	/* --- Preview Card (Step 5) --- */
	.preview-card {
		background: white;
		border: 1px solid var(--border);
		border-radius: 8px;
		overflow: hidden;
	}

	.preview-header {
		background: var(--primary-light);
		padding: 1.5rem;
		border-bottom: 1px solid var(--border);
	}

	.preview-tag {
		display: inline-block;
		background: white;
		color: var(--primary);
		padding: 0.25rem 0.75rem;
		border-radius: 20px;
		font-size: 0.8rem;
		font-weight: 700;
		margin-bottom: 0.5rem;
		text-transform: uppercase;
	}

	.preview-tagline { color: var(--text-muted); margin: 0.5rem 0 0 0; font-style: italic; }
	
	.preview-stats {
		display: grid;
		grid-template-columns: 1fr 1fr 1fr;
		padding: 1rem;
		border-bottom: 1px solid var(--border);
		text-align: center;
	}

	.stat span { display: block; font-size: 0.8rem; color: var(--text-muted); margin-bottom: 0.25rem; }
	.stat strong { display: block; font-size: 1.1rem; color: var(--text-main); }

	.preview-body { padding: 1.5rem; white-space: pre-wrap; color: #4a5568; line-height: 1.6; }

	/* --- Footer Buttons --- */
	.wizard-footer {
		margin-top: 3rem;
		padding-top: 1.5rem;
		border-top: 1px solid var(--border);
		display: flex;
		justify-content: space-between;
		align-items: center;
	}

	button.btn-primary, button.btn-secondary {
		padding: 0.8rem 1.5rem;
		border-radius: 8px;
		font-weight: 600;
		cursor: pointer;
		transition: all 0.2s;
		font-size: 1rem;
	}

	.btn-secondary {
		background: transparent;
		color: var(--text-muted);
		border: none;
	}

	.btn-secondary:hover:not(:disabled) { color: var(--text-main); background: #edf2f7; }
	.btn-secondary:disabled { opacity: 0.3; cursor: not-allowed; }

	.btn-primary {
		background: var(--primary);
		color: white;
		border: none;
		box-shadow: 0 4px 6px rgba(52, 162, 115, 0.3);
	}

	.btn-primary:hover { background: var(--primary-dark); transform: translateY(-1px); }

	.launch-btn { background: #1a202c; }
	.launch-btn:hover { background: #000; }

	/* --- Animations --- */
	.fade-in {
		animation: fadeIn 0.4s ease forwards;
	}

	@keyframes fadeIn {
		from { opacity: 0; transform: translateY(10px); }
		to { opacity: 1; transform: translateY(0); }
	}

	/* Mobile tweaks */
	@media (max-width: 600px) {
		.wizard-card { padding: 1.5rem; }
		.row-group { grid-template-columns: 1fr; gap: 1rem; }
		.category-grid { grid-template-columns: 1fr 1fr; }
	}
</style>