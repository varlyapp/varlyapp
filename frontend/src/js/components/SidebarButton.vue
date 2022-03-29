<script setup lang="ts">
import { inject } from 'vue'
import type { Varly } from '@root/plugins/varly'

const varly = inject<Varly>('varly')!
const props = defineProps(['emoji', 'text', 'selected', 'to', 'path'])

function handleClick(e: any) {
    if (props.to && typeof props.to === 'string' && props.to.length) {
        return varly.router.push({ name: props.to })
    }

    try {
        return props.to(e)
    } catch (error) {
        console.error(error)
        return false
    }
}
</script>

<template>
    <button
        type="button"
        class="flex min-w-full mt-2 py-2 px-4 items-center rounded hover:bg-slate-900 hover:bg-opacity-5 hover:dark:bg-slate-50 hover:dark:bg-opacity-5"
        :class="[props.selected ? 'bg-slate-900 bg-opacity-5 dark:bg-slate-50 dark:bg-opacity-5' : '']"
        @click="handleClick"
    >
        <span><slot></slot></span>
        <span v-text="text" />
    </button>
</template>
