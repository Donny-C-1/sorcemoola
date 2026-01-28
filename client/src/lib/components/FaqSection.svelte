<script>
  import Accordion from './ui/Accordion.svelte';
  import AccordionItem from './ui/AccordionItem.svelte';
  import { faqs } from '$lib/data/faq';
</script>

<section id="faq" class="faq-section">
  <div class="container">
    <div class="section-title animate-on-scroll">
      <h2>Frequently Asked Questions</h2>
      <p>Find answers to common questions about SorceMoola and how our platform works.</p>
    </div>
    
    <div class="faq-container animate-on-scroll">
      <Accordion>
        {#each faqs as faq, index}
          <AccordionItem 
            title={faq.question} 
            expanded={index === 0}
            transitionDelay={index * 100}
          >
            <p>{faq.answer}</p>
          </AccordionItem>
        {/each}
      </Accordion>
    </div>
    
    <div class="support-cta animate-on-scroll">
      <p>Still have questions? Our support team is here to help.</p>
      <a href="/contact" class="support-link">Contact Support</a>
    </div>
  </div>
</section>

<style>
  .faq-section {
    background-color: var(--primary-bg);
  }
  
  .faq-container {
    max-width: 800px;
    margin: 0 auto;
  }
  
  .support-cta {
    text-align: center;
    margin-top: var(--spacing-xl);
  }
  
  .support-cta p {
    font-size: 1.125rem;
    color: var(--text-medium);
    margin-bottom: var(--spacing-md);
  }
  
  .support-link {
    display: inline-block;
    color: var(--primary-color);
    font-weight: 600;
    text-decoration: none;
    padding-bottom: 2px;
    border-bottom: 2px solid var(--primary-color);
    transition: all 0.2s ease;
  }
  
  .support-link:hover {
    color: var(--primary-dark);
    border-bottom-color: var(--primary-dark);
  }
</style>