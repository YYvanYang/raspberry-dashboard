<script lang="ts">
    import { onMount } from "svelte";
    import { api } from "$lib/utils/api";

    let services = $state<any[]>([]);
    let loading = $state(true);

    async function fetchServices() {
        try {
            const response = await api.get("/system/health");
            services = Object.entries(response.data.services).map(([name, active]) => ({
                name,
                active
            }));
        } catch (error) {
            console.error("Failed to fetch services:", error);
        } finally {
            loading = false;
        }
    }

    onMount(fetchServices);
</script>

<div class="bg-white shadow rounded-lg overflow-hidden">
    <div class="px-4 py-5 sm:px-6">
        <h3 class="text-lg leading-6 font-medium text-gray-900">
            Services Status
        </h3>
    </div>
    <div class="border-t border-gray-200">
        {#if loading}
            <div class="text-center py-4">Loading services...</div>
        {:else}
            <ul class="divide-y divide-gray-200">
                {#each services as service}
                    <li class="px-4 py-4 sm:px-6">
                        <div class="flex items-center justify-between">
                            <div class="flex items-center">
                                <div class={`h-2.5 w-2.5 rounded-full ${
                                    service.active ? 'bg-green-500' : 'bg-red-500'
                                }`}></div>
                                <p class="ml-3 text-sm font-medium text-gray-900">
                                    {service.name}
                                </p>
                            </div>
                            <span class={`px-2 inline-flex text-xs leading-5 font-semibold rounded-full ${
                                service.active 
                                    ? 'bg-green-100 text-green-800'
                                    : 'bg-red-100 text-red-800'
                            }`}>
                                {service.active ? 'Active' : 'Inactive'}
                            </span>
                        </div>
                    </li>
                {/each}
            </ul>
        {/if}
    </div>
</div> 