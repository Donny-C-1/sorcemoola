<script>
	import { onMount } from 'svelte';

    let { data } = $props();
    // Svelte 5 Runes for state management
    let searchQuery = $state("");
    let filterStatus = $state("active");

    let campaigns = $state([
        {
            id: 1,
            name: "Emergency Flood Relief - Northern Region",
            description: "Providing clean water and food supplies to over 500 families affected by recent flooding.",
            raised: 8500,
            goal: 10000,
            donors: 124,
            daysLeft: 5,
            status: "active",
            category: "Emergency",
            image: "https://images.unsplash.com/photo-1547619292-8816ee7cdd50?q=80&w=400&h=250&auto=format&fit=crop",
            slug: "efrnr"
        },
        {
            id: 2,
            name: "Medical Supplies for City Hospital",
            description: "Funding essential surgical equipment and basic medicine for the pediatric ward.",
            raised: 12000,
            goal: 50000,
            donors: 45,
            daysLeft: 22,
            status: "active",
            category: "Medical",
            image: "https://images.unsplash.com/photo-1584515933487-759f398f5851?q=80&w=400&h=250&auto=format&fit=crop",
            slug: "msfch"
        },
        {
            id: 3,
            name: "Clean Water Well Construction",
            description: "Building three solar-powered water wells in the arid sub-districts.",
            raised: 15000,
            goal: 15000,
            donors: 310,
            daysLeft: 0,
            status: "completed",
            category: "Infrastructure",
            image: "https://images.unsplash.com/photo-1541810270631-4171616c6806?q=80&w=400&h=250&auto=format&fit=crop",
            slug: "cwwc"
        }
    ]);

    // Derived rune for filtering
    let filteredCampaigns = $derived(
        data.campaigns.filter(c => {
            const matchesSearch = c.name.toLowerCase().includes(searchQuery.toLowerCase());
            const matchesFilter = filterStatus === "all" || c.status === filterStatus;
            return matchesSearch && matchesFilter;
        })
    );

    const calculateProgress = (raised, goal) => Math.min((raised / goal) * 100, 100);

    onMount(() => $inspect(data));
</script>

