<script lang="ts">
    import { goto } from "$app/navigation";
    import { notifications } from "$lib/stores/notifications";
    import { api } from "$lib/utils/api";

    let username = $state("");
    let password = $state("");
    let error = $state("");

    async function handleLogin(event: SubmitEvent) {
        event.preventDefault();
        error = "";

        try {
            const response = await api.post('/api/auth/login', {
                username,
                password
            });

            localStorage.setItem('token', response.data.token);
            notifications.success('登录成功');
            goto('/dashboard');
        } catch (err) {
            error = '用户名或密码错误';
            notifications.error('登录失败');
        }
    }
</script>

<div class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8">
        <div>
            <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
                树莓派管理系统
            </h2>
        </div>

        {#if error}
            <div class="bg-red-50 p-4 rounded-md">
                <p class="text-sm text-red-700">{error}</p>
            </div>
        {/if}

        <form onsubmit={handleLogin} class="space-y-4">
            <div>
                <label for="username" class="block text-sm font-medium text-gray-700">用户名</label>
                <input
                    id="username"
                    type="text"
                    required
                    bind:value={username}
                    class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                />
            </div>

            <div>
                <label for="password" class="block text-sm font-medium text-gray-700">密码</label>
                <input
                    id="password"
                    type="password"
                    required
                    bind:value={password}
                    class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                />
            </div>

            <div>
                <button
                    type="submit"
                    class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                >
                    登录
                </button>
            </div>
        </form>
    </div>
</div> 