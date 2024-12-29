<script lang="ts">
    import { goto } from "$app/navigation";
    import { auth } from "$lib/stores/auth";
    import { api } from "$lib/utils/api";

    let username = "";
    let password = "";
    let error = "";

    async function handleLogin() {
        try {
            const response = await api.post("/auth/login", { username, password });
            const token = response.data.token;
            localStorage.setItem("token", token);
            auth.set({ token });
            goto("/dashboard");
        } catch (err: any) {
            error = err.response?.data?.error || "Login failed";
        }
    }
</script>

<div class="min-h-screen bg-gray-100 flex items-center justify-center">
    <div class="max-w-md w-full bg-white rounded-lg shadow-lg p-8">
        <h2 class="text-2xl font-bold text-center mb-6">Login</h2>
        
        {#if error}
            <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
                {error}
            </div>
        {/if}

        <form on:submit|preventDefault={handleLogin} class="space-y-4">
            <div>
                <label for="username" class="block text-sm font-medium text-gray-700">Username</label>
                <input
                    type="text"
                    id="username"
                    bind:value={username}
                    class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                    required
                />
            </div>

            <div>
                <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
                <input
                    type="password"
                    id="password"
                    bind:value={password}
                    class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                    required
                />
            </div>

            <button
                type="submit"
                class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            >
                Login
            </button>
        </form>
    </div>
</div> 