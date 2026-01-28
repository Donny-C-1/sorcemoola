<script>
	import Logo from "$lib/components/ui/Logo.svelte";
	import Button from "$lib/components/ui/Button.svelte";
	import googleIcon from "$lib/assets/google-icon.svg?raw";
	import Icon from "$lib/components/ui/Icon.svelte";
	import { onMount } from "svelte";
	import { enhance } from "$app/forms";
	import Snackbar from "$lib/components/ui/Snackbar.svelte";

	let { form } = $props();

	let email = $state(form?.email ?? "");
	let password = $state("");

	let isSubmitting = $state(false);
	let showPassword = $state(false);
	let isSnackbarVisible = $state(false);

	$effect(() => {
		if (form) {
			isSubmitting = false;
		}

		if (form?.error?.label === "server") {
			isSnackbarVisible = true;
		}
	});

	function handleSubmit(e) {
		if (isSubmitting) {
			e.preventDefault();
			return;
		}
		isSubmitting = true;
	}
</script>

<svelte:head>
	<title>Log in - SorceMoola</title>
</svelte:head>

<main>
	<form action="?/login" method="post" onsubmit={handleSubmit} use:enhance>
		<div class="form_header">
			<h1>Join <Logo size="1.8rem" /></h1>
			<p class="title">Welcome back</p>
		</div>

		<div class="form_field">
			<label for="email">Email</label>
			<input type="email" name="email" id="email" placeholder="example@mail.com" bind:value={email} autocomplete="email" maxlength="100" required />
			{#if form?.error.label === "email"}
				<p class="error_message">{form?.error?.message}</p>
			{/if}
		</div>

		<div class="form_field">
			<label for="password">Password</label>
			<div class="password_container">
				<input type={showPassword ? "text" : "password"} name="password" id="password" placeholder="Create a strong password" bind:value={password} autocomplete="new-password" maxlength="128" required />
				<button type="button" class="toggle_password" aria-label="show password" onclick={() => (showPassword = !showPassword)}><i class="fa-regular {showPassword ? 'fa-eye' : 'fa-eye-slash'}"></i></button>
			</div>
		</div>

		<Button type="submit" wide="true" large="true" disabled={isSubmitting}>
			{#if isSubmitting}
				<span class="spinner"></span>
				Logging in
			{:else}
				Login
			{/if}
		</Button>
		<div class="divider"><span>or</span></div>

		<Button wide={true} large={true} primary={false}><Icon icon={googleIcon} />Sign in with Google</Button>

		<p class="login_prompt">New to SorceMoola? <a href="/signup">Sign up</a></p>
	</form>

	<Snackbar bind:visible={isSnackbarVisible} type={"error"}>{form?.error?.message || "Cannot connect to server"}</Snackbar>
</main>

<style>
	main {
		display: flex;
		justify-content: center;
		font-family: Poppins;
    background-color: var(--primary-bg);
	}

	form {
		box-shadow: var(--shadow-md);
		width: 100%;
		max-width: 42rem;
		padding: var(--spacing-xl) var(--spacing-sm);
		border-radius: var(--radius-lg);
    background-color: var(--text-white);

		@media screen and (min-width: 45rem) {
			padding: var(--spacing-xl);
			margin-block: var(--spacing-lg);
		}
	}

	.form_header {
		text-align: center;
	}

	h1 {
		display: flex;
		justify-content: center;
		font-weight: 700;
		gap: 0.8rem;
		font-size: 1.8rem;
		color: var(--text-medium);
		margin: 0;
	}

	.title {
		font-size: 1.1rem;
		margin: var(--spacing-sm) 0;
	}

	.form_field {
		margin-bottom: var(--spacing-md);
	}

	label {
		display: block;
		margin-bottom: var(--spacing-xs);
	}

	input:is([type="text"], [type="email"], [type="password"]) {
		width: 100%;
		padding: var(--spacing-sm);
		border: 1px solid var(--neutral-color);
		border-radius: var(--radius-sm);
		transition: border 0.2s ease;
	}

	.error_message {
		display: block;
		color: #e53935;
		font-size: 0.875rem;
		margin-top: var(--spacing-xs);
	}

	.password_container {
		position: relative;
	}

	.toggle_password {
		position: absolute;
		right: var(--spacing-sm);
		top: 50%;
		translate: 0 -50%;
		border: 0;
	}

	.spinner {
		display: inline-block;
		width: 1rem;
		height: 1rem;
		margin-right: var(--spacing-sm);
		border: 2px solid rgba(255, 255, 255, 0.3);
		border-radius: 50%;
		border-top-color: var(--text-white);
		animation: spin 0.8s linear infinite;
	}

	@keyframes spin {
		to {
			transform: rotate(360deg);
		}
	}

	.divider {
		display: flex;
		align-items: center;
		text-align: center;
		margin: var(--spacing-sm) 0;
	}

	.divider::before,
	.divider::after {
		content: "";
		flex: 1;
		border-bottom: 1px solid var(--neutral-color);
	}

	.divider span {
		padding: 0 var(--spacing-sm);
		color: var(--text-light);
		font-size: 0.875rem;
	}

	.login_prompt {
		text-align: center;
		margin-top: var(--spacing-lg);
		color: var(--text-medium);
	}

	.login_prompt a {
		color: var(--primary-color);
		font-weight: 500;
	}

	.login_prompt a:hover {
		text-decoration: underline;
	}
</style>
