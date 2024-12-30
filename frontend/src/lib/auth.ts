const API_BASE_URL = import.meta.env.DEV ? '/api' : '/api';

// 创建一个简单的事件总线
const authEvents = new EventTarget();

export const AUTH_STATUS_CHANGE = 'auth_status_change';

// 通知登录状态变化
export function notifyAuthChange(isLoggedIn: boolean) {
    authEvents.dispatchEvent(new CustomEvent(AUTH_STATUS_CHANGE, { detail: isLoggedIn }));
}

// 监听登录状态变化
export function onAuthChange(callback: (isLoggedIn: boolean) => void) {
    authEvents.addEventListener(AUTH_STATUS_CHANGE, ((event: CustomEvent) => {
        callback(event.detail);
    }) as EventListener);
}

export async function logout() {
    try {
        const response = await fetch(`${API_BASE_URL}/auth/logout`, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            }
        });

        if (response.ok) {
            localStorage.removeItem('token');
            notifyAuthChange(false);
            window.location.href = '/login';
        } else {
            console.error('Logout failed:', await response.text());
        }
    } catch (error) {
        console.error('Logout failed:', error);
    }
} 