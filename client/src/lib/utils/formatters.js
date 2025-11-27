export function currencyFormat(value, currencyCode, locale = 'en-US') {
	try {
		return new Intl.NumberFormat(locale, {
			style: 'currency',
			currency: currencyCode,
			minimumFractionDigits: 0, // Show whole numbers for goals
			maximumFractionDigits: 0
		}).format(value);
	} catch (error) {
		console.error('Error formatting currency:', error);
		// Fallback to a simple string if formatting fails
		return `${currencyCode} ${value.toLocaleString()}`;
	}
}
