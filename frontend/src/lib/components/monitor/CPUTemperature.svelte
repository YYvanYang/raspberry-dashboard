<script lang="ts">
    export let temperature: number;

    function getTemperatureColor(temp: number): string {
        if (temp >= 80) return 'text-red-600';
        if (temp >= 70) return 'text-orange-500';
        if (temp >= 60) return 'text-yellow-500';
        return 'text-green-500';
    }
</script>

<div class="bg-white overflow-hidden shadow rounded-lg">
    <div class="p-5">
        <div class="flex items-center">
            <div class="flex-shrink-0">
                <!-- Temperature icon -->
                <svg class="h-6 w-6 {getTemperatureColor(temperature)}" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v3m0 0v3m0-3h3m-3 0H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
            </div>
            <div class="ml-5 w-0 flex-1">
                <dl>
                    <dt class="text-sm font-medium text-gray-500 truncate">
                        CPU Temperature
                    </dt>
                    <dd class="text-2xl font-semibold {getTemperatureColor(temperature)}">
                        {temperature?.toFixed(1) || '--'}°C
                    </dd>
                </dl>
            </div>
        </div>
    </div>
</div> 