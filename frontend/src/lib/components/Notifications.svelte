<script lang="ts">
    import { notifications } from "$lib/stores/notifications";
    import { fade, fly } from "svelte/transition";

    let notificationsList: any[] = [];

    notifications.subscribe(value => {
        notificationsList = value;
    });
</script>

<div class="fixed top-4 right-4 z-50 space-y-2">
    {#each notificationsList as notification (notification.id)}
        <div
            transition:fly={{ x: 20 }}
            class="bg-white shadow-lg rounded-lg p-4 max-w-sm"
            class:bg-red-50={notification.type === 'error'}
            class:bg-green-50={notification.type === 'success'}
        >
            <div class="flex items-start">
                <div class="ml-3 w-0 flex-1">
                    <p class={`text-sm font-medium ${
                        notification.type === 'error' ? 'text-red-800' : 'text-green-800'
                    }`}>
                        {notification.message}
                    </p>
                </div>
                <button
                    class="ml-4 text-gray-400 hover:text-gray-500"
                    on:click={() => notifications.remove(notification.id)}
                >
                    <span class="sr-only">Close</span>
                    <svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
                    </svg>
                </button>
            </div>
        </div>
    {/each}
</div> 