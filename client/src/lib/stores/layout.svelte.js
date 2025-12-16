import { BREAKPOINTS } from "$lib/config/constants.js";

let screenWidth = $state(0);

let isMobile = $derived(screenWidth <= BREAKPOINTS.mobile);
let isTablet = $derived(screenWidth <= BREAKPOINTS.tablet);
let isDesktop = $derived(screenWidth > BREAKPOINTS.tablet);

let screen = $derived(screenWidth <= BREAKPOINTS.mobile ? "mobile" : screenWidth <= BREAKPOINTS.tablet ? "tablet" : "desktop");

let components = $state({
	sidebar: false,
	messageBoard: false,
	panelVisible: true,
	viewVisible: true,
	panel: "conversations",
	view: "profile"
});

export const layout = {
	get screenWidth() {
		return screenWidth;
	},
	set screenWidth(width) {
		screenWidth = width;
	},
	get screen() {
		return screen;
	},
	get isMobile() {
		return isMobile;
	},
	get isTablet() {
		return isTablet;
	},
	get isDesktop() {
		return isDesktop;
	},
	get components() {
		return components;
	},
	updateComponents(newComponents) {
		Object.assign(components, newComponents);
	}
};
