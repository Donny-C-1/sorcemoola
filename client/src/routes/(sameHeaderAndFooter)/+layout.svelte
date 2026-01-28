<script>
	import FixedNavbar from '$lib/components/FixedNavbar.svelte';
	import Footer from '$lib/components/Footer.svelte';

    let { children, data } = $props();
</script>


<FixedNavbar isLoggedIn={data.user} />
{@render children?.()}
<Footer />