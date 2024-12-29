<script lang="ts">
    let { filter = $bindable() } = $props<{
        filter: {
            type: string;
            search: string;
            dateRange: string;
        }
    }>();

    const types = ["all", "info", "warning", "error"];
    const ranges = [
        { value: "1h", label: "Last hour" },
        { value: "24h", label: "Last 24 hours" },
        { value: "7d", label: "Last 7 days" },
        { value: "30d", label: "Last 30 days" }
    ];
</script>

<div class="flex items-center space-x-4 mt-4">
    <div class="flex-1">
        <input
            type="text"
            placeholder="Search logs..."
            bind:value={filter.search}
            class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm"
        />
    </div>
    
    <select
        bind:value={filter.type}
        class="block rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm"
    >
        {#each types as type}
            <option value={type}>
                {type.charAt(0).toUpperCase() + type.slice(1)}
            </option>
        {/each}
    </select>

    <select
        bind:value={filter.dateRange}
        class="block rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm"
    >
        {#each ranges as range}
            <option value={range.value}>{range.label}</option>
        {/each}
    </select>
</div> 