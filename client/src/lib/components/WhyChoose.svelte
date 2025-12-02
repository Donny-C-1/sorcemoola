<script>
  import { features } from "$lib/data/features";
</script>

<section id="why-choose" class="why-choose">
  <div class="container">
    <div class="section-title animate-on-scroll">
      <h2>Why Choose SorceMoola</h2>
      <p>Our platform offers unique benefits to help creators and backers connect and collaborate effectively.</p>
    </div>
    
    <div class="features-grid">
      {#each features as feature, index}
        <div class="feature animate-on-scroll" style="transition-delay: {index * 150}ms">
          <div class="feature-icon">
            <img src={`/images/icons/${feature.icon}.svg`} alt={feature.title} loading="lazy" />
          </div>
          <h3>{feature.title}</h3>
          <p>{feature.description}</p>
        </div>
      {/each}
    </div>
  </div>
</section>

<style>
  .why-choose {
    background-color: white;
  }
  
  .features-grid {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: var(--spacing-xl);
  }
  
  .feature {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
  }
  
  .feature-icon {
    width: 70px;
    height: 70px;
    background-color: var(--primary-bg);
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-bottom: var(--spacing-md);
    transition: all 0.3s ease;
  }
  
  .feature:hover .feature-icon {
    background-color: var(--primary-color);
  }
  
  .feature-icon img {
    width: 35px;
    height: 35px;
    transition: all 0.3s ease;
  }
  
  .feature h3 {
    margin-top: 0;
    margin-bottom: var(--spacing-sm);
    color: var(--text-dark);
  }
  
  .feature p {
    color: var(--text-medium);
    margin: 0;
  }
  
  @media (max-width: 768px) {
    .features-grid {
      grid-template-columns: 1fr;
      gap: var(--spacing-lg);
    }
  }
</style>