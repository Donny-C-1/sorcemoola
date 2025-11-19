<script>
  import { onMount } from 'svelte';
  import { testimonials } from '$lib/data/testimonials';
  
  let currentSlide = $state(0);
  let intervalId;
  
  function nextSlide() {
    currentSlide = (currentSlide + 1) % testimonials.length;
  }
  
  function prevSlide() {
    currentSlide = (currentSlide - 1 + testimonials.length) % testimonials.length;
  }
  
  function goToSlide(index) {
    currentSlide = index;
    resetInterval();
  }
  
  function resetInterval() {
    clearInterval(intervalId);
    intervalId = setInterval(nextSlide, 5000);
  }
  
  onMount(() => {
    resetInterval();
    return () => clearInterval(intervalId);
  });
</script>

<section id="success-stories" class="success-stories">
  <div class="container">
    <div class="section-title animate-on-scroll">
      <h2>Success Stories</h2>
      <p>Hear from creators and backers who have found success through SorceMoola.</p>
    </div>
    
    <div class="testimonial-slider animate-on-scroll">
      <button aria-label="something" class="slider-arrow prev" onclick={prevSlide}>
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="15 18 9 12 15 6"></polyline></svg>
      </button>
      
      <div class="testimonial-track">
        {#each testimonials as testimonial, index}
          <div 
            class="testimonial-slide" 
            class:active={index === currentSlide}
            style="transform: translateX({(index - currentSlide) * 100}%)"
          >
            <div class="testimonial-content">
              <div class="quote-mark">"</div>
              <p class="testimonial-text">{testimonial.quote}</p>
              <div class="testimonial-author">
                <img src={testimonial.avatar} alt={testimonial.name} loading="lazy" />
                <div class="author-details">
                  <div class="author-name">{testimonial.name}</div>
                  <div class="author-role">{testimonial.role}</div>
                  <div class="project-name">{testimonial.project}</div>
                </div>
              </div>
            </div>
          </div>
        {/each}
      </div>
      
      <button aria-label="something" class="slider-arrow next" onclick={nextSlide}>
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="9 18 15 12 9 6"></polyline></svg>
      </button>
    </div>
    
    <div class="testimonial-dots">
      {#each testimonials as _, index}
        <button 
          class="dot" 
          class:active={currentSlide === index}
          onclick={() => goToSlide(index)}
          aria-label={`Go to testimonial ${index + 1}`}
        ></button>
      {/each}
    </div>
  </div>
</section>

<style>
  .success-stories {
    background-color: var(--primary-bg);
    position: relative;
    overflow: hidden;
  }
  
  .success-stories::before {
    content: "";
    position: absolute;
    top: -30px;
    right: -30px;
    width: 200px;
    height: 200px;
    background-color: rgba(52, 162, 115, 0.1);
    border-radius: 50%;
    z-index: 0;
  }
  
  .testimonial-slider {
    position: relative;
    margin: 0 auto;
    max-width: 900px;
    overflow: hidden;
  }
  
  .testimonial-track {
    position: relative;
    height: 400px;
  }
  
  .testimonial-slide {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    transition: transform 0.5s ease;
  }
  
  .testimonial-content {
    background-color: white;
    border-radius: var(--radius-lg);
    box-shadow: var(--shadow-md);
    padding: var(--spacing-xl);
    height: 100%;
    display: flex;
    flex-direction: column;
    position: relative;
  }
  
  .quote-mark {
    font-size: 5rem;
    line-height: 1;
    position: absolute;
    top: var(--spacing-md);
    left: var(--spacing-md);
    color: rgba(52, 162, 115, 0.1);
    font-family: serif;
  }
  
  .testimonial-text {
    font-size: 1.25rem;
    line-height: 1.6;
    color: var(--text-dark);
    margin-top: var(--spacing-md);
    flex-grow: 1;
    position: relative;
    z-index: 1;
  }
  
  .testimonial-author {
    display: flex;
    align-items: center;
    margin-top: var(--spacing-lg);
  }
  
  .testimonial-author img {
    width: 60px;
    height: 60px;
    border-radius: 50%;
    object-fit: cover;
    margin-right: var(--spacing-md);
  }
  
  .author-name {
    font-weight: 600;
    color: var(--text-dark);
    font-size: 1.1rem;
  }
  
  .author-role {
    color: var(--primary-color);
    font-size: 0.875rem;
    margin-top: 2px;
  }
  
  .project-name {
    color: var(--text-medium);
    font-size: 0.875rem;
    margin-top: 2px;
  }
  
  .slider-arrow {
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
  
  .slider-arrow:hover {
    background-color: var(--primary-color);
    color: white;
  }
  
  .slider-arrow.prev {
    left: -20px;
  }
  
  .slider-arrow.next {
    right: -20px;
  }
  
  .testimonial-dots {
    display: flex;
    justify-content: center;
    margin-top: var(--spacing-lg);
    gap: 8px;
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
  
  @media (max-width: 768px) {
    .testimonial-track {
      height: 500px;
    }
    
    .slider-arrow {
      width: 30px;
      height: 30px;
    }
    
    .slider-arrow.prev {
      left: 10px;
    }
    
    .slider-arrow.next {
      right: 10px;
    }
  }
</style>