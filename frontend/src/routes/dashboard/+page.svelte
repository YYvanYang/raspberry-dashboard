<script lang="ts">
import ResourceCharts from '$lib/components/monitor/ResourceCharts.svelte';

// 使用 runes 声明状态
let systemInfo = $state({
  hostname: '',
  uptime: '',
  os: '',
  arch: ''
});

let services = $state([
  { name: 'nginx', status: 'running' },
  { name: 'mysql', status: 'stopped' },
  { name: 'redis', status: 'running' }
]);

// 获取系统信息
async function fetchSystemInfo() {
  // TODO: 实现从后端 API 获取系统信息
}

// 获取服务状态
async function fetchServices() {
  // TODO: 实现从后端 API 获取服务状态
}
</script>

<div class="container mx-auto p-4">
  <h1 class="text-2xl font-bold mb-4">系统仪表盘</h1>
  
  <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-8">
    <div class="bg-white p-4 rounded-lg shadow">
      <h2 class="text-xl font-semibold mb-4">系统信息</h2>
      <div class="space-y-2">
        <p>主机名: {systemInfo.hostname}</p>
        <p>运行时间: {systemInfo.uptime}</p>
        <p>操作系统: {systemInfo.os}</p>
        <p>架构: {systemInfo.arch}</p>
      </div>
    </div>
    
    <div class="bg-white p-4 rounded-lg shadow">
      <h2 class="text-xl font-semibold mb-4">服务状态</h2>
      <div class="space-y-2">
        {#each services as service}
          <div class="flex justify-between items-center">
            <span>{service.name}</span>
            <span class={service.status === 'running' ? 'text-green-500' : 'text-red-500'}>
              {service.status}
            </span>
          </div>
        {/each}
      </div>
    </div>
  </div>

  <div class="bg-white p-4 rounded-lg shadow">
    <h2 class="text-xl font-semibold mb-4">资源监控</h2>
    <ResourceCharts />
  </div>
</div> 