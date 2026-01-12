import { API_URL } from "$env/static/private";
import { error, fail } from "@sveltejs/kit";

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

		return { campaign: body };
	} catch (err) {
		console.error("Campaign load error:", err);

		error(404, {
			message: "Failed to load campaign data",
			id: slug
		});
	}
}

export const actions = {
	contribute: async ({ request, params, locals }) => {
		const data = await request.formData();

		const payload = {
			amount: parseFloat(data.get("amount")) * 100,
			campaignID: params.campaignID,
			userID: locals?.user?.id
		};

		console.log(payload);

		try {
			const response = await fetch(`${API_URL}/campaigns/fund`, {
				method: "post",
				headers: {
					"Content-Type": "application/json"
				},
				body: JSON.stringify(payload)
			});

			const body = await response.json();
			console.log("Got here", body);

			if (!response.ok) return fail(400, { error: "Payment failed" });

			return {
				success: true
			};
		} catch (err) {
			console.log(err);
			fail(500, { message: err });
		}
	}
};
