<script>
    let { primary = true, large = false, href = undefined, type = "button", children } = $props();
</script>

{#if href}
    <a {href} class:primary class:large>{@render children?.()}</a>
{:else}
    <button {type} class:primary class:large>{@render children?.()}</button>
{/if}


<style>
  button, a {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    font-family: 'Poppins', sans-serif;
    font-weight: 600;
    font-size: 1rem;
    padding: 12px 24px;
    border-radius: var(--radius-md);
    text-decoration: none;
    cursor: pointer;
    transition: all 0.2s ease-in-out;
    border: none;
    outline: none;
    position: relative;
    overflow: hidden;
    z-index: 1;
  }
  
  button::after, a::after {
    content: "";
    position: absolute;
    width: 100%;
    height: 100%;
    top: 0;
    left: -100%;
    background: linear-gradient(90deg, rgba(255,255,255,0) 0%, rgba(255,255,255,0.2) 50%, rgba(255,255,255,0) 100%);
    transition: left 0.5s ease;
    z-index: -1;
  }
  
  button:hover::after, a:hover::after {
    left: 100%;
  }
  
  .primary {
    background-color: var(--primary-color);
    color: white;
    box-shadow: 0 4px 15px rgba(52, 162, 115, 0.3);
  }
  
  .primary:hover {
    background-color: var(--primary-dark);
    transform: translateY(-2px);
    box-shadow: 0 6px 20px rgba(52, 162, 115, 0.4);
  }
  
  .primary:active {
    transform: translateY(1px);
    box-shadow: 0 2px 10px rgba(52, 162, 115, 0.3);
  }
  
  button:not(.primary), a:not(.primary) {
    background-color: transparent;
    color: var(--primary-color);
    border: 2px solid var(--primary-color);
  }
  
  button:not(.primary):hover, a:not(.primary):hover {
    background-color: rgba(52, 162, 115, 0.1);
    transform: translateY(-2px);
  }
  
  button:not(.primary):active, a:not(.primary):active {
    transform: translateY(1px);
  }
  
  .large {
    font-size: 1.125rem;
    padding: 14px 28px;
  }
  
  @media (max-width: 576px) {
    button, a {
      width: 100%;
    }
    
    .large {
      padding: 12px 24px;
      font-size: 1rem;
    }
  }
</style>