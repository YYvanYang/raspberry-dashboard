<script lang="ts">
    import '../app.css';
    import ErrorBoundary from '$lib/components/ErrorBoundary.svelte';
    import Notifications from '$lib/components/Notifications.svelte';
    import type { LayoutData } from './$types';
    import type { Snippet } from 'svelte';
    
    const props = $props<{
        data: LayoutData;
        children: Snippet;
    }>();
</script>

<ErrorBoundary>
    <Notifications />
    {@render props.children()}
</ErrorBoundary> 