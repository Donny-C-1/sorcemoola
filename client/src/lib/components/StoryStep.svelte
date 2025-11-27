<script lang="ts">
	/**
	 * Props:
	 * formData: A bindable object containing the campaign data.
	 */
	let { formData = $bindable({ video: "" }) } = $props();

	// Local state for UI feedback
	let videoFile = $state<File | null>(null);
	let isDragging = $state(false);

	// Handle Video Upload
	function handleVideoSelect(event: Event) {
		const input = event.target as HTMLInputElement;
		if (input.files && input.files[0]) {
			videoFile = input.files[0];
			// In a real app, upload to server here
			formData.video = videoFile.name; // Simulating data update
		}
	}

	// Drag and Drop Handlers
	function onDragOver(e: DragEvent) {
		e.preventDefault();
		isDragging = true;
	}
	
	function onDragLeave() {
		isDragging = false;
	}
	
	function onDrop(e: DragEvent) {
		e.preventDefault();
		isDragging = false;
		if (e.dataTransfer?.files && e.dataTransfer.files[0]) {
			const file = e.dataTransfer.files[0];
			if (file.type.startsWith('video/')) {
				videoFile = file;
				formData.video = file.name;
			}
		}
	}

	function removeVideo() {
		videoFile = null;
		formData.video = null;
	}

	// Mock Toolbar actions (visual only for this demo)
	function execCmd(cmd: string) {
		// In a real implementation, this would interact with a library like Tiptap
		console.log(`Execute command: ${cmd}`);
	}
</script>

