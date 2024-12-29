<script lang="ts">
    const { temperature } = $props<{
        temperature: number;
    }>();

    const formattedTemp = $derived(temperature.toFixed(1));
    const tempColor = $derived(
        temperature >= 80 ? 'text-red-600' :
        temperature >= 70 ? 'text-orange-500' :
        temperature >= 60 ? 'text-yellow-500' :
        'text-green-500'
    );
</script>

<div class="bg-white overflow-hidden shadow rounded-lg">
    <div class="p-5">
        <div class="flex flex-col gap-4">
            <div>
                <div class="text-sm font-medium text-gray-500">CPU Temperature</div>
                <div class="mt-1 flex items-center">
                    <div class={`text-2xl font-semibold ${tempColor}`}>{formattedTemp}°C</div>
                    <div class="ml-2 w-full bg-gray-200 rounded-full h-2.5">
                        <div class={`h-2.5 rounded-full ${tempColor.replace('text-', 'bg-')}`} style="width: {Math.min(temperature, 100)}%"></div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div> 