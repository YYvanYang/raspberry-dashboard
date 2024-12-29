<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import JSONEditor from "$lib/components/JSONEditor.svelte";
    
    export let service: string;
    export let config: any;
    
    const dispatch = createEventDispatcher();
    let editedConfig = JSON.stringify(config, null, 2);
    let error = "";
    
    function handleSave() {
        try {
            const parsedConfig = JSON.parse(editedConfig);
            dispatch("save", parsedConfig);
            error = "";
        } catch (e) {
            error = "Invalid JSON format";
        }
    }
</script>

<div class="h-full flex flex-col">
    <div class="p-4 border-b flex justify-between items-center">
        <h2 class="text-lg font-medium">{service} Configuration</h2>
        <button
            on:click={handleSave}
            class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700"
        >
            Save Changes
        </button>
    </div>
    
    {#if error}
        <div class="p-4 bg-red-50 text-red-700">
            {error}
        </div>
    {/if}
    
    <div class="flex-1 p-4">
        <JSONEditor bind:value={editedConfig} />
    </div>
</div> 