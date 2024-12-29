<script lang="ts">
    import { notifications } from '$lib/stores/notifications';
    import { dev } from '$app/environment';
    import type { Snippet } from 'svelte';
    
    let error = $state<Error | null>(null);
    let { children } = $props<{ children: Snippet }>();
    
    $effect(() => {
        const handleError = (event: ErrorEvent) => {
            error = event.error || new Error(event.message);
            console.error('Component Error:', error);
            notifications.error('An unexpected error occurred');
            return false;
        };

        window.addEventListener('error', handleError);
        
        return () => {
            window.removeEventListener('error', handleError);
        };
    });
</script>

{#if error}
    <div class="bg-red-50 p-4 rounded-md">
        <p class="text-red-700">Something went wrong</p>
        {#if dev}
            <pre class="mt-2 text-sm text-red-600">{error.stack}</pre>
        {/if}
    </div>
{:else}
    {@render children()}
{/if} 