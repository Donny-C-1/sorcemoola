<!-- src/components/ui/Counter.svelte -->
<script>
  import { onMount } from 'svelte';
  
  let { value = 0, duration = 2000, prefix = "", suffix = "" } = $props();
  
  
  let count = $state(0);
  
  onMount(() => {
    const startTime = Date.now();
    
    function updateCount() {
      const currentTime = Date.now();
      const elapsedTime = currentTime - startTime;
      const progress = Math.min(elapsedTime / duration, 1);
      
      count = Math.floor(progress * value);
      
      if (progress < 1) {
        requestAnimationFrame(updateCount);
      } else {
        count = value;
      }
    }
    
    const observer = new IntersectionObserver((entries) => {
      if (entries[0].isIntersecting) {
        requestAnimationFrame(updateCount);
        observer.disconnect();
      }
    });
    
    observer.observe(document.querySelector('.stats'));
  });
  
  let formattedCount = $derived(count.toLocaleString());
</script>

<span>{prefix}{formattedCount}{suffix}</span>