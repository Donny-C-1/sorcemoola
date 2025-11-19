<!-- src/components/FaqSection.svelte -->
<script>
  import Accordion from './ui/Accordion.svelte';
  import AccordionItem from './ui/AccordionItem.svelte';
  
  let faqs = [
    {
      question: "How does SorceMoola's funding work?",
      answer: "SorceMoola uses an all-or-nothing funding model. This means project creators set a funding goal and a deadline. If the project reaches its funding goal, backers' credit cards are charged when the deadline arrives. If the project falls short, no one is charged."
    },
    {
      question: "What fees does SorceMoola charge?",
      answer: "SorceMoola charges a 5% platform fee on the total amount raised, plus payment processing fees (typically 3-5%). These fees only apply if your project reaches its funding goal."
    },
    {
      question: "Can I launch a project from anywhere in the world?",
      answer: "Yes! SorceMoola is available globally. However, you'll need a bank account in a supported country to receive funds. Check our help center for the full list of supported countries."
    },
    {
      question: "What happens if a project doesn't reach its funding goal?",
      answer: "If a project doesn't reach its funding goal by the deadline, no backers are charged, and the project creator receives no funds. You can always relaunch your project with improvements based on feedback."
    },
    {
      question: "As a backer, when am I charged?",
      answer: "As a backer, your card is only charged if the project reaches its funding goal. The charge occurs when the project's funding deadline arrives, not when you make the pledge."
    },
    {
      question: "What kind of projects can be funded on SorceMoola?",
      answer: "SorceMoola welcomes creative projects in art, comics, crafts, dance, design, fashion, film, food, games, journalism, music, photography, publishing, technology, and theater. Projects should create something to share with others and have a clear plan with a defined beginning and end."
    },
    {
      question: "How do I get help with my project?",
      answer: "SorceMoola offers personalized support for project creators. You can contact our Creator Success team at any time for guidance on project setup, marketing strategies, and fulfillment planning."
    }
  ];
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