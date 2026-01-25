import { redirect, fail } from "@sveltejs/kit";

import { API_URL } from "$env/static/private";

import { AUTH_COOKIE_NAME } from "$lib/config/constants.js";

const isDefined = API_URL ? API_URL : "http://localhost:8080/api/v1";

export const actions = {
	login: async ({ request, cookies }) => {
		const data = await request.formData();

		const payload = {
			email: data.get("email"),
			password: data.get("password")
		};
		let user, token;

		try {
			// const response = await fetch(`${API_URL}/auth/login`, {
			// 	method: "POST",
			// 	body: JSON.stringify(payload),
			// 	headers: {
			// 		"Content-Type": "application/json"
			// 	}
			// });
			console.log("this is API_URL", API_URL);
			const response = await fetch(`${isDefined}/auth/login`, {
				method: "POST",
				body: JSON.stringify(payload),
				headers: {
					"Content-Type": "application/json"
				}
			});

			let body;
			try {
				body = await response.json();
			} catch {
				body = {};
			}

			if (response.ok === false) {
				console.log(body);
				return fail(422, {
					error: body,
					...payload
				});
			}

			({ user, token } = body);
		} catch (err) {
			console.log(err);
			return fail(500, {
				error: {
					label: "server",
					message: err.message
				},
				...payload
			});
		}

		cookies.set(AUTH_COOKIE_NAME, token, {
			httpOnly: true,
			secure: true,
			sameSite: "lax",
			path: "/",
			maxAge: 60 * 60 * 24 * 7
		});
		redirect(303, "/dashboard");
	}
};
