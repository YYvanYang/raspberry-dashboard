<script lang="ts">
    import { goto } from '$app/navigation';
    import { onMount } from 'svelte';

    onMount(() => {
        goto('/dashboard');
    });
</script>

<div class="flex h-screen items-center justify-center">
    <div class="text-gray-600">正在跳转到仪表盘...</div>
</div> 