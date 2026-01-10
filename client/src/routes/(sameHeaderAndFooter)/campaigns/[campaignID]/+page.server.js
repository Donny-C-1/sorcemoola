import { API_URL } from "$env/static/private";
import { error } from "@sveltejs/kit";

export async function load({ params }) {
	const slug = params.campaignID;

	try {
		const response = await fetch(`${API_URL}/campaigns/${slug}`);

		const body = await response.json();

		if (!response.ok) {
			throw error(response.status, {
				message: body.message || `Failed to load campaign: ${response.statusText}`,
				id: slug
			});
		}

		console.log("Body", body);

		return { campaign: body };
	} catch (err) {
		console.error("Campaign load error:", err);

		error(404, {
			message: "Failed to load campaign data",
			id: slug
		});
	}
}
