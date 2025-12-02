<script>
	import { onMount } from 'svelte';
	import PlaceholderCampaigns from '$lib/components/PlaceholderCampaigns.svelte';
	import ErrorComponent from '$lib/components/ErrorComponent.svelte';

	// State management
	let searchQuery = $state('');
	let activeCategory = $state('All Campaigns');
	let loading = $state(true);
	let error = $state(null);
	let campaigns = $state([]);

	let categories = [
		'All Campaigns',
		'Technology',
		'Creative',
		'Business',
		'Education',
		'Community',
		'Health'
	];

	let displayedCampaigns = $derived(campaigns.slice());

	// Filter campaigns based on category and search query
	function filterCampaigns() {
		if (activeCategory === 'All Campaigns' && !searchQuery) {
			displayedCampaigns = [...campaigns];
			return;
		}

		displayedCampaigns = campaigns.filter((campaign) => {
			const matchesCategory =
				activeCategory === 'All Campaigns' || campaign.category === activeCategory;
			const matchesSearch =
				!searchQuery ||
				campaign.title.toLowerCase().includes(searchQuery.toLowerCase()) ||
				campaign.description.toLowerCase().includes(searchQuery.toLowerCase());

			return matchesCategory && matchesSearch;
		});
	}

	// Format currency
	function formatCurrency(amount) {
		return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(amount);
	}

	// Handle category selection
	function selectCategory(category) {
		activeCategory = category;
		filterCampaigns();
	}

	// Handle search input
	function handleSearch() {
		filterCampaigns();
	}

	// Add smooth scroll behavior on mount
	onMount(() => {
		document.documentElement.style.scrollBehavior = 'smooth';

		fetchCampaigns();
	});

	async function fetchCampaigns() {
		loading = true;
		error = null;

		try {
			// Simulate API call with setTimeout
			const response = await fetch('/mock/campaigns.json');

			if (!response.ok) {
				throw new Error(`HTTP error! status: ${response.status}`);
			}

			const data = await response.json();
			console.log(data);
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
	<meta
		name="description"
		content="Discover and fund innovative campaigns on SorceMoola, the premier crowdfunding platform."
	/>
</svelte:head>

<main>
	<!-- Navigation Bar -->
	<nav class="navbar">
		<div class="container">
			<div class="logo">
				<a href="/">
					<span class="logo-text">Sorce<span class="highlight">Moola</span></span>
				</a>
			</div>

			<button class="nav-toggle" aria-label="toggle nav" tabindex="0">
				<span></span>
				<span></span>
				<span></span>
			</button>

			<div class="nav-links">
				<a href="/">Home</a>
				<a href="/explore" class="active">Explore</a>
				<a href="/start-campaign">Start Campaign</a>
				<a href="/about">About</a>
			</div>

			<div class="auth-buttons">
				<a href="/login" class="login-btn">Login</a>
				<a href="/signup" class="signup-btn">Sign Up</a>
			</div>
		</div>
	</nav>

	<!-- Hero Section with Search -->
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

	<!-- Campaign Cards Section -->
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

	<!-- Footer -->
	<footer>
		<div class="container">
			<div class="footer-grid">
				<div class="footer-about">
					<div class="footer-logo">
						<span class="logo-text">Sorce<span class="highlight">Moola</span></span>
					</div>
					<p class="tagline">Empowering innovators. Funding the future.</p>
					<div class="social-icons">
						<a href="./#" aria-label="Facebook"><i class="fab fa-facebook-f"></i></a>
						<a href="./#" aria-label="Twitter"><i class="fab fa-twitter"></i></a>
						<a href="./#" aria-label="Instagram"><i class="fab fa-instagram"></i></a>
						<a href="./#" aria-label="LinkedIn"><i class="fab fa-linkedin-in"></i></a>
					</div>
				</div>

				<div class="footer-links">
					<h3>Quick Links</h3>
					<ul>
						<li><a href="/about">About Us</a></li>
						<li><a href="/terms">Terms of Service</a></li>
						<li><a href="/privacy">Privacy Policy</a></li>
						<li><a href="/help">Help & Support</a></li>
					</ul>
				</div>

				<div class="footer-newsletter">
					<h3>Stay Updated</h3>
					<p>Subscribe to our newsletter for the latest campaigns and updates</p>
					<div class="newsletter-form">
						<input type="email" placeholder="Your email address" />
						<button type="submit">Subscribe</button>
					</div>
				</div>
			</div>

			<div class="footer-bottom">
				<p>© {new Date().getFullYear()} SorceMoola. All rights reserved.</p>
			</div>
		</div>
	</footer>
</main>

<style>

	.highlight {
		color: #3cb983;
	}

	/* Navigation Bar */
	.navbar {
		position: sticky;
		top: 0;
		background: white;
		box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
		padding: 1rem 0;
		z-index: 1000;
	}

	.navbar .container {
		display: flex;
		align-items: center;
		justify-content: space-between;
	}

	.logo {
		font-size: 1.8rem;
		font-weight: 700;
	}

	.logo-text {
		letter-spacing: -0.5px;
	}

	.nav-links {
		display: flex;
		gap: 2rem;
	}

	.nav-links a {
		font-weight: 500;
		padding-bottom: 5px;
		border-bottom: 2px solid transparent;
	}

	.nav-links a:hover,
	.nav-links a.active {
		color: #3cb983;
		border-bottom: 2px solid #3cb983;
	}

	.auth-buttons {
		display: flex;
		gap: 1rem;
	}

	.login-btn {
		padding: 0.5rem 1rem;
		border-radius: 4px;
		font-weight: 500;
	}

	.login-btn:hover {
		color: #3cb983;
	}

	.signup-btn {
		background: #3cb983;
		color: white;
		padding: 0.5rem 1.5rem;
		border-radius: 4px;
		font-weight: 500;
		transition: background-color 0.3s ease;
	}

	.signup-btn:hover {
		background: #34a273;
	}

	.nav-toggle {
		display: none;
		flex-direction: column;
		gap: 5px;
		cursor: pointer;
	}

	.nav-toggle span {
		display: block;
		width: 25px;
		height: 3px;
		background-color: #333;
		border-radius: 3px;
		transition: transform 0.3s ease;
	}

	/* Hero Section with Search */
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

	/* Campaign Cards Section */
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

	/* Footer */
	footer {
		background: #272a33;
		color: white;
		padding: 4rem 0 1.5rem;
	}

	.footer-grid {
		display: grid;
		grid-template-columns: 2fr 1fr 2fr;
		gap: 3rem;
		margin-bottom: 3rem;
	}

	.footer-about {
		display: flex;
		flex-direction: column;
		gap: 1rem;
	}

	.footer-logo {
		font-size: 1.5rem;
		font-weight: 700;
	}

	.tagline {
		color: #ccc;
	}

	.social-icons {
		display: flex;
		gap: 1rem;
		margin-top: 1rem;
	}

	.social-icons a {
		display: flex;
		align-items: center;
		justify-content: center;
		width: 36px;
		height: 36px;
		border-radius: 50%;
		background: rgba(255, 255, 255, 0.1);
		transition: background 0.3s ease;
	}

	.social-icons a:hover {
		background: #3cb983;
	}

	.footer-links h3,
	.footer-newsletter h3 {
		font-size: 1.2rem;
		margin-bottom: 1.5rem;
	}

	.footer-links ul {
		list-style: none;
		padding: 0;
		margin: 0;
	}

	.footer-links li {
		margin-bottom: 0.75rem;
	}

	.footer-links a {
		color: #ccc;
		transition: color 0.3s ease;
	}

	.footer-links a:hover {
		color: #3cb983;
	}

	.newsletter-form {
		display: flex;
		margin-top: 1rem;
	}

	.newsletter-form input {
		flex-grow: 1;
		padding: 0.75rem 1rem;
		border: none;
		border-radius: 4px 0 0 4px;
		outline: none;
	}

	.newsletter-form button {
		background: #3cb983;
		color: white;
		border: none;
		padding: 0.75rem 1.25rem;
		border-radius: 0 4px 4px 0;
		font-weight: 500;
		transition: background 0.3s ease;
	}

	.newsletter-form button:hover {
		background: #34a273;
	}

	.footer-bottom {
		text-align: center;
		padding-top: 1.5rem;
		border-top: 1px solid rgba(255, 255, 255, 0.1);
		color: #aaa;
		font-size: 0.9rem;
	}

	/* Responsive Styles */
	@media (max-width: 1024px) {
		.campaign-grid {
			grid-template-columns: repeat(2, 1fr);
		}
	}

	@media (max-width: 768px) {
		.nav-links,
		.auth-buttons {
			display: none;
		}

		.nav-toggle {
			display: flex;
		}

		.footer-grid {
			grid-template-columns: 1fr;
			gap: 2rem;
		}

		.social-icons {
			justify-content: center;
		}

		.footer-links {
			text-align: center;
		}

		.footer-links ul {
			display: flex;
			flex-wrap: wrap;
			justify-content: center;
			gap: 1rem;
		}

		.footer-newsletter {
			text-align: center;
		}

		/* When nav is opened */
		.navbar.open .nav-links,
		.navbar.open .auth-buttons {
			display: flex;
			flex-direction: column;
			width: 100%;
		}

		.navbar.open .container {
			flex-wrap: wrap;
		}

		.navbar.open .nav-links {
			order: 3;
			margin-top: 1rem;
		}

		.navbar.open .auth-buttons {
			order: 4;
			margin-top: 1rem;
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
