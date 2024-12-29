<script lang="ts">
    export let disk: {
        total: number;
        used: number;
        free: number;
        usage_rate: number;
    };

    function formatSize(bytes: number): string {
        const units = ['B', 'KB', 'MB', 'GB', 'TB'];
        let size = bytes;
        let unitIndex = 0;

        while (size >= 1024 && unitIndex < units.length - 1) {
            size /= 1024;
            unitIndex++;
        }

        return `${size.toFixed(2)} ${units[unitIndex]}`;
    }

    function getUsageColor(rate: number): string {
        if (rate >= 90) return 'text-red-600';
        if (rate >= 80) return 'text-orange-500';
        if (rate >= 70) return 'text-yellow-500';
        return 'text-green-500';
    }
</script>

<div class="bg-white overflow-hidden shadow rounded-lg">
    <div class="p-5">
        <div class="flex flex-col">
            <div class="flex items-center justify-between mb-4">
                <h3 class="text-lg font-medium text-gray-900">Disk Usage</h3>
                <span class={`text-sm font-semibold ${getUsageColor(disk.usage_rate)}`}>
                    {disk.usage_rate.toFixed(1)}%
                </span>
            </div>
            
            <div class="relative pt-1">
                <div class="overflow-hidden h-2 text-xs flex rounded bg-gray-200">
                    <div
                        style="width: {disk.usage_rate}%"
                        class={`shadow-none flex flex-col text-center whitespace-nowrap text-white justify-center ${
                            getUsageColor(disk.usage_rate).replace('text-', 'bg-')
                        }`}
                    ></div>
                </div>
            </div>

            <div class="mt-4 grid grid-cols-3 gap-4 text-sm">
                <div>
                    <span class="block text-gray-500">Total</span>
                    <span class="block font-medium">{formatSize(disk.total)}</span>
                </div>
                <div>
                    <span class="block text-gray-500">Used</span>
                    <span class="block font-medium">{formatSize(disk.used)}</span>
                </div>
                <div>
                    <span class="block text-gray-500">Free</span>
                    <span class="block font-medium">{formatSize(disk.free)}</span>
                </div>
            </div>
        </div>
    </div>
</div> 