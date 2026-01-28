<script>
	import { onMount } from "svelte";
	import { fade, fly, slide } from "svelte/transition";
	import { PUBLIC_SERVER_URL } from "$env/static/public";
	import { Splide, SplideSlide } from "@splidejs/svelte-splide";
	import "@splidejs/svelte-splide/css";
	import Button from "$lib/components/ui/Button.svelte";
	import { enhance } from "$app/forms";
	import Snackbar from "$lib/components/ui/Snackbar.svelte";
	import { invalidateAll } from "$app/navigation";
	import { formatCurrency } from "$lib/utils/formatters.js";
	import { formatTimestamp } from "little-timestamp";

	let { data, form } = $props();

	const PRESET_AMOUNTS = [10000, 50000, 200000];
	const STEP_AMOUNT = 100;
	const MIN_AMOUNT = 100;
	const MAX_AMOUNT = 1000000;

	let campaign = {
		creatorAvatar: "https://images.unsplash.com/photo-1760574772950-f37de9dce85c?ixid=M3w4MjcwNjd8MHwxfHNlYXJjaHwzfHx3aXphcmQlMjB3aWxsb3d8ZW58MHx8fHwxNzYzNjYxNTA4fDA&ixlib=rb-4.1.0&fit=max&q=80",
		daysLeft: 15,
		updates: [
			{
				date: "2023-10-25",
				title: "Legal progress!",
				content: "We've secured a temporary injunction to halt development while our case is being reviewed."
			},
			{
				date: "2023-10-10",
				title: "25% funded!",
				content: "Thank you to all our amazing supporters. We're a quarter of the way there!"
			}
		]
	};

	let options = {
		perPage: 1,
		type: "loop",
		gap: "1rem",
		autoplay: "true"
	};

	let pledgeAmount = $state(50000);
	let activeTab = $state("story");
	let isSnackbarVisible = $state(false);
	let isModalOpen = $state(false);

	const getProgressPercentage = () => {
		return Math.min(Math.floor((campaign.currentAmount / campaign.targetAmount) * 100), 100);
	};

	function toggleModal() {
		isModalOpen = !isModalOpen;
	}

	function handleInputChange() {
		if (pledgeAmount < MIN_AMOUNT) pledgeAmount = MIN_AMOUNT;
		if (pledgeAmount > MAX_AMOUNT) pledgeAmount = MAX_AMOUNT;
	}

	async function handlePayment() {
		return async ({ result, update }) => {
			const { default: PaystackPop } = await import("@paystack/inline-js");

			const popup = new PaystackPop();
			popup.resumeTransaction(result.data.access_code, {
				onSuccess: async(transaction) => {
					console.log("Transaction", transaction);

					setTimeout(async() => {
						await invalidateAll();
					}, 2000)
				}
			})

			isModalOpen = false;
		}
	}

	$effect(() => {
		if (form?.success) {
			isSnackbarVisible = true;
		}
	});
</script>

<svelte:head>
	<title>{data.campaign.name} | Sorcemoola</title>
	<meta name="description" content={data.campaign.description} />
</svelte:head>

