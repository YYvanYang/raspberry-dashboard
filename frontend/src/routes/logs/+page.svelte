<script lang="ts">
    import { onMount } from "svelte";
    import { api } from "$lib/utils/api";

    interface Log {
        timestamp: string;
        level: string;
        message: string;
    }

    // 使用 runes 声明状态
    let logs = $state<Log[]>([]);
    let loading = $state(true);

    let filter = $state({
        level: 'all',
        search: ''
    });

    // 使用 derived 计算过滤后的日志
    let filteredLogs = $derived(() => {
        return logs.filter(log => {
            if (filter.level !== 'all' && log.level !== filter.level) return false;
            if (filter.search && !log.message.toLowerCase().includes(filter.search.toLowerCase())) return false;
            return true;
        });
    });

    async function fetchLogs() {
        try {
            loading = true;
            const params = new URLSearchParams();
            if (filter.level !== "all") params.append("level", filter.level);
            if (filter.search) params.append("search", filter.search);

            const response = await api.get(`/system/logs?${params.toString()}`);
            logs = response.data;
        } catch (error) {
            console.error("Failed to fetch logs:", error);
        } finally {
            loading = false;
        }
    }

    function clearFilter() {
        filter.level = 'all';
        filter.search = '';
    }

    // 使用 effect 监听过滤器变化
    $effect(() => {
        fetchLogs();
    });

    onMount(fetchLogs);
</script>

<div class="min-h-screen bg-gray-100">
    <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div class="px-4 py-6 sm:px-0">
            <div class="bg-white rounded-lg shadow">
                <div class="p-4 border-b">
                    <h2 class="text-lg font-medium mb-4">系统日志</h2>
                    
                    <div class="flex gap-4 mb-4">
                        <select 
                            class="border rounded px-2 py-1"
                            bind:value={filter.level}
                        >
                            <option value="all">所有级别</option>
                            <option value="info">信息</option>
                            <option value="warning">警告</option>
                            <option value="error">错误</option>
                        </select>
                        
                        <input 
                            type="text"
                            placeholder="搜索日志..."
                            class="border rounded px-2 py-1 flex-1"
                            bind:value={filter.search}
                        />
                        
                        <button 
                            class="bg-gray-200 px-4 py-1 rounded hover:bg-gray-300"
                            onclick={clearFilter}
                        >
                            清除过滤器
                        </button>
                    </div>
                </div>

                {#if loading}
                    <div class="flex justify-center items-center p-8">
                        <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-500"></div>
                    </div>
                {:else}
                    <div class="overflow-x-auto">
                        <table class="w-full">
                            <thead>
                                <tr class="border-b">
                                    <th class="text-left p-2">时间</th>
                                    <th class="text-left p-2">级别</th>
                                    <th class="text-left p-2">消息</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#each filteredLogs as log}
                                    <tr class="border-b hover:bg-gray-50">
                                        <td class="p-2">{log.timestamp}</td>
                                        <td class="p-2">
                                            <span class={`px-2 py-1 rounded text-sm
                                                ${log.level === 'error' ? 'bg-red-100 text-red-800' : 
                                                    log.level === 'warning' ? 'bg-yellow-100 text-yellow-800' : 
                                                    'bg-blue-100 text-blue-800'}`}
                                            >
                                                {log.level}
                                            </span>
                                        </td>
                                        <td class="p-2">{log.message}</td>
                                    </tr>
                                {/each}
                            </tbody>
                        </table>
                    </div>
                {/if}
            </div>
        </div>
    </div>
</div> 