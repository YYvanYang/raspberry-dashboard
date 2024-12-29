<script lang="ts">
    import ChangePassword from '$lib/components/ChangePassword.svelte';
</script>

<div class="container mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold mb-8">系统设置</h1>
    
    <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div>
            <ChangePassword />
        </div>
        
        <!-- 其他设置选项可以在这里添加 -->
        <div class="bg-white rounded-lg shadow-md p-6">
            <h2 class="text-2xl font-bold mb-4 text-gray-800">其他设置</h2>
            <p class="text-gray-600">更多设置选项即将推出...</p>
        </div>
    </div>
</div> 