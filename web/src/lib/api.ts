async function fetchApi<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
	const url = endpoint.startsWith('/api')
		? endpoint
		: `/api${endpoint.startsWith('/') ? endpoint : `/${endpoint}`}`;
	console.log(`Fetching API: ${options.method || 'GET'} ${url}`);

	const response = await fetch(url, {
		...options,
		headers: {
			'Content-Type': 'application/json',
			...options.headers
		}
	});

	if (!response.ok) {
		let errorMessage = `API Error: ${response.status} ${response.statusText}`;
		try {
			const errorData = await response.json();
			if (errorData.error) {
				errorMessage = errorData.error;
			}
		} catch (e) {
			// Failed to parse error JSON, keep the default error message
		}
		throw new Error(errorMessage);
	}

	return response.json();
}

export const api = {
	get: <T>(endpoint: string) => fetchApi<T>(endpoint),

	post: <T>(endpoint: string, body: any) =>
		fetchApi<T>(endpoint, {
			method: 'POST',
			body: JSON.stringify(body)
		}),

	// Endpoints
	getStatus: () => fetchApi<{ status: string; system: string }>('/status')
};
