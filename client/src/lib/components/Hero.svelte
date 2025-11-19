<!-- src/components/Hero.svelte -->
<script>
  import Button from './ui/Button.svelte';
  
  function handleExplore() {
    // Scroll to featured projects section
    document.getElementById('featured-projects').scrollIntoView({ behavior: 'smooth' });
  }
</script>

<section class="hero-section">
  <div class="overlay"></div>
  <div class="parallax-bg"></div>
  <div class="container hero-content">
    <div class="hero-text">
      <h1>Fund Your Dreams, Support Others</h1>
      <p>SorceMoola connects innovative creators with passionate backers to bring amazing ideas to life. Join our thriving community today.</p>
      <div class="hero-buttons">
        <Button primary={true} large={true} href="/start-project">Start Your Project</Button>
        <Button primary={false} large={true} on:click={handleExplore}>Explore Projects</Button>
      </div>
    </div>
    <div class="hero-image">
      <img src="https://images.unsplash.com/photo-1761370981247-1dfd749ec96b?ixid=M3w4MjcwNjd8MHwxfHNlYXJjaHw5fHxCbGFjayUyMFBlb3BsZSUyMGNvbGxhYm9yYXRpbmclMjBvbiUyMHByb2plY3RzfGVufDB8fHx8MTc2MzA0MjI4NXww&ixlib=rb-4.1.0&fit=max&q=80" alt="People collaborating on projects" loading="eager" />
    </div>
  </div>
</section>

<style>
  .hero-section {
    position: relative;
    height: 90vh;
    min-height: 600px;
    display: flex;
    align-items: center;
    overflow: hidden;
    background-color: var(--primary-bg);
    margin-top: -1px; /* Fix for potential gap */
  }
  
  .overlay {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: linear-gradient(135deg, rgba(52, 162, 115, 0.1) 0%, rgba(52, 162, 115, 0.2) 100%);
    z-index: 1;
  }
  
  .parallax-bg {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-image: url('/images/dots-pattern.svg');
    opacity: 0.4;
    z-index: 0;
    transform: translateY(0);
    transition: transform 0.5s ease-out;
  }
  
  .hero-section:hover .parallax-bg {
    transform: translateY(-20px);
  }
  
  .hero-content {
    position: relative;
    z-index: 2;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  
  .hero-text {
    flex: 1;
    max-width: 600px;
    padding-right: var(--spacing-xl);
  }
  
  .hero-text h1 {
    color: var(--text-dark);
    margin-bottom: var(--spacing-md);
    font-size: 3.5rem;
    font-weight: 700;
    line-height: 1.2;
  }
  
  .hero-text p {
    font-size: 1.25rem;
    color: var(--text-medium);
    margin-bottom: var(--spacing-lg);
  }
  
  .hero-buttons {
    display: flex;
    gap: var(--spacing-md);
    flex-wrap: wrap;
  }
  
  .hero-image {
    flex: 1;
    max-width: 500px;
    display: flex;
    justify-content: flex-end;
    animation: float 6s ease-in-out infinite;
  }
  
  .hero-image img {
    width: 100%;
    height: auto;
  }
  
  @keyframes float {
    0% {
      transform: translateY(0px);
    }
    50% {
      transform: translateY(-20px);
    }
    100% {
      transform: translateY(0px);
    }
  }
  
  @media (max-width: 992px) {
    .hero-content {
      flex-direction: column-reverse;
      text-align: center;
    }
    
    .hero-text {
      max-width: 100%;
      padding-right: 0;
      margin-top: var(--spacing-xl);
    }
    
    .hero-buttons {
      justify-content: center;
    }
    
    .hero-image {
      max-width: 400px;
      margin-top: var(--spacing-xl);
      justify-content: center;
    }
  }
  
  @media (max-width: 576px) {
    .hero-section {
      height: auto;
      padding-top: var(--spacing-xl);
      padding-bottom: var(--spacing-xl);
    }
    
    .hero-text h1 {
      font-size: 2.5rem;
    }
    
    .hero-image {
      max-width: 280px;
    }
    
    .hero-buttons {
      flex-direction: column;
      gap: var(--spacing-sm);
    }
  }
</style>