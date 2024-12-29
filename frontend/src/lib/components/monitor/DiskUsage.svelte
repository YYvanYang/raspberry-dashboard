<script lang="ts">
    const { disk } = $props<{
        disk: {
            total: number;
            used: number;
            free: number;
            usage_rate: number;
            mount_point: string;
        };
    }>();

    const usagePercent = $derived(disk.usage_rate.toFixed(1));
</script>

<div class="bg-white overflow-hidden shadow rounded-lg">
    <div class="p-5">
        <div class="flex flex-col gap-4">
            <div>
                <div class="text-sm font-medium text-gray-500">Mount Point</div>
                <div class="mt-1 text-lg font-semibold text-gray-900">{disk.mount_point}</div>
            </div>
            <div>
                <div class="text-sm font-medium text-gray-500">Usage</div>
                <div class="mt-1 flex items-center">
                    <div class="text-2xl font-semibold text-gray-900">{usagePercent}%</div>
                    <div class="ml-2 w-full bg-gray-200 rounded-full h-2.5">
                        <div class="bg-purple-600 h-2.5 rounded-full" style="width: {usagePercent}%"></div>
                    </div>
                </div>
            </div>
            <div class="grid grid-cols-3 gap-4">
                <div>
                    <div class="text-sm font-medium text-gray-500">Total</div>
                    <div class="mt-1 text-lg font-semibold text-gray-900">{(disk.total / 1024 / 1024 / 1024).toFixed(1)} GB</div>
                </div>
                <div>
                    <div class="text-sm font-medium text-gray-500">Used</div>
                    <div class="mt-1 text-lg font-semibold text-gray-900">{(disk.used / 1024 / 1024 / 1024).toFixed(1)} GB</div>
                </div>
                <div>
                    <div class="text-sm font-medium text-gray-500">Free</div>
                    <div class="mt-1 text-lg font-semibold text-gray-900">{(disk.free / 1024 / 1024 / 1024).toFixed(1)} GB</div>
                </div>
            </div>
        </div>
    </div>
</div> 