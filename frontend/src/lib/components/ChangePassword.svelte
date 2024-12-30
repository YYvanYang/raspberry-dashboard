<script lang="ts">
    import { notifications } from '$lib/stores/notifications';
    import { api } from '$lib/utils/api';

    let currentPassword = $state('');
    let newPassword = $state('');
    let confirmPassword = $state('');
    let loading = $state(false);

    async function handleSubmit(event: Event) {
        event.preventDefault();
        
        if (newPassword !== confirmPassword) {
            notifications.add({
                type: 'error',
                message: '新密码与确认密码不匹配'
            });
            return;
        }

        try {
            loading = true;
            await api.post('/auth/change-password', {
                currentPassword,
                newPassword
            });
            
            notifications.add({
                type: 'success',
                message: '密码修改成功'
            });
            
            // 清空表单
            currentPassword = '';
            newPassword = '';
            confirmPassword = '';
        } catch (error) {
            notifications.add({
                type: 'error',
                message: '密码修改失败: ' + (error.response?.data?.message || error.message)
            });
        } finally {
            loading = false;
        }
    }
</script>

<div class="bg-white p-6 rounded-lg shadow-md max-w-md w-full">
    <h2 class="text-2xl font-bold mb-6 text-gray-800">修改密码</h2>
    
    <!-- 使用正确的事件处理语法 -->
    <form onsubmit={handleSubmit} class="space-y-4">
        <div>
            <label for="currentPassword" class="block text-sm font-medium text-gray-700">
                当前密码
            </label>
            <input
                id="currentPassword"
                type="password"
                bind:value={currentPassword}
                required
                autocomplete="current-password"
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
            />
        </div>

        <div>
            <label for="newPassword" class="block text-sm font-medium text-gray-700">
                新密码
            </label>
            <input
                id="newPassword"
                type="password"
                bind:value={newPassword}
                required
                minlength="8"
                autocomplete="new-password"
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
            />
        </div>

        <div>
            <label for="confirmPassword" class="block text-sm font-medium text-gray-700">
                确认新密码
            </label>
            <input
                id="confirmPassword"
                type="password"
                bind:value={confirmPassword}
                required
                minlength="8"
                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
            />
        </div>

        <button
            type="submit"
            disabled={loading}
            class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50"
        >
            {loading ? '处理中...' : '修改密码'}
        </button>
    </form>
</div> 