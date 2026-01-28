<script>
  import Button from './ui/Button.svelte';
  
  let email = $state('');
  let submitted = $state(false);
  
  function handleSubscribe(e) {
    e.preventDefault();

    if (email && email.includes('@')) {
      submitted = true;
      email = '';
      setTimeout(() => {
        submitted = false;
      }, 3000);
    }
  }
</script>

<section id="call-to-action" class="cta-section">
  <div class="container">
    <div class="cta-content animate-on-scroll">
      <h2>Ready to Bring Your Idea to Life?</h2>
      <p>Join thousands of creators and backers on SorceMoola and be part of our thriving community.</p>
      
      <div class="cta-buttons">
        <Button href="/start" primary={true} inverted={true} large={true}>Start a Project</Button>
        <Button href="/explore" primary={false} inverted={true} large={true}>Back a Project</Button>
      </div>
      
      <div class="newsletter">
        <h3>Stay Updated</h3>
        <p>Subscribe to our newsletter for the latest projects and updates.</p>
        
        <form class="newsletter-form" onsubmit={handleSubscribe}>
          <input 
            type="email" 
            placeholder="Enter your email address" 
            bind:value={email} 
            required 
          />
          <Button primary={true} type="submit">Subscribe</Button>
        </form>
        
        {#if submitted}
          <div class="success-message">Thank you for subscribing!</div>
        {/if}
      </div>
    </div>
  </div>
</section>

<style>
  .cta-section {
    background-color: var(--primary-color);
    color: white;
    position: relative;
    overflow: hidden;
  }
  
  .cta-section::before {
    content: "";
    position: absolute;
    top: -100px;
    right: -100px;
    width: 300px;
    height: 300px;
    background-color: rgba(255, 255, 255, 0.1);
    border-radius: 50%;
  }
  
  .cta-section::after {
    content: "";
    position: absolute;
    bottom: -50px;
    left: -50px;
    width: 200px;
    height: 200px;
    background-color: rgba(255, 255, 255, 0.1);
    border-radius: 50%;
  }
  
  .cta-content {
    position: relative;
    z-index: 1;
    text-align: center;
    max-width: 800px;
    margin: 0 auto;
  }
  
  .cta-content h2 {
    font-size: 2.5rem;
    margin-bottom: var(--spacing-md);
  }
  
  .cta-content p {
    font-size: 1.25rem;
    margin-bottom: var(--spacing-lg);
    opacity: 0.9;
  }
  
  .cta-buttons {
    display: flex;
    justify-content: center;
    gap: var(--spacing-md);
    margin-bottom: var(--spacing-xl);
  }
  
  .newsletter {
    background-color: rgba(255, 255, 255, 0.1);
    border-radius: var(--radius-lg);
    padding: var(--spacing-lg);
    margin-top: var(--spacing-xl);
  }
  
  .newsletter h3 {
    margin-top: 0;
    margin-bottom: var(--spacing-xs);
  }
  
  .newsletter p {
    font-size: 1rem;
    margin-bottom: var(--spacing-md);
  }
  
  .newsletter-form {
    display: flex;
    gap: var(--spacing-sm);
  }
  
  .newsletter-form input {
    flex: 1;
    padding: var(--spacing-sm);
    border-radius: var(--radius-md);
    border: none;
    font-size: 1rem;
    font-family: 'Poppins', sans-serif;
  }
  
  .success-message {
    margin-top: var(--spacing-md);
    padding: var(--spacing-sm);
    background-color: rgba(255, 255, 255, 0.2);
    border-radius: var(--radius-md);
    animation: fadeIn 0.5s ease;
  }
  
  @keyframes fadeIn {
    from {
      opacity: 0;
      transform: translateY(10px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }
  
  @media (max-width: 768px) {
    .cta-content h2 {
      font-size: 2rem;
    }
    
    .cta-buttons {
      flex-direction: column;
    }
    
    .newsletter-form {
      flex-direction: column;
    }
  }
</style>