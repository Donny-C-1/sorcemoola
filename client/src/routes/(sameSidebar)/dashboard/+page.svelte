<script>
	// Mock Data - In a real app, these would be props or fetched from your Go backend
	let stats = $state({
		totalRaised: 45280.5,
		activeCampaigns: 4,
		totalDonors: 892,
		livesImpacted: 1240,
		availableBalance: 12450.0,
		kycStatus: "verified" // 'pending', 'verified', 'action_required'
	});

	let recentDonations = $state([
		{ id: 1, donor: "Sarah J.", amount: 500, campaign: "Flood Relief", time: "2m ago" },
		{ id: 2, donor: "Anonymous", amount: 1200, campaign: "Medical Supplies", time: "1h ago" },
		{ id: 3, donor: "TechCorp Ltd", amount: 5000, campaign: "Education Fund", time: "3h ago" },
		{ id: 4, donor: "Michael O.", amount: 50, campaign: "Flood Relief", time: "5h ago" }
	]);

	const formatCurrency = (val) => new Intl.NumberFormat("en-US", { style: "currency", currency: "USD" }).format(val);
</script>

<div class="overview-container">
	<header class="dashboard-header">
		<div class="welcome">
			<h1>Welcome back, SorceMoola Partner</h1>
			<p>Here’s what’s happening with your relief efforts today.</p>
		</div>
		<button class="btn-primary">+ Create Campaign</button>
	</header>

	<section class="stats-grid">
		<div class="card stat-card">
			<span class="label">Total Funds Raised</span>
			<h2 class="value text-primary">{formatCurrency(stats.totalRaised)}</h2>
			<div class="trend positive">↑ 12% from last month</div>
		</div>
		<div class="card stat-card">
			<span class="label">Active Campaigns</span>
			<h2 class="value">{stats.activeCampaigns}</h2>
			<div class="trend">Across 2 regions</div>
		</div>
		<div class="card stat-card">
			<span class="label">Total Donors</span>
			<h2 class="value">{stats.totalDonors}</h2>
			<div class="trend positive">↑ 48 new this week</div>
		</div>
		<div class="card stat-card impact">
			<span class="label">Lives Impacted</span>
			<h2 class="value text-accent">{stats.livesImpacted}</h2>
			<div class="trend">Direct aid delivered</div>
		</div>
	</section>

	<div class="main-content-split">
		<section class="card chart-container">
			<div class="card-header">
				<h3>Donation Trends</h3>
				<select class="small-select">
					<option>Last 7 Days</option>
					<option>Last 30 Days</option>
				</select>
			</div>
			<div class="visual-placeholder">
				<div class="bar-chart">
					{#each [40, 70, 45, 90, 65, 80, 50] as height}
						<div class="bar" style="height: {height}%"></div>
					{/each}
				</div>
				<p class="text-light">Visualization of daily contributions</p>
			</div>
		</section>

		<aside class="side-widgets">
			<div class="card financial-card">
				<h3>Available Balance</h3>
				<p class="balance-amount">{formatCurrency(stats.availableBalance)}</p>
				<button class="btn-secondary w-full">Withdraw Funds</button>
				<p class="sub-text">Next scheduled payout: Oct 24th</p>
			</div>

			<div class="card kyc-widget {stats.kycStatus}">
				<div class="kyc-icon">!</div>
				<div>
					<h4>Compliance Status</h4>
					<p>{stats.kycStatus === "verified" ? "Account fully verified" : "Action Required"}</p>
				</div>
			</div>
		</aside>
	</div>

	<section class="card table-card">
		<div class="card-header">
			<h3>Recent Donations</h3>
			<button class="btn-text">View All</button>
		</div>
		<table class="data-table">
			<thead>
				<tr>
					<th>Donor</th>
					<th>Campaign</th>
					<th>Amount</th>
					<th>Time</th>
				</tr>
			</thead>
			<tbody>
				{#each recentDonations as donation}
					<tr>
						<td><strong>{donation.donor}</strong></td>
						<td>{donation.campaign}</td>
						<td class="text-primary font-bold">+{formatCurrency(donation.amount)}</td>
						<td class="text-light">{donation.time}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</section>
</div>

<style>
	.overview-container {
		display: flex;
		flex-direction: column;
		gap: var(--spacing-lg);
		max-width: var(--container-width);
		margin: 0 auto;
		padding: var(--container-padding);
		color: var(--text-dark);
	}

	.dashboard-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
	}

	.dashboard-header h1 {
		font-size: 1.75rem;
		margin-bottom: var(--spacing-xs);
	}

	.dashboard-header p {
		color: var(--text-medium);
	}

	/* Card Base */
	.card {
		background: var(--text-white);
		border-radius: var(--radius-lg);
		padding: var(--spacing-md);
		box-shadow: var(--shadow-sm);
		border: 1px solid rgba(0, 0, 0, 0.05);
	}

	/* Stats Grid */
	.stats-grid {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
		gap: var(--spacing-md);
	}

	.stat-card .label {
		font-size: 0.875rem;
		color: var(--text-medium);
		display: block;
		margin-bottom: var(--spacing-xs);
	}

	.stat-card .value {
		font-size: 1.5rem;
		font-weight: 700;
		margin-bottom: var(--spacing-xs);
	}

	.text-primary {
		color: var(--primary-color);
	}
	.text-accent {
		color: var(--accent-color);
	}

	.trend {
		font-size: 0.75rem;
		color: var(--text-light);
	}

	.trend.positive {
		color: var(--primary-dark);
		font-weight: 600;
	}

	/* Main Split */
	.main-content-split {
		display: grid;
		grid-template-columns: 2fr 1fr;
		gap: var(--spacing-md);
	}

	.card-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: var(--spacing-md);
	}

	/* Visual Placeholder */
	.visual-placeholder {
		height: 200px;
		background: var(--primary-bg);
		border-radius: var(--radius-md);
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: flex-end;
		padding: var(--spacing-md);
	}

	.bar-chart {
		display: flex;
		align-items: flex-end;
		gap: var(--spacing-sm);
		width: 100%;
		height: 120px;
		margin-bottom: var(--spacing-sm);
	}

	.bar {
		flex: 1;
		background: var(--primary-light);
		border-radius: 4px 4px 0 0;
		transition: height 0.3s ease;
	}

	/* Side Widgets */
	.side-widgets {
		display: flex;
		flex-direction: column;
		gap: var(--spacing-md);
	}

	.financial-card {
		background: linear-gradient(135deg, var(--text-white) 0%, var(--primary-bg) 100%);
	}

	.balance-amount {
		font-size: 2rem;
		font-weight: 800;
		margin: var(--spacing-sm) 0;
	}

	.kyc-widget {
		display: flex;
		align-items: center;
		gap: var(--spacing-sm);
		border-left: 4px solid var(--neutral-color);
	}

	.kyc-widget.verified {
		border-left-color: var(--primary-color);
		background: #f0fff4;
	}

	.kyc-icon {
		width: 32px;
		height: 32px;
		background: var(--primary-color);
		color: white;
		border-radius: 50%;
		display: flex;
		align-items: center;
		justify-content: center;
		font-weight: bold;
	}

	/* Table Styles */
	.data-table {
		width: 100%;
		border-collapse: collapse;
		text-align: left;
	}

	.data-table th {
		padding: var(--spacing-smr);
		border-bottom: 2px solid var(--primary-bg);
		color: var(--text-light);
		font-size: 0.8rem;
		text-transform: uppercase;
	}

	.data-table td {
		padding: var(--spacing-sm);
		border-bottom: 1px solid var(--primary-bg);
		font-size: 0.95rem;
	}

	/* Buttons */
	.btn-primary {
		background: var(--primary-color);
		color: white;
		padding: var(--spacing-smr) var(--spacing-md);
		border-radius: var(--radius-md);
		border: none;
		font-weight: 600;
		cursor: pointer;
	}

	.btn-secondary {
		background: var(--text-dark);
		color: white;
		padding: var(--spacing-smr);
		border-radius: var(--radius-md);
		border: none;
		cursor: pointer;
	}

	.w-full {
		width: 100%;
	}
	.font-bold {
		font-weight: 700;
	}
	.text-light {
		color: var(--text-light);
	}

	@media (max-width: 768px) {
		.main-content-split {
			grid-template-columns: 1fr;
		}
	}
</style>
