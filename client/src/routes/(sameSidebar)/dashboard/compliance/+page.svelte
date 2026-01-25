<script>
	// Svelte 5 Runes
	let currentStep = $state(1);
	let entityType = $state("ngo"); // 'individual' or 'ngo'
	let kycStatus = $state("unverified"); // 'unverified', 'pending', 'verified'

	// Form State
	let formData = $state({
		fullName: "",
		registrationNumber: "",
		address: "",
		documentType: "passport",
		files: {
			idFront: null,
			idBack: null,
			registrationDoc: null
		}
	});

	const nextStep = () => (currentStep += 1);
	const prevStep = () => (currentStep -= 1);

	const handleFileChange = (e, key) => {
		formData.files[key] = e.target.files[0];
	};

	const submitKYC = () => {
		// Here you would send the Multipart form to your Go backend
		kycStatus = "pending";
	};
</script>

<div class="kyc-page">
	<header class="page-header">
		<h1>KYC & Compliance</h1>
		<p>Verify your identity to enable payouts and increase donor trust.</p>
	</header>

	{#if kycStatus === "pending"}
		<div class="card status-card pending">
			<div class="icon">⏳</div>
			<h2>Verification in Progress</h2>
			<p>Our compliance team is reviewing your documents. This usually takes 24-48 hours.</p>
			<button class="btn-outline" onclick={() => (kycStatus = "unverified")}>Update Submission</button>
		</div>
	{:else if kycStatus === "verified"}
		<div class="card status-card verified">
			<div class="icon">✅</div>
			<h2>Account Verified</h2>
			<p>Your identity has been confirmed. You now have full access to all SorceMoola features.</p>
		</div>
	{:else}
		<div class="stepper">
			<div class="step" class:active={currentStep >= 1}><span>1</span> Entity</div>
			<div class="step-line" class:active={currentStep >= 2}></div>
			<div class="step" class:active={currentStep >= 2}><span>2</span> Details</div>
			<div class="step-line" class:active={currentStep >= 3}></div>
			<div class="step" class:active={currentStep >= 3}><span>3</span> Documents</div>
		</div>

		<div class="card kyc-form-card">
			{#if currentStep === 1}
				<div class="step-content">
					<h3>Select Entity Type</h3>
					<p class="step-desc">Are you raising funds as an individual or a registered organization?</p>

					<div class="selection-grid">
						<button class="select-box" class:selected={entityType === "individual"} onclick={() => (entityType = "individual")}>
							<div class="icon-circle">👤</div>
							<strong>Individual</strong>
							<span>Personal aid, medical bills, or individual relief.</span>
						</button>

						<button class="select-box" class:selected={entityType === "ngo"} onclick={() => (entityType = "ngo")}>
							<div class="icon-circle">🏢</div>
							<strong>NGO / Organization</strong>
							<span>Non-profits, charities, and community groups.</span>
						</button>
					</div>
				</div>
			{:else if currentStep === 2}
				<div class="step-content">
					<h3>Basic Information</h3>
					<div class="form-group">
						<label for="name">{entityType === "ngo" ? "Organization Name" : "Full Name"}</label>
						<input type="text" id="name" bind:value={formData.fullName} placeholder="Legal name as on ID" />
					</div>

					{#if entityType === "ngo"}
						<div class="form-group">
							<label for="reg">Registration Number</label>
							<input type="text" id="reg" bind:value={formData.registrationNumber} placeholder="NGO-12345678" />
						</div>
					{/if}

					<div class="form-group">
						<label for="addr">Physical Address</label>
						<textarea id="addr" bind:value={formData.address} placeholder="Street, City, State, Country"></textarea>
					</div>
				</div>
			{:else if currentStep === 3}
				<div class="step-content">
					<h3>Upload Documents</h3>
					<p class="step-desc">Please provide clear photos of the following documents.</p>

					<div class="upload-section">
						<div class="upload-item">
							<label>Government Issued ID (Front)</label>
							<div class="file-drop">
								<input type="file" onchange={(e) => handleFileChange(e, "idFront")} />
								<span class="file-label">{formData.files.idFront ? formData.files.idFront.name : "Click to upload ID Front"}</span>
							</div>
						</div>

						{#if entityType === "ngo"}
							<div class="upload-item">
								<label>NGO Registration Certificate</label>
								<div class="file-drop">
									<input type="file" onchange={(e) => handleFileChange(e, "registrationDoc")} />
									<span class="file-label">{formData.files.registrationDoc ? formData.files.registrationDoc.name : "Click to upload Certificate"}</span>
								</div>
							</div>
						{/if}
					</div>

					<div class="compliance-note">
						<p>🛡️ Your data is encrypted and stored according to global privacy standards.</p>
					</div>
				</div>
			{/if}

			<footer class="form-footer">
				{#if currentStep > 1}
					<button class="btn-text" onclick={prevStep}>Back</button>
				{/if}

				{#if currentStep < 3}
					<button class="btn-primary" onclick={nextStep}>Continue</button>
				{:else}
					<button class="btn-primary" onclick={submitKYC}>Submit for Review</button>
				{/if}
			</footer>
		</div>
	{/if}
</div>

<style>
	.kyc-page {
		max-width: 800px;
		margin: 0 auto;
		padding: var(--spacing-lg);
	}

	.page-header {
		margin-bottom: var(--spacing-xl);
		text-align: center;
	}
	.page-header h1 {
		font-size: 2rem;
		color: var(--text-dark);
		margin-bottom: var(--spacing-xs);
	}
	.page-header p {
		color: var(--text-medium);
	}

	/* Stepper Styles */
	.stepper {
		display: flex;
		align-items: center;
		justify-content: center;
		gap: var(--spacing-sm);
		margin-bottom: var(--spacing-xl);
	}

	.step {
		display: flex;
		align-items: center;
		gap: 8px;
		color: var(--text-light);
		font-weight: 600;
		font-size: 0.9rem;
	}

	.step span {
		width: 28px;
		height: 28px;
		border-radius: 50%;
		background: var(--neutral-color);
		color: white;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.step.active {
		color: var(--accent-color);
	}
	.step.active span {
		background: var(--accent-color);
	}

	.step-line {
		width: 60px;
		height: 2px;
		background: var(--primary-bg);
	}
	.step-line.active {
		background: var(--accent-light);
	}

	/* Card & Content */
	.card {
		background: white;
		border-radius: var(--radius-lg);
		padding: var(--spacing-xl);
		box-shadow: var(--shadow-md);
		border: 1px solid rgba(0, 0, 0, 0.05);
	}

	.step-content h3 {
		font-size: 1.5rem;
		margin-bottom: var(--spacing-xs);
		color: var(--text-dark);
	}
	.step-desc {
		color: var(--text-medium);
		margin-bottom: var(--spacing-lg);
	}

	/* Entity Selection */
	.selection-grid {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: var(--spacing-md);
	}

	.select-box {
		padding: var(--spacing-lg);
		border: 2px solid var(--primary-bg);
		border-radius: var(--radius-md);
		background: white;
		cursor: pointer;
		display: flex;
		flex-direction: column;
		align-items: center;
		text-align: center;
		transition: all 0.2s;
	}

	.select-box:hover {
		border-color: var(--primary-light);
		background: var(--primary-bg);
	}
	.select-box.selected {
		border-color: var(--primary-color);
		background: var(--primary-bg);
	}

	.icon-circle {
		width: 50px;
		height: 50px;
		background: white;
		border-radius: 50%;
		display: flex;
		align-items: center;
		justify-content: center;
		font-size: 1.5rem;
		margin-bottom: var(--spacing-sm);
		box-shadow: var(--shadow-sm);
	}

	.select-box strong {
		display: block;
		margin-bottom: 4px;
		color: var(--text-dark);
	}
	.select-box span {
		font-size: 0.8rem;
		color: var(--text-medium);
		line-height: 1.4;
	}

	/* Form Inputs */
	.form-group {
		margin-bottom: var(--spacing-md);
		display: flex;
		flex-direction: column;
		gap: 8px;
	}
	.form-group label {
		font-weight: 600;
		font-size: 0.9rem;
		color: var(--text-dark);
	}
	input,
	textarea {
		padding: var(--spacing-smr);
		border: 1px solid var(--neutral-color);
		border-radius: var(--radius-sm);
		font-family: inherit;
	}

	/* Uploads */
	.file-drop {
		position: relative;
		border: 2px dashed var(--neutral-color);
		padding: var(--spacing-lg);
		text-align: center;
		border-radius: var(--radius-md);
		cursor: pointer;
		background: var(--primary-bg);
	}

	.file-drop input {
		position: absolute;
		inset: 0;
		opacity: 0;
		cursor: pointer;
	}

	.file-label {
		color: var(--primary-dark);
		font-weight: 500;
	}

	.compliance-note {
		margin-top: var(--spacing-lg);
		padding: var(--spacing-sm);
		background: #f8fafc;
		border-radius: var(--radius-sm);
		font-size: 0.85rem;
		color: var(--text-medium);
	}

	/* Status Cards */
	.status-card {
		text-align: center;
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: var(--spacing-sm);
	}
	.status-card .icon {
		font-size: 4rem;
		margin-bottom: var(--spacing-sm);
	}
	.status-card.verified h2 {
		color: var(--primary-color);
	}

	.form-footer {
		display: flex;
		justify-content: flex-end;
		gap: var(--spacing-md);
		margin-top: var(--spacing-xl);
		padding-top: var(--spacing-md);
		border-top: 1px solid var(--primary-bg);
	}

	.btn-primary {
		background: var(--primary-color);
		color: white;
		border: none;
		padding: var(--spacing-smr) var(--spacing-xl);
		border-radius: var(--radius-sm);
		font-weight: 600;
		cursor: pointer;
	}

	.btn-text {
		background: none;
		border: none;
		color: var(--text-medium);
		cursor: pointer;
		font-weight: 600;
	}
	.btn-outline {
		border: 1px solid var(--neutral-color);
		background: none;
		padding: 10px 20px;
		cursor: pointer;
		border-radius: 4px;
	}

	@media (max-width: 600px) {
		.selection-grid {
			grid-template-columns: 1fr;
		}
	}
</style>
