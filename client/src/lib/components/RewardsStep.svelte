<script>
	/**
	 * Props:
	 * formData: A bindable object containing the campaign data.
	 */
	let { formData = $bindable({ rewards: "" }) } = $props();

	// --- Component State ---

	// Mock rewards data, linked to formData for persistence simulation
	if (!formData.rewards) {
		formData.rewards = [
			{ id: 1, amount: 5, title: "Digital Thank You", description: "A digital wallpaper and a warm thank you.", delivery: "Sep 2025", limit: null, shippingRequired: false, shippingCost: 0, shipsTo: "N/A" },
			{ id: 2, amount: 35, title: "The Standard Kit", description: "Get one unit of the main product at a special early-bird price.", delivery: "Oct 2025", limit: 500, shippingRequired: true, shippingCost: 10, shipsTo: "Worldwide" },
			{ id: 3, amount: 150, title: "Ultimate Creator Pack (Limited)", description: "Two units of the product plus exclusive signed artwork.", delivery: "Nov 2025", limit: 50, shippingRequired: true, shippingCost: 15, shipsTo: "US/EU Only" }
		];
	}

	let rewards = $state(formData.rewards);

	let isAdding = $state(false);
	let currentReward = $state(null);
	let nextId = $state(rewards.length > 0 ? Math.max(...rewards.map((r) => r.id)) + 1 : 1);

	// --- Actions ---

	function openAddForm() {
		currentReward = {
			id: nextId,
			amount: 10,
			title: "New Reward Tier",
			description: "",
			delivery: "Dec 2025",
			limit: null,
			shippingRequired: true,
			shippingCost: 10,
			shipsTo: "Worldwide"
		};
		isAdding = true;
	}

	function openEditForm(reward) {
		// Deep clone for editing
		currentReward = JSON.parse(JSON.stringify(reward));
		isAdding = true;
	}

	function saveReward() {
		if (!currentReward) return;

		// Validation (simple)
		if (!currentReward.title || currentReward.amount <= 0 || !currentReward.delivery) {
			console.error("Please fill in all required fields.");
			return;
		}

		const index = rewards.findIndex((r) => r.id === !currentReward.id);

		if (index !== -1) {
			// Update existing reward
			rewards[index] = currentReward;
		} else {
			// Add new reward
			rewards = [...rewards, currentReward];
			nextId++;
		}

		// Ensure the main form data is updated
		formData.rewards = rewards;

		// Close form
		isAdding = false;
		currentReward = null;
	}

	function deleteReward(id) {
		if (window.confirm("Are you sure you want to delete this reward tier?")) {
			rewards = rewards.filter((r) => r.id !== id);
			formData.rewards = rewards;
		}
	}

	function cancelEdit() {
		isAdding = false;
		currentReward = null;
	}

	// --- Helper Functions ---
	function formatDelivery(date) {
		return `Est. ${date}`;
	}
</script>

