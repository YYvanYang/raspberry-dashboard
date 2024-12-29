<script lang="ts">
    import { api } from "$lib/utils/api";
    import { notifications } from "$lib/stores/notifications";
    import StatusCard from "$lib/components/StatusCard.svelte";
    import SystemChart from "$lib/components/SystemChart.svelte";

    interface SystemStatus {
        cpu: {
            usage: number;
        };
        memory: {
            usage: number;
        };
        disk: {
            usage: number;
        };
        history: {
            labels: string[];
            datasets: Array<{
                data: number[];
            }>;
        };
    }

    let systemStatus = $state<SystemStatus | null>(null);
    let loading = $state(true);

    async function fetchStatus() {
        try {
            const response = await api.get('/api/system/status');
            systemStatus = response.data;
        } catch (error) {
            notifications.error('获取系统状态失败');
            console.error('Error fetching status:', error);
        } finally {
            loading = false;
        }
    }

    $effect(() => {
        fetchStatus();
        const interval = setInterval(fetchStatus, 5000);
        return () => clearInterval(interval);
    });
</script>

<div class="min-h-screen bg-gray-100">
    <nav class="bg-white shadow-sm">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="flex justify-between h-16">
                <div class="flex">
                    <div class="flex-shrink-0 flex items-center">
                        <h1 class="text-xl font-bold">Raspberry Dashboard</h1>
                    </div>
                </div>
            </div>
        </div>
    </nav>

    <main class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        {#if loading}
            <div class="flex justify-center items-center h-64">
                <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-500"></div>
            </div>
        {:else if systemStatus}
            <div class="space-y-6">
                <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
                    <StatusCard
                        title="CPU 使用率"
                        value={`${systemStatus.cpu.usage}%`}
                        icon="cpu"
                    />
                    <StatusCard
                        title="内存使用率"
                        value={`${systemStatus.memory.usage}%`}
                        icon="memory"
                    />
                    <StatusCard
                        title="磁盘使用率"
                        value={`${systemStatus.disk.usage}%`}
                        icon="disk"
                    />
                </div>

                <div class="bg-white rounded-lg shadow">
                    <SystemChart data={systemStatus.history} />
                </div>
            </div>
        {/if}
    </main>
</div> 