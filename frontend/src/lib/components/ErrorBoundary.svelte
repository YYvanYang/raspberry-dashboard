<script lang="ts">
    import { onError } from 'svelte';
    import { notifications } from '$lib/stores/notifications';
    
    let error: Error | null = null;
    
    onError((e) => {
        error = e;
        console.error('Component Error:', e);
        notifications.error('An unexpected error occurred');
    });
</script>

{#if error}
    <div class="bg-red-50 p-4 rounded-md">
        <p class="text-red-700">Something went wrong</p>
        {#if import.meta.env.DEV}
            <pre class="mt-2 text-sm text-red-600">{error.stack}</pre>
        {/if}
    </div>
{:else}
    <slot />
{/if} 