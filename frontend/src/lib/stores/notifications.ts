import { writable } from 'svelte/store';

export type NotificationType = 'success' | 'error' | 'warning' | 'info';

export interface Notification {
    id: string;
    type: NotificationType;
    message: string;
    duration?: number;
}

function createNotificationStore() {
    const { subscribe, update } = writable<Notification[]>([]);

    function addNotification(message: string, type: NotificationType = 'info', duration = 3000) {
        const id = crypto.randomUUID();
        update(n => [...n, { id, type, message }]);

        if (duration > 0) {
            setTimeout(() => {
                removeNotification(id);
            }, duration);
        }
    }

    function removeNotification(id: string) {
        update(n => n.filter(notification => notification.id !== id));
    }

    return {
        subscribe,
        success: (message: string, duration?: number) => addNotification(message, 'success', duration),
        error: (message: string, duration?: number) => addNotification(message, 'error', duration),
        warning: (message: string, duration?: number) => addNotification(message, 'warning', duration),
        info: (message: string, duration?: number) => addNotification(message, 'info', duration),
        remove: removeNotification
    };
}

export const notifications = createNotificationStore(); 