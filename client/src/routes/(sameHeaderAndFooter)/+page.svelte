<script>
    import Hero from "$lib/components/Hero.svelte";
    import HowItWorks from "$lib/components/HowItWorks.svelte";
    import FeaturedProjects from "$lib/components/FeaturedProjects.svelte";
    import Categories from "$lib/components/Categories.svelte";
    import SuccessStories from "$lib/components/SuccessStories.svelte";
    import WhyChoose from "$lib/components/WhyChoose.svelte";
    import CallToAction from "$lib/components/CallToAction.svelte";
    import FaqSection from "$lib/components/FaqSection.svelte";
    import { onMount } from "svelte";

     onMount(() => {
    // Add intersection observer for scroll animations
    const observer = new IntersectionObserver((entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          entry.target.classList.add('animate-in');
          observer.unobserve(entry.target);
        }
      });
    }, { threshold: 0.1 });
    
    document.querySelectorAll('.animate-on-scroll').forEach((el) => {
      observer.observe(el);
    });
  });
</script>

<main>
    <Hero />
    <HowItWorks />
    <FeaturedProjects />
    <Categories />
    <SuccessStories />
    <WhyChoose />
    <CallToAction />
    <FaqSection />
</main>