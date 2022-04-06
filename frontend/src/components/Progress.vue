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
            <h1 v-else class="flex items-center justify-center">
                <svg width="38" height="38" viewBox="0 0 38 38" xmlns="http://www.w3.org/2000/svg">
                    <defs>
                        <linearGradient x1="8.042%" y1="0%" x2="65.682%" y2="23.865%" id="a">
                            <stop stop-color="#fff" stop-opacity="0" offset="0%" />
                            <stop stop-color="#fff" stop-opacity=".631" offset="63.146%" />
                            <stop stop-color="#fff" offset="100%" />
                        </linearGradient>
                    </defs>
                    <g fill="none" fill-rule="evenodd">
                        <g transform="translate(1 1)">
                            <path
                                d="M36 18c0-9.94-8.06-18-18-18"
                                id="Oval-2"
                                stroke="url(#a)"
                                stroke-width="2"
                            >
                                <animateTransform
                                    attributeName="transform"
                                    type="rotate"
                                    from="0 18 18"
                                    to="360 18 18"
                                    dur="0.9s"
                                    repeatCount="indefinite"
                                />
                            </path>
                            <circle fill="#fff" cx="36" cy="18" r="1">
                                <animateTransform
                                    attributeName="transform"
                                    type="rotate"
                                    from="0 18 18"
                                    to="360 18 18"
                                    dur="0.9s"
                                    repeatCount="indefinite"
                                />
                            </circle>
                        </g>
                    </g>
                </svg>
            </h1>
            <div class="flex flex-1 mt-8 xl:mt-16 h-4 bg-fuchsia-600 bg-opacity-20 rounded-full">
                <div
                    class="h-4 bg-fuchsia-600 rounded-full shadow"
                    :style="{ width: `${completed}%` }"
                ></div>
            </div>
        </div>
    </div>
</template>
