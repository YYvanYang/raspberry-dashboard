<script lang="ts">
    import { onMount, onDestroy } from "svelte";
    import { api } from "$lib/utils/api";
    import SystemChart from "$lib/components/SystemChart.svelte";

    let systemData = $state({
        labels: [] as string[],
        datasets: [
            {
                label: "CPU使用率",
                data: [] as number[]
            },
            {
                label: "内存使用率", 
                data: [] as number[]
            }
        ]
    });

    let interval = $state<number>();

    async function fetchSystemStatus() {
        try {
            const response = await api.get("/system/status");
            const time = new Date().toLocaleTimeString();
            
            systemData = {
                labels: [...systemData.labels.slice(-20), time],
                datasets: [
                    {
                        ...systemData.datasets[0],
                        data: [...systemData.datasets[0].data.slice(-20), response.data.cpu]
                    },
                    {
                        ...systemData.datasets[1],
                        data: [...systemData.datasets[1].data.slice(-20), response.data.memory.usageRate]
                    }
                ]
            };
        } catch (error) {
            console.error("Failed to fetch system status:", error);
        }
    }

    onMount(() => {
        fetchSystemStatus();
        interval = setInterval(fetchSystemStatus, 5000);
    });

    onDestroy(() => {
        if (interval) {
            clearInterval(interval);
        }
    });
</script>

<div class="space-y-6">
    <h1 class="text-2xl font-semibold text-gray-900">系统监控</h1>
    <SystemChart data={systemData} />
</div> 