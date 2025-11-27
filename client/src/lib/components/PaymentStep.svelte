<script>
	import { writable } from 'svelte/store';
	
	// Props: formData (bindable) is used to save the selected Payout Profile ID or new details.
	let { formData = $bindable({}) } = $props();


	const existingProfiles = [
		{ 
			id: 'prof_8a7d3s4x', 
			isVerified: true, 
			legalName: 'Acme Corp, LLC', 
			bankName: 'Fidelity Bank', 
			accountMask: '****1234', 
			currency: 'USD' 
		},
	];

	
	let selectedOption = $state(existingProfiles.length > 0 ? 'existing' : 'new');
	let selectedProfileId = $state(existingProfiles.length > 0 ? existingProfiles[0].id : null);

	// New profile state (for simulation purposes)
	let newProfile = $state({
		legalName: '',
		taxId: '',
		accountNumber: '',
		routingNumber: '',
	});
	
	// Initialize formData persistence
	if (!formData.payment) {
		formData.payment = {
			selection: selectedOption,
			profileId: selectedProfileId,
			newDetails: null
		};
	}
	
	// Effect to update formData when selection changes
	$effect(() => {
		formData.payment.selection = selectedOption;
		
		if (selectedOption === 'existing') {
			formData.payment.profileId = selectedProfileId;
			formData.payment.newDetails = null;
		} else {
			// In 'new' mode, we'd typically wait for submission, but we'll bind for realism.
			formData.payment.profileId = null;
			formData.payment.newDetails = newProfile;
		}
	});

	function handleSubmitNewProfile(e) {
        e.preventDefault();
		// In a real app: 
		// 1. Call API to securely save banking/tax info.
		// 2. Wait for verification status (usually takes minutes to hours).
		// 3. If successful, set the new profile as the selected one.
		
		if (newProfile.legalName && newProfile.taxId && newProfile.accountNumber && newProfile.routingNumber) {
			alert('New Payout Profile submitted for verification. Please continue.');
			// Simulate adding the new profile to the existing list and selecting it
			existingProfiles.push({
				id: 'prof_' + Math.random().toString(36).substring(2, 10),
				isVerified: false,
				legalName: newProfile.legalName,
				bankName: 'New Draft Bank',
				accountMask: '****' + newProfile.accountNumber.slice(-4),
				currency: 'USD (Pending KYC)',
			});
			selectedProfileId = existingProfiles[existingProfiles.length - 1].id;
			selectedOption = 'existing';
		} else {
			alert('Please fill out all required fields for the new profile.');
		}
	}

</script>

