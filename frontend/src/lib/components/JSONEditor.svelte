<script lang="ts">
    export let value: string;
    
    let textarea: HTMLTextAreaElement;
    
    function handleInput() {
        // 自动调整高度
        textarea.style.height = "auto";
        textarea.style.height = textarea.scrollHeight + "px";
    }
</script>

<textarea
    bind:this={textarea}
    bind:value
    on:input={handleInput}
    class="w-full h-full min-h-[400px] font-mono text-sm p-4 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
    spellcheck="false"
></textarea> 