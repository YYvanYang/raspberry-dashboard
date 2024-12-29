<script lang="ts">
    import { onMount } from "svelte";
    import { api } from "$lib/utils/api";
    import { notifications } from "$lib/stores/notifications";

    let currentPassword = $state("");
    let newPassword = $state("");
    let confirmPassword = $state("");
    let updating = $state(false);

    async function handlePasswordChange() {
        if (newPassword !== confirmPassword) {
            notifications.error("New passwords do not match");
            return;
        }

        updating = true;
        try {
            await api.post("/auth/change-password", {
                current_password: currentPassword,
                new_password: newPassword
            });
            notifications.success("Password updated successfully");
            currentPassword = "";
            newPassword = "";
            confirmPassword = "";
        } catch (error) {
            notifications.error("Failed to update password");
        } finally {
            updating = false;
        }
    }
</script>

<div class="min-h-screen bg-gray-100">
    <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div class="px-4 py-6 sm:px-0">
            <div class="bg-white shadow rounded-lg">
                <div class="px-4 py-5 sm:p-6">
                    <h3 class="text-lg font-medium leading-6 text-gray-900">
                        Change Password
                    </h3>
                    
                    <div class="mt-6 space-y-6">
                        <div>
                            <label for="current-password" class="block text-sm font-medium text-gray-700">
                                Current Password
                            </label>
                            <input
                                type="password"
                                id="current-password"
                                bind:value={currentPassword}
                                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm"
                            />
                        </div>

                        <div>
                            <label for="new-password" class="block text-sm font-medium text-gray-700">
                                New Password
                            </label>
                            <input
                                type="password"
                                id="new-password"
                                bind:value={newPassword}
                                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm"
                            />
                        </div>

                        <div>
                            <label for="confirm-password" class="block text-sm font-medium text-gray-700">
                                Confirm New Password
                            </label>
                            <input
                                type="password"
                                id="confirm-password"
                                bind:value={confirmPassword}
                                class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm"
                            />
                        </div>

                        <button
                            type="button"
                            disabled={updating}
                            onclick={handlePasswordChange}
                            class="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50"
                        >
                            {updating ? 'Updating...' : 'Update Password'}
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div> 