<div class="story-container">

	<!-- LEFT COLUMN: Content Form -->
	<div class="form-section">
		
		<!-- 1. Video Section -->
		<section class="form-group">
			<div class="label-row">
				<label for="video-upload" class="label-lg">Main Campaign Video</label>
				<span class="badge-recommended">Highly Recommended</span>
			</div>
			<p class="helper-text-lg">Projects with a video have a much higher success rate. Keep it under 3 minutes.</p>

			<div 
				class="video-dropzone {isDragging ? 'dragging' : ''} {videoFile ? 'has-file' : ''}"
				role="button"
				tabindex="0"
				ondragover={onDragOver}
				ondragleave={onDragLeave}
				ondrop={onDrop}
				aria-label="Upload video area"
			>
				{#if !videoFile}
					<input 
						type="file" 
						id="video-upload" 
						accept="video/mp4,video/webm,video/mov" 
						onchange={handleVideoSelect}
						class="hidden-input"
					/>
					<div class="dropzone-empty">
						<div class="icon-circle">
							<svg viewBox="0 0 24 24" width="32" height="32" stroke="currentColor" stroke-width="2" fill="none"><polygon points="23 7 16 12 23 17 23 7"></polygon><rect x="1" y="5" width="15" height="14" rx="2" ry="2"></rect></svg>
						</div>
						<span class="upload-title"><strong>Click to upload</strong> or drag and drop</span>
						<span class="upload-meta">MP4, MOV, WEBM (Max 50MB)</span>
					</div>
				{:else}
					<div class="dropzone-filled">
						<div class="file-info">
							<div class="file-icon">
								<svg viewBox="0 0 24 24" width="24" height="24" stroke="currentColor" stroke-width="2" fill="none"><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path><polyline points="17 8 12 3 7 8"></polyline><line x1="12" y1="3" x2="12" y2="15"></line></svg>
							</div>
							<div class="file-details">
								<span class="filename">{videoFile.name}</span>
								<span class="filesize">{(videoFile.size / (1024 * 1024)).toFixed(2)} MB</span>
							</div>
						</div>
						<button class="btn-remove" onclick={removeVideo} aria-label="Remove video">
							Remove
						</button>
					</div>
					<!-- Fake progress bar -->
					<div class="upload-progress">
						<div class="progress-bar" style="width: 100%"></div>
					</div>
				{/if}
			</div>
		</section>

		<!-- 2. Story / Description Section -->
		<section class="form-group">
			<label for="story-editor" class="label-lg">Campaign Description</label>
			<p class="helper-text-lg">Tell your full story. Use images to show, not just tell.</p>

			<div class="rich-text-wrapper">
				<!-- Toolbar -->
				<div class="rte-toolbar" role="toolbar" aria-label="Text formatting">
					<div class="btn-group">
						<button type="button" class="toolbar-btn font-bold" onclick={() => execCmd('bold')} aria-label="Bold">B</button>
						<button type="button" class="toolbar-btn italic" onclick={() => execCmd('italic')} aria-label="Italic">I</button>
						<button type="button" class="toolbar-btn underline" onclick={() => execCmd('underline')} aria-label="Underline">U</button>
					</div>
					<div class="divider"></div>
					<div class="btn-group">
						<button type="button" class="toolbar-btn" onclick={() => execCmd('h1')}>H1</button>
						<button type="button" class="toolbar-btn" onclick={() => execCmd('h2')}>H2</button>
					</div>
					<div class="divider"></div>
					<div class="btn-group">
						<button type="button" class="toolbar-btn" onclick={() => execCmd('image')} aria-label="Insert Image">
							<svg viewBox="0 0 24 24" width="16" height="16" stroke="currentColor" stroke-width="2" fill="none"><rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect><circle cx="8.5" cy="8.5" r="1.5"></circle><polyline points="21 15 16 10 5 21"></polyline></svg>
						</button>
						<button type="button" class="toolbar-btn" onclick={() => execCmd('link')} aria-label="Insert Link">
							<svg viewBox="0 0 24 24" width="16" height="16" stroke="currentColor" stroke-width="2" fill="none"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"></path><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"></path></svg>
						</button>
					</div>
				</div>

				<!-- Editor Area -->
				<textarea 
					id="story-editor"
					bind:value={formData.story}
					placeholder="Start writing your story here... Introduce your team, explain your budget, and share your passion."
					class="rte-content"
				></textarea>
			</div>
		</section>

		<!-- 3. Risks Section -->
		<section class="form-group">
			<label for="risks-input" class="label-lg">Risks & Challenges</label>
			
			<div class="info-box">
				<div class="info-icon">
					<svg viewBox="0 0 24 24" width="20" height="20" stroke="currentColor" stroke-width="2" fill="none"><circle cx="12" cy="12" r="10"></circle><line x1="12" y1="8" x2="12" y2="12"></line><line x1="12" y1="16" x2="12.01" y2="16"></line></svg>
				</div>
				<div class="info-content">
					<strong>Transparency builds trust.</strong>
					<p>Be honest about potential delays in shipping or manufacturing. Backers appreciate honesty over perfection.</p>
				</div>
			</div>

			<textarea 
				id="risks-input" 
				bind:value={formData.risks}
				rows="5"
				placeholder="e.g. As with any manufacturing project, there may be delays in shipping due to supply chain issues..."
				class="standard-textarea"
			></textarea>
		</section>

	</div>

	<!-- RIGHT COLUMN: Sidebar / Guidelines -->
	<aside class="sidebar-section">
		<div class="sticky-wrapper">
			
			<!-- Tips Card -->
			<div class="sidebar-card">
				<div class="card-header">
					<svg class="icon-primary" viewBox="0 0 24 24" width="20" height="20" stroke="currentColor" stroke-width="2" fill="none"><path d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2z"></path><path d="M22 3h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7z"></path></svg>
					<h3>Storytelling Tips</h3>
				</div>
				<ul class="tips-list">
					<li>
						<span class="tip-number">1</span>
						<span><strong>The Hook:</strong> Explain what you are making in the first 2 sentences.</span>
					</li>
					<li>
						<span class="tip-number">2</span>
						<span><strong>Visuals:</strong> Use high-quality GIFs to show your product in action.</span>
					</li>
					<li>
						<span class="tip-number">3</span>
						<span><strong>Who are you?</strong> Introduce the team. People fund people, not just ideas.</span>
					</li>
				</ul>
			</div>

			<!-- Specs Card -->
			<div class="specs-card">
				<h4>Video Specs</h4>
				<div class="spec-row">
					<span>Aspect Ratio</span>
					<span class="spec-val">16:9</span>
				</div>
				<div class="spec-row">
					<span>Resolution</span>
					<span class="spec-val">1080p Rec.</span>
				</div>
				<div class="spec-row">
					<span>Format</span>
					<span class="spec-val">MP4, MOV</span>
				</div>
			</div>

		</div>
	</aside>

</div>

<style>
	/* --- VARIABLES --- */
	.story-container {
		--primary: #34a273;
		--primary-hover: #2d8f65;
		--primary-light: rgba(52, 162, 115, 0.1);
		--primary-super-light: rgba(52, 162, 115, 0.05);
		--text-dark: #111827;
		--text-gray: #4b5563;
		--text-light: #9ca3af;
		--border: #e5e7eb;
		--bg-white: #ffffff;
		--bg-page: #f9fafb;
		--radius: 8px;
		--radius-lg: 12px;
		--shadow-sm: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
	}

	/* --- LAYOUT --- */
	.story-container {
		display: grid;
		grid-template-columns: 1fr;
		gap: 3rem;
		width: 100%;
	}

	@media (min-width: 1024px) {
		.story-container {
			grid-template-columns: 2fr 1fr;
		}
	}

	.form-section {
		display: flex;
		flex-direction: column;
		gap: 3.5rem;
	}

	.form-group {
		display: flex;
		flex-direction: column;
		gap: 0.75rem;
	}

	/* Typography */
	.label-lg {
		font-size: 1.125rem;
		font-weight: 700;
		color: var(--text-dark);
		margin: 0;
	}

	.helper-text-lg {
		font-size: 0.95rem;
		color: var(--text-gray);
		margin: 0;
		line-height: 1.5;
	}

	.label-row {
		display: flex;
		align-items: center;
		justify-content: space-between;
	}

	.badge-recommended {
		background-color: var(--primary-light);
		color: var(--primary);
		font-size: 0.75rem;
		font-weight: 700;
		padding: 0.25rem 0.5rem;
		border-radius: 4px;
		text-transform: uppercase;
		letter-spacing: 0.05em;
	}

	/* --- VIDEO DROPZONE --- */
	.video-dropzone {
		position: relative;
		border: 2px dashed var(--border);
		border-radius: var(--radius-lg);
		background-color: var(--bg-white);
		min-height: 200px;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		transition: all 0.2s ease;
		cursor: pointer;
		overflow: hidden;
	}

	.video-dropzone:hover, 
	.video-dropzone.dragging {
		border-color: var(--primary);
		background-color: var(--primary-super-light);
	}

	.video-dropzone.has-file {
		border-style: solid;
		background-color: #f0fdf4; /* Green tint */
		border-color: #bbf7d0;
	}

	.hidden-input {
		position: absolute;
		inset: 0;
		opacity: 0;
		cursor: pointer;
		width: 100%;
		height: 100%;
		z-index: 10;
	}

	.dropzone-empty {
		text-align: center;
		pointer-events: none;
		display: flex;
		flex-direction: column;
		align-items: center;
		padding: 2rem;
	}

	.icon-circle {
		width: 4rem;
		height: 4rem;
		background-color: var(--bg-page);
		border-radius: 50%;
		display: flex;
		align-items: center;
		justify-content: center;
		color: var(--text-light);
		margin-bottom: 1rem;
		transition: color 0.2s;
	}

	.video-dropzone:hover .icon-circle {
		color: var(--primary);
	}

	.upload-title {
		font-size: 1rem;
		color: var(--text-dark);
		margin-bottom: 0.25rem;
	}
	
	.upload-title strong {
		color: var(--primary);
	}

	.upload-meta {
		font-size: 0.85rem;
		color: var(--text-light);
	}

	/* Filled State */
	.dropzone-filled {
		width: 100%;
		height: 100%;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		padding: 2rem;
		z-index: 20; /* Above input */
	}

	.file-info {
		display: flex;
		align-items: center;
		gap: 1rem;
		margin-bottom: 1rem;
	}

	.file-icon {
		color: var(--primary);
	}

	.file-details {
		display: flex;
		flex-direction: column;
	}

	.filename {
		font-weight: 600;
		color: var(--text-dark);
	}

	.filesize {
		font-size: 0.85rem;
		color: var(--text-gray);
	}

	.btn-remove {
		padding: 0.5rem 1rem;
		font-size: 0.875rem;
		color: #ef4444;
		background: #fef2f2;
		border: 1px solid #fee2e2;
		border-radius: var(--radius);
		cursor: pointer;
		transition: background 0.2s;
	}
	
	.btn-remove:hover {
		background: #fee2e2;
	}

	.upload-progress {
		position: absolute;
		bottom: 0;
		left: 0;
		width: 100%;
		height: 4px;
		background: #bbf7d0;
	}
	
	.progress-bar {
		height: 100%;
		background: var(--primary);
		width: 0%;
		animation: load 0.5s ease-out forwards;
	}
	
	@keyframes load {
		to { width: 100%; }
	}

	/* --- RICH TEXT EDITOR MOCK --- */
	.rich-text-wrapper {
		border: 1px solid var(--border);
		border-radius: var(--radius);
		background: var(--bg-white);
		overflow: hidden;
		box-shadow: var(--shadow-sm);
		transition: border 0.2s, box-shadow 0.2s;
	}

	.rich-text-wrapper:focus-within {
		border-color: var(--primary);
		box-shadow: 0 0 0 3px rgba(52, 162, 115, 0.2);
	}

	.rte-toolbar {
		background-color: #f9fafb;
		border-bottom: 1px solid var(--border);
		padding: 0.5rem;
		display: flex;
		align-items: center;
		gap: 0.5rem;
	}

	.btn-group {
		display: flex;
		gap: 0.25rem;
	}

	.toolbar-btn {
		width: 2rem;
		height: 2rem;
		display: flex;
		align-items: center;
		justify-content: center;
		border: none;
		background: transparent;
		border-radius: 4px;
		color: var(--text-gray);
		cursor: pointer;
		font-family: serif;
		font-size: 1rem;
	}

	.toolbar-btn:hover {
		background-color: #e5e7eb;
		color: var(--text-dark);
	}

	.font-bold { font-weight: bold; }
	.italic { font-style: italic; }
	.underline { text-decoration: underline; }

	.divider {
		width: 1px;
		height: 1.25rem;
		background-color: #d1d5db;
		margin: 0 0.25rem;
	}

	.rte-content {
		width: 100%;
		min-height: 400px;
		padding: 1.5rem;
		border: none;
		resize: vertical;
		outline: none;
		font-family: inherit;
		font-size: 1rem;
		line-height: 1.6;
		color: var(--text-dark);
	}

	/* --- INFO BOX & RISKS --- */
	.info-box {
		background-color: #eff6ff; /* Blue tint */
		border: 1px solid #dbeafe;
		border-radius: var(--radius);
		padding: 1rem;
		display: flex;
		gap: 1rem;
		margin-bottom: 0.5rem;
	}

	.info-icon {
		color: #2563eb;
		margin-top: 0.25rem;
	}

	.info-content {
		font-size: 0.875rem;
		color: #1e40af;
		line-height: 1.5;
	}

	.standard-textarea {
		width: 100%;
		padding: 1rem;
		border: 1px solid var(--border);
		border-radius: var(--radius);
		font-size: 1rem;
		font-family: inherit;
		color: var(--text-dark);
		outline: none;
		transition: border 0.2s, box-shadow 0.2s;
	}

	.standard-textarea:focus {
		border-color: var(--primary);
		box-shadow: 0 0 0 3px rgba(52, 162, 115, 0.2);
	}

	/* --- SIDEBAR --- */
	.sidebar-section {
		display: none;
	}

	@media (min-width: 1024px) {
		.sidebar-section {
			display: block;
		}
	}

	.sticky-wrapper {
		position: sticky;
		top: 6rem;
		display: flex;
		flex-direction: column;
		gap: 1.5rem;
	}

	.sidebar-card {
		background: var(--bg-white);
		border: 1px solid var(--border);
		border-radius: var(--radius);
		padding: 1.5rem;
		box-shadow: var(--shadow-sm);
	}

	.card-header {
		display: flex;
		align-items: center;
		gap: 0.5rem;
		margin-bottom: 1rem;
	}

	.icon-primary {
		color: var(--primary);
	}

	.card-header h3 {
		font-size: 1rem;
		font-weight: 700;
		margin: 0;
		color: var(--text-dark);
	}

	.tips-list {
		list-style: none;
		padding: 0;
		margin: 0;
		display: flex;
		flex-direction: column;
		gap: 1rem;
	}

	.tips-list li {
		display: flex;
		gap: 0.75rem;
		font-size: 0.9rem;
		color: var(--text-gray);
		line-height: 1.4;
	}

	.tip-number {
		font-weight: 800;
		color: #e5e7eb;
		font-size: 1.25rem;
		line-height: 1;
	}

	/* Specs Card */
	.specs-card {
		background: #f9fafb;
		border: 1px solid var(--border);
		border-radius: var(--radius);
		padding: 1.5rem;
	}

	.specs-card h4 {
		font-size: 0.75rem;
		text-transform: uppercase;
		letter-spacing: 0.05em;
		color: var(--text-gray);
		margin: 0 0 1rem 0;
		font-weight: 700;
	}

	.spec-row {
		display: flex;
		justify-content: space-between;
		font-size: 0.85rem;
		margin-bottom: 0.5rem;
		color: var(--text-light);
	}

	.spec-val {
		font-weight: 600;
		color: var(--text-dark);
	}
</style>