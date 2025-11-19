<script>
  let { category, style = '' } = $props();
</script>

<a href={`/categories/${category.slug}`} class="category-card" {style}>
  <div class="category-icon">
    <img src={`/images/icons/${category.icon}.svg`} alt={category.name} loading="lazy" />
  </div>
  <h3>{category.name}</h3>
  <div class="project-count">{category.projectCount} projects</div>
</a>

<style>
  .category-card {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: var(--spacing-lg);
    background-color: white;
    border-radius: var(--radius-lg);
    box-shadow: var(--shadow-md);
    transition: all 0.3s ease;
    text-decoration: none;
    color: var(--text-dark);
    text-align: center;
  }
  
  .category-card:hover {
    transform: translateY(-5px);
    box-shadow: var(--shadow-lg);
    background-color: var(--primary-bg);
  }
  
  .category-icon {
    width: 60px;
    height: 60px;
    background-color: var(--primary-bg);
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-bottom: var(--spacing-md);
    transition: all 0.3s ease;
  }
  
  .category-card:hover .category-icon {
    background-color: var(--primary-light);
  }
  
  .category-icon img {
    width: 30px;
    height: 30px;
    transition: all 0.3s ease;
  }
  
  .category-card:hover .category-icon img {
    filter: brightness(5);
  }
  
  .category-card h3 {
    margin-top: 0;
    margin-bottom: var(--spacing-xs);
    font-size: 1.25rem;
  }
  
  .project-count {
    color: var(--text-medium);
    font-size: 0.875rem;
  }
</style>