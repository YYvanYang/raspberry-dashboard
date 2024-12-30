<script lang="ts">
    import '../app.css';
    import ErrorBoundary from '$lib/components/ErrorBoundary.svelte';
    import Notifications from '$lib/components/Notifications.svelte';
    import type { LayoutData } from './$types';
    import type { Snippet } from 'svelte';
    import { logout, onAuthChange } from '$lib/auth';
    import { clickOutside } from '$lib/actions/clickOutside';
    
    const props = $props<{
        data: LayoutData;
        children: Snippet;
    }>();
    
    let showUserMenu = $state(false);
    let isLoggedIn = $state(!!localStorage.getItem('token'));
    
    // 监听登录状态变化
    onAuthChange((status) => {
        isLoggedIn = status;
    });
    
    function toggleUserMenu() {
        showUserMenu = !showUserMenu;
    }
    
    async function handleLogout() {
        await logout();
    }
    
    function closeUserMenu() {
        showUserMenu = false;
    }
</script>

<ErrorBoundary>
    <div class="bg-gray-100 h-screen flex flex-col">
        {#if isLoggedIn}
        <nav class="bg-gray-800 text-white">
            <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
                <div class="flex h-16 items-center justify-between">
                    <!-- 左侧 Logo -->
                    <div class="flex-shrink-0">
                        <a href="/" class="text-white font-bold">树莓派管理面板</a>
                    </div>
                    
                    <!-- 右侧用户菜单 -->
                    <div class="relative ml-3">
                        <button 
                            type="button"
                            class="flex rounded-full bg-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800"
                            onclick={toggleUserMenu}
                        >
                            <span class="sr-only">打开用户菜单</span>
                            <svg class="h-8 w-8 rounded-full p-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                            </svg>
                        </button>

                        {#if showUserMenu}
                        <div 
                            use:clickOutside={closeUserMenu}
                            class="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none"
                            role="menu"
                        >
                            <a 
                                href="/settings" 
                                class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                                role="menuitem"
                            >
                                设置
                            </a>
                            <button
                                class="block w-full px-4 py-2 text-left text-sm text-gray-700 hover:bg-gray-100"
                                role="menuitem"
                                onclick={handleLogout}
                            >
                                退出
                            </button>
                        </div>
                        {/if}
                    </div>
                </div>
            </div>
        </nav>
        {/if}

        <main class="flex-1">
            <Notifications />
            {@render props.children()}
        </main>
    </div>
</ErrorBoundary> 

<style lang="postcss">
    /* ... 其他样式 ... */
</style> 