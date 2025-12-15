<script>
    import { page } from '$app/stores';
    import { ArrowRight, Settings, Lock, Coins, BarChart3, Users, Briefcase } from 'lucide-svelte';

    let { children, data } = $props();

    // Access user data from the layout.server.js (which must redirect if user is not logged in)
    const user = data.user;

    const navLinks = [
        { label: 'Dashboard Overview', href: '/dashboard', icon: BarChart3, type: 'primary' },
        { label: 'My Campaigns', href: '/dashboard/campaigns', icon: Briefcase, type: 'primary' },
        { label: 'Launch Campaign', href: '/dashboard/create', icon: ArrowRight, type: 'primary' },
        { label: 'Funds & Payouts', href: '/dashboard/funds', icon: Coins, type: 'primary' },
        { label: 'Compliance & KYC', href: '/dashboard/compliance', icon: Users, type: 'primary' },
    ];

    const accountLinks = [
        { label: 'Profile Settings', href: '/account/profile', icon: Settings, type: 'account' },
        { label: 'Security & Login', href: '/account/security', icon: Lock, type: 'account' },
    ];
</script>

<div class="dashboard-layout">
    <nav class="sidebar">
        <div class="user-greeting">
            <p>Welcome, <strong>{user?.fullName || user?.email.split('@')[0]}</strong></p>
        </div>

        <ul class="nav-list">
            {#each navLinks as link}
                <li class:active={$page.url.pathname === link.href}>
                    <a href={link.href}>
                        <link.icon />
                        <span>{link.label}</span>
                    </a>
                </li>
            {/each}
        </ul>

        <hr class="divider">

        <h3 class="account-header">Account Management</h3>
        <ul class="nav-list account-nav">
            {#each accountLinks as link}
                <li class:active={$page.url.pathname === link.href}>
                    <a href={link.href}>
                        <link.icon />
                        <span>{link.label}</span>
                    </a>
                </li>
            {/each}
        </ul>
    </nav>

    <main class="content-area">
        {@render children?.()}
    </main>
</div>

<style>
    /* --- Layout Container --- */
    .dashboard-layout {
        display: flex;
        min-height: 100vh;
        background-color: var(--primary-bg);
    }

    /* --- Sidebar Styles --- */
    .sidebar {
        width: 250px;
        flex-shrink: 0;
        background-color: var(--text-white);
        padding: var(--spacing-md);
        box-shadow: var(--shadow-md); /* Stronger shadow to separate from content */
    }

    .user-greeting {
        padding: var(--spacing-sm) 0;
        margin-bottom: var(--spacing-md);
        border-bottom: 1px solid var(--neutral-color);
        font-size: 0.9rem;
    }

    .user-greeting strong {
        color: var(--primary-dark);
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
    }

    /* --- Navigation Links --- */
    .nav-list {
        list-style: none;
        padding: 0;
        margin: 0;
    }

    .nav-list li {
        margin-bottom: 4px;
        border-radius: var(--radius-sm);
    }

    .nav-list a {
        display: flex;
        align-items: center;
        padding: var(--spacing-smr) var(--spacing-sm);
        text-decoration: none;
        color: var(--text-dark);
        font-weight: 500;
        transition: all 0.2s;
        gap: var(--spacing-sm);
    }

    .nav-list li:hover {
        background-color: var(--primary-bg);
    }

    /* Active Link State */
    .nav-list li.active {
        background-color: var(--primary-color);
        box-shadow: var(--shadow-sm);
    }

    .nav-list li.active a {
        color: var(--text-white);
        font-weight: 600;
    }

    /* --- Main Content Area --- */
    .content-area {
        flex-grow: 1;
        padding: var(--spacing-lg);
        overflow-y: auto; /* Ensure content can scroll */
    }
</style>