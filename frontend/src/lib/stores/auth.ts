import { writable } from "svelte/store";

interface AuthState {
    token: string | null;
}

function createAuthStore() {
    const { subscribe, set } = writable<AuthState>({
        token: null,
    });

    return {
        subscribe,
        set: (state: AuthState) => set(state),
        reset: () => set({ token: null }),
    };
}

export const auth = createAuthStore(); 