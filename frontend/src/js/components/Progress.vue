<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps(['steps', 'currentStep', 'loadingText'])

const completed = computed(() => {
    let completed = parseInt(props.currentStep, 10) / parseInt(props.steps, 10) * 100

    if (isNaN(completed)) {
        completed = 0
    }

    return parseInt(completed.toString(), 10)
})
</script>

<template>
    <div class="h-full w-full flex items-center justify-center p-16">
        <div class="flex-1 text-center">
            <h1 v-if="completed > 0" class="text-4xl">{{ currentStep }}/{{ steps }}</h1>
            <h1 v-else class="text-4xl animate-pulse" v-text="loadingText"></h1>
            <div class="flex flex-1 mt-16  h-4 bg-fuchsia-600 bg-opacity-20 rounded-full">
                <div
                    class="h-4 bg-fuchsia-600 rounded-full shadow"
                    :style="{ width: `${completed}%` }"
                ></div>
            </div>
        </div>
    </div>
</template>
