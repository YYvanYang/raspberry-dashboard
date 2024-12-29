<script lang="ts">
    import { api } from "$lib/utils/api";
    import ConfigEditor from "$lib/components/ConfigEditor.svelte";
    import { notifications } from "$lib/stores/notifications";

    interface Config {
        [key: string]: Record<string, any>;
    }

    let configs = $state<Config>({});
    let selectedService = $state<string | null>(null);
    let loading = $state(true);

    async function fetchConfigs() {
        try {
            const response = await api.get('/api/configs');
            configs = response.data;
            if (!selectedService && Object.keys(configs).length > 0) {
                selectedService = Object.keys(configs)[0];
            }
        } catch (error) {
            notifications.error('获取配置失败');
            console.error('Error fetching configs:', error);
        } finally {
            loading = false;
        }
    }

    async function handleSaveConfig(newConfig: Record<string, any>) {
        if (!selectedService) return;
        try {
            await api.put(`/api/configs/${selectedService}`, newConfig);
            configs[selectedService] = newConfig;
            notifications.success('配置保存成功');
        } catch (error) {
            notifications.error('保存配置失败');
            console.error('Error saving config:', error);
        }
    }

    $effect(() => {
        fetchConfigs();
    });
</script>

<div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
    {#if loading}
        <div class="flex justify-center items-center h-64">
            <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-500"></div>
        </div>
    {:else}
        <div class="bg-white shadow rounded-lg">
            <div class="grid grid-cols-4 min-h-[600px]">
                <!-- Service List -->
                <div class="col-span-1 border-r">
                    <div class="p-4 border-b">
                        <h2 class="text-lg font-medium text-gray-900">服务列表</h2>
                    </div>
                    <nav class="flex-1 min-h-0 overflow-y-auto">
                        <div class="py-2">
                            {#each Object.keys(configs) as service}
                                <button
                                    class="w-full text-left px-4 py-2 rounded hover:bg-gray-100"
                                    class:bg-blue-50={selectedService === service}
                                    onclick={() => selectedService = service}
                                >
                                    {service}
                                </button>
                            {/each}
                        </div>
                    </nav>
                </div>

                <!-- Config Editor -->
                <div class="col-span-3">
                    {#if selectedService}
                        <ConfigEditor 
                            service={selectedService}
                            config={configs[selectedService]}
                            onSave={handleSaveConfig}
                        />
                    {:else}
                        <div class="p-4 text-center text-gray-500">
                            请选择一个服务来编辑配置
                        </div>
                    {/if}
                </div>
            </div>
        </div>
    {/if}
</div> 