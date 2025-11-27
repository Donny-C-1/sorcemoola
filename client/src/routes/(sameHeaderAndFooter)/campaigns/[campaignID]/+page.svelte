<script>
    import { onMount } from 'svelte';
    import { fade, fly } from 'svelte/transition';
    import { PUBLIC_SERVER_URL } from "$env/static/public";

    let PaystackPop;
    
    // Campaign data - would typically come from an API
    let campaign = {
        id: 'cam123456',
        title: 'Save the Enchanted Forest',
        creator: 'Wizard Willow',
        creatorAvatar: 'https://images.unsplash.com/photo-1760574772950-f37de9dce85c?ixid=M3w4MjcwNjd8MHwxfHNlYXJjaHwzfHx3aXphcmQlMjB3aWxsb3d8ZW58MHx8fHwxNzYzNjYxNTA4fDA&ixlib=rb-4.1.0&fit=max&q=80',
        targetAmount: 50000,
        currentAmount: 32750,
        backers: 428,
        daysLeft: 15,
        description: 'The Enchanted Forest is home to magical creatures and rare plants with healing properties. Industrial development threatens to destroy this unique ecosystem. Your contributions will help us purchase the land and establish it as a protected sanctuary.',
        story: `<p>For centuries, the Enchanted Forest has been a haven for magical creatures and a source of powerful ingredients for potions and spells. The crystal-clear streams that flow through it are said to have rejuvenating properties, and the rare flora that grows here can't be found anywhere else in the world.</p>
                <p>Unfortunately, a large corporation has acquired the land and plans to clear it for development. Once this ecosystem is destroyed, we can never get it back.</p>
                <p>With your help, we aim to purchase the land and establish it as a protected magical sanctuary. Every contribution brings us one step closer to saving this irreplaceable natural wonder.</p>`,
        images: [
            'https://images.unsplash.com/photo-1619982870053-1545857b7eb4?ixid=M3w4MjcwNjd8MHwxfHNlYXJjaHwzfHxlbmNoYW50ZWQlMjBmb3Jlc3R8ZW58MHx8fHwxNzYzNjYxMjc5fDA&ixlib=rb-4.1.0&fit=max&q=80',
            'https://images.unsplash.com/photo-1597201423947-3e0028337902?ixid=M3w4MjcwNjd8MHwxfHNlYXJjaHw1fHxlbmNoYW50ZWQlMjBmb3Jlc3R8ZW58MHx8fHwxNzYzNjYxMjc5fDA&ixlib=rb-4.1.0&fit=max&q=80',
            'https://images.unsplash.com/photo-1518562180175-34a163b1a9a6?ixid=M3w4MjcwNjd8MHwxfHNlYXJjaHwxMHx8ZW5jaGFudGVkJTIwZm9yZXN0fGVufDB8fHx8MTc2MzY2MTI3OXww&ixlib=rb-4.1.0&fit=max&q=80',
        ],
        updates: [
            { date: '2023-10-25', title: 'Legal progress!', content: 'We\'ve secured a temporary injunction to halt development while our case is being reviewed.' },
            { date: '2023-10-10', title: '25% funded!', content: 'Thank you to all our amazing supporters. We\'re a quarter of the way there!' }
        ],
        rewards: [
            { 
                id: 'r1',
                amount: 25,
                title: 'Magical Seeds',
                description: 'A packet of seeds from the Enchanted Forest. Plant them in your garden to attract friendly sprites.',
                backers: 215,
                delivery: 'December 2023'
            },
            {
                id: 'r2',
                amount: 100,
                title: 'Crystal Vial',
                description: 'A small vial of water from the Enchanted Stream, known for its restorative properties.',
                backers: 132,
                delivery: 'December 2023'
            },
            {
                id: 'r3',
                amount: 500,
                title: 'Guardian Status',
                description: 'Your name will be engraved on the Sanctuary Guardian plaque. Includes all previous rewards.',
                backers: 45,
                delivery: 'January 2024'
            }
        ]
    };

    let currentImage = 0;
    let pledgeAmount = 0;
    let customAmount = '';
    let selectedReward = null;
    let showDonationForm = false;
    let activeTab = 'story';

    const formatCurrency = (amount) => {
        return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(amount);
    };

    const getProgressPercentage = () => {
        return Math.min(Math.floor((campaign.currentAmount / campaign.targetAmount) * 100), 100);
    };

    const nextImage = () => {
        currentImage = (currentImage + 1) % campaign.images.length;
    };

    const prevImage = () => {
        currentImage = (currentImage - 1 + campaign.images.length) % campaign.images.length;
    };

    const selectReward = (reward) => {
        selectedReward = reward;
        pledgeAmount = reward.amount;
        customAmount = '';
        showDonationForm = true;
    };

    const handleCustomPledge = () => {
        selectedReward = null;
        pledgeAmount = parseFloat(customAmount);
        showDonationForm = true;
    };

    const submitDonation = () => {
        // In a real app, this would connect to a payment processor
        alert(`Thank you for your ${formatCurrency(pledgeAmount)} pledge to save the Enchanted Forest!`);

        initPayment();
        showDonationForm = false;
        selectedReward = null;
        pledgeAmount = 0;
        customAmount = '';
    };

    onMount(async() => {
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

<div class="campaign-page">
    <section class="campaign-header">
        <div class="image-gallery">
            <button class="gallery-nav prev" onclick={prevImage} aria-label="Previous image">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <polyline points="15 18 9 12 15 6"></polyline>
                </svg>
            </button>
            <div class="gallery-container">
                {#each campaign.images as image, i}
                    <div class="gallery-image" class:active={i === currentImage}>
                        <img src={image} alt={`${campaign.title} - image ${i+1}`} />
                    </div>
                {/each}
            </div>
            <button class="gallery-nav next" onclick={nextImage} aria-label="Next image">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <polyline points="9 18 15 12 9 6"></polyline>
                </svg>
            </button>
            <div class="gallery-indicators">
                {#each campaign.images as _, i}
                    <button 
                        class="indicator" 
                        class:active={i === currentImage}
                        onclick={() => currentImage = i}
                        aria-label={`View image ${i+1}`}>
                    </button>
                {/each}
            </div>
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
        <main class="campaign-details">
            <div class="campaign-tabs">
                <button 
                    class:active={activeTab === 'story'} 
                    onclick={() => activeTab = 'story'}>
                    Campaign Story
                </button>
                <button 
                    class:active={activeTab === 'updates'} 
                    onclick={() => activeTab = 'updates'}>
                    Updates ({campaign.updates.length})
                </button>
            </div>

            <div class="tab-content">
                {#if activeTab === 'story'}
                    <section class="story-tab" transition:fade={{ duration: 200 }}>
                        <p class="campaign-summary">{campaign.description}</p>
                        <div class="campaign-story">
                            {@html campaign.story}
                        </div>
                    </section>
                {:else if activeTab === 'updates'}
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
        </main>

        <aside class="campaign-sidebar">
            <div class="funding-status">
                <div class="funding-amount">
                    <h2>{formatCurrency(campaign.currentAmount)}</h2>
                    <p>raised of {formatCurrency(campaign.targetAmount)} goal</p>
                </div>

                <div class="progress-bar">
                    <div class="progress" style="width: {getProgressPercentage()}%"></div>
                </div>

                <div class="funding-stats">
                    <div class="stat">
                        <div class="stat-value">{campaign.backers}</div>
                        <div class="stat-label">Backers</div>
                    </div>
                    <div class="stat">
                        <div class="stat-value">{campaign.daysLeft}</div>
                        <div class="stat-label">Days Left</div>
                    </div>
                </div>
            </div>

            <div class="donation-options">
                <div class="custom-pledge">
                    <label for="custom-amount">Enter a donation amount:</label>
                    <div class="amount-input">
                        <span class="currency-symbol">$</span>
                        <input 
                            type="number" 
                            id="custom-amount"
                            min="1"
                            placeholder="Enter amount"
                            bind:value={customAmount}
                        />
                    </div>
                    <button 
                        class="donate-button"
                        disabled={!customAmount || isNaN(parseFloat(customAmount)) || parseFloat(customAmount) <= 0}
                        onclick={handleCustomPledge}>
                        Continue
                    </button>
                </div>

                <div class="reward-options">
                    <h3>Select a reward</h3>
                    {#each campaign.rewards as reward}
                        <div 
                            class="reward-card"
                            class:selected={selectedReward && selectedReward.id === reward.id}
                            onclick={() => selectReward(reward)}
                            transition:fade
                        >
                            <div class="reward-amount">{formatCurrency(reward.amount)}</div>
                            <h4>{reward.title}</h4>
                            <p>{reward.description}</p>
                            <div class="reward-details">
                                <span class="backers">{reward.backers} backers</span>
                                <span class="delivery">Est. delivery: {reward.delivery}</span>
                            </div>
                        </div>
                    {/each}
                </div>
            </div>
        </aside>
    </div>

    {#if showDonationForm}
        <div class="donation-modal-overlay" onclick={() => showDonationForm = false}>
            <div class="donation-modal" onclick={e => e.stopPropagation()}>
                <button class="close-modal" onclick={() => showDonationForm = false}>×</button>
                <h2>Complete Your Donation</h2>
                
                <div class="donation-summary">
                    <p>You're pledging: <strong>{formatCurrency(pledgeAmount)}</strong></p>
                    {#if selectedReward}
                        <p>Selected reward: <strong>{selectedReward.title}</strong></p>
                    {/if}
                </div>
                
                <form class="donation-form" onsubmit={submitDonation}>
                    <div class="form-group">
                        <label for="donor-name">Full Name</label>
                        <input type="text" id="donor-name" required />
                    </div>
                    
                    <div class="form-group">
                        <label for="donor-email">Email</label>
                        <input type="email" id="donor-email" required />
                    </div>
                    
                    <div class="form-group">
                        <label for="card-number">Card Number</label>
                        <input type="text" id="card-number" placeholder="1234 5678 9012 3456" required />
                    </div>
                    
                    <div class="form-row">
                        <div class="form-group half">
                            <label for="card-expiry">Expiry Date</label>
                            <input type="text" id="card-expiry" placeholder="MM/YY" required />
                        </div>
                        
                        <div class="form-group half">
                            <label for="card-cvv">CVV</label>
                            <input type="text" id="card-cvv" placeholder="123" required />
                        </div>
                    </div>
                    
                    <button type="submit" class="donate-button">Complete {formatCurrency(pledgeAmount)} Donation</button>
                    
                    <p class="secure-notice">
                        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
                            <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
                        </svg>
                        Secure payment processing
                    </p>
                </form>
            </div>
        </div>
    {/if}
</div>

<style>
    /* Campaign Page */
    .campaign-page {
        max-width: 1200px;
        margin: 0 auto;
        padding: 0 20px;
    }

    /* Campaign Header */
    .campaign-header {
        margin-bottom: 40px;
    }

    .campaign-title {
        padding: 20px 0;
    }

    .campaign-title h1 {
        font-size: 2.5rem;
        margin-bottom: 10px;
        color: #333;
    }

    .creator-info {
        display: flex;
        align-items: center;
        gap: 10px;
    }

    .creator-info img {
        width: 40px;
        height: 40px;
        border-radius: 50%;
        object-fit: cover;
    }

    /* Image Gallery */
    .image-gallery {
        position: relative;
        border-radius: 8px;
        overflow: hidden;
        height: 400px;
        margin-bottom: 20px;
        box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    }

    .gallery-container {
        height: 100%;
        position: relative;
    }

    .gallery-image {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        opacity: 0;
        transition: opacity 0.5s ease-in-out;
    }

    .gallery-image.active {
        opacity: 1;
        z-index: 1;
    }

    .gallery-image img {
        width: 100%;
        height: 100%;
        object-fit: cover;
    }

    .gallery-nav {
        position: absolute;
        top: 50%;
        transform: translateY(-50%);
        background: rgba(255, 255, 255, 0.8);
        border: none;
        border-radius: 50%;
        width: 40px;
        height: 40px;
        display: flex;
        align-items: center;
        justify-content: center;
        cursor: pointer;
        z-index: 2;
        box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
    }

    .gallery-nav:hover {
        background: white;
    }

    .gallery-nav.prev {
        left: 15px;
    }

    .gallery-nav.next {
        right: 15px;
    }

    .gallery-indicators {
        position: absolute;
        bottom: 15px;
        left: 50%;
        transform: translateX(-50%);
        display: flex;
        gap: 8px;
        z-index: 2;
    }

    .indicator {
        width: 10px;
        height: 10px;
        border-radius: 50%;
        background: rgba(255, 255, 255, 0.5);
        border: none;
        padding: 0;
        cursor: pointer;
    }

    .indicator.active {
        background: white;
    }

    /* Campaign Content Layout */
    .campaign-content {
        display: grid;
        grid-template-columns: 1fr 350px;
        gap: 40px;
    }

    /* Campaign Details Section */
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
        line-height: 1.6;
        color: #555;
        margin-bottom: 30px;
    }

    .campaign-story {
        font-size: 1rem;
        line-height: 1.7;
        color: #333;
    }

    .campaign-story p {
        margin-bottom: 20px;
    }

    /* Updates Tab */
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
        color: #888;
        font-size: 0.9rem;
        margin-bottom: 8px;
    }

    .update-content h3 {
        margin-bottom: 8px;
        color: #444;
    }

    .no-updates {
        color: #888;
        font-style: italic;
    }

    /* Campaign Sidebar */
    .campaign-sidebar {
        display: flex;
        flex-direction: column;
        gap: 30px;
    }

    /* Funding Status */
    .funding-status {
        background: white;
        border-radius: 8px;
        padding: 25px;
        box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
    }

    .funding-amount h2 {
        font-size: 2rem;
        color: #34a273;
        margin-bottom: 5px;
    }

    .funding-amount p {
        color: #666;
        margin-bottom: 15px;
    }

    .progress-bar {
        height: 8px;
        background: #e0e0e0;
        border-radius: 4px;
        margin-bottom: 20px;
        overflow: hidden;
    }

    .progress {
        height: 100%;
        background: #34a273;
    }

    .funding-stats {
        display: flex;
        justify-content: space-between;
    }

    .stat {
        text-align: center;
    }

    .stat-value {
        font-size: 1.5rem;
        font-weight: 700;
        color: #333;
    }

    .stat-label {
        font-size: 0.9rem;
        color: #666;
    }

    /* Donation Options */
    .donation-options {
        background: white;
        border-radius: 8px;
        padding: 25px;
        box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
        display: flex;
        flex-direction: column;
        gap: 25px;
    }

    .custom-pledge {
        margin-bottom: 10px;
    }

    .custom-pledge label {
        display: block;
        margin-bottom: 10px;
        color: #555;
        font-weight: 500;
    }

    .amount-input {
        position: relative;
        margin-bottom: 15px;
    }

    .currency-symbol {
        position: absolute;
        top: 50%;
        left: 15px;
        transform: translateY(-50%);
        color: #666;
    }

    #custom-amount {
        width: 100%;
        padding: 12px 12px 12px 30px;
        border: 1px solid #ddd;
        border-radius: 4px;
        font-size: 1rem;
    }

    .donate-button {
        width: 100%;
        background-color: #34a273;
        color: white;
        border: none;
        padding: 12px 20px;
        border-radius: 4px;
        font-weight: 600;
        cursor: pointer;
        transition: background-color 0.2s;
    }

    .donate-button:hover {
        background-color: #2c8b62;
    }

    .donate-button:disabled {
        background-color: #a0a0a0;
        cursor: not-allowed;
    }

    /* Reward Options */
    .reward-options h3 {
        margin-bottom: 15px;
        color: #444;
    }

    .reward-card {
        border: 2px solid #eee;
        border-radius: 8px;
        padding: 20px;
        margin-bottom: 15px;
        cursor: pointer;
        transition: all 0.2s;
    }

    .reward-card:hover {
        border-color: #34a273;
        transform: translateY(-2px);
        box-shadow: 0 5px 15px rgba(0, 0, 0, 0.05);
    }

    .reward-card.selected {
        border-color: #34a273;
        background-color: rgba(52, 162, 115, 0.05);
    }

    .reward-amount {
        font-size: 1.2rem;
        color: #34a273;
        font-weight: 700;
        margin-bottom: 10px;
    }

    .reward-card h4 {
        margin-bottom: 8px;
        color: #333;
    }

    .reward-card p {
        color: #666;
        margin-bottom: 15px;
        line-height: 1.5;
    }

    .reward-details {
        display: flex;
        justify-content: space-between;
        font-size: 0.9rem;
        color: #888;
    }

    /* Donation Modal */
    .donation-modal-overlay {
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background: rgba(0, 0, 0, 0.7);
        display: flex;
        align-items: center;
        justify-content: center;
        z-index: 1000;
        padding: 20px;
    }

    .donation-modal {
        background: white;
        border-radius: 8px;
        max-width: 500px;
        width: 100%;
        padding: 30px;
        position: relative;
        max-height: 90vh;
        overflow-y: auto;
    }

    .close-modal {
        position: absolute;
        top: 15px;
        right: 15px;
        background: none;
        border: none;
        font-size: 1.5rem;
        cursor: pointer;
        color: #666;
    }

    .donation-modal h2 {
        margin-bottom: 20px;
        color: #333;
    }

    .donation-summary {
        background: #f8f8f8;
        padding: 15px;
        border-radius: 8px;
        margin-bottom: 25px;
    }

    .donation-summary p {
        margin-bottom: 8px;
    }

    .donation-form {
        display: flex;
        flex-direction: column;
        gap: 15px;
    }

    .form-group {
        display: flex;
        flex-direction: column;
    }

    .form-row {
        display: flex;
        gap: 15px;
    }

    .form-group.half {
        width: 50%;
    }

    .form-group label {
        margin-bottom: 8px;
        color: #555;
        font-weight: 500;
    }

    .form-group input {
        padding: 12px;
        border: 1px solid #ddd;
        border-radius: 4px;
        font-size: 1rem;
    }

    .secure-notice {
        display: flex;
        align-items: center;
        gap: 8px;
        font-size: 0.9rem;
        color: #666;
        margin-top: 15px;
        justify-content: center;
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
        
        .form-row {
            flex-direction: column;
            gap: 15px;
        }
        
        .form-group.half {
            width: 100%;
        }
        
        .funding-amount h2 {
            font-size: 1.75rem;
        }
    }
</style>