<div class="payment-container">
	<div class="header">
		<h2 class="label-lg">Select Campaign Payout</h2>
		<p class="helper-text-lg">This step ensures we know where to send the funds if your campaign is successful. We comply with all legal requirements and keep your data safe.</p>
	</div>

	<!-- --- SELECTION TABS --- -->
	<div class="selection-tabs">
		
		{#if existingProfiles.length > 0}
			<button 
				class="tab {selectedOption === 'existing' ? 'active' : ''}"
				onclick={() => selectedOption = 'existing'}
			>
				Use Existing Profile ({existingProfiles.length})
			</button>
		{/if}
		
		<button 
			class="tab {selectedOption === 'new' ? 'active' : ''}"
			onclick={() => selectedOption = 'new'}>
			Set Up New Payout
		</button>
	</div>

	<!-- --- CONTENT BASED ON SELECTION --- -->
	<div class="content-panel">
		
		<!-- 1. EXISTING PROFILE SELECTION -->
		{#if selectedOption === 'existing'}
			<div class="profile-selection">
				<p class="section-title">Verified Payout Profiles</p>
				<p class="helper-text-sm">Select the profile you wish to use for this campaign's funds.</p>

				{#each existingProfiles as profile}
					<label class="profile-card {selectedProfileId === profile.id ? 'selected' : ''}">
						<input 
							type="radio" 
							name="payoutProfile" 
							value={profile.id} 
							bind:group={selectedProfileId}
							checked={selectedProfileId === profile.id}
						/>
						<div class="profile-details">
							<div class="status-badge {profile.isVerified ? 'verified' : 'pending'}">
								<svg viewBox="0 0 24 24" width="16" height="16" stroke="currentColor" stroke-width="2" fill="none"><path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path><polyline points="22 4 12 14.01 9 11.01"></polyline></svg>
								<span>{profile.isVerified ? 'Verified' : 'Pending'}</span>
							</div>
							
							<div class="detail-row">
								<p class="label">Legal Name:</p>
								<p class="value">{profile.legalName}</p>
							</div>
							
							<div class="detail-row">
								<p class="label">Bank Account:</p>
								<p class="value">{profile.bankName} ({profile.accountMask})</p>
							</div>
							
							<div class="detail-row">
								<p class="label">Currency:</p>
								<p class="value">{profile.currency}</p>
							</div>
						</div>
					</label>
				{/each}
				
				<p class="footer-note">Need to update or add a new profile? Use the "Set Up New Payout" tab.</p>
			</div>
		{/if}

		<!-- 2. NEW PROFILE SETUP -->
		{#if selectedOption === 'new'}
			<form onsubmit={handleSubmitNewProfile}>
				<p class="section-title">Legal Entity & Tax Information</p>
				<p class="helper-text-sm">This must match the owner of the bank account and the legal entity that will report the income.</p>
				
				<!-- Legal Name -->
				<div class="form-group">
					<label for="legalName">Full Legal Name / Business Name</label>
					<input 
						id="legalName" 
						type="text" 
						bind:value={newProfile.legalName}
						placeholder="John D. Creator or Creator Labs Inc."
						required
					/>
				</div>
				
				<!-- Tax ID -->
				<div class="form-group">
					<label for="taxId">Tax ID (SSN / EIN)</label>
					<input 
						id="taxId" 
						type="text" 
						bind:value={newProfile.taxId}
						placeholder="For tax reporting purposes (Required for verification)"
						required
					/>
				</div>

				<p class="section-title mt-8">Payout Account Details</p>
				<p class="helper-text-sm">We use a secure third-party processor to facilitate all bank transfers.</p>

				<!-- Bank Details -->
				<div class="grid-cols-2">
					<div class="form-group">
						<label for="routingNumber">Routing Number</label>
						<input 
							id="routingNumber" 
							type="text" 
							bind:value={newProfile.routingNumber}
							placeholder="9 digits"
							maxlength="9"
							required
						/>
					</div>
					<div class="form-group">
						<label for="accountNumber">Account Number</label>
						<input 
							id="accountNumber" 
							type="text" 
							bind:value={newProfile.accountNumber}
							placeholder="Your bank account number"
							required
						/>
					</div>
				</div>
				
				<div class="form-actions-bottom">
					<button type="submit" class="btn-save-payout">
						<svg viewBox="0 0 24 24" width="20" height="20" stroke="currentColor" stroke-width="2" fill="none"><path d="M10 20H4a2 2 0 0 1-2-2V5c0-1.1.9-2 2-2h12l5 5v3"></path><path d="M16 12L12 16L16 20M12 16H22"></path></svg>
						Save & Verify Payout
					</button>
				</div>
			</form>
		{/if}
	</div>
</div>

<style>
	/* --- VARIABLES --- */
	.payment-container {
		--primary: #34a273;
		--primary-hover: #2d8f65;
		--primary-light: #f0fdf4;
		--text-dark: #111827;
		--text-gray: #4b5563;
		--text-light: #9ca3af;
		--border: #e5e7eb;
		--bg-white: #ffffff;
		--radius: 8px;
		--shadow-md: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -2px rgba(0, 0, 0, 0.1);
		--verified: #10b981; /* Green 500 */
		--pending: #f59e0b; /* Amber 500 */
	}

	/* --- LAYOUT & TYPOGRAPHY --- */
	.payment-container {
		width: 100%;
		max-width: 800px;
		margin: 0 auto;
		font-family: 'Inter', sans-serif;
	}

	.header {
		margin-bottom: 2rem;
	}

	.label-lg {
		font-size: 1.5rem;
		font-weight: 800;
		color: var(--text-dark);
		margin: 0 0 0.5rem 0;
	}

	.helper-text-lg {
		font-size: 0.95rem;
		color: var(--text-gray);
		margin: 0;
		line-height: 1.5;
	}
	
	.helper-text-sm {
		font-size: 0.875rem;
		color: var(--text-gray);
		margin-bottom: 1.5rem;
	}

	.section-title {
		font-size: 1.125rem;
		font-weight: 700;
		color: var(--text-dark);
		margin: 2rem 0 0.5rem 0;
		padding-bottom: 0.5rem;
		border-bottom: 1px solid var(--border);
	}
	
	.mt-8 {
		margin-top: 2rem !important;
	}

	/* --- TABS --- */
	.selection-tabs {
		display: flex;
		border-bottom: 1px solid var(--border);
		margin-bottom: 2rem;
	}

	.tab {
		background: none;
		border: none;
		padding: 1rem 1.5rem;
		font-size: 1rem;
		font-weight: 600;
		color: var(--text-gray);
		cursor: pointer;
		border-bottom: 3px solid transparent;
		transition: all 0.2s;
	}

	.tab.active {
		color: var(--primary);
		border-bottom: 3px solid var(--primary);
	}
	
	.tab:not(.active):hover {
		color: var(--text-dark);
	}

	/* --- CONTENT PANEL (shared structure) --- */
	.content-panel {
		background: var(--bg-white);
		border: 1px solid var(--border);
		border-radius: var(--radius);
		padding: 2rem;
		box-shadow: var(--shadow-md);
	}

	/* --- EXISTING PROFILE VIEW --- */
	.profile-card {
		display: flex;
		align-items: flex-start;
		gap: 1.5rem;
		padding: 1.5rem;
		border: 2px solid var(--border);
		border-radius: var(--radius);
		margin-bottom: 1rem;
		cursor: pointer;
		transition: all 0.2s;
	}

	.profile-card:hover {
		border-color: #d1d5db;
	}

	.profile-card.selected {
		border-color: var(--primary);
		background: var(--primary-light);
	}
	
	.profile-card input[type="radio"] {
		margin-top: 0.25rem;
		appearance: none;
		width: 1.25rem;
		height: 1.25rem;
		border: 2px solid var(--text-gray);
		border-radius: 50%;
		transition: border-color 0.2s, background-color 0.2s;
		flex-shrink: 0;
	}
	
	.profile-card input[type="radio"]:checked {
		border-color: var(--primary);
		background: var(--primary);
		box-shadow: inset 0 0 0 4px var(--primary-light);
	}

	.profile-details {
		flex-grow: 1;
	}
	
	.detail-row {
		display: flex;
		justify-content: space-between;
		font-size: 0.95rem;
		margin-bottom: 0.25rem;
	}

	.detail-row .label {
		font-weight: 600;
		color: var(--text-dark);
		margin: 0;
	}
	
	.detail-row .value {
		color: var(--text-gray);
		margin: 0;
	}
	
	.status-badge {
		display: inline-flex;
		align-items: center;
		gap: 0.3rem;
		padding: 0.25rem 0.75rem;
		border-radius: 9999px;
		font-size: 0.8rem;
		font-weight: 600;
		margin-bottom: 0.75rem;
	}
	
	.status-badge.verified {
		background: #d1fae5; /* Green 100 */
		color: var(--verified);
	}
	
	.status-badge.pending {
		background: #fef9c3; /* Amber 100 */
		color: var(--pending);
	}
	
	.status-badge svg {
		stroke-width: 2.5;
	}

	.footer-note {
		text-align: center;
		font-size: 0.85rem;
		color: var(--text-light);
		margin-top: 1rem;
	}
	
	/* --- NEW PROFILE FORM --- */
	.form-group {
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
		margin-bottom: 1.5rem;
	}

	.form-group label {
		font-size: 0.9rem;
		font-weight: 600;
		color: var(--text-dark);
	}

	input[type="text"], input[type="number"] {
		width: 100%;
		padding: 0.75rem 1rem;
		border: 1px solid var(--border);
		border-radius: 6px;
		font-size: 1rem;
		font-family: inherit;
		color: var(--text-dark);
		outline: none;
		transition: border 0.2s, box-shadow 0.2s;
	}

	input:focus {
		border-color: var(--primary);
		box-shadow: 0 0 0 3px rgba(52, 162, 115, 0.2);
	}
	
	.grid-cols-2 {
		display: grid;
		grid-template-columns: 1fr;
		gap: 1.5rem;
	}
	
	@media (min-width: 640px) {
		.grid-cols-2 {
			grid-template-columns: repeat(2, 1fr);
		}
	}

	/* --- FORM ACTIONS --- */
	.form-actions-bottom {
		display: flex;
		justify-content: flex-end;
		margin-top: 2rem;
		padding-top: 1rem;
		border-top: 1px solid var(--border);
	}

	.btn-save-payout {
		padding: 0.75rem 1.5rem;
		border-radius: 8px;
		font-size: 1rem;
		font-weight: 600;
		cursor: pointer;
		display: flex;
		align-items: center;
		gap: 0.5rem;
		transition: background-color 0.15s;
		border: none;
		background: var(--primary);
		color: var(--bg-white);
	}
	
	.btn-save-payout:hover {
		background: var(--primary-hover);
	}
	
</style>