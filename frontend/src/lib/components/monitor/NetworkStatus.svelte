<script lang="ts">
    const { network } = $props<{
        network: {
            interfaces: {
                name: string;
                ip: string;
                rx_bytes: number;
                tx_bytes: number;
                status: string;
            }[];
        };
    }>();

    function formatSpeed(bytes: number): string {
        const units = ['B/s', 'KB/s', 'MB/s', 'GB/s'];
        let speed = bytes;
        let unitIndex = 0;

        while (speed >= 1024 && unitIndex < units.length - 1) {
            speed /= 1024;
            unitIndex++;
        }

        return `${speed.toFixed(2)} ${units[unitIndex]}`;
    }
</script>

<div class="bg-white shadow rounded-lg overflow-hidden">
    <div class="px-4 py-5 sm:px-6">
        <h3 class="text-lg leading-6 font-medium text-gray-900">
            Network Interfaces
        </h3>
    </div>
    <div class="border-t border-gray-200">
        <ul class="divide-y divide-gray-200">
            {#each network.interfaces as iface}
                <li class="px-4 py-4 sm:px-6">
                    <div class="flex flex-col space-y-2">
                        <div class="flex items-center justify-between">
                            <div class="flex items-center">
                                <div class={`h-2.5 w-2.5 rounded-full ${
                                    iface.status === 'up' ? 'bg-green-500' : 'bg-red-500'
                                }`}></div>
                                <p class="ml-3 text-sm font-medium text-gray-900">
                                    {iface.name}
                                </p>
                            </div>
                            <span class="text-sm text-gray-500">{iface.ip}</span>
                        </div>
                        <div class="grid grid-cols-2 gap-4 text-sm">
                            <div>
                                <span class="text-gray-500">↓ RX</span>
                                <span class="ml-2 font-medium">{formatSpeed(iface.rx_bytes)}</span>
                            </div>
                            <div>
                                <span class="text-gray-500">↑ TX</span>
                                <span class="ml-2 font-medium">{formatSpeed(iface.tx_bytes)}</span>
                            </div>
                        </div>
                    </div>
                </li>
            {/each}
        </ul>
    </div>
</div> 