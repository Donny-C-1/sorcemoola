<!-- src/components/ui/AccordionItem.svelte (continued) -->
<script>
  import { slide } from 'svelte/transition';
  
  let { title, expanded = false, transitionDelay = 0, children } = $props();
  
  let isExpanded = $state(expanded);
  
  function toggle() {
    isExpanded = !isExpanded;
  }
</script>

<div 
  class="accordion-item" 
  class:expanded={isExpanded}
  style="transition-delay: {transitionDelay}ms"
>
  <button class="accordion-header" onclick={toggle}>
    {title}
    <svg 
      class="icon" 
      xmlns="http://www.w3.org/2000/svg" 
      width="24" 
      height="24" 
      viewBox="0 0 24 24" 
      fill="none" 
      stroke="currentColor" 
      stroke-width="2" 
      stroke-linecap="round" 
      stroke-linejoin="round"
    >
      <polyline points="6 9 12 15 18 9"></polyline>
    </svg>
  </button>
  
  {#if isExpanded}
    <div class="accordion-content" transition:slide={{ duration: 300 }}>
      {@render children?.()}
    </div>
  {/if}
</div>

<style>
  .accordion-item {
    border-bottom: 1px solid #e0e0e0;
  }
  
  .accordion-item:last-child {
    border-bottom: none;
  }
  
  .accordion-header {
    width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: var(--spacing-md) var(--spacing-lg);
    background-color: white;
    border: none;
    cursor: pointer;
    font-family: 'Poppins', sans-serif;
    font-size: 1.125rem;
    font-weight: 600;
    color: var(--text-dark);
    text-align: left;
    transition: background-color 0.2s ease;
  }
  
  .accordion-header:hover {
    background-color: rgba(52, 162, 115, 0.05);
  }
  
  .expanded .accordion-header {
    background-color: rgba(52, 162, 115, 0.1);
    color: var(--primary-color);
  }
  
  .icon {
    transition: transform 0.3s ease;
    color: var(--primary-color);
  }
  
  .expanded .icon {
    transform: rotate(180deg);
  }
  
  .accordion-content {
    padding: var(--spacing-md) var(--spacing-lg);
    padding-top: 0;
    color: var(--text-medium);
    line-height: 1.6;
  }
  
  .accordion-content :global(p) {
    margin-top: 0;
  }
</style>