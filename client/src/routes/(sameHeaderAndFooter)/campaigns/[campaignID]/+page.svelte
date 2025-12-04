<script>
	import { onMount } from "svelte";
	import { fade, fly } from "svelte/transition";
	import { PUBLIC_SERVER_URL } from "$env/static/public";
	import { Splide, SplideSlide } from "@splidejs/svelte-splide";
	import "@splidejs/svelte-splide/css";
	import Button from "$lib/components/ui/Button.svelte";

	let PaystackPop;

	let campaign = {
		id: "cam123456",
		title: "Save the Enchanted Forest",
		creator: "Wizard Willow",
		creatorAvatar:
			"https://images.unsplash.com/photo-1760574772950-f37de9dce85c?ixid=M3w4MjcwNjd8MHwxfHNlYXJjaHwzfHx3aXphcmQlMjB3aWxsb3d8ZW58MHx8fHwxNzYzNjYxNTA4fDA&ixlib=rb-4.1.0&fit=max&q=80",
		targetAmount: 50000,
		currentAmount: 32750,
		backers: 428,
		daysLeft: 15,
		donors: [
			{
				name: "John Doe",
				amount: 150,
				timeAgo: "2h ago",
				avatar: "/images/avatar1.jpg"
			},
			{
				name: "Sarah Smith",
				amount: 75,
				timeAgo: "5h ago",
				avatar: null
			},
			{
				name: "Mike Johnson",
				amount: 250,
				timeAgo: "1d ago",
				avatar: "/images/avatar3.jpg"
			}
		],
		description:
			"The Enchanted Forest is home to magical creatures and rare plants with healing properties. Industrial development threatens to destroy this unique ecosystem. Your contributions will help us purchase the land and establish it as a protected sanctuary.",
		story: `<p>For centuries, the Enchanted Forest has been a haven for magical creatures and a source of powerful ingredients for potions and spells. The crystal-clear streams that flow through it are said to have rejuvenating properties, and the rare flora that grows here can't be found anywhere else in the world.</p>
                <p>Unfortunately, a large corporation has acquired the land and plans to clear it for development. Once this ecosystem is destroyed, we can never get it back.</p>
                <p>With your help, we aim to purchase the land and establish it as a protected magical sanctuary. Every contribution brings us one step closer to saving this irreplaceable natural wonder.</p>`,
		images: [
			"https://images.unsplash.com/photo-1619982870053-1545857b7eb4?ixid=M3w4MjcwNjd8MHwxfHNlYXJjaHwzfHxlbmNoYW50ZWQlMjBmb3Jlc3R8ZW58MHx8fHwxNzYzNjYxMjc5fDA&ixlib=rb-4.1.0&fit=max&q=80",
			"https://images.unsplash.com/photo-1597201423947-3e0028337902?ixid=M3w4MjcwNjd8MHwxfHNlYXJjaHw1fHxlbmNoYW50ZWQlMjBmb3Jlc3R8ZW58MHx8fHwxNzYzNjYxMjc5fDA&ixlib=rb-4.1.0&fit=max&q=80",
			"https://images.unsplash.com/photo-1518562180175-34a163b1a9a6?ixid=M3w4MjcwNjd8MHwxfHNlYXJjaHwxMHx8ZW5jaGFudGVkJTIwZm9yZXN0fGVufDB8fHx8MTc2MzY2MTI3OXww&ixlib=rb-4.1.0&fit=max&q=80"
		],
		updates: [
			{
				date: "2023-10-25",
				title: "Legal progress!",
				content:
					"We've secured a temporary injunction to halt development while our case is being reviewed."
			},
			{
				date: "2023-10-10",
				title: "25% funded!",
				content: "Thank you to all our amazing supporters. We're a quarter of the way there!"
			}
		],
		backersCount: 187,
		totalBackers: 187,
		rewards: [
			{
				id: "r1",
				amount: 25,
				title: "Magical Seeds",
				description:
					"A packet of seeds from the Enchanted Forest. Plant them in your garden to attract friendly sprites.",
				backers: 215,
				delivery: "December 2023"
			},
			{
				id: "r2",
				amount: 100,
				title: "Crystal Vial",
				description:
					"A small vial of water from the Enchanted Stream, known for its restorative properties.",
				backers: 132,
				delivery: "December 2023"
			},
			{
				id: "r3",
				amount: 500,
				title: "Guardian Status",
				description:
					"Your name will be engraved on the Sanctuary Guardian plaque. Includes all previous rewards.",
				backers: 45,
				delivery: "January 2024"
			}
		]
	};

	let options = {
		perPage: 1,
		type: "loop",
		gap: "1rem",
		autoplay: "true"
	}

	let pledgeAmount = 0;
	let customAmount = "";
	let selectedReward = null;
	let showDonationForm = false;
	let activeTab = "story";

	const formatCurrency = (amount) => {
		return new Intl.NumberFormat("en-US", { style: "currency", currency: "USD" }).format(amount);
	};

	const getProgressPercentage = () => {
		return Math.min(Math.floor((campaign.currentAmount / campaign.targetAmount) * 100), 100);
	};

	const selectReward = (reward) => {
		selectedReward = reward;
		pledgeAmount = reward.amount;
		customAmount = "";
		showDonationForm = true;
	};

	const handleCustomPledge = () => {
		selectedReward = null;
		pledgeAmount = parseFloat(customAmount);
		showDonationForm = true;
	};

	const submitDonation = () => {
		// In a real app, this would connect to a payment processor
		alert(
			`Thank you for your ${formatCurrency(pledgeAmount)} pledge to save the Enchanted Forest!`
		);

		initPayment();
		showDonationForm = false;
		selectedReward = null;
		pledgeAmount = 0;
		customAmount = "";
	};

	onMount(async () => {
		const PaystackModule = await import("@paystack/inline-js");
		PaystackPop = PaystackModule.default || PaystackModule;
		// Image slideshow automation
		const interval = setInterval(nextImage, 5000);
		return () => clearInterval(interval);
	});

	async function initPayment() {
		let accessCode;
		try {
			const response = await fetch(`${PUBLIC_SERVER_URL}/campaigns/fund`);

			if (!response.ok) throw new Error("Server Error");

			const data = await response.json();

			accessCode = data.data.access_code;
			console.log(data);
		} catch (err) {
			console.log(err);
			return;
		}

		const popup = new PaystackPop();
		popup.resumeTransaction(accessCode);
	}
