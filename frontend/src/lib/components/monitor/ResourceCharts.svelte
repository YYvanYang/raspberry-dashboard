<script lang="ts">
    import { Line } from 'svelte-chartjs';
    import {
        Chart as ChartJS,
        CategoryScale,
        LinearScale,
        PointElement,
        LineElement,
        Title,
        Tooltip,
        Legend,
        Filler
    } from 'chart.js';

    ChartJS.register(
        CategoryScale,
        LinearScale,
        PointElement,
        LineElement,
        Title,
        Tooltip,
        Legend,
        Filler
    );

    export let data: any[] = [];

    $: chartData = {
        labels: data.map(() => ''),  // 使用空标签，因为我们有足够的数据点
        datasets: [
            {
                label: 'CPU Usage',
                data: data.map(d => d.cpu),
                borderColor: 'rgb(59, 130, 246)',
                backgroundColor: 'rgba(59, 130, 246, 0.1)',
                fill: true,
                tension: 0.3
            },
            {
                label: 'Memory Usage',
                data: data.map(d => d.memory.usage_rate),
                borderColor: 'rgb(34, 197, 94)',
                backgroundColor: 'rgba(34, 197, 94, 0.1)',
                fill: true,
                tension: 0.3
            }
        ]
    };

    const options = {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
            legend: {
                position: 'top' as const,
            },
            title: {
                display: true,
                text: 'System Resources History'
            }
        },
        scales: {
            y: {
                beginAtZero: true,
                max: 100,
                grid: {
                    display: true,
                    color: 'rgba(0, 0, 0, 0.1)'
                }
            },
            x: {
                grid: {
                    display: false
                }
            }
        }
    };
</script>

<div class="bg-white shadow rounded-lg p-4">
    <div class="h-80">
        <Line data={chartData} {options} />
    </div>
</div> 