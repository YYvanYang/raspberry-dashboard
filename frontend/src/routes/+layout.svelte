<script lang="ts">
    import "../app.css";
    import { onMount } from "svelte";
    import { page } from "$app/stores";
    import { goto } from "$app/navigation";
    import { auth } from "$lib/stores/auth";
    import Navigation from "$lib/components/Navigation.svelte";
    import ErrorBoundary from "$lib/components/ErrorBoundary.svelte";
    import Notifications from "$lib/components/Notifications.svelte";

    onMount(() => {
        const token = localStorage.getItem("token");
        const publicPaths = ['/login'];
        
        if (!token && !publicPaths.includes($page.url.pathname)) {
            goto("/login");
        }
    });

    $: isLoginPage = $page.url.pathname === '/login';
</script>

{#if !isLoginPage}
    <div class="min-h-screen bg-gray-100">
        <Navigation />
        <main class="py-6">
            <ErrorBoundary>
                <Notifications />
                <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                    <slot />
                </div>
            </ErrorBoundary>
        </main>
    </div>
{:else}
    <ErrorBoundary>
        <Notifications />
        <slot />
    </ErrorBoundary>
{/if} 