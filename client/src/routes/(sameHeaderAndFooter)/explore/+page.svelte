<script>
    import PlaceholderCampaigns from "$lib/components/PlaceholderCampaigns.svelte";
    import ErrorComponent from "$lib/components/ErrorComponent.svelte";
	import { onMount } from "svelte";

    let searchQuery = $state("");
    let activeCategory = $state("All Campaigns");
    let loading = $state(true);
    let error = $state(null);
    let campaigns = $state([]);

    let categories = [
        "All Campaigns",
        "Technology",
        "Creative",
        "Business",
        "Education",
        "Community",
        "Health"
    ];

    let displayedCampaigns = $derived(campaigns.slice());

    onMount(() => {
        fetchCampaigns();
    })

    function formatCurrency(amount) {
		return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(amount);
	}

    function filterCampaigns() {
        if (activeCategory === "All Campaigns" && !searchQuery) {
            displayedCampaigns = [...campaigns];
            return;
        }

        displayedCampaigns = campaigns.filter(campaign => {
            const matchesCategory = activeCategory === "All Campaigns" || campaign.category === activeCategory;
            const matchesSearch = !searchQuery || campaign.title.toLowerCase().includes(searchQuery.toLowerCase()) || campaign.description.toLowerCase().includes(searchQuery.toLowerCase());
            return matchesCategory && matchesSearch;
        })
    }

    function handleSearch() {
        filterCampaigns();
    }

    function selectCategory(category) {
        activeCategory = category;
        filterCampaigns();
    }

    async function fetchCampaigns() {
		loading = true;
		error = null;

		try {
			const response = await fetch('/mock/campaigns.json');

			if (!response.ok) {
				throw new Error(`HTTP error! status: ${response.status}`);
			}

			const data = await response.json();
			campaigns = data;
		} catch (err) {
			error = 'Failed to load campaigns. Please try again.';
		} finally {
			loading = false;
		}
	}
</script>

<svelte:head>
    <title>Explore Campaigns | SorceMoola</title>
    <meta name="description" content="Explore the latest crowdfunding campaigns on SorceMoola. Discover innovative projects and support the creators you love."/>
</svelte:head>

