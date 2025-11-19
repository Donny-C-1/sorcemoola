<!-- src/components/ui/ProgressBar.svelte -->
<script>
  let { value = 0, height = '6px' } = $props();
  
  let displayValue = $derived(Math.min(100, Math.max(0, value)));
</script>

<div class="progress-container" style="--bar-height: {height}">
  <div class="progress-bar" style="width: {displayValue}%"></div>
</div>

<style>
  .progress-container {
    width: 100%;
    height: var(--bar-height);
    background-color: #e0e0e0;
    border-radius: 100px;
    overflow: hidden;
  }
  
  .progress-bar {
    height: 100%;
    background-color: var(--primary-color);
    border-radius: 100px;
    transition: width 1s ease-in-out;
  }
</style>