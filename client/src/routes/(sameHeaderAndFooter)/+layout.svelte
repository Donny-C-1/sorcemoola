<script>
	import FixedNavbar from '$lib/components/FixedNavbar.svelte';
	import Footer from '$lib/components/Footer.svelte';

    let { children } = $props();
</script>

<FixedNavbar />
{@render children?.()}
<Footer />