<div class="rewards-container">
	<!-- LEFT COLUMN: Rewards List & Add Button -->
	<div class="list-section">
		<div class="header">
			<h2 class="label-lg">Pledge Rewards</h2>
			<p class="helper-text-lg">Define the tiers backers can pledge toward. Be clear about what they receive and when.</p>
		</div>

		{#if rewards.length === 0}
			<div class="empty-state">
				<div class="icon-box">
					<svg viewBox="0 0 24 24" width="32" height="32" stroke="currentColor" stroke-width="2" fill="none"><path d="M12 2L2 7l10 5 10-5-10-5z"></path><path d="M2 17l10 5 10-5M2 12l10 5 10-5"></path></svg>
				</div>
				<p>No rewards defined yet. Start by creating your first pledge tier.</p>
			</div>
		{:else}
			<!-- Rewards Grid -->
			<div class="rewards-grid">
				{#each rewards as reward}
					<div class="reward-card" key={reward.id}>
						<div class="card-content">
							<span class="pledge-amount">${reward.amount}</span>
							<h3 class="reward-title">{reward.title}</h3>
							<p class="reward-description">{reward.description}</p>
						</div>

						<div class="card-meta">
							<div class="meta-item">
								<svg viewBox="0 0 24 24" width="16" height="16" stroke="currentColor" stroke-width="2" fill="none"><rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect><line x1="16" y1="2" x2="16" y2="6"></line><line x1="8" y1="2" x2="8" y2="6"></line><line x1="3" y1="10" x2="21" y2="10"></line></svg>
								<span>{formatDelivery(reward.delivery)}</span>
							</div>

							<div class="meta-item">
								<svg viewBox="0 0 24 24" width="16" height="16" stroke="currentColor" stroke-width="2" fill="none"><path d="M22 10s-5 0-7 2c-3 1-5 2-7 2-3 0-4-1-4-1v10h18V10z"></path><path d="M2 10L12 2l10 8"></path></svg>
								{#if reward.limit}
									<span>{reward.limit} available</span>
								{:else}
									<span>Unlimited</span>
								{/if}
							</div>
						</div>

						<div class="card-actions">
							<button class="btn-edit" onclick={() => openEditForm(reward)}>Edit</button>
							<button class="btn-delete" onclick={() => deleteReward(reward.id)}>Delete</button>
						</div>
					</div>
				{/each}
			</div>
		{/if}

		<button class="btn-add-reward" onclick={openAddForm} disabled={isAdding}>
			<svg viewBox="0 0 24 24" width="20" height="20" stroke="currentColor" stroke-width="2" fill="none"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
			Add New Reward Tier
		</button>
	</div>

	<!-- RIGHT COLUMN: Add/Edit Reward Form -->
	<div class="form-panel {isAdding ? 'active' : ''}">
		{#if isAdding && currentReward}
			<form
				onsubmit={(e) => {
					e.preventDefault();
					saveReward();
				}}
			>
				<h3>{currentReward.id ? "Edit Reward" : "Add New Reward"}</h3>
				<p class="helper-text-sm">Define the experience and physical goods for this tier.</p>

				<!-- Pledge Amount -->
				<div class="form-group-sm">
					<label for="amount">Pledge Amount (USD)</label>
					<div class="input-group">
						<span class="input-prefix">$</span>
						<input id="amount" type="number" min="1" bind:value={currentReward.amount} placeholder="Min $10" required />
					</div>
				</div>

				<!-- Title -->
				<div class="form-group-sm">
					<label for="title">Title</label>
					<input id="title" type="text" bind:value={currentReward.title} maxlength="40" placeholder="e.g., Early Bird Special" required />
					<span class="char-count">{currentReward.title.length}/40 characters</span>
				</div>

				<!-- Description -->
				<div class="form-group-sm">
					<label for="description">Description</label>
					<textarea id="description" rows="4" bind:value={currentReward.description} placeholder="Briefly describe what the backer receives." required></textarea>
				</div>

				<!-- Estimated Delivery -->
				<div class="form-group-sm">
					<label for="delivery">Estimated Delivery</label>
					<input id="delivery" type="text" bind:value={currentReward.delivery} placeholder="e.g., Oct 2025" required />
				</div>

				<!-- Inventory Limit -->
				<div class="form-group-sm">
					<label for="limit">Inventory Limit</label>
					<input id="limit" type="number" min="1" bind:value={currentReward.limit} placeholder="Leave blank for unlimited" />
					<div class="checkbox-group">
						<input
							type="checkbox"
							id="unlimited"
							checked={currentReward.limit === null}
							onchange={(e) => {
								currentReward.limit = e.target.checked ? null : 100;
							}}
						/>
						<label for="unlimited" class="text-sm">Unlimited inventory</label>
					</div>
				</div>

				<!-- Shipping Requirements -->
				<div class="form-group-sm">
					<label>Shipping</label>
					<div class="checkbox-group">
						<input type="checkbox" id="shipping" bind:checked={currentReward.shippingRequired} />
						<label for="shipping" class="text-sm">Physical goods require shipping</label>
					</div>

					{#if currentReward.shippingRequired}
						<div class="shipping-details">
							<div class="form-group-inline">
								<label for="shipCost">Shipping Cost (USD)</label>
								<input id="shipCost" type="number" min="0" bind:value={currentReward.shippingCost} />
							</div>
							<div class="form-group-inline">
								<label for="shipsTo">Ships To</label>
								<select id="shipsTo" bind:value={currentReward.shipsTo}>
									<option value="Worldwide">Worldwide</option>
									<option value="US Only">US Only</option>
									<option value="US/EU Only">US/EU Only</option>
									<option value="Custom">Custom Region</option>
								</select>
							</div>
						</div>
					{/if}
				</div>

				<!-- Form Actions -->
				<div class="form-actions">
					<button type="button" class="btn-cancel" onclick={cancelEdit}>Cancel</button>
					<button type="submit" class="btn-save">
						<svg viewBox="0 0 24 24" width="20" height="20" stroke="currentColor" stroke-width="2" fill="none"><polyline points="20 6 9 17 4 12"></polyline></svg>
						Save Reward
					</button>
				</div>
			</form>
		{:else}
			<div class="panel-placeholder">
				<svg viewBox="0 0 24 24" width="40" height="40" stroke="currentColor" stroke-width="1.5" fill="none"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path><polyline points="14 2 14 8 20 8"></polyline><line x1="16" y1="13" x2="8" y2="13"></line><line x1="16" y1="17" x2="8" y2="17"></line><line x1="10" y1="9" x2="8" y2="9"></line></svg>
				<p>Select an existing reward to edit or click "Add New Reward Tier" below to get started.</p>
			</div>
		{/if}
	</div>
</div>

<style>
	/* --- VARIABLES --- */
	.rewards-container {
		--primary: #34a273;
		--primary-hover: #2d8f65;
		--primary-light: #f0fdf4; /* Green 50 */
		--text-dark: #111827;
		--text-gray: #4b5563;
		--text-light: #9ca3af;
		--border: #e5e7eb;
		--bg-white: #ffffff;
		--radius: 8px;
		--shadow-md: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -2px rgba(0, 0, 0, 0.1);
	}

	/* --- LAYOUT --- */
	.rewards-container {
		display: grid;
		grid-template-columns: 1fr;
		gap: 3rem;
		width: 100%;
		font-family: "Inter", sans-serif;
	}

	@media (min-width: 1024px) {
		.rewards-container {
			grid-template-columns: 3fr 2fr;
		}
	}

	.list-section {
		display: flex;
		flex-direction: column;
		gap: 1.5rem;
	}

	.header {
		margin-bottom: 1rem;
	}

	/* Typography */
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

	/* --- REWARDS GRID --- */
	.rewards-grid {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
		gap: 1.5rem;
	}

	.reward-card {
		background: var(--bg-white);
		border: 1px solid var(--border);
		border-radius: var(--radius);
		box-shadow: var(--shadow-md);
		overflow: hidden;
		display: flex;
		flex-direction: column;
		transition: all 0.2s;
	}

	.reward-card:hover {
		border-color: var(--primary);
		transform: translateY(-2px);
		box-shadow: 0 6px 10px -2px rgba(0, 0, 0, 0.15);
	}

	.card-content {
		padding: 1.5rem;
		flex-grow: 1;
	}

	.pledge-amount {
		font-size: 2rem;
		font-weight: 800;
		color: var(--primary);
		display: block;
		margin-bottom: 0.5rem;
	}

	.reward-title {
		font-size: 1.125rem;
		font-weight: 700;
		color: var(--text-dark);
		margin: 0 0 0.5rem 0;
	}

	.reward-description {
		font-size: 0.9rem;
		color: var(--text-gray);
		line-height: 1.4;
		margin: 0;
	}

	.card-meta {
		border-top: 1px solid var(--border);
		padding: 1rem 1.5rem;
		display: flex;
		justify-content: space-between;
		font-size: 0.85rem;
		color: var(--text-gray);
	}

	.meta-item {
		display: flex;
		align-items: center;
		gap: 0.5rem;
	}

	.card-actions {
		border-top: 1px solid var(--border);
		padding: 0.75rem 1.5rem;
		display: flex;
		gap: 0.5rem;
		background: #f9fafb;
	}

	.btn-edit,
	.btn-delete {
		padding: 0.5rem 1rem;
		border-radius: 6px;
		font-size: 0.85rem;
		font-weight: 600;
		cursor: pointer;
		transition: background-color 0.15s;
		border: 1px solid transparent;
	}

	.btn-edit {
		background: #d1fae5; /* Green 100 */
		color: var(--primary);
	}

	.btn-edit:hover {
		background: #a7f3d0; /* Green 200 */
	}

	.btn-delete {
		background: #fee2e2; /* Red 100 */
		color: #ef4444; /* Red 500 */
	}

	.btn-delete:hover {
		background: #fecaca; /* Red 200 */
	}

	/* --- ADD BUTTON --- */
	.btn-add-reward {
		background: var(--primary-light);
		color: var(--primary);
		border: 2px dashed var(--primary);
		border-radius: var(--radius);
		padding: 1.5rem;
		font-size: 1rem;
		font-weight: 600;
		cursor: pointer;
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 0.75rem;
		transition: all 0.2s;
	}

	.btn-add-reward:hover {
		background: #d1fae5; /* Slightly darker green light */
		color: var(--primary-hover);
	}

	.btn-add-reward:disabled {
		opacity: 0.6;
		cursor: not-allowed;
	}

	/* --- FORM PANEL --- */
	.form-panel {
		background: var(--bg-white);
		border: 1px solid var(--border);
		border-radius: var(--radius);
		padding: 2rem;
		box-shadow: var(--shadow-md);
		position: sticky;
		top: 6rem;
		height: fit-content;
	}

	.form-panel h3 {
		font-size: 1.25rem;
		font-weight: 700;
		color: var(--text-dark);
		margin: 0 0 0.5rem 0;
	}

	.panel-placeholder {
		text-align: center;
		padding: 3rem 1rem;
		color: var(--text-light);
	}

	.panel-placeholder svg {
		color: var(--border);
		margin-bottom: 1rem;
	}

	.form-group-sm {
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
		margin-bottom: 1.5rem;
	}

	.form-group-sm label {
		font-size: 0.9rem;
		font-weight: 600;
		color: var(--text-dark);
	}

	input[type="text"],
	input[type="number"],
	textarea,
	select {
		width: 100%;
		padding: 0.75rem 1rem;
		border: 1px solid var(--border);
		border-radius: 6px;
		font-size: 1rem;
		font-family: inherit;
		color: var(--text-dark);
		outline: none;
		transition:
			border 0.2s,
			box-shadow 0.2s;
	}

	input:focus,
	textarea:focus,
	select:focus {
		border-color: var(--primary);
		box-shadow: 0 0 0 3px rgba(52, 162, 115, 0.2);
	}

	.input-group {
		display: flex;
		align-items: center;
		border: 1px solid var(--border);
		border-radius: 6px;
		transition:
			border 0.2s,
			box-shadow 0.2s;
	}

	.input-group:focus-within {
		border-color: var(--primary);
		box-shadow: 0 0 0 3px rgba(52, 162, 115, 0.2);
	}

	.input-group input {
		border: none;
		flex-grow: 1;
		padding-left: 0.5rem;
		box-shadow: none;
	}

	.input-prefix {
		padding: 0.75rem 0.5rem 0.75rem 1rem;
		color: var(--text-gray);
		font-weight: 600;
		background: #f9fafb;
		border-right: 1px solid var(--border);
		border-radius: 6px 0 0 6px;
	}

	.char-count {
		font-size: 0.75rem;
		color: var(--text-light);
		align-self: flex-end;
	}

	.checkbox-group {
		display: flex;
		align-items: center;
		gap: 0.5rem;
	}

	.checkbox-group input[type="checkbox"] {
		width: auto;
		height: auto;
	}

	.checkbox-group label {
		font-weight: 400;
		color: var(--text-gray);
	}

	.shipping-details {
		border: 1px solid #d1fae5;
		background: #f0fdf4;
		border-radius: 6px;
		padding: 1rem;
		margin-top: 0.5rem;
		display: flex;
		flex-direction: column;
		gap: 1rem;
	}

	.form-group-inline {
		display: flex;
		flex-direction: column;
		gap: 0.25rem;
	}

	/* --- FORM ACTIONS --- */
	.form-actions {
		display: flex;
		justify-content: flex-end;
		gap: 1rem;
		margin-top: 2rem;
	}

	.btn-cancel,
	.btn-save {
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
	}

	.btn-cancel {
		background: #e5e7eb;
		color: var(--text-gray);
	}

	.btn-cancel:hover {
		background: #d1d5db;
	}

	.btn-save {
		background: var(--primary);
		color: var(--bg-white);
	}

	.btn-save:hover {
		background: var(--primary-hover);
	}
</style>
