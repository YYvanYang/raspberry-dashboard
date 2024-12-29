/// <reference types="svelte" />

declare module "svelte" {
    interface ComponentProps {
        warning?: boolean;
    }
}

declare global {
    const $state: <T>(value: T) => T;
    const $derived: <T>(fn: () => T) => T;
}

export {}; 