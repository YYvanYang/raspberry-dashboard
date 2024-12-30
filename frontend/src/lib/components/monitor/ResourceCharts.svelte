<script lang="ts">
import * as echarts from 'echarts';
import { onMount } from 'svelte';

interface ChartData {
  cpu: number;
  memory: number;
  disk: number;
  timestamp: number;
}

// 使用 runes 声明状态和属性
let { maxPoints = 50 } = $props<{ maxPoints?: number }>();
let cpuChart = $state<echarts.ECharts | null>(null);
let memoryChart = $state<echarts.ECharts | null>(null);
let diskChart = $state<echarts.ECharts | null>(null);

// 图表数据
const chartData = $state<{
  cpu: [number, number][];
  memory: [number, number][];
  disk: [number, number][];
}>({
  cpu: [],
  memory: [],
  disk: []
});

// 图表配置
const baseOptions = {
  tooltip: { trigger: 'axis' },
  xAxis: { 
    type: 'time',
    axisLabel: { formatter: '{HH:mm:ss}' }
  },
  yAxis: { 
    type: 'value',
    max: 100,
    min: 0
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  }
};

const chartOptions = $derived({
  cpu: {
    ...baseOptions,
    title: { text: 'CPU 使用率' },
    series: [{
      name: 'CPU',
      type: 'line',
      data: chartData.cpu,
      areaStyle: {},
      smooth: true
    }]
  },
  memory: {
    ...baseOptions,
    title: { text: '内存使用率' },
    series: [{
      name: '内存',
      type: 'line',
      data: chartData.memory,
      areaStyle: {},
      smooth: true
    }]
  },
  disk: {
    ...baseOptions,
    title: { text: '磁盘使用率' },
    series: [{
      name: '磁盘',
      type: 'line',
      data: chartData.disk,
      areaStyle: {},
      smooth: true
    }]
  }
});

// 初始化图表
onMount(() => {
  const cpuDom = document.getElementById('cpuChart');
  const memoryDom = document.getElementById('memoryChart');
  const diskDom = document.getElementById('diskChart');
  
  if (cpuDom) cpuChart = echarts.init(cpuDom);
  if (memoryDom) memoryChart = echarts.init(memoryDom);
  if (diskDom) diskChart = echarts.init(diskDom);
  
  cpuChart?.setOption(chartOptions.cpu);
  memoryChart?.setOption(chartOptions.memory);
  diskChart?.setOption(chartOptions.disk);

  const handleResize = () => {
    cpuChart?.resize();
    memoryChart?.resize();
    diskChart?.resize();
  };

  window.addEventListener('resize', handleResize);
  
  return () => {
    window.removeEventListener('resize', handleResize);
    cpuChart?.dispose();
    memoryChart?.dispose();
    diskChart?.dispose();
  };
});

// 更新数据方法
export function updateData(data: ChartData) {
  const { cpu, memory, disk, timestamp } = data;
  
  chartData.cpu.push([timestamp, cpu]);
  chartData.memory.push([timestamp, memory]); 
  chartData.disk.push([timestamp, disk]);

  // 保持固定数量的数据点
  if (chartData.cpu.length > maxPoints) {
    chartData.cpu.shift();
    chartData.memory.shift();
    chartData.disk.shift();
  }
  
  cpuChart?.setOption(chartOptions.cpu);
  memoryChart?.setOption(chartOptions.memory);
  diskChart?.setOption(chartOptions.disk);
}
</script>

<div class="grid grid-cols-1 md:grid-cols-3 gap-4 p-4">
  <div class="h-[300px] bg-white rounded-lg shadow" id="cpuChart"></div>
  <div class="h-[300px] bg-white rounded-lg shadow" id="memoryChart"></div>
  <div class="h-[300px] bg-white rounded-lg shadow" id="diskChart"></div>
</div> 