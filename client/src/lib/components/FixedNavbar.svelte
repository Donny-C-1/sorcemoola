<script>
	import { page } from "$app/state";
	import { slide } from "svelte/transition";
	import { spring } from "svelte/motion";
	import Button from "./ui/Button.svelte";

	let { isLoggedIn } = $props();

	let isMenuOpen = $state(false);

	function toggleMenu() {
		isMenuOpen = !isMenuOpen;
	}

	function closeMenu() {
		isMenuOpen = false;
	}
</script>

<nav class="navbar">
	<div class="container">
		<div class="logo">
			<a href="/"><span class="logo-text">Sorce<span class="highlight">Moola</span></span></a>
		</div>

		<button class="nav-toggle" aria-label="toggle nav" aria-expanded={isMenuOpen} tabindex="0" onclick={toggleMenu}>
			<span class:first={isMenuOpen}></span>
			<span class:middle={isMenuOpen}></span>
			<span class:last={isMenuOpen}></span>
		</button>

		<div class="nav-wrapper" class:open={isMenuOpen}>
			<div class="nav-links">
				<a href="/about" class:active={page.url.pathname === "/about"}>About</a>
				<a href="/explore" class="active" class:active={page.url.pathname === "/explore"}>Explore</a>
				<a href="/start" class:active={page.url.pathname === "/start"}>Start Campaign</a>
			</div>

			{#if isLoggedIn}
				<Button href="/dashboard">Profile</Button>
			{:else}
				<div class="auth-buttons">
					<a href="/login" class="login-btn">Login</a>
					<Button primary={true} href="/signup">Sign Up</Button>
				</div>
			{/if}
		</div>
	</div>
</nav>

<style>
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
		position: relative;
	}

	.logo {
		font-size: 1.8rem;
		font-weight: 700;
	}

	.logo-text {
		letter-spacing: -0.5px;
	}

	.highlight {
		color: #3cb983;
	}

	.nav-wrapper {
		display: flex;
		align-items: center;
		gap: 2rem;
	}

	.nav-links {
		display: flex;
		align-items: center;
		gap: 2rem;
		position: absolute;
		left: 50%;
		translate: -50%;
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

	.nav-toggle {
		display: none;
		flex-direction: column;
		gap: 5px;
		z-index: 1001;
		padding: 0.5rem;
		cursor: pointer;
	}

	.nav-toggle span {
		display: block;
		width: 25px;
		height: 3px;
		background-color: #333;
		border-radius: 3px;
		transition: 0.3s ease-in-out;
	}

	.nav-toggle .first {
		transform: translateY(8px) rotate(45deg);
	}

	.nav-toggle .middle {
		opacity: 0;
	}

	.nav-toggle .last {
		transform: translateY(-8px) rotate(-45deg);
	}

	@media (max-width: 54rem) {
		.nav-wrapper {
			position: fixed;
			top: 0;
			right: -100%;
			width: 80%;
			max-width: 320px;
			height: 100vh;
			flex-direction: column;
			justify-content: flex-start;
			align-items: flex-start;
			background-color: white;
			box-shadow: -2px 0 10px rgba(0, 0, 0, 0.1);
			padding: 6rem 2rem 2rem;
			transition: right 0.3s ease-in-out;
			z-index: 1000;
		}

		.nav-wrapper.open {
			right: 0;
		}

		.nav-links {
			position: static;
			left: unset;
			translate: unset;
			flex-direction: column;
			width: 100%;
			margin-bottom: 2rem;
		}

		.nav-links a {
			width: 100%;
			padding: 1rem 0;
			border-bottom: 1px solid #eee;
		}

		.auth-buttons {
			flex-direction: column;
			width: 100%;
			gap: 1rem;
		}

		.auth-buttons a {
			width: 100%;
			text-align: center;
		}

		.nav-toggle {
			display: flex;
		}

		.nav-wrapper.open::before {
			content: "";
			position: fixed;
			top: 0;
			left: 0;
			right: 0;
			bottom: 0;
			background: rgba(0, 0, 0, 0.5);
			z-index: -1;
		}
	}
</style>
