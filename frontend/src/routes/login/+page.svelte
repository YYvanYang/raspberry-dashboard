<script lang="ts">
    import { goto } from "$app/navigation";
    import { notifications } from "$lib/stores/notifications";
    import { api } from "$lib/utils/api";
    import { notifyAuthChange } from "$lib/auth";

    let username = $state("");
    let password = $state("");
    let error = $state("");
    let isLoading = $state(false);

    async function handleLogin(event: SubmitEvent) {
        event.preventDefault();
        error = "";
        isLoading = true;

        try {
            const response = await api.post('/api/auth/login', {
                username,
                password
            });

            localStorage.setItem('token', response.token);
            notifyAuthChange(true);
            notifications.success('登录成功');
            goto('/dashboard');
        } catch (err) {
            console.error('Login error:', err);
            error = '用户名或密码错误';
            notifications.error('登录失败');
        } finally {
            isLoading = false;
        }
    }
</script>

<div class="h-full flex items-center justify-center bg-gray-50 py-8 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full">
        <div>
            <h2 class="mt-2 text-center text-3xl font-extrabold text-gray-900">
                树莓派管理系统
            </h2>
        </div>

        {#if error}
            <div class="mt-8 bg-red-50 p-4 rounded-md">
                <p class="text-sm text-red-700">{error}</p>
            </div>
        {/if}

        <form onsubmit={handleLogin} class="mt-8 space-y-4">
            <div>
                <label for="username" class="block text-sm font-medium text-gray-700">用户名</label>
                <input
                    id="username"
                    type="text"
                    required
                    disabled={isLoading}
                    bind:value={username}
                    autocomplete="username"
                    class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-100 disabled:cursor-not-allowed"
                />
            </div>

            <div>
                <label for="password" class="block text-sm font-medium text-gray-700">密码</label>
                <input
                    id="password"
                    type="password"
                    required
                    disabled={isLoading}
                    bind:value={password}
                    autocomplete="current-password"
                    class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 disabled:bg-gray-100 disabled:cursor-not-allowed"
                />
            </div>

            <div>
                <button
                    type="submit"
                    disabled={isLoading}
                    class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:bg-blue-400 disabled:cursor-not-allowed"
                >
                    {#if isLoading}
                        <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                        </svg>
                        登录中...
                    {:else}
                        登录
                    {/if}
                </button>
            </div>
        </form>
    </div>
</div> 