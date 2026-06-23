<script lang="ts">
	import logo from '$lib/assets/kemkes.png';
	import './global.css';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={logo} />
</svelte:head>

{@render children()}
