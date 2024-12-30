<script lang="ts">
import ResourceCharts from '$lib/components/monitor/ResourceCharts.svelte';
import { onMount } from 'svelte';

let resourceCharts: ResourceCharts;

// 模拟数据更新
onMount(() => {
  const interval = setInterval(() => {
    resourceCharts?.updateData({
      cpu: Math.random() * 100,
      memory: Math.random() * 100,
      disk: Math.random() * 100,
      timestamp: Date.now()
    });
  }, 2000);

  return () => clearInterval(interval);
});
</script>

<div class="container mx-auto">
  <h1 class="text-2xl font-bold mb-4">系统资源监控</h1>
  <ResourceCharts bind:this={resourceCharts} />
</div> 