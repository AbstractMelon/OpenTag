<script lang="ts">
	import { onMount } from 'svelte';
	import { api } from '$lib/api';

	let serverStatus = $state('Disconnected');
	let serverMessage = $state('Waiting for connection...');

	onMount(async () => {
		try {
			const data = await api.getStatus();
			serverStatus = 'Connected';
			serverMessage = data.system || data.status;
		} catch (e) {
			serverStatus = 'Disconnected';
			serverMessage = 'Cannot reach backend server. Is it running?';
		}
	});
</script>

<main class="container">
	<h1>Opentag</h1>

	<div class="status-panel">
		<h2>
			Server Status: <span class={serverStatus === 'Connected' ? 'success' : 'error'}
				>{serverStatus}</span
			>
		</h2>
		<p>{serverMessage}</p>
	</div>
</main>
