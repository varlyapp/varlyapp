<script setup lang="ts">
import { useVarly } from '@/varly'

const varly = useVarly()
const props = defineProps(['name', 'text', 'click', 'size'])

const src = `assets/icons/${varly.isDark ? 'dark/' : 'light/'}${props.name}.svg`
</script>

<template>
    <div class="text-center">
        <button @click="$emit('click')" type="button">
            <img class="h-auto" :style="`width: ${size ?? '64pt'}`" :src="src" alt="" />
        </button>
        <h1 v-if="text" v-text="text" class="mt-4 text-sm"></h1>
    </div>
</template>
