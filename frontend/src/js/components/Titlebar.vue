<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps({
    title: String
})

const titleText = computed(() => {
    if (props.title) {
        return `${props.title}`
    }
    return ''
})
</script>

<template>
    <header
        data-wails-drag
        class="sticky bg-slate-900 bg-opacity-10 top-0 border-b border-slate-900 border-opacity-10 dark:border-slate-50 dark:border-opacity-10"
    >
        <div class="p-3">
            <h1
                v-text="titleText"
                class="px-4 text-sm font-semibold text-slate-900 text-opacity-80 dark:text-slate-50 dark:text-opacity-80"
            />
        </div>
    </header>
</template>
