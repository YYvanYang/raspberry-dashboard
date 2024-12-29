import { writable } from "svelte/store";

interface Notification {
    id: number;
    type: "success" | "error";
    message: string;
}

function createNotificationsStore() {
    const { subscribe, update } = writable<Notification[]>([]);

    let nextId = 1;

    function addNotification(type: "success" | "error", message: string) {
        const notification: Notification = {
            id: nextId++,
            type,
            message
        };

        update(notifications => [...notifications, notification]);

        // 3秒后自动移除
        setTimeout(() => {
            update(notifications => notifications.filter(n => n.id !== notification.id));
        }, 3000);
    }

    return {
        subscribe,
        success: (message: string) => addNotification("success", message),
        error: (message: string) => addNotification("error", message),
        remove: (id: number) => update(notifications => 
            notifications.filter(n => n.id !== id)
        )
    };
}

export const notifications = createNotificationsStore(); 