import { API_URL } from "$env/static/private";
import { AUTH_COOKIE_NAME } from "$lib/config/constants";
import { fail, redirect } from "@sveltejs/kit";

export const actions = {
	createCampaign: async ({ request, locals, cookies }) => {
		if (!locals.user) {
			return fail(400, {
				message: "You must be logged in to do this"
			});
		}

		const data = await request.formData();
		const token = cookies.get(AUTH_COOKIE_NAME);

		const payload = {
			campaignName: data.get("campaign_name"),
			campaignStory: data.get("campaign_story"),
			createdBy: locals.user.id,
			fundGoal: Math.round(parseFloat(data.get("fund_goal")) * 100)
		};

		// Validation
		if (!payload.campaignName || !payload.campaignStory) {
			return fail(422, {
				error: {
					label: "input",
					message: "Missing required input"
				},
				...payload
			});
		}

		// Fetch
		let campaign;
		try {
			const response = await fetch(`${API_URL}/campaigns`, {
				method: "POST",
				body: JSON.stringify(payload),
				headers: {
					"Content-Type": "application/json",
					Authorization: "Bearer " + token
				}
			});

			const body = await response.json();

			if (!response.ok) {
				return fail(422, {
					error: {
						message: "Error creating campaign",
						...body
					},
					...payload
				});
			}

			campaign = body;
		} catch (err) {
			return fail(500, {
				error: {
					label: "server",
					message: "Failed to create campaign",
					...err
				},
				...payload
			});
		}

		redirect(303, `/campaigns/${campaign.id}`);
	}
};
