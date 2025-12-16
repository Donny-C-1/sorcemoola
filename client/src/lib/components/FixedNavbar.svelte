<script>
	import { page } from "$app/state";
	import { fade, slide } from "svelte/transition";
	import { spring } from "svelte/motion";
	import Button from "./ui/Button.svelte";
	import { ChevronDown, ChevronRight, CircleUserRound, Globe, LayoutDashboard, LogOut, Settings, User } from "lucide-svelte";
	import { onMount } from "svelte";
	import { layout } from "$lib/stores/layout.svelte";
	import Logo from "./ui/Logo.svelte";
	import { fly } from "svelte/transition";
	import { cubicInOut } from "svelte/easing";

	let { isLoggedIn, user = { slug: "", avatar: "/images/avatar.png", username: "Ugwuenuvuenvuennoovuvuenvuenosas" } } = $props();

	let isMenuOpen = $state(false);
	let isDropdownOpen = $state(false);

	function toggleMenu() {
		isMenuOpen = !isMenuOpen;
	}

	function toggleDropdown() {
		isDropdownOpen = !isDropdownOpen;
	}

	function closeMenu() {
		isMenuOpen = false;
	}

	function logout() {
		//
	}
</script>

<svelte:window
	onclick={(e) => {
		let target = e.target;
		if (isDropdownOpen && !target.closest(".dropdown") && !target.closest(".dropdown_toggle")) {
			isDropdownOpen = false;
		}
	}}
/>

