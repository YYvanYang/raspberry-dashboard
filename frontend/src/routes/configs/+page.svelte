<script lang="ts">
    import { onMount } from "svelte";
    import { api } from "$lib/utils/api";
    import ConfigEditor from "$lib/components/ConfigEditor.svelte";
    import { notifications } from "$lib/stores/notifications";

    let configs: Record<string, any> = {};
    let selectedService: string | null = null;
    let loading = true;

    async function fetchConfigs() {
        try {
            const response = await api.get("/configs");
            configs = response.data;
        } catch (error) {
            console.error("Failed to fetch configs:", error);
        } finally {
            loading = false;
        }
    }

    async function handleSaveConfig(service: string, config: any) {
        try {
            await api.put(`/configs/${service}`, config);
            notifications.success(`Successfully updated ${service} config`);
            await fetchConfigs();
        } catch (error) {
            notifications.error(`Failed to update ${service} config`);
        }
    }

    onMount(fetchConfigs);
</script>

<div class="min-h-screen bg-gray-100">
    <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div class="px-4 py-6 sm:px-0">
            <div class="flex space-x-4">
                <!-- 服务列表 -->
                <div class="w-64 bg-white rounded-lg shadow">
                    <div class="p-4 border-b">
                        <h2 class="text-lg font-medium">Services</h2>
                    </div>
                    <div class="p-2">
                        {#each Object.keys(configs) as service}
                            <button
                                class="w-full text-left px-4 py-2 rounded hover:bg-gray-100"
                                class:bg-blue-50={selectedService === service}
                                on:click={() => selectedService = service}
                            >
                                {service}
                            </button>
                        {/each}
                    </div>
                </div>

                <!-- 配置编辑器 -->
                <div class="flex-1 bg-white rounded-lg shadow">
                    {#if selectedService && configs[selectedService]}
                        <ConfigEditor
                            service={selectedService}
                            config={configs[selectedService]}
                            on:save={(e) => handleSaveConfig(selectedService, e.detail)}
                        />
                    {:else}
                        <div class="p-4 text-center text-gray-500">
                            Select a service to edit its configuration
                        </div>
                    {/if}
                </div>
            </div>
        </div>
    </div>
</div> 