// src/routes/dashboard/+layout.server.js

import { redirect } from "@sveltejs/kit";

export async function load({ locals }) {
	// 🛑 SECURITY GATE: Checks if user data was attached by the server hook
	if (!locals.user) {
		// If not authenticated, redirect to the login page
		throw redirect(302, "/login");
	}

	// Pass the user object to the +layout.svelte and all child pages (+page.svelte)
	return {
		user: locals.user
	};
}
