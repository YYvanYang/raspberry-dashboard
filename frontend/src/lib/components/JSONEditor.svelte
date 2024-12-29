<script lang="ts">
    import { onMount } from 'svelte';
    import { createEventDispatcher } from 'svelte';
    import * as monaco from 'monaco-editor';

    const dispatch = createEventDispatcher<{
        update: { value: string };
    }>();

    const { value } = $props<{
        value: string;
    }>();

    let editor: monaco.editor.IStandaloneCodeEditor;
    let editorContainer: HTMLDivElement;

    onMount(() => {
        if (editorContainer) {
            editor = monaco.editor.create(editorContainer, {
                value: value,
                language: 'json',
                theme: 'vs-dark',
                automaticLayout: true,
                minimap: {
                    enabled: false
                }
            });

            editor.onDidChangeModelContent(() => {
                const newValue = editor.getValue();
                if (newValue !== value) {
                    dispatch('update', { value: newValue });
                }
            });
        }

        return () => {
            if (editor) {
                editor.dispose();
            }
        };
    });

    $effect(() => {
        if (editor && value !== editor.getValue()) {
            editor.setValue(value);
        }
    });
</script>

<div class="h-full" bind:this={editorContainer}></div>

<style>
    div {
        min-height: 300px;
    }
</style> 