</script>

<svelte:head>
	<title>{campaign.title} | Sorcemoola</title>
	<meta name="description" content={campaign.description} />
</svelte:head>

<main class="container">
	<section class="campaign-header">
		<div class="image-gallery">
			<Splide {options}>
				{#each campaign.images as image, i}
					<SplideSlide>
						<img
							class="gallery_image"
							src={image}
							alt={`${campaign.title} - image ${i + 1}`}
							width="100%"
							height="auto"
						/>
					</SplideSlide>
				{/each}
			</Splide>
		</div>

		<div class="campaign-title">
			<h1>{campaign.title}</h1>
			<div class="creator-info">
				<img src={campaign.creatorAvatar} alt={campaign.creator} />
				<p>by <strong>{campaign.creator}</strong></p>
			</div>
		</div>
	</section>

	<div class="campaign-content">
		<div class="campaign-details">
			<div class="campaign-tabs">
				<button class:active={activeTab === "story"} onclick={() => (activeTab = "story")}>
					Campaign Story
				</button>
				<button class:active={activeTab === "updates"} onclick={() => (activeTab = "updates")}>
					Updates ({campaign.updates.length})
				</button>
			</div>

			<div class="tab-content">
				{#if activeTab === "story"}
					<section class="story-tab" transition:fade={{ duration: 200 }}>
						<blockquote class="campaign-summary">{campaign.description}</blockquote>
						<div class="campaign-story">
							{@html campaign.story}
						</div>
					</section>
				{:else if activeTab === "updates"}
					<section class="updates-tab" transition:fade={{ duration: 200 }}>
						<h2>Campaign Updates</h2>
						{#if campaign.updates.length > 0}
							<div class="updates-list">
								{#each campaign.updates as update, i}
									<div class="update-item" in:fly={{ y: 20, delay: i * 100, duration: 300 }}>
										<div class="update-date">{update.date}</div>
										<div class="update-content">
											<h3>{update.title}</h3>
											<p>{update.content}</p>
										</div>
									</div>
								{/each}
							</div>
						{:else}
							<p class="no-updates">No updates posted yet.</p>
						{/if}
					</section>
				{/if}
			</div>
		</div>

		<aside class="campaign_aside">
			<div class="funding_status">
				<div class="funding_amount">
					<h2 class="amount_raised">{formatCurrency(campaign.currentAmount)}</h2>
					<p>raised of {formatCurrency(campaign.targetAmount)}</p>
				</div>

				<div class="progress_container">
					<div class="progress_bar">
						<div class="progress_fill" style:width="{getProgressPercentage()}%"></div>
					</div>
				</div>

				<div class="backer_count">
					<h3>{campaign.backersCount.toLocaleString()}</h3>
					<p>backers</p>
				</div>
			</div>

			<div class="action_buttons">
				<Button large={true} handler={() => showDonationForm = true}>Dontate Now</Button>
				<Button primary={false} large={true}>Share</Button>
			</div>

			<div class="recent_donors">
				<p class="title_text">43 people have donated</p>
				<ul class="donor_list">
					{#each campaign.donors.slice(0, 3) as donor}
						<li class="donor_item">
							<div class="donor_avatar">
								{#if donor.avatar}
									<img src={donor.avatar} alt={donor.name} />
								{:else}
									<div class="avatar_placeholder">{donor.name.charAt(0)}</div>
								{/if}
							</div>

							<div class="donor_info">
								<strong>{donor.name}</strong>
								<div class="donor_details">
									<span class="amount">{formatCurrency(donor.amount)}</span>
									<span class="time">{donor.timeAgo}</span>
								</div>
							</div>
						</li>
					{/each}
				</ul>

				{#if campaign.totalBackers > campaign.donors.length}
					<Button neutral={true} primary={false}>See all {campaign.totalBackers} backers</Button>
				{/if}
			</div>

			<div class="transparency-note">
				<p>Your contribution supports a vetted and verified campaign.</p>
				<span class="icon-lock"></span>
			</div>
		</aside>
	</div>
</main>

<style>
	.campaign-header {
		margin-bottom: var(--spacing-lg);
	}

	:global(.splide__arrow) {
		background: #fff;
		box-shadow: var(--shadow-md);
		transition: all 0.2s ease;
		width: 2.5rem;
		height: 2.5rem;
	}

	:global(.splide__arrow:hover) {
		background-color: var(--primary-color);
		box-shadow: var(--shadow-lg);
	}

	:global(.splide__arrow svg) {
		fill: var(--primary-color);
	}

	:global(.splide__arrow:hover svg) {
		fill: white;
	}

	.image-gallery {
		border-radius: var(--radius-md);
		overflow: hidden;
		height: 25rem;
		margin-bottom: var(--spacing-md);
		box-shadow: var(--shadow-md);
	}

	.gallery_image {
		width: 100%;
		height: 25rem;
		object-fit: cover;
	}

	.campaign-title {
		padding: 1.25rem 0;
	}

	.campaign-title h1 {
		font-size: 2.5rem;
		margin-bottom: var(--spacing-xs);
	}

	.creator-info {
		display: flex;
		align-items: center;
		gap: var(--spacing-smr);
	}

	.creator-info img {
		width: 2.5rem;
		height: 2.5rem;
		border-radius: 50%;
		object-fit: cover;
	}

	.campaign-content {
		display: grid;
		grid-template-columns: 1fr 22rem;
		gap: var(--spacing-lg);
	}

	.campaign-details {
		background: white;
		border-radius: 8px;
		box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
		overflow: hidden;
	}

	.campaign-tabs {
		display: flex;
		border-bottom: 1px solid #eee;
	}

	.campaign-tabs button {
		padding: 15px 20px;
		background: none;
		border: none;
		cursor: pointer;
		font-weight: 500;
		color: #666;
		border-bottom: 2px solid transparent;
		transition: all 0.2s;
	}

	.campaign-tabs button:hover {
		color: #34a273;
	}

	.campaign-tabs button.active {
		color: #34a273;
		border-bottom-color: #34a273;
	}

	.tab-content {
		padding: 30px;
	}

	.campaign-summary {
		font-size: 1.1rem;
		line-height: 1.8;
		color: #555;
		margin-bottom: var(--spacing-lg);
		font-weight: 500;
		padding-bottom: var(--spacing-sm);
		border-bottom: 1px solid #eee;
		font-style: italic;
	}

	.campaign-story {
		font-size: 1rem;
		line-height: 1.7;
		color: #333;
		padding-top: 15px;
	}

	.campaign-story :global(p) {
		margin-bottom: var(--spacing-md);
	}

	.updates-tab h2 {
		margin-bottom: 25px;
		color: #333;
	}

	.updates-list {
		display: flex;
		flex-direction: column;
		gap: 20px;
	}

	.update-item {
		border-bottom: 1px solid #eee;
		padding-bottom: 20px;
	}

	.update-date {
		color: var(--text-light);
		font-size: 0.9rem;
		margin-bottom: 8px;
	}

	.update-content h3 {
		margin-bottom: 8px;
		color: #444;
	}

	.no-updates {
		color: var(--text-medium);
		font-style: italic;
	}

	.campaign_aside {
		display: flex;
		flex-direction: column;
		gap: var(--spacing-md);
	}

	.amount_raised {
		font-size: 2.5rem;
		color: var(--primary-color);
		font-weight: 500;
	}

	.funding_amount p {
		color: var(--text-medium);
		margin-bottom: 20px;
	}

	.progress_bar {
		height: 6px;
		background: #e0e0e0;
		border-radius: 3px;
		margin-bottom: 10px;
		overflow: hidden;
	}

	.progress_fill {
		height: 100%;
		background: var(--primary-light);
	}

	.backer_count {
		display: flex;
		align-items: baseline;
		gap: var(--spacing-xs);
	}

	.backer_count h3 {
		font-size: 1.5rem;
		font-weight: 700;
		color: var(--text-dark);
		margin: 0;
	}

	.backer_count p {
		color: var(--text-medium);
		margin: 0;
	}

	.action_buttons {
		display: flex;
		flex-direction: column;
		gap: var(--spacing-sm);
		margin-bottom: var(--spacing-lg);
	}

	/* --- Recent Donors --- */
	.recent_donors {
		display: flex;
		flex-direction: column;
		border-top: 1px solid rgba(0, 0, 0, 0.1);
		padding-top: var(--spacing-md);
		margin-bottom: var(--spacing-lg);
	}

	.title_text {
		font-size: 1.1rem;
		color: var(--text-dark);
		margin: 0 0 var(--spacing-sm);
	}

	.donor_list {
		list-style: none;
		padding: 0;
		margin: 0 0 var(--spacing-sm);
	}

	.donor_item {
		display: flex;
		align-items: center;
		gap: var(--spacing-sm);
		padding: var(--spacing-sm) 0;
		border-bottom: 1px solid rgba(0, 0, 0, 0.05);
	}

	.donor_item:last-child {
		border-bottom: none;
	}

	.donor_avatar {
		width: 40px;
		height: 40px;
		border-radius: 50%;
		overflow: hidden;
		flex-shrink: 0;
	}

	.donor_avatar img {
		width: 100%;
		height: 100%;
		object-fit: cover;
	}

	.avatar_placeholder {
		width: 100%;
		height: 100%;
		display: flex;
		align-items: center;
		justify-content: center;
		background-color: var(--accent-light);
		color: var(--text-white);
		font-weight: 600;
		font-size: 1.2rem;
	}

	.donor_info {
		flex: 1;
	}

	.donor_info strong {
		display: block;
		font-size: 0.95rem;
		color: var(--text-dark);
		margin-bottom: 2px;
	}

	.donor_details {
		display: flex;
		justify-content: space-between;
		font-size: 0.85rem;
	}

	.amount {
		color: var(--primary-dark);
		font-weight: 600;
	}

	.time {
		color: var(--text-light);
	}

	.transparency-note {
		text-align: center;
		font-size: 0.85rem;
		color: var(--text-light);
		padding: 10px 0;
		display: flex;
		justify-content: center;
		align-items: center;
		gap: 8px;
	}

	/* Responsive Design */
	@media (max-width: 900px) {
		.campaign-content {
			grid-template-columns: 1fr;
		}

		.image-gallery {
			height: 300px;
		}
	}

	@media (max-width: 600px) {
		.campaign-title h1 {
			font-size: 2rem;
		}

		.tab-content {
			padding: 20px;
		}
	}
</style>
