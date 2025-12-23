// src/routes/logout/+server.js

import { redirect } from "@sveltejs/kit";
import { AUTH_COOKIE_NAME } from "$lib/config/constants";

export async function GET({ cookies, locals }) {
	cookies.delete(AUTH_COOKIE_NAME, { path: "/" });

	locals.user = undefined;

	throw redirect(302, "/");
}
