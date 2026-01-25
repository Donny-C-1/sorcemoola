<script>
	import "../app.css";
	import favicon from "$lib/assets/favicon.svg";
	import { onMount } from "svelte";
	import { layout } from "$lib/stores/layout.svelte";

	let { children } = $props();

	onMount(() => {
		if (typeof window === "undefined") return;

		const update = () => (layout.screenWidth = window.innerWidth);

		update();

		window.addEventListener("resize", update);

		return () => {
			window.removeEventListener("resize", update);
		};
	});
</script>

<svelte:head>
	<title>SorceMoola - Fund Anything, Effortlessly</title>
	<link rel="icon" href={favicon} />
	<meta name="description" content="Discover and fund innovative campaigns on SorceMoola, the premier crowdfunding platform." />
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css" />
	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorgin />
	<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600;700&display=swap" />
	<link href="https://fonts.googleapis.com/css2?family=Aleo:ital,wght@0,100..900;1,100..900&family=Nunito+Sans:ital,opsz,wght@0,6..12,200..1000;1,6..12,200..1000&family=Open+Sans:ital,wght@0,300..800;1,300..800&display=swap" rel="stylesheet" />
	<link href="https://fonts.googleapis.com/css2?family=Quicksand:wght@300..700&display=swap" rel="stylesheet" />
</svelte:head>

{@render children?.()}