<main>
    <section class="hero-section">
		<div class="container">
			<h1>Discover Amazing Campaigns</h1>
			<p>Find innovative projects to back and help bring creative ideas to life</p>

			<div class="search-container">
				<div class="search-bar">
					<i class="fas fa-search"></i>
					<input
						type="text"
						placeholder="Search for campaigns..."
						bind:value={searchQuery}
						oninput={handleSearch}
					/>
				</div>

				<div class="categories">
					{#each categories as category}
						<button
							class="category-btn {activeCategory === category ? 'active' : ''}"
							onclick={() => selectCategory(category)}
						>
							{category}
						</button>
					{/each}
				</div>
			</div>
		</div>
	</section>

    <section class="campaigns-section">
		<div class="container">
			{#if loading}
				<PlaceholderCampaigns count={6} layout="grid" />
			{:else if error}
				<ErrorComponent
					message="Couldn't load campaigns"
					subtext="We're having trouble connecting to our servers. Please check your connection and try again."
					icon="wifi-off"
					retryFunction={fetchCampaigns}
				/>
			{:else if displayedCampaigns.length > 0}
				<div class="campaign-grid">
					{#each displayedCampaigns as campaign (campaign.id)}
						<div class="campaign-card">
							<div class="card-image">
								<img src={campaign.image} alt={campaign.title} />
								<div class="category-tag">{campaign.category}</div>
							</div>

							<div class="card-content">
								<h3 class="card-title">{campaign.title}</h3>
								<p class="card-description">{campaign.description}</p>

								<div class="progress-container">
									<div class="progress-bar">
										<div class="progress" style="width: {campaign.percentFunded}%"></div>
									</div>
									<div class="progress-stats">
										<span>{formatCurrency(campaign.currentAmount)}</span>
										<span>{campaign.percentFunded}%</span>
									</div>
									<div class="campaign-details">
										<div class="detail">
											<span class="label">Goal</span>
											<span class="value">{formatCurrency(campaign.goalAmount)}</span>
										</div>
										<div class="detail">
											<span class="label">Days Left</span>
											<span class="value">{campaign.daysRemaining}</span>
										</div>
									</div>
								</div>

								<button class="fund-button">Fund this project</button>
							</div>
						</div>
					{/each}
				</div>
			{:else}
				<div class="no-results">
					<i class="fas fa-search fa-3x"></i>
					<h2>No campaigns found</h2>
					<p>Try adjusting your search or category selection</p>
				</div>
			{/if}
		</div>
	</section>
</main>

<style>
    .hero-section {
		background-color: #f8f9fa;
		padding: 3rem 0;
		text-align: center;
	}

	.hero-section h1 {
		font-size: 2.5rem;
		margin-bottom: 1rem;
	}

	.hero-section p {
		font-size: 1.1rem;
		color: #666;
		margin-bottom: 2rem;
		max-width: 700px;
		margin-left: auto;
		margin-right: auto;
	}

	.search-container {
		max-width: 800px;
		margin: 0 auto;
	}

	.search-bar {
		position: relative;
		margin-bottom: 1.5rem;
	}

	.search-bar input {
		width: 100%;
		padding: 1rem 1rem 1rem 3rem;
		font-size: 1.1rem;
		border: 2px solid #e1e1e1;
		border-radius: 8px;
		transition: all 0.3s ease;
	}

	.search-bar input:focus {
		outline: none;
		border-color: #3cb983;
		box-shadow: 0 0 0 3px rgba(60, 185, 131, 0.2);
	}

	.search-bar i {
		position: absolute;
		left: 1rem;
		top: 50%;
		transform: translateY(-50%);
		color: #888;
	}

	.categories {
		display: flex;
		flex-wrap: wrap;
		gap: 0.75rem;
		justify-content: center;
		margin-top: 1.5rem;
	}

	.category-btn {
		background: transparent;
		border: 1px solid #ddd;
		padding: 0.5rem 1rem;
		border-radius: 20px;
		font-size: 0.95rem;
		transition: all 0.3s ease;
	}

	.category-btn:hover {
		border-color: #3cb983;
		color: #3cb983;
	}

	.category-btn.active {
		background: #3cb983;
		color: white;
		border-color: #3cb983;
	}

    .campaigns-section {
		padding: 3rem 0;
	}

	.campaign-grid {
		display: grid;
		grid-template-columns: repeat(3, 1fr);
		gap: 2rem;
	}

	.campaign-card {
		border-radius: 8px;
		overflow: hidden;
		box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
		transition:
			transform 0.3s ease,
			box-shadow 0.3s ease;
		background: white;
	}

	.campaign-card:hover {
		transform: translateY(-5px);
		box-shadow: 0 6px 20px rgba(0, 0, 0, 0.15);
	}

	.card-image {
		position: relative;
		height: 200px;
		overflow: hidden;
	}

	.card-image img {
		width: 100%;
		height: 100%;
		object-fit: cover;
		transition: transform 0.5s ease;
	}

	.campaign-card:hover .card-image img {
		transform: scale(1.05);
	}

	.category-tag {
		position: absolute;
		top: 1rem;
		right: 1rem;
		background: rgba(0, 0, 0, 0.7);
		color: white;
		padding: 0.25rem 0.75rem;
		border-radius: 20px;
		font-size: 0.8rem;
		font-weight: 500;
	}

	.card-content {
		padding: 1.5rem;
	}

	.card-title {
		font-size: 1.2rem;
		margin: 0 0 0.75rem 0;
		display: -webkit-box;
		-webkit-line-clamp: 2;
		line-clamp: 2;
		-webkit-box-orient: vertical;
		overflow: hidden;
		text-overflow: ellipsis;
		height: 3rem; /* Approx 2 lines */
	}

	.card-description {
		color: #666;
		margin-bottom: 1.25rem;
		display: -webkit-box;
		-webkit-line-clamp: 3;
		line-clamp: 3;
		-webkit-box-orient: vertical;
		overflow: hidden;
		text-overflow: ellipsis;
		height: 4.8rem; /* Approx 3 lines */
	}

	.progress-container {
		margin-bottom: 1.25rem;
	}

	.progress-bar {
		height: 6px;
		background-color: #eee;
		border-radius: 3px;
		margin-bottom: 0.5rem;
		overflow: hidden;
	}

	.progress {
		height: 100%;
		background-color: #3cb983;
		border-radius: 3px;
		transition: width 0.5s ease-in-out;
	}

	.progress-stats {
		display: flex;
		justify-content: space-between;
		font-weight: 500;
		margin-bottom: 0.75rem;
	}

	.campaign-details {
		display: flex;
		justify-content: space-between;
		margin-bottom: 1rem;
	}

	.detail {
		display: flex;
		flex-direction: column;
	}

	.label {
		font-size: 0.85rem;
		color: #888;
	}

	.value {
		font-weight: 600;
	}

	.fund-button {
		width: 100%;
		background: #3cb983;
		color: white;
		border: none;
		padding: 0.75rem 0;
		border-radius: 4px;
		font-weight: 600;
		transition: background 0.3s ease;
	}

	.fund-button:hover {
		background: #34a273;
	}

	.no-results {
		text-align: center;
		padding: 4rem 0;
		color: #666;
	}

	.no-results i {
		color: #ccc;
		margin-bottom: 1.5rem;
	}

	.no-results h2 {
		margin-bottom: 0.5rem;
	}

	@media (max-width: 1024px) {
		.campaign-grid {
			grid-template-columns: repeat(2, 1fr);
		}
	}
	
	@media (max-width: 640px) {
		.campaign-grid {
			grid-template-columns: 1fr;
		}

		.hero-section h1 {
			font-size: 2rem;
		}

		.categories {
			gap: 0.5rem;
		}

		.category-btn {
			font-size: 0.85rem;
			padding: 0.4rem 0.8rem;
		}
	}
</style>