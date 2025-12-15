import { API_URL } from "$env/static/private";

const AUTH_COOKIE_NAME = "jwt";

export async function handle({ event, resolve }) {
	event.locals.user = undefined;

	let token = event.cookies.get(AUTH_COOKIE_NAME);

	if (!token) {
		return resolve(event);
	}

	try {
		const response = await fetch(`${API_URL}/auth/verify`, {
			headers: {
				Authorization: "Bearer " + token
			}
		});

		const body = await response.json();

		if (!response.ok) {
			console.warn(`[AUTH HOOK] Token invalid or expired. Status: ${response.status}`);
			event.cookies.delete(AUTH_COOKIE_NAME);
		}

		event.locals.user = {
			id: body.user_id,
			email: body.email
		};
	} catch (err) {
		console.error(`[AUTH HOOK] Verification failed (Network/Other): ${err}`);

		event.cookies.delete(AUTH_COOKIE_NAME, { path: "/" });
		event.locals.user = undefined;
	}

	return resolve(event);
}
