import { API_URL } from "$env/static/private";

export async function load({ locals }) {
	try {
		const response = await fetch(`${API_URL}/users/${locals.user.id}/campaigns`);

		let body = await response.json();

		if (!response.ok) {
			body = null;
		}

		return { campaigns: body.campaigns };
	} catch (err) {
		console.log(err);
		return null;
	}
}
