<script>
	import { currencyFormat } from "$lib/utils"; // Placeholder for a utility function

	// Props: formData (read-only in this step)
	let {
		formData = {
			basics: { title: "", category: "", duration: "" },
			funding: { goal: 0, currency: "naira", perkCount: 0 },
			payment: { selection: "", profileId: ",", newDetails: "" }
		}
	} = $props();

	let termsAccepted = $state(false);

	// Simulated Function to handle navigation back to a specific step
	function editStep(stepName) {
		console.log(`Simulating navigation back to the ${stepName} step for editing.`);
		// In a real SvelteKit app, this would use an anchor link or a routing function to change the current step state.
		alert(`Navigating back to the "${stepName}" step.`);
	}

	function launchCampaign() {
		if (termsAccepted) {
			console.log("Campaign Data:", JSON.stringify(formData, null, 2));
			alert("Congratulations! Your campaign is now launched (simulated submission).");
			// In a real app: call API to submit formData and redirect to campaign dashboard
		} else {
			alert("You must accept the terms and conditions to launch your campaign.");
		}
	}

	// --- Simulated Payout Data based on previous step's selection ---
	// In a real application, you would fetch the full profile details using the profileId.
	const getPayoutDetails = () => {
		if (formData.payment.selection === "existing") {
			return {
				status: "Verified",
				legalName: "Acme Corp, LLC",
				bankInfo: "Fidelity Bank (****1234)",
				notes: "Using existing, verified profile."
			};
		} else if (formData.payment.selection === "new" && formData.payment.newDetails) {
			const details = formData.payment.newDetails;
			return {
				status: "Pending Verification",
				legalName: details.legalName || "N/A",
				bankInfo: `Account ending in ****${details.accountNumber?.slice(-4) || "XXXX"}`,
				notes: "New account details submitted for KYC and banking verification. Funds will be held until verification is complete."
			};
		}
		return { status: "Missing", legalName: "N/A", bankInfo: "N/A", notes: "No payout profile selected." };
	};

	const payoutDetails = getPayoutDetails();
</script>

