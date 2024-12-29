<script lang="ts">
    const { logs, loading } = $props<{
        logs: {
            type: string;
            message: string;
            created_at: string;
        }[];
        loading: boolean;
    }>();

    function getLogTypeColor(type: string): string {
        switch (type.toLowerCase()) {
            case 'error':
                return 'text-red-600 bg-red-50';
            case 'warning':
                return 'text-yellow-600 bg-yellow-50';
            case 'info':
                return 'text-blue-600 bg-blue-50';
            default:
                return 'text-gray-600 bg-gray-50';
        }
    }

    function formatDate(date: string): string {
        return new Date(date).toLocaleString();
    }
</script>

<div class="overflow-x-auto">
    {#if loading}
        <div class="p-4 text-center text-gray-500">
            Loading logs...
        </div>
    {:else if logs.length === 0}
        <div class="p-4 text-center text-gray-500">
            No logs found
        </div>
    {:else}
        <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
                <tr>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        Time
                    </th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        Type
                    </th>
                    <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        Message
                    </th>
                </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
                {#each logs as log}
                    <tr>
                        <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                            {formatDate(log.created_at)}
                        </td>
                        <td class="px-6 py-4 whitespace-nowrap">
                            <span class={`px-2 inline-flex text-xs leading-5 font-semibold rounded-full ${getLogTypeColor(log.type)}`}>
                                {log.type}
                            </span>
                        </td>
                        <td class="px-6 py-4 text-sm text-gray-500">
                            {log.message}
                        </td>
                    </tr>
                {/each}
            </tbody>
        </table>
    {/if}
</div> 