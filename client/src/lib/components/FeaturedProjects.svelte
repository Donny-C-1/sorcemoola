<!-- src/components/FeaturedProjects.svelte -->
<script>
  import { onMount } from 'svelte';
  import Button from './ui/Button.svelte';
  import ProjectCard from './ui/ProjectCard.svelte';
  import { featuredProjects } from '$lib/data/featured-projects';
  
  let currentIndex = $state(0);
  let isDesktop = $state(true);
  let cardsPerPage = $derived(isDesktop ? 3 : 1);
  let totalPages = $derived(Math.ceil(featuredProjects.length / cardsPerPage));
  
  function nextSlide() {
    currentIndex = (currentIndex + 1) % totalPages;
  }
  
  function prevSlide() {
    currentIndex = (currentIndex - 1 + totalPages) % totalPages;
  }
  
  function goToSlide(index) {
    currentIndex = index;
  }
  
  onMount(() => {
    const handleResize = () => {
      isDesktop = window.innerWidth >= 768;
    };
    
    window.addEventListener('resize', handleResize);
    handleResize();
    
    return () => {
      window.removeEventListener('resize', handleResize);
    };
  });
</script>

<section id="featured-projects" class="featured-projects">
  <div class="container">
    <div class="section-title animate-on-scroll">
      <h2>Featured Projects</h2>
      <p>Discover innovative projects from creators around the world that are currently seeking funding on SorceMoola.</p>
    </div>
    
    <div class="carousel-container animate-on-scroll">
      <button class="carousel-button prev" onclick={prevSlide} aria-label="Previous projects">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="15 18 9 12 15 6"></polyline></svg>
      </button>
      
      <div class="carousel">
        <div class="carousel-track" style="transform: translateX(-{currentIndex * 100}%)">
          {#each featuredProjects as project, index}
            <div class="carousel-slide" style="flex: 0 0 {100 / cardsPerPage}%">
              <ProjectCard {project} />
            </div>
          {/each}
        </div>
      </div>
      
      <button class="carousel-button next" onclick={nextSlide} aria-label="Next projects">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="9 18 15 12 9 6"></polyline></svg>
      </button>
    </div>
    
    <div class="carousel-dots">
      {#each Array(totalPages) as _, index}
        <button 
          class="dot" 
          class:active={currentIndex === index} 
          onclick={() => goToSlide(index)}
          aria-label={`Go to slide ${index + 1}`}
        ></button>
      {/each}
    </div>
    
    <div class="view-all-container animate-on-scroll">
      <Button href="/projects" primary={false}>View All Projects</Button>
    </div>
  </div>
</section>

<style>
  .featured-projects {
    background-color: var(--primary-bg);
  }
  
  .carousel-container {
    position: relative;
    margin-bottom: var(--spacing-lg);
  }
  
  .carousel {
    overflow: hidden;
    position: relative;
    margin: 0 40px;
  }
  
  .carousel-track {
    display: flex;
    transition: transform 0.5s ease-in-out;
  }
  
  .carousel-slide {
    padding: 0 var(--spacing-sm);
    box-sizing: border-box;
  }
  
  .carousel-button {
    position: absolute;
    top: 50%;
    transform: translateY(-50%);
    background-color: white;
    border: none;
    width: 40px;
    height: 40px;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    box-shadow: var(--shadow-md);
    z-index: 10;
    color: var(--primary-color);
    transition: all 0.2s ease;
  }
  
  .carousel-button:hover {
    background-color: var(--primary-color);
    color: white;
    box-shadow: var(--shadow-lg);
  }
  
  .carousel-button.prev {
    left: 0;
  }
  
  .carousel-button.next {
    right: 0;
  }
  
  .carousel-dots {
    display: flex;
    justify-content: center;
    gap: 8px;
    margin-bottom: var(--spacing-lg);
  }
  
  .dot {
    width: 12px;
    height: 12px;
    border-radius: 50%;
    background-color: var(--text-light);
    border: none;
    cursor: pointer;
    transition: all 0.2s ease;
    padding: 0;
  }
  
  .dot.active {
    background-color: var(--primary-color);
    transform: scale(1.2);
  }
  
  .view-all-container {
    text-align: center;
    margin-top: var(--spacing-xl);
  }
  
  @media (max-width: 576px) {
    .carousel {
      margin: 0 30px;
    }
    
    .carousel-button {
      width: 30px;
      height: 30px;
    }
  }
</style>