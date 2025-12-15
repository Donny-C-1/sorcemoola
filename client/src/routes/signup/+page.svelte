<script>
	import Logo from "$lib/components/ui/Logo.svelte";
	import Button from "$lib/components/ui/Button.svelte";
	import googleIcon from "$lib/assets/google-icon.svg?raw";
	import Icon from "$lib/components/ui/Icon.svelte";
	import { onMount } from "svelte";
	import { enhance } from "$app/forms";
	import Snackbar from "$lib/components/ui/Snackbar.svelte";

	let { form } = $props();

	let fullname = $state(form?.fullName ?? "");
	let email = $state(form?.email ?? "");
	let password = $state("");
	let accountType = $state(form?.accountType ?? "Individual");

	let isSubmitting = $state(false);
	let showPassword = $state(false);
	let isSnackbarVisible = $state(false);

	onMount(() => {
		document.getElementById("individual").click();
	});

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

	function changeAccountType(e) {
		let label = this.nextElementSibling;
		let slide = this.parentElement;

		slide.style.setProperty("--left", label.offsetLeft + "px");
		slide.style.setProperty("--width", label.offsetWidth + "px");
	}
</script>

<svelte:head>
	<title>Sign up - SorceMoola</title>
</svelte:head>

<main>
	<form action="?/create" method="post" onsubmit={handleSubmit} use:enhance>
		<div class="form_header">
			<h1>Join <Logo size="1.8rem" /></h1>
			<p class="title">Start your crowdfunding journey today</p>
		</div>

		<div class="account_type">
			<input type="radio" name="accountType" id="individual" value="Individual" oninput={changeAccountType} bind:group={accountType} hidden />
			<label for="individual">Individual</label>
			<input type="radio" name="accountType" id="ngo" value="NGO" oninput={changeAccountType} bind:group={accountType} hidden />
			<label for="ngo">NGO</label>
			<input type="radio" name="accountType" id="organization" value="Organization" oninput={changeAccountType} bind:group={accountType} hidden />
			<label for="organization">Organization</label>
		</div>
		{#if accountType === "Individual"}
			<p>Grow Your Vision. Perfect for personal skill development, funding a creative project, or overcoming a temporary personal challenge. Supports self-growth and immediate needs.</p>
		{:else if accountType === "NGO"}
			<p>Fund Your Mission. Dedicated to officially registered non-profits, charities, and community aid organizations. Enables high-impact campaigns and tax-deductible giving.</p>
		{:else if accountType === "Organization"}
			<p>Expand Your Impact. For socially conscious enterprises, educational institutions, or businesses seeking funds to develop community-focused products or services.</p>
		{/if}

		<div class="form_field">
			<label for="name">{accountType === "Individual" ? "Full Name" : "Organization Name"}</label>
			<input type="text" name="fullName" id="name" placeholder={accountType === "Individual" ? "John Doe" : "Acme Corp"} bind:value={fullname} autocomplete="name" maxlength="255" required />
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
				Creating Account
			{:else}
				Create Account
			{/if}
		</Button>
		<div class="divider"><span>or</span></div>

		<Button wide={true} large={true} primary={false}><Icon icon={googleIcon} />Sign up with Google</Button>

		<p class="login_prompt">Already have an account? <a href="/login">Log In</a></p>
	</form>

	<Snackbar bind:visible={isSnackbarVisible} type={"error"}>{form?.error?.message || "Cannot connect to server"}</Snackbar>
</main>

<style>
	main {
		display: flex;
		justify-content: center;
		font-family: Poppins;
	}

	form {
		box-shadow: var(--shadow-md);
		width: 100%;
		max-width: 42rem;
		padding: var(--spacing-xl) var(--spacing-sm);
		border-radius: var(--radius-lg);

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

	.account_type {
		display: flex;
		justify-content: space-between;
		position: relative;
		background-color: var(--primary-bg);
		padding: var(--spacing-xs);
		margin-bottom: var(--spacing-md);
		border-radius: var(--radius-md);
	}

	.account_type::before {
		content: "";
		position: absolute;
		z-index: 0;
		background-color: var(--text-white);
		box-shadow: var(--shadow-sm);
		border-radius: var(--radius-sm);
		left: var(--left, 1%);
		width: var(--width, 31%);
		top: var(--spacing-xs);
		height: calc(100% - var(--spacing-xs) * 2);
		transition: left 0.3s ease;
	}

	.account_type label {
		flex: 1;
		text-align: center;
		margin: 0;
		padding: var(--spacing-xs);
		z-index: 1;
		cursor: pointer;
		font-weight: 600;
		font-size: 1.1rem;
		transition: 0.3s ease;
	}

	.account_type input:checked + label {
		color: var(--primary-color);
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
