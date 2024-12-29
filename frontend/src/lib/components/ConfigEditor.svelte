<script lang="ts">
    const { service, config, onSave } = $props<{
        service: string;
        config: Record<string, any>;
        onSave: (config: Record<string, any>) => Promise<void>;
    }>();

    let editingConfig = $state(JSON.stringify(config, null, 2));
    let error = $state('');

    async function handleSave() {
        try {
            const parsedConfig = JSON.parse(editingConfig);
            await onSave(parsedConfig);
            error = '';
        } catch (err) {
            if (err instanceof SyntaxError) {
                error = '配置格式错误，请检查 JSON 格式';
            } else {
                error = '保存失败';
            }
        }
    }
</script>

<div class="p-4">
    <div class="flex justify-between items-center mb-4">
        <h3 class="text-lg font-medium text-gray-900">{service} 配置</h3>
        <button
            onclick={handleSave}
            class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
        >
            保存
        </button>
    </div>

    {#if error}
        <div class="mb-4 p-4 bg-red-50 text-red-700 rounded-md">
            {error}
        </div>
    {/if}

    <textarea
        bind:value={editingConfig}
        class="w-full h-[500px] font-mono text-sm p-4 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
        spellcheck="false"
    ></textarea>
</div> 