<main class="container">
	<section class="campaign-header">
		<div class="image-gallery">
			<Splide {options}>
				{#each [{ mobile: "/images/share-mobile.png", desktop: "/images/share.png" }] as image, i}
					<SplideSlide>
						<picture>
							<source media="(max-width: 650px)" srcset={image.mobile} />
							<img class="gallery_image" src={image.desktop} alt={`${campaign.title} - image ${i + 1}`} width="100%" height="auto" />
						</picture>
					</SplideSlide>
				{/each}
			</Splide>
		</div>

		<div class="campaign-title">
			<h1>{data.campaign.name}</h1>
			<div class="creator-info">
				<img src={campaign.creatorAvatar} alt={campaign.creator} />
				<p>by <a href="../p/{data.campaign.creator.id}"><strong>{data.campaign.creator.name}</strong></a></p>
			</div>
		</div>
	</section>

	<div class="campaign-content">
		<div class="campaign-details">
			<div class="campaign-tabs">
				<button class:active={activeTab === "story"} onclick={() => (activeTab = "story")}> Campaign Story </button>
				<button class:active={activeTab === "updates"} onclick={() => (activeTab = "updates")}>
					Updates ({campaign.updates.length})
				</button>
			</div>

			<div class="tab-content">
				{#if activeTab === "story"}
					<section class="story-tab" transition:fade={{ duration: 200 }}>
						<blockquote class="campaign-summary">{data.campaign.description}</blockquote>
						<div class="campaign-story">
							{#each data.campaign.story.split("\n") as paragraph}
								{#if paragraph.trim() !== ""}
									<p>{paragraph}</p>
								{/if}
							{/each}
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
					<h2 class="amount_raised">{formatCurrency(data.campaign.amountRaised / 100)}</h2>
					<p>raised of {formatCurrency(data.campaign.fundGoal / 100)}</p>
				</div>

				<div class="progress_container">
					<div class="progress_bar">
						<div class="progress_fill" style:width="{getProgressPercentage()}%"></div>
					</div>
				</div>

				<div class="backer_count">
					<h3>{data.campaign.backersCount}</h3>
					<p>backers</p>
				</div>
			</div>

			<div class="action_buttons">
				<form
					method="post"
					action="?/initiatePayment"
					use:enhance={() =>
						async ({ result, update }) => {
							const popup = new PaystackPop();
							popup.resumeTransaction(result.data.access_code, {
								onSuccess: async (transaction) => {
									console.log("Transaction", transaction);

									setTimeout(async () => {
										await invalidateAll();
									}, 2000);
								}
							});
							console.log(result);
						}}
				>
					<input type="hidden" name="amount" bind:value={pledgeAmount} />
				</form>
				<Button large={true} handler={toggleModal}>Dontate Now</Button>
				<Button primary={false} large={true}>Share</Button>
			</div>

			{#if data.campaign.backersCount > 0}
				<div class="recent_donors">
					<p class="title_text">{data.campaign.backersCount} people have contributed</p>
					<ul class="donor_list">
						{#each data.campaign.recentContributions as donor}
							<li class="donor_item">
								<div class="donor_avatar">
									{#if donor.avatar}
										<img src={donor.avatar} alt={donor.name} />
									{:else}
										<div class="avatar_placeholder">{donor.User.name.charAt(0)}</div>
									{/if}
								</div>

								<div class="donor_info">
									<strong>{donor.User.name}</strong>
									<div class="donor_details">
										<span class="amount">{formatCurrency(donor.amount / 100)}</span>
										<span class="time">{formatTimestamp(new Date(donor.createdAt))}</span>
									</div>
								</div>
							</li>
						{/each}
					</ul>

					{#if data.campaign.backersCount > data.campaign.recentContributions.length}
						<Button neutral={true} primary={false}>See all {data.campaign.backersCount} backers</Button>
					{/if}
				</div>
			{/if}

			<div class="transparency-note">
				<p>Your contribution supports a vetted and verified campaign.</p>
				<span class="icon-lock"></span>
			</div>
		</aside>
	</div>

	<Snackbar bind:visible={isSnackbarVisible} type="info">Donation successful. <span class="amount">₦{pledgeAmount}</span></Snackbar>

	{#if isModalOpen}
		<div class="modal_backdrop" transition:fade>
			<form method="post" action="?/initiatePayment" use:enhance={handlePayment} class="modal_container" in:slide>
				<div class="modal_header">
					<h2>Back this Project</h2>
					<button class="close_btn" type="button" onclick={toggleModal}>&times;</button>
				</div>

				<div class="modal_body">
					<p class="label">Select an amount</p>
					<div class="preset_grid">
						{#each PRESET_AMOUNTS as preset}
							<button class="preset_btn" type="button" class:active={pledgeAmount === preset} onclick={_ => pledgeAmount = preset}>{formatCurrency(preset)}</button>
						{/each}
					</div>

					<p class="label label_custom">Or enter custom amount</p>
					<div class="input_wrapper">
						<span class="currency_symbol">₦</span>
						<input type="number" name="amount" class="amount_input" bind:value={pledgeAmount} step={STEP_AMOUNT} min={MIN_AMOUNT} max={MAX_AMOUNT} onchange={handleInputChange} />
					</div>

					<div class="slider_wrapper">
						<input type="range" name="amount" id="range_amount" class="range_slider" bind:value={pledgeAmount} min={MIN_AMOUNT} max={MAX_AMOUNT} step={STEP_AMOUNT} style:--progres={(pledgeAmount / MAX_AMOUNT) * 100}%>
						<div class="slider_labels">
							<span>{formatCurrency(MIN_AMOUNT)}</span>
							<span>{formatCurrency(MAX_AMOUNT)}</span>
						</div>
					</div>
					<p class="min_text">Minimum donation is {formatCurrency(MIN_AMOUNT)}</p>

					<Button wide={true} large={true} type="submit">Pay {formatCurrency(pledgeAmount)}</Button>
				</div>
			</form>
		</div>
	{/if}
</main>

<style>
	.campaign-header {
		margin-bottom: var(--spacing-lg);
		padding-bottom: 0;
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
		height: 27rem;
		margin-bottom: var(--spacing-md);
		box-shadow: var(--shadow-md);
	}

	.gallery_image {
		width: 100%;
		height: 27rem;
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

	.amount {
		color: var(--secondary-color);
	}

	.modal_backdrop {
		position: fixed;
		inset: 0;
		background-color: rgba(0, 0, 0, .5);
		display: flex;
		justify-content: center;
		align-items: center;
		z-index: 5;
		backdrop-filter: blur(2px);
	}

	.modal_container {
		background-color: var(--text-white);
		border-radius: var(--radius-lg);
		width: 90%;
		max-width: 32rem;
		box-shadow: 0 10px 15px -3px rgba(0, 0, 0, .1), 0 4px 6px -2px rgba(0, 0, 0, .05);
		overflow: hidden;
	}

	.modal_header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 1.5rem;
		border-bottom: 1px solid var(--neutral-color);
	}

	.modal_header h2 {
		font-size: 1.5rem;
		font-weight: 500;
		margin: 0;
	}

	.close_btn {
		background-color: transparent;
		border: 0;
		font-size: 2rem;
		line-height: 1;
		color: var(--neutral-color);
	}

	.modal_body {
		padding: var(--spacing-md);
	}

	.label {
		font-weight: 600;
		margin-bottom: var(--spacing-sm);
		display: block;
	}

	.label_custom {
		margin-top: var(--spacing-md);
	}

	.preset_grid {
		display: grid;
		grid-template-columns: repeat(3, 1fr);
		gap: 1rem;
	}

	.preset_btn {
		padding: var(--spacing-smr);
		border: 2px solid var(--neutral-color);
		background-color: var(--text-white);
		border-radius: var(--radius-md);
		font-size: 1rem;
		font-weight: 600;
		transition: .2s ease;
	}

	.preset_btn:is(:hover, :focus-visible, :active) {
		border-color: var(--primary-color);
		color: var(--primary-color);
	}

	.preset_btn.active {
		background-color: var(--primary-color);
		border-color: var(--primary-color);
		color: var(--text-white);
	}

	.input_wrapper {
		position: relative;
		display: flex;
		align-items: center;
		border: 2px solid var(--neutral-color);
		border-radius: var(--radius-md);
		padding: var(--spacing-smr);
		margin-bottom: var(--spacing-md);
		transition: border-color .2s;
	}

	.input_wrapper:focus-within {
		border-color: var(--primary-color)
	}

	.currency_symbol {
		font-size: 1.5rem;
		font-weight: 600;
		color: #6b7280;
		margin-right: .5rem;
	}

	.amount_input {
		border: none;
		font-size: 1.5rem;
		font-weight: 600;
		width: 100%;
		outline: none;
	}

	.slider_wrapper {
		margin-bottom: var(--spacing-xs);
	}

	.range_slider {
		width: 100%;
		cursor: pointer;
		accent-color: var(--primary-color);
		height: 6px;
		background-color: var(--neutral-color);
		border-radius: var(--radius-sm);
		outline: none;
	}

	.slider_labels {
		display: flex;
		justify-content: space-between;
		font-size: .875rem;
		color: var(--neutral-dark);
		margin-top: .5rem;
	}

	.min_text {
		font-size: .875rem;
		text-align: center;
		color: var(--neutral-dark);
		margin-bottom: var(--spacing-md);
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
