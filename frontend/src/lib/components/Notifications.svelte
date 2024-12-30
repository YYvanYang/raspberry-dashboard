<script lang="ts">
    import { notifications } from '$lib/stores/notifications';
    import { fly } from 'svelte/transition';
    import { flip } from 'svelte/animate';

    const icons = {
        success: `<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
        </svg>`,
        error: `<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>`,
        warning: `<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
        </svg>`,
        info: `<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>`
    };
</script>

<div class="fixed top-4 right-4 z-50 flex flex-col gap-2 min-w-[320px] max-w-[420px]">
    {#each $notifications as notification (notification.id)}
        <div
            animate:flip={{ duration: 200 }}
            transition:fly={{ x: 100, duration: 300 }}
            class="p-4 rounded-lg shadow-lg"
            class:bg-green-100={notification.type === 'success'}
            class:bg-red-100={notification.type === 'error'}
            class:bg-yellow-100={notification.type === 'warning'}
            class:bg-blue-100={notification.type === 'info'}
            role="alert"
        >
            <div class="flex items-start">
                <div class="flex-shrink-0">
                    {@html icons[notification.type]}
                </div>
                <div class="ml-3 w-0 flex-1">
                    <p 
                        class="text-sm font-medium"
                        class:text-green-800={notification.type === 'success'}
                        class:text-red-800={notification.type === 'error'}
                        class:text-yellow-800={notification.type === 'warning'}
                        class:text-blue-800={notification.type === 'info'}
                    >
                        {notification.message}
                    </p>
                </div>
                <div class="ml-4 flex-shrink-0 flex">
                    <button
                        type="button"
                        class="inline-flex text-gray-400 hover:text-gray-500 focus:outline-none"
                        on:click={() => notifications.remove(notification.id)}
                    >
                        <span class="sr-only">关闭</span>
                        <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                        </svg>
                    </button>
                </div>
            </div>
        </div>
    {/each}
</div> 