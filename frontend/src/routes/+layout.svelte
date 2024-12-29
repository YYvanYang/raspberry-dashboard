<script lang="ts">
    import Notifications from '$lib/components/Notifications.svelte';
    
    const slots = $props<{
        default?: () => unknown;
    }>();
</script>

<Notifications />
{@render slots.default?.()} 