<script>
	import { page } from "$app/stores";
	import Logo from "$lib/components/ui/Logo.svelte";
	import MenuButton from "$lib/components/ui/MenuButton.svelte";
	import { ArrowRight, Settings, Lock, Coins, BarChart3, Users, Briefcase, ExternalLink } from "lucide-svelte";

	let { children, data } = $props();

	const user = data.user;

	let sidebarOpen = $state(false);

	function toggleSidebar() {
		sidebarOpen = !sidebarOpen;
	}

	function closeSidebar() {
		sidebarOpen = false;
	}

	const navLinks = [
		{ label: "Dashboard Overview", href: "/dashboard", icon: BarChart3 },
		{ label: "My Campaigns", href: "/dashboard/campaigns", icon: Briefcase },
		{ label: "Launch Campaign", href: "/dashboard/create", icon: ArrowRight },
		{ label: "Funds & Payouts", href: "/dashboard/funds", icon: Coins },
		{ label: "Compliance & KYC", href: "/dashboard/compliance", icon: Users }
	];

	const accountLinks = [
		{ label: "Profile Settings", href: "/account/profile", icon: Settings },
		{ label: "Security & Login", href: "/account/security", icon: Lock }
	];
</script>

<div class="dashboard-layout">
	<nav class="sidebar" class:open={sidebarOpen}>
		<div class="nav_header">
			<Logo />
			<a class="profile_link" href="/profile/{user.slug}"><span>View Profile</span> <ExternalLink size={16} /></a>
			<span class="close_nav"><MenuButton active={true} clickHandler={closeSidebar} /></span>
		</div>

		<ul class="nav-list">
			{#each navLinks as link}
				<li class:active={$page.url.pathname === link.href}>
					<a href={link.href} onclick={closeSidebar}>
						<link.icon />
						<span>{link.label}</span>
					</a>
				</li>
			{/each}
		</ul>

		<hr class="divider" />

		<h3 class="account-header">Account Management</h3>
		<ul class="nav-list account-nav">
			{#each accountLinks as link}
				<li class:active={$page.url.pathname === link.href}>
					<a href={link.href} onclick={closeSidebar}>
						<link.icon />
						<span>{link.label}</span>
					</a>
				</li>
			{/each}
		</ul>
	</nav>

	<div class="overlay" class:visible={sidebarOpen} onclick={closeSidebar} role="none"></div>

	<main>
		<header class="mobile_header">
			<MenuButton active={sidebarOpen} clickHandler={toggleSidebar} />
			<Logo />
		</header>
		<div class="content-area">
			{@render children?.()}
		</div>
	</main>
</div>

<style>
	.dashboard-layout {
		display: flex;
		height: 100vh;
		background-color: var(--primary-bg);
	}

	.sidebar {
		min-width: 18rem;
		background-color: var(--text-white);
		box-shadow: var(--shadow-md);
		flex-shrink: 0;
		overflow-y: auto;
	}

	.nav_header {
		position: relative;
		border-bottom: 1px solid var(--neutral-color);
		padding: var(--spacing-sm) var(--spacing-md);
	}

	.profile_link {
		font-weight: 600;
		transition: 0.3s ease;
		color: var(--text-medium);
	}

	.profile_link:is(:hover, :focus-visible, :active) {
		color: var(--primary-color);
	}

	.close_nav {
		display: none;
		position: absolute;
		top: var(--spacing-md);
		right: var(--spacing-md);
	}

	.nav-list {
		list-style: none;
		padding: var(--spacing-sm);
		margin: 0;
	}

	.nav-list li {
		margin-bottom: var(--spacing-xs);
		border-radius: var(--radius-sm);
	}

	.nav-list a {
		display: flex;
		align-items: center;
		padding: var(--spacing-smr) var(--spacing-sm);
		font-weight: 500;
		gap: var(--spacing-sm);
		transition: background 0.2s ease;
	}

	.nav-list li:is(:hover, :focus-within, :active) {
		background-color: var(--primary-bg);
	}

	.nav-list li.active {
		background-color: var(--primary-color);
		box-shadow: var(--shadow-sm);
	}

	.nav-list li.active a {
		color: var(--text-white);
		font-weight: 600;
	}

	.divider {
		border: none;
		border-top: 1px solid var(--neutral-color);
		margin: var(--spacing-md) 0;
	}

	.account-header {
		font-size: 0.8rem;
		color: var(--text-medium);
		text-transform: uppercase;
		margin-bottom: var(--spacing-xs);
		padding: 0 var(--spacing-md);
	}

	.overlay {
		display: none;
		position: fixed;
		inset: 0;
		background: rgba(0, 0, 0, 0.4);
		z-index: 2;
		backdrop-filter: blur(2px);
	}

	main {
		flex-grow: 1;
		overflow-y: auto;
	}
	.content-area {
		padding: var(--spacing-lg);
	}

	.mobile_header {
		display: none;
		align-items: center;
		padding: var(--spacing-sm);
		gap: var(--spacing-sm);
		background-color: var(--text-white);
		box-shadow: var(--shadow-sm);
		position: sticky;
		top: 0;
		z-index: 1;
	}

	@media (max-width: 48rem) {
		.dashboard-layout {
			flex-direction: column;
		}

		.mobile_header {
			display: flex;
		}

		.sidebar {
			position: fixed;
			top: 0;
			left: 0;
			height: 100vh;
			transform: translateX(-100%);
			transition: transform 0.25s ease;
			z-index: 3;
		}

		.sidebar.open {
			transform: translateX(0);
		}

		.close_nav {
			display: initial;
		}

		.overlay.visible {
			display: block;
		}

		.content-area {
			padding: var(--spacing-md);
		}
	}
</style>
