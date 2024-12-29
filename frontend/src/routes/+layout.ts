import { redirect } from '@sveltejs/kit';
import type { LayoutLoad } from './$types';
import { browser } from '$app/environment';

export const load: LayoutLoad = async ({ url }) => {
    const token = browser ? localStorage.getItem('token') : null;
    const isLoginPage = url.pathname === '/login';

    // 如果未登录且不在登录页面，重定向到登录页
    if (!token && !isLoginPage) {
        throw redirect(302, '/login');
    }

    // 如果已登录且在登录页面，重定向到仪表盘
    if (token && isLoginPage) {
        throw redirect(302, '/dashboard');
    }

    return {
        isAuthenticated: !!token
    };
}; 