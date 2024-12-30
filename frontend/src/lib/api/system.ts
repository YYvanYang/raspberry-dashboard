import { API_BASE_URL } from '$lib/config';

// 获取系统状态
export async function getSystemStatus(): Promise<SystemStatus> {
  const response = await fetch(`${API_BASE_URL}/system/status`, {
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    }
  });
  
  if (!response.ok) {
    throw new Error('Failed to fetch system status');
  }
  
  return response.json();
}

// 获取系统指标
export async function getSystemMetrics(): Promise<SystemMetrics> {
  const response = await fetch(`${API_BASE_URL}/monitor/metrics`, {
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    }
  });

  if (!response.ok) {
    throw new Error('Failed to fetch system metrics');
  }

  return response.json();
}

// 控制服务
export async function controlService(name: string, action: 'start' | 'stop' | 'restart'): Promise<void> {
  const response = await fetch(`${API_BASE_URL}/system/services/${name}?action=${action}`, {
    method: 'POST',
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    }
  });

  if (!response.ok) {
    throw new Error(`Failed to ${action} service ${name}`);
  }
} 