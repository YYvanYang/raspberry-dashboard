<script lang="ts">
    import { api } from "$lib/utils/api";

    const { services } = $props<{
        services: {
            name: string;
            running: boolean;
        }[];
    }>();

    async function handleServiceAction(service: string, action: string) {
        try {
            await api.post(`/system/services/${service}?action=${action}`);
            // 触发刷新
        } catch (error) {
            console.error(`Failed to ${action} service:`, error);
        }
    }
</script>

<div class="bg-white shadow overflow-hidden sm:rounded-lg">
    <div class="px-4 py-5 sm:px-6">
        <h3 class="text-lg leading-6 font-medium text-gray-900">
            Services Status
        </h3>
    </div>
    <div class="border-t border-gray-200">
        <ul class="divide-y divide-gray-200">
            {#each services as service}
                <li class="px-4 py-4 sm:px-6">
                    <div class="flex items-center justify-between">
                        <div class="flex items-center">
                            <div class={`w-2.5 h-2.5 rounded-full ${
                                service.running ? 'bg-green-400' : 'bg-red-400'
                            }`}></div>
                            <p class="ml-3 text-sm font-medium text-gray-900">
                                {service.name}
                            </p>
                        </div>
                        <div class="flex space-x-2">
                            <button
                                on:click={() => handleServiceAction(service.name, 'restart')}
                                class="inline-flex items-center px-3 py-1 border border-transparent text-sm leading-4 font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700"
                            >
                                Restart
                            </button>
                            {#if service.running}
                                <button
                                    on:click={() => handleServiceAction(service.name, 'stop')}
                                    class="inline-flex items-center px-3 py-1 border border-transparent text-sm leading-4 font-medium rounded-md text-white bg-red-600 hover:bg-red-700"
                                >
                                    Stop
                                </button>
                            {:else}
                                <button
                                    on:click={() => handleServiceAction(service.name, 'start')}
                                    class="inline-flex items-center px-3 py-1 border border-transparent text-sm leading-4 font-medium rounded-md text-white bg-green-600 hover:bg-green-700"
                                >
                                    Start
                                </button>
                            {/if}
                        </div>
                    </div>
                </li>
            {/each}
        </ul>
    </div>
</div> 