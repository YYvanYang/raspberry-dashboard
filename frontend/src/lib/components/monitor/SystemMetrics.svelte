<script lang="ts">
    const { metrics } = $props<{
        metrics: {
            cpu?: number;
            memory?: {
                usage_rate?: number;
            };
            disk?: {
                usage_rate?: number;
            };
        };
    }>();

    const cpuUsage = $derived(metrics?.cpu?.toFixed(1) || 0);
    const memoryUsage = $derived(metrics?.memory?.usage_rate?.toFixed(1) || 0);
    const diskUsage = $derived(metrics?.disk?.usage_rate?.toFixed(1) || 0);
</script>

<div class="bg-white overflow-hidden shadow rounded-lg">
    <div class="p-5">
        <div class="flex flex-col gap-4">
            <div>
                <div class="text-sm font-medium text-gray-500">CPU Usage</div>
                <div class="mt-1 flex items-center">
                    <div class="text-2xl font-semibold text-gray-900">{cpuUsage}%</div>
                    <div class="ml-2 w-full bg-gray-200 rounded-full h-2.5">
                        <div class="bg-blue-600 h-2.5 rounded-full" style="width: {cpuUsage}%"></div>
                    </div>
                </div>
            </div>

            <div>
                <div class="text-sm font-medium text-gray-500">Memory Usage</div>
                <div class="mt-1 flex items-center">
                    <div class="text-2xl font-semibold text-gray-900">{memoryUsage}%</div>
                    <div class="ml-2 w-full bg-gray-200 rounded-full h-2.5">
                        <div class="bg-green-600 h-2.5 rounded-full" style="width: {memoryUsage}%"></div>
                    </div>
                </div>
            </div>

            <div>
                <div class="text-sm font-medium text-gray-500">Disk Usage</div>
                <div class="mt-1 flex items-center">
                    <div class="text-2xl font-semibold text-gray-900">{diskUsage}%</div>
                    <div class="ml-2 w-full bg-gray-200 rounded-full h-2.5">
                        <div class="bg-purple-600 h-2.5 rounded-full" style="width: {diskUsage}%"></div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div> 