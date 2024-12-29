<script lang="ts">
    import { onMount } from 'svelte';
    import * as echarts from 'echarts/core';
    import { LineChart } from 'echarts/charts';
    import {
        TitleComponent,
        TooltipComponent,
        GridComponent,
        LegendComponent
    } from 'echarts/components';
    import { UniversalTransition } from 'echarts/features';
    import { CanvasRenderer } from 'echarts/renderers';

    echarts.use([
        TitleComponent,
        TooltipComponent,
        GridComponent,
        LegendComponent,
        LineChart,
        CanvasRenderer,
        UniversalTransition
    ]);

    interface ChartData {
        labels?: string[];
        datasets?: Array<{
            data: number[];
        }>;
    }

    const { data } = $props<{
        data: ChartData;
    }>();

    let chartDom = $state<HTMLDivElement>();
    let myChart = $state<echarts.ECharts>();

    const option = {
        title: {
            text: '系统资源监控',
            left: 'center',
            textStyle: {
                fontSize: 16,
                fontWeight: 'normal'
            }
        },
        tooltip: {
            trigger: 'axis',
            formatter: function(params: any) {
                let result = params[0].axisValue + '<br/>';
                params.forEach((item: any) => {
                    result += `${item.seriesName}: ${item.value}%<br/>`;
                });
                return result;
            }
        },
        legend: {
            data: ['CPU使用率', '内存使用率'],
            bottom: '0'
        },
        grid: {
            left: '3%',
            right: '4%',
            bottom: '60px',
            top: '60px',
            containLabel: true
        },
        xAxis: {
            type: 'category',
            boundaryGap: false,
            data: [],
            axisLabel: {
                rotate: 45
            }
        },
        yAxis: {
            type: 'value',
            min: 0,
            max: 100,
            splitLine: {
                lineStyle: {
                    type: 'dashed'
                }
            }
        },
        series: [
            {
                name: 'CPU使用率',
                type: 'line',
                smooth: true,
                data: [],
                itemStyle: {
                    color: '#67C23A'
                },
                areaStyle: {
                    color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                        { offset: 0, color: 'rgba(103,194,58,0.3)' },
                        { offset: 1, color: 'rgba(103,194,58,0.1)' }
                    ])
                }
            },
            {
                name: '内存使用率',
                type: 'line',
                smooth: true,
                data: [],
                itemStyle: {
                    color: '#409EFF'
                },
                areaStyle: {
                    color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                        { offset: 0, color: 'rgba(64,158,255,0.3)' },
                        { offset: 1, color: 'rgba(64,158,255,0.1)' }
                    ])
                }
            }
        ]
    };

    function initChart() {
        if (chartDom) {
            myChart = echarts.init(chartDom);
            if (myChart) {
                myChart.setOption(option);
                
                const resizeObserver = new ResizeObserver(() => {
                    myChart?.resize();
                });
                resizeObserver.observe(chartDom);

                return () => {
                    resizeObserver.disconnect();
                    myChart?.dispose();
                };
            }
        }
        return () => {};
    }

    function updateChart() {
        if (!data || !myChart) return;

        const timeLabels = data.labels || [];
        const cpuData = data.datasets?.[0].data || [];
        const memoryData = data.datasets?.[1].data || [];

        myChart.setOption({
            xAxis: {
                data: timeLabels
            },
            series: [
                {
                    data: cpuData
                },
                {
                    data: memoryData
                }
            ]
        });
    }

    onMount(() => {
        const cleanup = initChart();
        return cleanup;
    });

    $effect(() => {
        if (data) {
            updateChart();
        }
    });
</script>

<div class="bg-white p-4 rounded-lg shadow">
    <div class="h-[400px]" bind:this={chartDom}></div>
</div> 