<div class="review-container">
	<header class="header">
		<h1 class="title">Review & Launch</h1>
		<p class="subtitle">This is your final check. Please confirm all details are correct before hitting "Launch Campaign."</p>
	</header>

	<div class="review-grid">
		<!-- 1. Basics Summary -->
		<section class="card summary-card">
			<div class="card-header">
				<h2 class="card-title">1. Basic Information</h2>
				<button class="edit-btn" onclick={() => editStep("Basics")}>Edit</button>
			</div>
			<div class="card-body">
				<div class="detail-item">
					<span class="label">Campaign Title:</span>
					<span class="value">{formData.basics.title}</span>
				</div>
				<div class="detail-item">
					<span class="label">Category:</span>
					<span class="value">{formData.basics.category}</span>
				</div>
				<div class="detail-item">
					<span class="label">Duration:</span>
					<span class="value">{formData.basics.duration}</span>
				</div>
			</div>
		</section>

		<!-- 2. Funding Summary -->
		<section class="card summary-card">
			<div class="card-header">
				<h2 class="card-title">2. Funding Goal & Perks</h2>
				<button class="edit-btn" onclick={() => editStep("Funding")}>Edit</button>
			</div>
			<div class="card-body">
				<div class="detail-item">
					<span class="label">Funding Goal:</span>
					<!-- Using a simulated utility for formatting -->
					<span class="value large-value">{formData.funding.currency} {new Intl.NumberFormat("en-US").format(formData.funding.goal)}</span>
				</div>
				<div class="detail-item">
					<span class="label">Perks/Rewards Offered:</span>
					<span class="value">{formData.funding.perkCount} unique tiers</span>
				</div>
				<div class="detail-item">
					<span class="label">Funding Type:</span>
					<span class="value">All-or-Nothing</span>
				</div>
			</div>
		</section>

		<!-- 3. Payment/Payout Summary -->
		<section class="card full-width summary-card payout-card">
			<div class="card-header">
				<h2 class="card-title">3. Payout Profile (Fund Disbursement)</h2>
				<button class="edit-btn" onclick={() => editStep("Payment")}>Edit</button>
			</div>
			<div class="card-body">
				<div class="payout-details">
					<div class="detail-item">
						<span class="label">Verification Status:</span>
						<span class="value status-{payoutDetails.status.toLowerCase().replace(' ', '-')}">{payoutDetails.status}</span>
					</div>
					<div class="detail-item">
						<span class="label">Legal Entity:</span>
						<span class="value">{payoutDetails.legalName}</span>
					</div>
					<div class="detail-item">
						<span class="label">Bank Account:</span>
						<span class="value">{payoutDetails.bankInfo}</span>
					</div>
				</div>
				<p class="payout-notes">
					<strong>Note:</strong>
					{payoutDetails.notes}
				</p>
			</div>
		</section>
	</div>

	<!-- Agreement and Launch -->
	<div class="agreement-box card">
		<label class="agreement-checkbox">
			<input type="checkbox" bind:checked={termsAccepted} />
			I certify that all the information provided above is true and accurate. I have read and agree to the platform's <a href="#" target="_blank">Terms & Conditions</a> and <a href="#" target="_blank">Service Fees</a>.
		</label>

		<button class="launch-btn" disabled={!termsAccepted} onclick={launchCampaign}>
			<svg viewBox="0 0 24 24" width="24" height="24" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><path d="M22 2L11 13"></path><path d="M22 2L15 22L11 13L2 9L22 2"></path></svg>
			Launch Campaign
		</button>

		{#if !termsAccepted}
			<p class="disclaimer">You must agree to the terms to launch.</p>
		{/if}
	</div>
</div>

<style>
	/* --- VARIABLES --- */
	.review-container {
		--primary-color: #34a273; /* Green */
		--primary-hover: #2d8f65;
		--text-dark: #111827;
		--text-gray: #4b5563;
		--bg-white: #ffffff;
		--border-light: #e5e7eb;
		--shadow-md: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -2px rgba(0, 0, 0, 0.05);
		--radius: 12px;
		font-family: "Inter", sans-serif;
		max-width: 1000px;
		margin: 0 auto;
		padding: 2rem 1rem;
	}

	/* --- HEADER --- */
	.header {
		margin-bottom: 3rem;
		border-bottom: 2px solid var(--border-light);
		padding-bottom: 1rem;
	}

	.title {
		font-size: 2.25rem;
		font-weight: 800;
		color: var(--text-dark);
		margin: 0;
	}

	.subtitle {
		font-size: 1.125rem;
		color: var(--text-gray);
		margin: 0.5rem 0 0 0;
	}

	/* --- CARD & GRID LAYOUT --- */
	.review-grid {
		display: grid;
		grid-template-columns: 1fr;
		gap: 2rem;
		margin-bottom: 3rem;
	}

	@media (min-width: 768px) {
		.review-grid {
			grid-template-columns: repeat(2, 1fr);
		}
		.full-width {
			grid-column: span 2;
		}
	}

	.card {
		background: var(--bg-white);
		border: 1px solid var(--border-light);
		border-radius: var(--radius);
		box-shadow: var(--shadow-md);
		transition: transform 0.2s;
	}

	.card:hover {
		transform: translateY(-2px);
	}

	/* --- SUMMARY CARD DETAILS --- */
	.summary-card {
		padding: 1.5rem;
	}

	.card-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding-bottom: 1rem;
		margin-bottom: 1rem;
		border-bottom: 1px dashed var(--border-light);
	}

	.card-title {
		font-size: 1.25rem;
		font-weight: 700;
		color: var(--text-dark);
		margin: 0;
	}

	.edit-btn {
		background: none;
		border: 1px solid var(--border-light);
		color: var(--text-gray);
		padding: 0.4rem 0.8rem;
		border-radius: 6px;
		cursor: pointer;
		font-weight: 600;
		font-size: 0.875rem;
		transition: all 0.2s;
	}

	.edit-btn:hover {
		background-color: #f3f4f6;
		color: var(--text-dark);
	}

	.card-body {
		display: flex;
		flex-direction: column;
		gap: 1rem;
	}

	.detail-item {
		display: flex;
		justify-content: space-between;
		padding: 0.5rem 0;
		border-bottom: 1px dotted #f3f4f6;
	}

	.detail-item:last-child {
		border-bottom: none;
	}

	.label {
		font-weight: 600;
		color: var(--text-dark);
		flex-shrink: 0;
		margin-right: 1rem;
	}

	.value {
		color: var(--text-gray);
		text-align: right;
		font-weight: 500;
	}

	.large-value {
		font-size: 1.1rem;
		font-weight: 700;
		color: var(--primary-color);
	}

	/* --- PAYOUT SPECIFIC STYLING --- */
	.payout-card .card-body {
		gap: 0.5rem;
	}

	.payout-details {
		display: grid;
		grid-template-columns: repeat(3, 1fr);
		gap: 1rem;
	}

	@media (max-width: 600px) {
		.payout-details {
			grid-template-columns: 1fr;
		}
	}

	.payout-details .detail-item {
		flex-direction: column;
		align-items: flex-start;
		border: none;
		padding: 0;
	}

	.payout-details .label {
		font-size: 0.9rem;
		color: var(--text-gray);
	}

	.payout-details .value {
		font-size: 1rem;
		font-weight: 700;
		text-align: left;
		color: var(--text-dark);
	}

	.status-verified {
		color: #10b981; /* Green */
	}

	.status-pending-verification {
		color: #f59e0b; /* Amber */
	}

	.payout-notes {
		font-size: 0.9rem;
		color: var(--text-gray);
		background-color: #f9fafb;
		border-radius: 8px;
		padding: 1rem;
		margin-top: 1rem;
	}

	.payout-notes strong {
		color: var(--text-dark);
	}

	/* --- AGREEMENT AND LAUNCH --- */
	.agreement-box {
		padding: 2rem;
		display: flex;
		flex-direction: column;
		align-items: center;
		text-align: center;
		border: 2px solid var(--border-light);
	}

	.agreement-checkbox {
		display: flex;
		align-items: flex-start;
		font-size: 0.95rem;
		color: var(--text-gray);
		margin-bottom: 1.5rem;
		max-width: 600px;
		text-align: left;
		cursor: pointer;
	}

	.agreement-checkbox input[type="checkbox"] {
		margin-top: 0.2rem;
		margin-right: 1rem;
		width: 1.25rem;
		height: 1.25rem;
		accent-color: var(--primary-color);
		flex-shrink: 0;
	}

	.agreement-checkbox a {
		color: var(--primary-color);
		text-decoration: none;
		font-weight: 600;
	}

	.agreement-checkbox a:hover {
		text-decoration: underline;
	}

	.launch-btn {
		display: flex;
		align-items: center;
		gap: 0.75rem;
		padding: 1rem 2.5rem;
		font-size: 1.15rem;
		font-weight: 700;
		color: var(--bg-white);
		background-color: var(--primary-color);
		border: none;
		border-radius: var(--radius);
		cursor: pointer;
		transition:
			background-color 0.2s,
			transform 0.1s;
		box-shadow: 0 4px 10px rgba(52, 162, 115, 0.4);
	}

	.launch-btn:hover:not(:disabled) {
		background-color: var(--primary-hover);
		transform: translateY(-1px);
	}

	.launch-btn:disabled {
		background-color: #9ca3af;
		cursor: not-allowed;
		box-shadow: none;
	}

	.disclaimer {
		margin-top: 1rem;
		font-size: 0.9rem;
		color: #ef4444; /* Red */
		font-weight: 600;
	}
</style>
