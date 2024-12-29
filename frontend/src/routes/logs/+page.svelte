<script lang="ts">
    import { onMount } from "svelte";
    import { api } from "$lib/utils/api";
    import LogViewer from "$lib/components/LogViewer.svelte";
    import LogFilter from "$lib/components/LogFilter.svelte";

    let logs = $state<any[]>([]);
    let loading = $state(true);
    let filter = $state({
        type: "all",
        search: "",
        dateRange: "24h"
    });

    async function fetchLogs() {
        try {
            const params = new URLSearchParams();
            if (filter.type !== "all") params.append("type", filter.type);
            if (filter.search) params.append("search", filter.search);
            if (filter.dateRange) params.append("range", filter.dateRange);

            const response = await api.get(`/system/logs?${params.toString()}`);
            logs = response.data;
        } catch (error) {
            console.error("Failed to fetch logs:", error);
        } finally {
            loading = false;
        }
    }

    $effect(() => {
        if (!loading) fetchLogs();
    });

    onMount(fetchLogs);
</script>

<div class="min-h-screen bg-gray-100">
    <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div class="px-4 py-6 sm:px-0">
            <div class="bg-white rounded-lg shadow">
                <div class="p-4 border-b">
                    <h2 class="text-lg font-medium">System Logs</h2>
                    <LogFilter bind:filter />
                </div>
                <LogViewer {logs} {loading} />
            </div>
        </div>
    </div>
</div> 