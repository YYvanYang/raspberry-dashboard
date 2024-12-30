<script lang="ts">
    // 按需引入 ECharts 核心模块
    import * as echarts from 'echarts/core';
    import {
        TitleComponent,
        ToolboxComponent,
        TooltipComponent,
        GridComponent,
        LegendComponent
    } from 'echarts/components';
    import { LineChart } from 'echarts/charts';
    import { UniversalTransition } from 'echarts/features';
    import { CanvasRenderer } from 'echarts/renderers';

    // 注册必需的组件
    echarts.use([
        TitleComponent,
        ToolboxComponent,
        TooltipComponent,
        GridComponent,
        LegendComponent,
        LineChart,
        CanvasRenderer,
        UniversalTransition
    ]);

    import { onMount, onDestroy } from 'svelte';

    let cpuChart: echarts.ECharts | null = null;
    let memoryChart: echarts.ECharts | null = null;
    let diskChart: echarts.ECharts | null = null;

    // 图表配置和数据
    let cpuData = $state<number[]>([]);
    let memoryData = $state<number[]>([]);
    let diskData = $state<number[]>([]);
    let timeData = $state<string[]>([]);

    // 基础配置 - 移除对 timeData 的直接引用
    const baseOption = {
        title: {
            left: 'center'
        },
        tooltip: {
            trigger: 'axis',
            axisPointer: {
                type: 'cross',
                label: {
                    backgroundColor: '#6a7985'
                }
            }
        },
        toolbox: {
            feature: {
                saveAsImage: {}
            }
        },
        grid: {
            left: '3%',
            right: '4%',
            bottom: '3%',
            containLabel: true
        },
        xAxis: {
            type: 'category',
            boundaryGap: false,
            data: [] // 初始为空数组
        },
        yAxis: {
            type: 'value',
            max: 100,
            min: 0
        }
    };

    // 获取图表选项的函数
    function getChartOption(title: string, data: number[]) {
        return {
            ...baseOption,
            title: { text: title },
            xAxis: {
                ...baseOption.xAxis,
                data: timeData // 在闭包中引用 timeData
            },
            series: [{
                name: title,
                type: 'line',
                data: data,
                areaStyle: {},
                smooth: true
            }]
        };
    }

    // 初始化图表
    function initCharts() {
        const cpuDom = document.getElementById('cpuChart');
        const memoryDom = document.getElementById('memoryChart');
        const diskDom = document.getElementById('diskChart');

        if (cpuDom) {
            cpuChart = echarts.init(cpuDom);
            cpuChart.setOption(getChartOption('CPU 使用率 (%)', cpuData));
        }

        if (memoryDom) {
            memoryChart = echarts.init(memoryDom);
            memoryChart.setOption(getChartOption('内存使用率 (%)', memoryData));
        }

        if (diskDom) {
            diskChart = echarts.init(diskDom);
            diskChart.setOption(getChartOption('磁盘使用率 (%)', diskData));
        }
    }

    // 更新图表数据
    $effect(() => {
        if (cpuChart && memoryChart && diskChart) {
            cpuChart.setOption({
                xAxis: { data: timeData },
                series: [{ data: cpuData }]
            });
            memoryChart.setOption({
                xAxis: { data: timeData },
                series: [{ data: memoryData }]
            });
            diskChart.setOption({
                xAxis: { data: timeData },
                series: [{ data: diskData }]
            });
        }
    });

    // 窗口大小改变时重绘图表
    function handleResize() {
        cpuChart?.resize();
        memoryChart?.resize();
        diskChart?.resize();
    }

    onMount(() => {
        initCharts();
        window.addEventListener('resize', handleResize);
    });

    onDestroy(() => {
        window.removeEventListener('resize', handleResize);
        cpuChart?.dispose();
        memoryChart?.dispose();
        diskChart?.dispose();
    });

    // 导出更新数据的方法
    export function updateData(newData: {
        cpu: number;
        memory: number;
        disk: number;
        time: string;
    }) {
        const MAX_DATA_POINTS = 20;

        cpuData = [...cpuData.slice(-MAX_DATA_POINTS + 1), newData.cpu];
        memoryData = [...memoryData.slice(-MAX_DATA_POINTS + 1), newData.memory];
        diskData = [...diskData.slice(-MAX_DATA_POINTS + 1), newData.disk];
        timeData = [...timeData.slice(-MAX_DATA_POINTS + 1), newData.time];
    }
</script>

<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
    <div class="bg-white p-4 rounded-lg shadow">
        <div id="cpuChart" class="h-64"></div>
    </div>
    <div class="bg-white p-4 rounded-lg shadow">
        <div id="memoryChart" class="h-64"></div>
    </div>
    <div class="bg-white p-4 rounded-lg shadow">
        <div id="diskChart" class="h-64"></div>
    </div>
</div> 