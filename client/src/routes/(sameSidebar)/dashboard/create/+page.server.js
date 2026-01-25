import { API_URL } from "$env/static/private";
import { PUBLIC_CLOUDINARY_CLOUD_NAME, PUBLIC_CLOUDINARY_UPLOAD_PRESET } from "$env/static/public";
import { AUTH_COOKIE_NAME } from "$lib/config/constants";
import { fail, redirect } from "@sveltejs/kit";

async function uploadImageToCloudinary(file) {
	if (!(file instanceof File)) {
		console.error("Invalid file passed to Cloudinary:", file);
		return null;
	}

	const formData = new FormData();
	formData.append("file", file);
	formData.append("upload_preset", PUBLIC_CLOUDINARY_UPLOAD_PRESET);
	formData.append("folder", "campaigns");

	const response = await fetch({ method: "POST", body: formData });

	if (!response.ok) {
		const errorText = await response.text();
		throw new Error(errorText);
	}

	const data = await response.json();
	return data.secure_url;
}

export const actions = {
	createCampaign: async ({ request, locals, cookies }) => {
		console.log("=== ACTION STARTED ===");

		if (!locals.user) {
			console.log("No user found");
			return fail(401, { error: { message: "You must be logged in" } });
		}
		console.log("User authenticated:", locals.user.id);

		const data = await request.formData();
		console.log("FormData received");

		const token = cookies.get(AUTH_COOKIE_NAME);
		const imageFile = data.get("campaign_banner");
		console.log("Image file:", imageFile?.name, imageFile?.size);

		const payload = {
			campaignName: data.get("campaign_name"),
			description: data.get("description"),
			campaignStory: data.get("campaign_story"),
			category: data.get("category"),
			createdBy: locals.user.id,
			fundGoal: Math.round(parseFloat(data.get("fund_goal")) * 100),
			imageUrl: null
		};
		console.log("Payload built:", payload);

		if (!payload.campaignName || !payload.campaignStory || !payload.category || !payload.description) {
			console.log("Validation failed");
			return fail(422, {
				error: { label: "input", message: "Missing required fields" },
				campaignName: payload.campaignName,
				description: payload.description,
				campaignStory: payload.campaignStory,
				category: payload.category
			});
		}
		console.log("Validation passed");

		let campaignId;
		try {
			if (imageFile instanceof File && imageFile.size > 0) {
				console.log("Starting image upload...");
				const imageUrl = await uploadImageToCloudinary(imageFile);
				console.log("Image uploaded:", imageUrl);
				if (imageUrl) payload.imageUrl = imageUrl;
			} else {
				console.log("No image to upload");
			}

			console.log("Calling backend API...", payload);
			const response = await fetch(`${API_URL}/campaigns`, {
				method: "POST",
				body: JSON.stringify(payload),
				headers: {
					"Content-Type": "application/json",
					Authorization: `Bearer ${token}`
				}
			});
			console.log("Backend response status:", response.status);

			const body = await response.json();
			console.log("Backend response body:", body);

			if (!response.ok) {
				return fail(response.status, {
					error: { message: body.message || "Backend error" },
					campaignName: payload.campaignName,
					description: payload.description,
					campaignStory: payload.campaignStory,
					category: payload.category,
					fundGoal: payload.fundGoal
				});
			}

			campaignId = body.id;
		} catch (err) {
			console.error("Action Error:", err);
			return fail(500, {
				error: { message: "Server error. Please try again later." },
				campaignName: payload.campaignName,
				description: payload.description,
				campaignStory: payload.campaignStory,
				category: payload.category,
				fundGoal: payload.fundGoal
			});
		}

		console.log("Redirecting to:", `/campaigns/${campaignId}`);
		redirect(303, `/campaigns/${campaignId}`);
	}
};
