<script setup lang="ts">
const props = defineProps(['variants'])

function getRowClasses(index) {
    // Odd/even support
    if (index % 2 === 0) {
        return `bg-slate-200 dark:bg-slate-800 bg-opacity-10 dark:bg-opacity-5`
    }

    return `bg-slate-800 dark:bg-slate-400 bg-opacity-5 dark:bg-opacity-5`
}
</script>

<template>
    <div :key="i" v-for="(variant, i) in props.variants"
        class="min-w-full flex items-center justify-between border-t border-slate-900 dark:border-slate-50 border-opacity-20 dark:border-opacity-20"
        :class="getRowClasses(i)">
        <div class="w-1/">
            <img class="w-16 p-2" :src="variant.path" alt="" />
        </div>

        <div
            class="w-4/6 border-r border-l border-dotted text-left whitespace-nowrap px-4 text-sm font-medium sm:px-6 lg:px-8"
            v-text="variant.name" />
        <div class="w-1/6 pr-2">
            <input
                type="text"
                class="field grow-0 text-right appearance-none bg-transparent border-0"
                v-model="variant.weight" />
        </div>
    </div>
</template>
