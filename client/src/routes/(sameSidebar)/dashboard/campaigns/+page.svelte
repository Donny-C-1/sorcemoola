<script>
	import { stringToHSL } from "$lib/utils/color.js";

	let { data } = $props();
	let searchQuery = $state("");
	let filterStatus = $state("active");

	let filteredCampaigns = $derived(
		data.campaigns.filter((c) => {
			const matchesSearch = c.name.toLowerCase().includes(searchQuery.toLowerCase());
			const matchesFilter = filterStatus === "all" || c.status === filterStatus;
			return matchesSearch && matchesFilter;
		})
	);

	const calculateProgress = (raised, goal) => Math.min((raised / goal) * 100, 100);

	const formatCurrency = (amount) => {
		return new Intl.NumberFormat("en-NG", {
			style: "currency",
			currency: "NGN",
			minimumFractionDigits: 0
		}).format(amount / 100);
	};
</script>

<div class="campaigns-page">
	<div class="page-hero">
		<div class="hero-content">
			<div class="hero-badge">Your Impact</div>
			<h1 class="hero-title">Campaign Management</h1>
			<p class="hero-subtitle">Every campaign is a story of hope and change</p>
		</div>
		<div class="hero-decoration"></div>
	</div>

	<div class="controls-bar">
		<div class="search-wrapper">
			<svg class="search-icon" width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
				<circle cx="11" cy="11" r="8" stroke="currentColor" stroke-width="2" />
				<path d="M21 21L16.65 16.65" stroke="currentColor" stroke-width="2" stroke-linecap="round" />
			</svg>
			<input type="text" placeholder="Search by name or cause..." bind:value={searchQuery} />
		</div>

		<div class="filter-pills">
			<button class="pill" class:active={filterStatus === "active"} onclick={() => (filterStatus = "active")}>
				<span class="pill-dot active-dot"></span>
				Active
			</button>
			<button class="pill" class:active={filterStatus === "completed"} onclick={() => (filterStatus = "completed")}>
				<span class="pill-dot completed-dot"></span>
				Completed
			</button>
			<button class="pill" class:active={filterStatus === "all"} onclick={() => (filterStatus = "all")}> All Campaigns </button>
		</div>

		<a href="/dashboard/create" class="btn-create">
			<svg width="18" height="18" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
				<path d="M12 5V19M5 12H19" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" />
			</svg>
			New Campaign
		</a>
	</div>

	<div class="campaigns-grid">
		{#each filteredCampaigns as campaign, index (campaign.id)}
			<article class="campaign-card" style="--delay: {index * 50}ms">
				<a href="/campaigns/{campaign.slug}" class="card-link">
					<div class="card-visual">
						{#if campaign.imageUrl}
							<img src={campaign.imageUrl} alt={campaign.name} />
							<div class="image-overlay"></div>
						{:else}
							<div class="placeholder-visual" style:background={stringToHSL(campaign.name)}>
								<div class="placeholder-pattern"></div>
								<h3 class="placeholder-title">{campaign.name}</h3>
							</div>
						{/if}
						<div class="category-tag">{campaign.category}</div>
						<div class="status-indicator {campaign.status}">
							<span class="status-pulse"></span>
							{campaign.status === "active" ? "Live" : "Ended"}
						</div>
					</div>

					<div class="card-content">
						<h2 class="campaign-title">{campaign.name}</h2>
						<p class="campaign-description">{campaign.description || campaign.story?.substring(0, 100) + "..."}</p>

						<div class="funding-section">
							<div class="funding-header">
								<span class="amount-raised">{formatCurrency(campaign.amountRaised)}</span>
								<span class="amount-divider">/</span>
								<span class="amount-goal">{formatCurrency(campaign.fundGoal)}</span>
							</div>

							<div class="progress-track">
								<div class="progress-bar" style="width: {calculateProgress(campaign.amountRaised, campaign.fundGoal)}%">
									<div class="progress-shine"></div>
								</div>
							</div>

							<div class="funding-stats">
								<div class="stat-item">
									<svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
										<path d="M17 21V19C17 17.9391 16.5786 16.9217 15.8284 16.1716C15.0783 15.4214 14.0609 15 13 15H5C3.93913 15 2.92172 15.4214 2.17157 16.1716C1.42143 16.9217 1 17.9391 1 19V21" stroke="currentColor" stroke-width="2" stroke-linecap="round" />
										<circle cx="9" cy="7" r="4" stroke="currentColor" stroke-width="2" />
										<path d="M23 21V19C22.9993 18.1137 22.7044 17.2528 22.1614 16.5523C21.6184 15.8519 20.8581 15.3516 20 15.13" stroke="currentColor" stroke-width="2" stroke-linecap="round" />
										<path d="M16 3.13C16.8604 3.35031 17.623 3.85071 18.1676 4.55232C18.7122 5.25392 19.0078 6.11683 19.0078 7.005C19.0078 7.89318 18.7122 8.75608 18.1676 9.45769C17.623 10.1593 16.8604 10.6597 16 10.88" stroke="currentColor" stroke-width="2" stroke-linecap="round" />
									</svg>
									<span>{campaign.backersCount} supporters</span>
								</div>
								<div class="stat-item">
									<span class="percentage-badge">{Math.round(calculateProgress(campaign.amountRaised, campaign.fundGoal))}%</span>
								</div>
							</div>
						</div>

						<div class="card-footer">
							<button
								class="action-btn"
								onclick={(e) => {
									e.preventDefault();
								}}
							>
								<svg width="18" height="18" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
									<path d="M11 4H4C3.46957 4 2.96086 4.21071 2.58579 4.58579C2.21071 4.96086 2 5.46957 2 6V20C2 20.5304 2.21071 21.0391 2.58579 21.4142C2.96086 21.7893 3.46957 22 4 22H18C18.5304 22 19.0391 21.7893 19.4142 21.4142C19.7893 21.0391 20 20.5304 20 20V13" stroke="currentColor" stroke-width="2" stroke-linecap="round" />
									<path d="M18.5 2.50001C18.8978 2.10219 19.4374 1.87869 20 1.87869C20.5626 1.87869 21.1022 2.10219 21.5 2.50001C21.8978 2.89784 22.1213 3.4374 22.1213 4.00001C22.1213 4.56262 21.8978 5.10219 21.5 5.50001L12 15L8 16L9 12L18.5 2.50001Z" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
								</svg>
								Manage
							</button>
							<button
								class="action-btn secondary"
								onclick={(e) => {
									e.preventDefault();
								}}
							>
								<svg width="18" height="18" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
									<path d="M4 12V20C4 20.5304 4.21071 21.0391 4.58579 21.4142C4.96086 21.7893 5.46957 22 6 22H18C18.5304 22 19.0391 21.7893 19.4142 21.4142C19.7893 21.0391 20 20.5304 20 20V12" stroke="currentColor" stroke-width="2" stroke-linecap="round" />
									<path d="M16 6L12 2L8 6" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
									<path d="M12 2V15" stroke="currentColor" stroke-width="2" stroke-linecap="round" />
								</svg>
								Share
							</button>
						</div>
					</div>
				</a>
			</article>
		{/each}

		{#if filteredCampaigns.length === 0}
			<div class="empty-state">
				<svg width="64" height="64" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
					<circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="1.5" />
					<path d="M8 15C8.91212 16.2144 10.3643 17 12 17C13.6357 17 15.0879 16.2144 16 15" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
					<path d="M9 9H9.01M15 9H15.01" stroke="currentColor" stroke-width="2" stroke-linecap="round" />
				</svg>
				<h3>No campaigns found</h3>
				<p>Try adjusting your search or filters</p>
			</div>
		{/if}
	</div>
</div>

<style>
	@import url("https://fonts.googleapis.com/css2?family=Fraunces:ital,opsz,wght@0,9..144,300..900;1,9..144,300..900&family=Inter:wght@300..700&display=swap");

	.campaigns-page {
		min-height: 100vh;
		background: linear-gradient(165deg, #f8fffe 0%, #fafcfb 100%);
		padding: 0 0 var(--spacing-3xl);
		position: relative;
	}

	.campaigns-page::before {
		content: "";
		position: absolute;
		top: 0;
		right: 0;
		width: 600px;
		height: 600px;
		background: radial-gradient(circle, rgba(52, 162, 115, 0.08) 0%, transparent 70%);
		pointer-events: none;
	}

	/* Hero Section */
	.page-hero {
		position: relative;
		padding: var(--spacing-3xl) var(--container-padding) var(--spacing-2xl);
		max-width: var(--container-width);
		margin: 0 auto;
		overflow: hidden;
	}

	.hero-content {
		position: relative;
		z-index: 2;
	}

	.hero-badge {
		display: inline-block;
		padding: 6px 16px;
		background: linear-gradient(135deg, var(--primary-color), var(--primary-light));
		color: white;
		font-size: 0.75rem;
		font-weight: 700;
		letter-spacing: 0.5px;
		text-transform: uppercase;
		border-radius: 20px;
		margin-bottom: var(--spacing-md);
		box-shadow: 0 4px 12px rgba(52, 162, 115, 0.3);
	}

	.hero-title {
		font-family: "Fraunces", serif;
		font-size: 4rem;
		font-weight: 700;
		line-height: 1.1;
		color: var(--text-dark);
		margin: 0 0 var(--spacing-sm) 0;
		letter-spacing: -0.02em;
	}

	.hero-subtitle {
		font-size: 1.25rem;
		color: var(--text-medium);
		margin: 0;
		font-weight: 400;
	}

	.hero-decoration {
		position: absolute;
		top: 20%;
		right: -100px;
		width: 400px;
		height: 400px;
		background: linear-gradient(135deg, var(--secondary-color) 0%, var(--accent-color) 100%);
		opacity: 0.06;
		border-radius: 50%;
		filter: blur(80px);
		pointer-events: none;
	}

	/* Controls Bar */
	.controls-bar {
		max-width: var(--container-width);
		margin: 0 auto var(--spacing-2xl);
		padding: 0 var(--container-padding);
		display: flex;
		gap: var(--spacing-md);
		align-items: center;
		flex-wrap: wrap;
	}

	.search-wrapper {
		position: relative;
		flex: 1;
		min-width: 280px;
	}

	.search-icon {
		position: absolute;
		left: 16px;
		top: 50%;
		transform: translateY(-50%);
		color: var(--text-light);
		pointer-events: none;
	}

	.search-wrapper input {
		width: 100%;
		padding: 14px 16px 14px 48px;
		border: 2px solid rgba(52, 162, 115, 0.12);
		border-radius: 12px;
		font-size: 0.95rem;
		background: white;
		transition: all 0.3s ease;
	}

	.search-wrapper input:focus {
		outline: none;
		border-color: var(--primary-color);
		box-shadow: 0 4px 20px rgba(52, 162, 115, 0.15);
	}

	.filter-pills {
		display: flex;
		flex-wrap: wrap;
		gap: 8px;
		background: white;
		padding: 6px;
		border-radius: 12px;
		border: 2px solid rgba(52, 162, 115, 0.08);
	}

	.pill {
		display: flex;
		align-items: center;
		gap: 8px;
		padding: 10px 20px;
		border: none;
		background: transparent;
		border-radius: 8px;
		font-size: 0.9rem;
		font-weight: 600;
		color: var(--text-medium);
		cursor: pointer;
		transition: all 0.2s ease;
	}

	.pill:hover {
		background: var(--primary-bg);
	}

	.pill.active {
		background: var(--primary-color);
		color: white;
	}

	.pill-dot {
		width: 8px;
		height: 8px;
		border-radius: 50%;
	}

	.active-dot {
		background: var(--primary-color);
		box-shadow: 0 0 8px var(--primary-light);
	}

	.pill.active .active-dot {
		background: white;
		box-shadow: 0 0 8px rgba(255, 255, 255, 0.5);
	}

	.completed-dot {
		background: var(--text-light);
	}

	.pill.active .completed-dot {
		background: white;
	}

	.btn-create {
		display: flex;
		align-items: center;
		gap: 8px;
		padding: 14px 24px;
		background: var(--primary-color);
		color: white;
		border: none;
		border-radius: 12px;
		font-size: 0.95rem;
		font-weight: 700;
		cursor: pointer;
		transition: all 0.3s ease;
		text-decoration: none;
	}

	.btn-create:hover {
		background: var(--primary-dark);
		transform: translateY(-2px);
		box-shadow: 0 8px 24px rgba(52, 162, 115, 0.3);
	}

	/* Campaign Grid */
	.campaigns-grid {
		max-width: var(--container-width);
		margin: 0 auto;
		padding: 0 var(--container-padding);
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(380px, 1fr));
		gap: var(--spacing-xl);
	}

	.campaign-card {
		background: white;
		border-radius: 20px;
		overflow: hidden;
		box-shadow: 0 4px 24px rgba(0, 0, 0, 0.06);
		transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
		border: 1px solid rgba(52, 162, 115, 0.08);
		animation: fadeInUp 0.6s ease forwards;
		animation-delay: var(--delay);
		opacity: 0;
	}

	@keyframes fadeInUp {
		from {
			opacity: 0;
			transform: translateY(30px);
		}
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}

	.campaign-card:hover {
		transform: translateY(-8px);
		box-shadow: 0 12px 48px rgba(0, 0, 0, 0.12);
		border-color: var(--primary-color);
	}

	.card-link {
		text-decoration: none;
		color: inherit;
		display: block;
	}

	/* Card Visual */
	.card-visual {
		position: relative;
		height: 240px;
		overflow: hidden;
	}

	.card-visual img {
		width: 100%;
		height: 100%;
		object-fit: cover;
		transition: transform 0.6s ease;
	}

	.campaign-card:hover .card-visual img {
		transform: scale(1.08);
	}

	.image-overlay {
		position: absolute;
		inset: 0;
		background: linear-gradient(180deg, transparent 0%, rgba(0, 0, 0, 0.4) 100%);
		opacity: 0;
		transition: opacity 0.3s ease;
	}

	.campaign-card:hover .image-overlay {
		opacity: 1;
	}

	.placeholder-visual {
		width: 100%;
		height: 100%;
		position: relative;
		display: flex;
		align-items: center;
		justify-content: center;
		padding: var(--spacing-xl);
		overflow: hidden;
	}

	.placeholder-pattern {
		position: absolute;
		inset: 0;
		background-image: repeating-linear-gradient(45deg, transparent, transparent 10px, rgba(255, 255, 255, 0.1) 10px, rgba(255, 255, 255, 0.1) 20px);
	}

	.placeholder-title {
		position: relative;
		color: white;
		text-align: center;
		font-size: 1.5rem;
		font-weight: 700;
		text-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
		line-height: 1.3;
		z-index: 1;
	}

	.category-tag {
		position: absolute;
		top: 16px;
		left: 16px;
		padding: 8px 16px;
		background: rgba(255, 255, 255, 0.95);
		backdrop-filter: blur(10px);
		border-radius: 8px;
		font-size: 0.75rem;
		font-weight: 700;
		color: var(--primary-dark);
		text-transform: uppercase;
		letter-spacing: 0.5px;
	}

	.status-indicator {
		position: absolute;
		top: 16px;
		right: 16px;
		display: flex;
		align-items: center;
		gap: 6px;
		padding: 8px 14px;
		background: rgba(255, 255, 255, 0.95);
		backdrop-filter: blur(10px);
		border-radius: 20px;
		font-size: 0.75rem;
		font-weight: 700;
		text-transform: uppercase;
	}

	.status-indicator.active {
		color: var(--primary-color);
	}

	.status-indicator.completed {
		color: var(--text-medium);
	}

	.status-pulse {
		width: 8px;
		height: 8px;
		border-radius: 50%;
		background: currentColor;
		animation: pulse 2s ease infinite;
	}

	@keyframes pulse {
		0%,
		100% {
			opacity: 1;
			transform: scale(1);
		}
		50% {
			opacity: 0.5;
			transform: scale(1.2);
		}
	}

	/* Card Content */
	.card-content {
		padding: var(--spacing-lg);
	}

	.campaign-title {
		font-family: "Fraunces", serif;
		font-size: 1.5rem;
		font-weight: 700;
		line-height: 1.3;
		margin: 0 0 var(--spacing-sm) 0;
		color: var(--text-dark);
		transition: color 0.3s ease;
	}

	.campaign-card:hover .campaign-title {
		color: var(--primary-color);
	}

	.campaign-description {
		font-size: 0.95rem;
		line-height: 1.6;
		color: var(--text-medium);
		margin: 0 0 var(--spacing-lg) 0;
		display: -webkit-box;
		-webkit-line-clamp: 2;
		-webkit-box-orient: vertical;
		overflow: hidden;
	}

	/* Funding Section */
	.funding-section {
		margin-bottom: var(--spacing-lg);
		padding: var(--spacing-md);
		background: var(--primary-bg);
		border-radius: 12px;
	}

	.funding-header {
		display: flex;
		align-items: baseline;
		gap: 8px;
		margin-bottom: var(--spacing-sm);
	}

	.amount-raised {
		font-size: 1.75rem;
		font-weight: 800;
		color: var(--primary-dark);
		font-family: "Fraunces", serif;
	}

	.amount-divider {
		font-size: 1.25rem;
		color: var(--text-light);
	}

	.amount-goal {
		font-size: 1.1rem;
		color: var(--text-medium);
		font-weight: 600;
	}

	.progress-track {
		height: 10px;
		background: rgba(52, 162, 115, 0.15);
		border-radius: 20px;
		overflow: hidden;
		margin-bottom: var(--spacing-sm);
		position: relative;
	}

	.progress-bar {
		height: 100%;
		background: linear-gradient(90deg, var(--primary-color), var(--primary-light));
		border-radius: 20px;
		position: relative;
		transition: width 0.8s cubic-bezier(0.4, 0, 0.2, 1);
	}

	.progress-shine {
		position: absolute;
		inset: 0;
		background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.3), transparent);
		animation: shine 2s ease infinite;
	}

	@keyframes shine {
		0% {
			transform: translateX(-100%);
		}
		100% {
			transform: translateX(100%);
		}
	}

	.funding-stats {
		display: flex;
		justify-content: space-between;
		align-items: center;
	}

	.stat-item {
		display: flex;
		align-items: center;
		gap: 6px;
		font-size: 0.85rem;
		color: var(--text-medium);
		font-weight: 500;
	}

	.stat-item svg {
		color: var(--primary-color);
	}

	.percentage-badge {
		padding: 4px 12px;
		background: white;
		border-radius: 12px;
		font-weight: 700;
		color: var(--primary-color);
		font-size: 0.85rem;
	}

	/* Card Footer */
	.card-footer {
		display: flex;
		gap: var(--spacing-xs);
		padding-top: var(--spacing-md);
		border-top: 1px solid rgba(52, 162, 115, 0.1);
	}

	.action-btn {
		flex: 1;
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 8px;
		padding: 12px;
		border: 2px solid var(--primary-color);
		background: var(--primary-color);
		color: white;
		border-radius: 10px;
		font-size: 0.9rem;
		font-weight: 600;
		cursor: pointer;
		transition: all 0.3s ease;
	}

	.action-btn:hover {
		background: var(--primary-dark);
		border-color: var(--primary-dark);
		transform: translateY(-2px);
	}

	.action-btn.secondary {
		background: white;
		color: var(--primary-color);
	}

	.action-btn.secondary:hover {
		background: var(--primary-bg);
	}

	/* Empty State */
	.empty-state {
		grid-column: 1 / -1;
		text-align: center;
		padding: var(--spacing-3xl);
		color: var(--text-medium);
	}

	.empty-state svg {
		margin-bottom: var(--spacing-md);
		color: var(--text-light);
	}

	.empty-state h3 {
		font-family: "Fraunces", serif;
		font-size: 1.5rem;
		margin: 0 0 var(--spacing-xs) 0;
		color: var(--text-dark);
	}

	.empty-state p {
		margin: 0;
		color: var(--text-light);
	}

	/* Responsive */
	@media (max-width: 768px) {
		.hero-title {
			font-size: 2.5rem;
		}

		.controls-bar {
			flex-direction: column;
			align-items: stretch;
		}

		.search-wrapper {
			min-width: 100%;
		}

		.campaigns-grid {
			grid-template-columns: 1fr;
			gap: var(--spacing-lg);
		}

		.card-footer {
			flex-direction: column;
		}
	}
</style>
