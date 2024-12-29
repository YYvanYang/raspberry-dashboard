<script lang="ts">
    import { onMount } from "svelte";
    import { api } from "$lib/utils/api";
    import StatusCard from "$lib/components/StatusCard.svelte";
    import ServiceList from "$lib/components/ServiceList.svelte";
    import SystemChart from "$lib/components/SystemChart.svelte";
    
    let systemStatus: any = null;
    let loading = true;
    
    async function fetchStatus() {
        try {
            const response = await api.get("/system/status");
            systemStatus = response.data;
        } catch (error) {
            console.error("Failed to fetch status:", error);
        } finally {
            loading = false;
        }
    }
    
    onMount(() => {
        fetchStatus();
        const interval = setInterval(fetchStatus, 30000); // 每30秒更新一次
        
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
            <div class="text-center">Loading...</div>
        {:else if systemStatus}
            <div class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4">
                <StatusCard
                    title="CPU Usage"
                    value="{systemStatus.cpu.toFixed(1)}%"
                    icon="cpu"
                />
                <StatusCard
                    title="Memory Usage"
                    value="{systemStatus.memory.usageRate.toFixed(1)}%"
                    icon="memory"
                />
                <StatusCard
                    title="Disk Usage"
                    value="{systemStatus.disk.usageRate.toFixed(1)}%"
                    icon="disk"
                />
                <StatusCard
                    title="Uptime"
                    value={systemStatus.uptime}
                    icon="clock"
                />
            </div>

            <div class="mt-8">
                <ServiceList services={systemStatus.services} />
            </div>

            <div class="mt-8">
                <SystemChart data={systemStatus} />
            </div>
        {/if}
    </main>
</div> 