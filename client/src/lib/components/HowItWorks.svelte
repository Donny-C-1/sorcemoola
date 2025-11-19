<!-- src/components/HowItWorks.svelte -->
<script>
  import { onMount } from 'svelte';
  import Counter from './ui/Counter.svelte';
  import { stats } from '$lib/data/stats';
  
  let steps = [
    {
      icon: 'create',
      title: 'Create',
      description: 'Share your idea with our community. Create a detailed project page with goals, rewards, and your story.'
    },
    {
      icon: 'fund',
      title: 'Fund',
      description: 'Connect with backers who believe in your vision. Receive funding to bring your project to life.'
    },
    {
      icon: 'launch',
      title: 'Launch',
      description: 'Deliver on your promises. Launch your project and stay connected with your supportive community.'
    }
  ];
</script>

<section id="how-it-works" class="how-it-works">
  <div class="container">
    <div class="section-title animate-on-scroll">
      <h2>How It Works</h2>
      <p>SorceMoola makes it easy to bring your creative ideas to life through our simple three-step process.</p>
    </div>
    
    <div class="steps">
      {#each steps as step, index}
        <div class="step animate-on-scroll" style="transition-delay: {index * 150}ms">
          <div class="step-icon">
            <img src="/images/icons/{step.icon}.svg" alt={step.title} />
            <div class="step-number">{index + 1}</div>
          </div>
          <h3>{step.title}</h3>
          <p>{step.description}</p>
        </div>
      {/each}
    </div>
    
    <div class="stats animate-on-scroll">
      {#each stats as stat}
        <div class="stat-item">
          <Counter value={stat.value} duration={2000} />
          <div class="stat-label">{stat.label}</div>
        </div>
      {/each}
    </div>
  </div>
</section>

<style>
  .how-it-works {
    background-color: white;
  }
  
  .steps {
    display: flex;
    justify-content: space-between;
    gap: var(--spacing-lg);
    margin-bottom: var(--spacing-2xl);
  }
  
  .step {
    flex: 1;
    text-align: center;
    padding: var(--spacing-lg);
    border-radius: var(--radius-lg);
    background-color: white;
    box-shadow: var(--shadow-md);
    transition: transform 0.3s ease, box-shadow 0.3s ease;
  }
  
  .step:hover {
    transform: translateY(-10px);
    box-shadow: var(--shadow-lg);
  }
  
  .step-icon {
    position: relative;
    width: 100px;
    height: 100px;
    margin: 0 auto var(--spacing-md);
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: var(--primary-bg);
    border-radius: 50%;
  }
  
  .step-icon img {
    width: 50px;
    height: 50px;
  }
  
  .step-number {
    position: absolute;
    top: 0;
    right: 0;
    width: 30px;
    height: 30px;
    background-color: var(--primary-color);
    color: white;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-weight: bold;
  }
  
  .step h3 {
    margin-bottom: var(--spacing-sm);
    color: var(--primary-dark);
  }
  
  .step p {
    color: var(--text-medium);
    font-size: 1rem;
  }
  
  .stats {
    display: flex;
    justify-content: space-around;
    background-color: var(--primary-bg);
    padding: var(--spacing-xl);
    border-radius: var(--radius-lg);
    margin-top: var(--spacing-xl);
  }
  
  .stat-item {
    text-align: center;
  }
  
  .stat-item :global(span) {
    font-size: 2.5rem;
    font-weight: 700;
    color: var(--primary-dark);
    display: block;
  }
  
  .stat-label {
    font-size: 1.125rem;
    color: var(--text-medium);
    margin-top: var(--spacing-xs);
  }
  
  @media (max-width: 992px) {
    .steps {
      flex-direction: column;
      gap: var(--spacing-md);
    }
    
    .step {
      max-width: 500px;
      margin: 0 auto;
    }
  }
  
  @media (max-width: 768px) {
    .stats {
      flex-direction: column;
      gap: var(--spacing-lg);
    }
    
    .stat-item {
      margin-bottom: var(--spacing-md);
    }
  }
</style>