<script>
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';

	let { children } = $props();
</script>

<svelte:head>
	<title>SorceMoola - Fund Anything, Effortlessly</title>
	<link rel="icon" href={favicon} />
	<meta name="description" content="Discover and fund innovative campaigns on SorceMoola, the premier crowdfunding platform." />
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css">
	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorgin>
	<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600;700&display=swap">
</svelte:head>

{@render children?.()}