<nav class="navbar">
	<div class="container">
		<Logo />

		<button class="nav-toggle" aria-label="toggle nav" aria-expanded={isMenuOpen} tabindex="0" onclick={toggleMenu}>
			<span class:first={isMenuOpen}></span>
			<span class:middle={isMenuOpen}></span>
			<span class:last={isMenuOpen}></span>
		</button>

		{#if !layout.isDesktop && isMenuOpen}
			<div class="menu_overlay" onclick={closeMenu} role="none"></div>
		{/if}

		{#if layout.isDesktop || isMenuOpen}
			<div class="nav-wrapper" class:open={isMenuOpen} transition:fly={{ x: 200, duration: 300, easing: cubicInOut }}>
				{#if !layout.isDesktop && isLoggedIn}
					<div class="mini_profile">
						<img src={user.avatar} alt={user.username} class="img" width="100" height="auto" />
						<p class="username">{user.username}</p>
						<a href="/p/{user.slug}">
							<span>Profile</span>
							<ChevronRight size={16} />
						</a>
					</div>
				{/if}
				<div class="nav-links">
					<a href="/about" class:active={page.url.pathname === "/about"}>About</a>
					<a href="/explore" class="active" class:active={page.url.pathname === "/explore"}>Explore</a>
					<a href="/dashboard/create" class:active={page.url.pathname === "/start"}>Start Campaign</a>
				</div>

				{#if isLoggedIn}
					<div class="dropdown">
						{#if layout.isDesktop}
							<button class="profile_button" onclick={toggleDropdown}>
								<img src={user.avatar} alt={user.username} class="img" width="40" height="auto" />
								<div class="text">
									<span class="username">{user.username}</span>
									<ChevronDown size={18} />
								</div>
							</button>
						{/if}

						{#if !layout.isDesktop || isDropdownOpen}
							<div class="dropdown_content">
								<a href="/dashboard" class="dropdown_item">
									<LayoutDashboard size={18} />
									<span>Dashboard</span>
								</a>
								{#if layout.isDesktop}
									<a href="/p/{user.slug}" class="dropdown_item">
										<Globe size={18} />
										<span>Profile</span>
									</a>
								{/if}
								<a href="/account/profile" class="dropdown_item">
									<Settings size={18} />
									<span>Account Setting</span>
								</a>
								<a href="/logout" class="dropdown_item logout_btn" onclick={logout}>
									<LogOut size={18} />
									<span>Log Out</span>
								</a>
							</div>
						{/if}
					</div>
				{:else}
					<div class="auth-buttons">
						<a href="/login" class="login-btn">Login</a>
						<Button primary={true} href="/signup">Sign Up</Button>
					</div>
				{/if}
			</div>
		{/if}
	</div>
</nav>

<style>
	.navbar {
		position: sticky;
		top: 0;
		background: var(--text-white);
		box-shadow: var(--shadow-md);
		padding: var(--spacing-sm) 0;
		z-index: 5;
	}

	.navbar .container {
		display: flex;
		align-items: center;
		justify-content: space-between;
		position: relative;
	}
	
	.nav-toggle {
		display: none;
		flex-direction: column;
		gap: 5px;
		border: 0;
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

	.nav-wrapper {
		display: flex;
		align-items: center;
		gap: var(--spacing-lg);
	}

	.mini_profile {
		display: flex;
		flex-direction: column;
		align-items: center;
		width: 100%;
	}

	.mini_profile img {
		border-radius: 50%;
		margin-bottom: 0.5rem;
	}

	.mini_profile .username {
		max-width: 60%;
		overflow: hidden;
		text-overflow: ellipsis;
		font-weight: 700;
		font-size: 1.2rem;
		font-family: "quicksand";
		margin: 0;
	}

	.mini_profile a {
		display: flex;
		align-items: center;
		gap: 0.25rem;
	}

	.mini_profile a:is(:hover, :focus-visible) {
		color: var(--primary-color);
	}

	.nav-links {
		display: flex;
		align-items: center;
		gap: var(--spacing-lg);
		position: absolute;
		left: 50%;
		translate: -50%;
	}

	.nav-links a {
		font-weight: 500;
		border-bottom: 2px solid transparent;
		transition: 0.3s ease;
	}

	.nav-links a:is(:hover, :focus-visible, :active) {
		color: var(--primary-color);
		border-bottom: 2px solid var(--primary-color);
	}

	.profile_button {
		display: flex;
		align-items: center;
		border: 0;
	}

	.profile_button img {
		z-index: 2;
	}

	.profile_button .text {
		display: flex;
		align-items: center;
		gap: 0.25rem;
		border-radius: 0 var(--radius-lg) var(--radius-lg) 0;
		padding: 0.35rem 0.375rem 0.35rem 0.75rem;
		padding-left: 0.75rem;
		translate: -0.35rem;
		background-color: var(--primary-bg);
		border: 1px solid var(--primary-light);
	}

	.profile_button .text span {
		max-width: 5rem;
		overflow: hidden;
	}

	.dropdown_content {
		position: absolute;
		top: 100%;
		right: 1rem;
		margin-top: 0.5rem;
		min-width: max-content;
		background-color: var(--text-white);
		box-shadow: var(--shadow-lg);
		border-radius: var(--radius-md);
		padding: var(--spacing-sm);
		border: 1px solid var(--neutral-color);
	}

	.dropdown_item {
		display: flex;
		align-items: center;
		gap: 0.75rem;
		padding: 10px 15px;
		text-decoration: none;
		font-weight: 500;
		transition: background-color 0.2s;
		width: 100%;
		text-align: left;
		border: none;
		background-color: none;
		cursor: pointer;
		transition: 0.3s ease;
	}

	.dropdown_item:hover {
		background-color: var(--primary-bg, #f1f9f6);
		color: var(--primary-dark, #2a8960);
	}

	.logout_btn {
		color: var(--error-color);
	}

	.logout_btn:hover {
		background-color: #f5292533;
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
		color: var(--primary-color);
	}

	@media (max-width: 54rem) {
		.nav-wrapper {
			position: fixed;
			top: 0;
			right: 0;
			width: 80%;
			max-width: 20rem;
			height: 100vh;
			flex-direction: column;
			overflow-y: auto;
			background-color: var(--text-white);
			box-shadow: var(--shadow-lg);
			padding: var(--spacing-2xl) var(--spacing-lg) var(--spacing-lg);
			z-index: 1000;
		}

		.menu_overlay {
			position: fixed;
			inset: 0;
			background-color: rgba(0, 0, 0, 0.5);
			backdrop-filter: blur(5px);
			z-index: 5;
		}

		.nav-links {
			position: static;
			left: unset;
			translate: unset;
			flex-direction: column;
			width: 100%;
			gap: var(--spacing-lg);
		}

		.nav-links a,
		.dropdown_item {
			width: 100%;
			padding: var(--spacing-xs) 0;
			border-bottom: 1px solid var(--text-light);
		}

		.dropdown_item:is(:hover, :focus-visible, :active) {
			color: var(--primary-color);
			border-bottom: 2px solid var(--primary-color);
			background-color: transparent;
		}

		.logout_btn {
			color: var(--error-color);
			border-bottom: 2px solid var(--error-color);
		}

		.logout_btn:is(:hover, :focus-visible, :active) {
			color: var(--error-color);
			border-bottom: 2px solid var(--error-color);
		}

		.dropdown {
			display: flex;
			flex-direction: column;
			position: static;
			width: 100%;
		}

		.dropdown_content {
			display: flex;
			flex-direction: column;
			position: static;
			width: 100%;
			margin-top: 0;
			box-shadow: none;
			border: none;
			border-radius: 0;
			padding: 0;
			gap: var(--spacing-lg);
		}

		.nav-toggle {
			display: flex;
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
	}
</style>
