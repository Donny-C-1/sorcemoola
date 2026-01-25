<script>
	import { onMount } from "svelte";
	import Button from "$lib/components/ui/Button.svelte";
	// import defaultAvatar from "$lib/assets/default-avatar.svg?raw";

	// User data
	let user = {
		fullName: "John Doe",
		email: "john@example.com",
		accountType: "Individual",
		avatar: "/images/avatar.png",
		campaignsCreated: 3,
		donationsMade: 5,
		fundsRaised: 1200000
	};

	// Edit mode state
	let editMode = false;
	let fullname = user.fullName;
	let email = user.email;
	let accountType = user.accountType;

	// Mock campaigns
	let campaigns = [
		{
			title: "School Supplies for Children",
			fundsRaised: 50000,
			target: 100000
		},
		{
			title: "Support Local Artist",
			fundsRaised: 200000,
			target: 250000
		},
		{
			title: "Community Health Initiative",
			fundsRaised: 800000,
			target: 1000000
		}
	];

	// Mock recent donations
	let donations = [
		{ campaign: "Orphanage Fund", amount: 10000, date: "2025-12-10" },
		{ campaign: "Clean Water Project", amount: 5000, date: "2025-12-08" },
		{ campaign: "School Supplies", amount: 15000, date: "2025-12-05" }
	];

	function toggleEdit() {
		editMode = !editMode;
	}

	function saveProfile() {
		user.fullName = fullname;
		user.email = email;
		user.accountType = accountType;
		editMode = false;
	}
</script>

<main>
	<section class="profile-header">
		<div class="avatar">
			<img src={user.avatar} alt={user.fullName} />
		</div>
		<div class="user-info">
			{#if editMode}
				<input type="text" bind:value={fullname} placeholder="Full Name" />
				<input type="email" bind:value={email} placeholder="Email" />
				<select bind:value={accountType}>
					<option value="Individual">Individual</option>
					<option value="NGO">NGO</option>
					<option value="Organization">Organization</option>
				</select>
				<div class="buttons">
					<Button on:click={saveProfile}>Save</Button>
					<Button on:click={toggleEdit} primary={false}>Cancel</Button>
				</div>
			{:else}
				<h1>{user.fullName}</h1>
				<p class="email">{user.email}</p>
				<p class="account-type">{user.accountType}</p>
				<Button on:click={toggleEdit}>Edit Profile</Button>
			{/if}
		</div>
	</section>

	<section class="profile-stats">
		<div class="stat-card">
			<p class="label">Campaigns Created</p>
			<p class="value">{user.campaignsCreated}</p>
		</div>
		<div class="stat-card">
			<p class="label">Funds Raised</p>
			<p class="value">₦{user.fundsRaised.toLocaleString()}</p>
		</div>
		<div class="stat-card">
			<p class="label">Donations Made</p>
			<p class="value">{user.donationsMade}</p>
		</div>
	</section>

	<section class="campaigns">
		<h2>Your Campaigns</h2>
		{#each campaigns as campaign}
			<div class="campaign-card">
				<p class="title">{campaign.title}</p>
				<div class="progress-bar">
					<div class="progress" style="width: {Math.min((campaign.fundsRaised / campaign.target) * 100, 100)}%"></div>
				</div>
				<p class="progress-text">
					₦{campaign.fundsRaised.toLocaleString()} raised of ₦{campaign.target.toLocaleString()}
				</p>
			</div>
		{/each}
	</section>

	<section class="donations">
		<h2>Recent Donations</h2>
		{#each donations as donation}
			<div class="donation-card">
				<p>{donation.campaign}</p>
				<p>₦{donation.amount.toLocaleString()}</p>
				<p class="date">{donation.date}</p>
			</div>
		{/each}
	</section>
</main>

<style>
	main {
		font-family: Poppins, sans-serif;
		max-width: var(--container-width);
		margin: 0 auto;
		padding: var(--spacing-lg);
		display: flex;
		flex-direction: column;
		gap: var(--spacing-lg);
	}

	.profile-header {
		display: flex;
		flex-direction: column;
		align-items: center;
		background-color: var(--primary-bg);
		padding: var(--spacing-lg);
		border-radius: var(--radius-lg);
		box-shadow: var(--shadow-md);
		text-align: center;
		gap: var(--spacing-md);
	}

	@media screen and (min-width: 768px) {
		.profile-header {
			flex-direction: row;
			text-align: left;
			gap: var(--spacing-xl);
		}
	}

	.avatar img {
		width: 120px;
		height: 120px;
		border-radius: 50%;
		object-fit: cover;
		border: 3px solid var(--primary-color);
		box-shadow: var(--shadow-sm);
	}

	.user-info h1 {
		font-size: 1.8rem;
		color: var(--text-dark);
		margin: 0;
	}

	.user-info .email {
		color: var(--text-medium);
		margin: 0.2rem 0;
	}

	.user-info .account-type {
		font-weight: 600;
		color: var(--primary-color);
		margin-bottom: var(--spacing-sm);
	}

	.user-info input,
	.user-info select {
		display: block;
		width: 100%;
		max-width: 300px;
		margin-bottom: var(--spacing-sm);
		padding: var(--spacing-sm);
		border: 1px solid var(--neutral-color);
		border-radius: var(--radius-sm);
	}

	.buttons {
		display: flex;
		gap: var(--spacing-sm);
		justify-content: center;
		margin-top: var(--spacing-sm);
	}

	.profile-stats {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(160px, 1fr));
		gap: var(--spacing-md);
	}

	.stat-card {
		background-color: var(--primary-bg);
		padding: var(--spacing-md);
		border-radius: var(--radius-md);
		text-align: center;
		box-shadow: var(--shadow-sm);
	}

	.stat-card .label {
		font-weight: 500;
		color: var(--text-medium);
		margin-bottom: var(--spacing-xs);
	}

	.stat-card .value {
		font-size: 1.4rem;
		font-weight: 700;
		color: var(--primary-color);
	}

	.campaigns h2,
	.donations h2 {
		font-size: 1.4rem;
		font-weight: 600;
		color: var(--text-dark);
		margin-bottom: var(--spacing-md);
	}

	.campaign-card,
	.donation-card {
		background-color: var(--primary-bg);
		padding: var(--spacing-md);
		border-radius: var(--radius-md);
		box-shadow: var(--shadow-sm);
		margin-bottom: var(--spacing-md);
	}

	.campaign-card .title {
		font-weight: 600;
		color: var(--text-dark);
		margin-bottom: var(--spacing-xs);
	}

	.progress-bar {
		background-color: var(--neutral-color);
		border-radius: var(--radius-sm);
		height: 10px;
		width: 100%;
		overflow: hidden;
		margin-bottom: var(--spacing-xs);
	}

	.progress {
		background-color: var(--primary-color);
		height: 100%;
		border-radius: var(--radius-sm);
		transition: width 0.5s ease;
	}

	.progress-text {
		font-size: 0.85rem;
		color: var(--text-medium);
	}

	.donation-card {
		display: flex;
		justify-content: space-between;
		align-items: center;
		font-size: 0.95rem;
	}

	.donation-card .date {
		color: var(--text-light);
		font-size: 0.8rem;
	}
</style>
