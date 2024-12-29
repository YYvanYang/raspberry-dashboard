export type ApiError = {
    error: string;
    status: number;
};

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:3001';

async function fetchWithAuth(endpoint: string, options: RequestInit = {}) {
    const token = localStorage.getItem('token');
    
    const headers = {
        'Content-Type': 'application/json',
        ...(token ? { 'Authorization': `Bearer ${token}` } : {}),
        ...options.headers,
    };

    try {
        const response = await fetch(`${API_URL}${endpoint}`, {
            ...options,
            headers,
        });

        if (response.status === 401) {
            localStorage.removeItem('token');
            window.location.href = '/login';
            throw new Error('Unauthorized');
        }

        if (!response.ok) {
            const error = await response.json();
            throw { error: error.error || 'An error occurred', status: response.status };
        }

        return response.json();
    } catch (err) {
        if (err instanceof Error) {
            throw { error: err.message, status: 500 };
        }
        throw err;
    }
}

export const api = {
    get: (endpoint: string) => fetchWithAuth(endpoint),
    
    post: (endpoint: string, data?: any) => fetchWithAuth(endpoint, {
        method: 'POST',
        body: data ? JSON.stringify(data) : undefined,
    }),
    
    put: (endpoint: string, data?: any) => fetchWithAuth(endpoint, {
        method: 'PUT',
        body: data ? JSON.stringify(data) : undefined,
    }),
    
    delete: (endpoint: string) => fetchWithAuth(endpoint, {
        method: 'DELETE',
    }),
}; 