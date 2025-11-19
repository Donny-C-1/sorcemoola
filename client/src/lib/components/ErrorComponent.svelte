<script>
  // Props
  export let message = "Sorry, we couldn't load the campaigns"; // Default error message
  export let subtext = "There was an issue connecting to our servers. Please try again."; // Additional context
  export let retryFunction = () => {}; // Function to call when retry button is clicked
  export let icon = "alert-circle"; // Default icon (you can use 'wifi-off', 'server-off', etc.)
</script>

<div class="error-container">
  <div class="error-card">
    <div class="error-icon">
      {#if icon === 'alert-circle'}
        <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="12" cy="12" r="10"></circle>
          <line x1="12" y1="8" x2="12" y2="12"></line>
          <line x1="12" y1="16" x2="12.01" y2="16"></line>
        </svg>
      {:else if icon === 'wifi-off'}
        <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <line x1="1" y1="1" x2="23" y2="23"></line>
          <path d="M16.72 11.06A10.94 10.94 0 0 1 19 12.55"></path>
          <path d="M5 12.55a10.94 10.94 0 0 1 5.17-2.39"></path>
          <path d="M10.71 5.05A16 16 0 0 1 22.58 9"></path>
          <path d="M1.42 9a15.91 15.91 0 0 1 4.7-2.88"></path>
          <path d="M8.53 16.11a6 6 0 0 1 6.95 0"></path>
          <line x1="12" y1="20" x2="12.01" y2="20"></line>
        </svg>
      {:else if icon === 'server-off'}
        <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M1 1l22 22"></path>
          <path d="M8 4h12a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H8"></path>
          <path d="M8 14h.01"></path>
          <path d="M8 20h.01"></path>
          <path d="M4 6a2 2 0 0 1 2-2"></path>
          <path d="M4 12a2 2 0 0 0 2 2h6"></path>
          <path d="M4 18a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2"></path>
        </svg>
      {:else}
        <!-- Default fallback icon -->
        <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="12" cy="12" r="10"></circle>
          <line x1="12" y1="8" x2="12" y2="12"></line>
          <line x1="12" y1="16" x2="12.01" y2="16"></line>
        </svg>
      {/if}
    </div>
    
    <h2 class="error-title">{message}</h2>
    <p class="error-message">{subtext}</p>
    
    <button class="retry-button" on:click={retryFunction}>
      <span class="retry-icon">
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M21.5 2v6h-6"></path>
          <path d="M2.5 12a10 10 0 0 1 5.5-9"></path>
          <path d="M2.5 12a10 10 0 0 0 10 10 10 10 0 0 0 4-0.8"></path>
          <path d="M20.3 14a10 10 0 0 0 1.2-2"></path>
          <path d="M21.5 8L17 12.5"></path>
        </svg>
      </span>
      Try Again
    </button>
  </div>
</div>

<style>
  .error-container {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 32px 16px;
    min-height: 400px;
    width: 100%;
  }
  
  .error-card {
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;
    max-width: 480px;
    padding: 40px 24px;
    background-color: white;
    border-radius: 16px;
    box-shadow: 0 8px 30px rgba(0, 0, 0, 0.06);
    animation: fadeIn 0.3s ease-out;
  }
  
  .error-icon {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 96px;
    height: 96px;
    border-radius: 50%;
    background-color: rgba(60, 185, 131, 0.1);
    color: #3cb983;
    margin-bottom: 24px;
  }
  
  .error-title {
    font-size: 24px;
    font-weight: 600;
    color: #333;
    margin: 0 0 12px 0;
    line-height: 1.3;
  }
  
  .error-message {
    font-size: 16px;
    line-height: 1.6;
    color: #666;
    margin: 0 0 32px 0;
  }
  
  .retry-button {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 8px;
    padding: 12px 24px;
    background-color: #3cb983;
    color: white;
    border: none;
    border-radius: 8px;
    font-size: 16px;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.2s ease;
    box-shadow: 0 4px 12px rgba(60, 185, 131, 0.2);
  }
  
  .retry-button:hover {
    background-color: #35a675;
    transform: translateY(-2px);
    box-shadow: 0 6px 16px rgba(60, 185, 131, 0.3);
  }
  
  .retry-button:active {
    transform: translateY(0);
    box-shadow: 0 2px 8px rgba(60, 185, 131, 0.2);
  }
  
  .retry-icon {
    display: flex;
    align-items: center;
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
  
  /* Responsive adjustments */
  @media (max-width: 600px) {
    .error-card {
      padding: 32px 16px;
    }
    
    .error-icon {
      width: 80px;
      height: 80px;
    }
    
    .error-icon svg {
      width: 40px;
      height: 40px;
    }
    
    .error-title {
      font-size: 20px;
    }
    
    .error-message {
      font-size: 15px;
    }
  }
</style>