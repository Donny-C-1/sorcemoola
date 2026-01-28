<script>
	let { clickHandler, active, visible = true } = $props();
</script>

<button class="nav-toggle" class:visible aria-label="toggle nav" aria-expanded={active} tabindex="0" onclick={clickHandler}>
	<span class:first={active}></span>
	<span class:middle={active}></span>
	<span class:last={active}></span>
</button>

<style>
    button {
        display: none;
        flex-direction: column;
        gap: 5px;
        border: 0;
        z-index: 10000;
        padding: .5rem;
        cursor: pointer;
    }

    button.visible {
        display: flex;
    }

    button span {
        display: block;
        width: 25px;
        height: 3px;
        background-color: var(--text-dark);
        border-radius: 3px;
        transition: .3s ease-in-out;
    }

    button .first {
		transform: translateY(8px) rotate(45deg);
	}

	button .middle {
		opacity: 0;
	}

	button .last {
		transform: translateY(-8px) rotate(-45deg);
	}

</style>