<div class="campaigns-page">
    <header class="page-header">
        <div class="title-section">
            <h1>Campaign Management</h1>
            <p>Monitor and manage your active relief efforts.</p>
        </div>

        <div class="action-bar">
            <div class="search-box">
                <input 
                    type="text" 
                    placeholder="Search campaigns..." 
                    bind:value={searchQuery}
                />
            </div>
            
            <div class="filter-tabs">
                <button 
                    class:active={filterStatus === "active"} 
                    onclick={() => filterStatus = "active"}>Active</button>
                <button 
                    class:active={filterStatus === "completed"} 
                    onclick={() => filterStatus = "completed"}>Completed</button>
                <button 
                    class:active={filterStatus === "all"} 
                    onclick={() => filterStatus = "all"}>All</button>
            </div>

            <button class="btn-primary">+ New Campaign</button>
        </div>
    </header>

    <div class="campaign-grid">
        {#each filteredCampaigns as campaign (campaign.id)}
            <article class="campaign-card">
                <div class="card-image">
                    <img src={campaign.image} alt={campaign.name} />
                    <span class="category-badge">{campaign.category}</span>
                </div>

                <div class="card-body">
                    <div class="card-title-row">
                        <h3><a href="/campaigns/{campaign.slug}">{campaign.name}</a></h3>
                        <div class="status-dot {campaign.status}"></div>
                    </div>
                    <p class="description">{campaign.description}</p>

                    <div class="progress-container">
                        <div class="progress-labels">
                            <span class="raised-amt">${campaign.amountRaised.toLocaleString()}</span>
                            <span class="goal-amt">of ${campaign.fundGoal.toLocaleString()}</span>
                        </div>
                        <div class="progress-bar-bg">
                            <div 
                                class="progress-fill" 
                                style="width: {calculateProgress(campaign.amountRaised, campaign.fundGoal)}%">
                            </div>
                        </div>
                        <span class="percentage-text">{Math.round(calculateProgress(campaign.amountRaised, campaign.fundGoal))}% funded</span>
                    </div>

                    <div class="card-footer">
                        <div class="stat">
                            <span class="stat-val">{campaign.donors}</span>
                            <span class="stat-lbl">Donors</span>
                        </div>
                        <div class="stat border-left">
                            <span class="stat-val">{campaign.daysLeft}</span>
                            <span class="stat-lbl">Days Left</span>
                        </div>
                        <div class="card-actions">
                            <button title="Edit Campaign" class="icon-btn">✎</button>
                            <button title="Share" class="icon-btn">↗</button>
                            <button title="View Details" class="btn-outline">Manage</button>
                        </div>
                    </div>
                </div>
            </article>
        {/each}

        {#if filteredCampaigns.length === 0}
            <div class="empty-state">
                <p>No campaigns found matching your criteria.</p>
            </div>
        {/if}
    </div>
</div>

<style>
    .campaigns-page {
        max-width: var(--container-width);
        margin: 0 auto;
        padding: var(--container-padding);
    }

    .page-header {
        display: flex;
        flex-direction: column;
        gap: var(--spacing-md);
        margin-bottom: var(--spacing-xl);
    }

    .title-section h1 {
        font-size: 2rem;
        color: var(--text-dark);
    }

    .title-section p {
        color: var(--text-medium);
    }

    /* Action Bar Styles */
    .action-bar {
        display: flex;
        justify-content: space-between;
        align-items: center;
        gap: var(--spacing-md);
        background: white;
        padding: var(--spacing-sm);
        border-radius: var(--radius-md);
        box-shadow: var(--shadow-sm);
    }

    .search-box input {
        padding: var(--spacing-smr) var(--spacing-md);
        border: 1px solid var(--primary-bg);
        border-radius: var(--radius-sm);
        min-width: 250px;
        background: var(--primary-bg);
    }

    .filter-tabs {
        display: flex;
        gap: var(--spacing-xs);
        background: var(--primary-bg);
        padding: 4px;
        border-radius: var(--radius-sm);
    }

    .filter-tabs button {
        padding: 6px 16px;
        border: none;
        background: transparent;
        cursor: pointer;
        border-radius: var(--radius-sm);
        font-weight: 500;
        color: var(--text-medium);
        transition: all 0.2s;
    }

    .filter-tabs button.active {
        background: white;
        color: var(--primary-color);
        box-shadow: var(--shadow-sm);
    }

    /* Card Grid Styles */
    .campaign-grid {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
        gap: var(--spacing-lg);
    }

    .campaign-card {
        background: white;
        border-radius: var(--radius-lg);
        overflow: hidden;
        box-shadow: var(--shadow-sm);
        transition: transform 0.2s, box-shadow 0.2s;
        border: 1px solid rgba(0,0,0,0.05);
    }

    .campaign-card:hover {
        transform: translateY(-4px);
        box-shadow: var(--shadow-md);
    }

    .card-image {
        height: 180px;
        position: relative;
    }

    .card-image img {
        width: 100%;
        height: 100%;
        object-fit: cover;
    }

    .category-badge {
        position: absolute;
        top: 12px;
        left: 12px;
        background: var(--accent-color);
        color: white;
        padding: 4px 12px;
        border-radius: 20px;
        font-size: 0.75rem;
        font-weight: 600;
    }

    .card-body {
        padding: var(--spacing-md);
    }

    .card-title-row {
        display: flex;
        justify-content: space-between;
        align-items: flex-start;
        margin-bottom: var(--spacing-xs);
    }

    .card-title-row h3 {
        font-size: 1.15rem;
        font-weight: 700;
        line-height: 1.3;
    }

    .status-dot {
        width: 10px;
        height: 10px;
        border-radius: 50%;
        margin-top: 6px;
    }

    .status-dot.active { background: var(--primary-color); box-shadow: 0 0 8px var(--primary-light); }
    .status-dot.completed { background: var(--text-light); }

    .description {
        font-size: 0.9rem;
        color: var(--text-medium);
        margin-bottom: var(--spacing-md);
        display: -webkit-box;
        -webkit-line-clamp: 2;
        line-clamp: 2;
        -webkit-box-orient: vertical;
        overflow: hidden;
    }

    /* Progress Styling */
    .progress-container {
        margin-bottom: var(--spacing-md);
    }

    .progress-labels {
        display: flex;
        align-items: baseline;
        gap: 6px;
        margin-bottom: 8px;
    }

    .raised-amt { font-weight: 800; font-size: 1.1rem; color: var(--text-dark); }
    .goal-amt { font-size: 0.85rem; color: var(--text-light); }

    .progress-bar-bg {
        height: 8px;
        background: var(--primary-bg);
        border-radius: 4px;
        overflow: hidden;
        margin-bottom: 6px;
    }

    .progress-fill {
        height: 100%;
        background: linear-gradient(90deg, var(--primary-color), var(--primary-light));
        border-radius: 4px;
    }

    .percentage-text {
        font-size: 0.75rem;
        font-weight: 600;
        color: var(--primary-dark);
    }

    /* Footer Stats */
    .card-footer {
        display: flex;
        align-items: center;
        padding-top: var(--spacing-md);
        border-top: 1px solid var(--primary-bg);
    }

    .stat {
        display: flex;
        flex-direction: column;
        padding-right: var(--spacing-md);
    }

    .stat.border-left {
        padding-left: var(--spacing-md);
        border-left: 1px solid var(--primary-bg);
    }

    .stat-val { font-weight: 700; font-size: 0.95rem; }
    .stat-lbl { font-size: 0.7rem; color: var(--text-light); text-transform: uppercase; }

    .card-actions {
        margin-left: auto;
        display: flex;
        gap: 8px;
    }

    .icon-btn {
        background: var(--primary-bg);
        border: none;
        width: 32px;
        height: 32px;
        border-radius: var(--radius-sm);
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
        color: var(--text-medium);
    }

    .btn-outline {
        border: 1px solid var(--primary-color);
        background: transparent;
        color: var(--primary-color);
        padding: 4px 12px;
        border-radius: var(--radius-sm);
        font-size: 0.85rem;
        font-weight: 600;
        cursor: pointer;
    }

    .btn-primary {
        background: var(--primary-color);
        color: white;
        border: none;
        padding: var(--spacing-smr) var(--spacing-md);
        border-radius: var(--radius-sm);
        font-weight: 600;
        cursor: pointer;
    }

    .empty-state {
        grid-column: 1 / -1;
        text-align: center;
        padding: var(--spacing-2xl);
        color: var(--text-light);
        background: var(--primary-bg);
        border-radius: var(--radius-lg);
        border: 2px dashed var(--neutral-color);
    }

    @media (max-width: 600px) {
        .action-bar {
            flex-direction: column;
            align-items: stretch;
        }
        .search-box input {
            min-width: 100%;
        }
    }
</style>