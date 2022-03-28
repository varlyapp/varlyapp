<script setup lang="ts">
import { inject } from 'vue'
import type { Varly } from '@root/plugins/varly'

const varly = inject<Varly>('varly')!
const props = defineProps(['emoji', 'text', 'selected', 'to'])

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
        class="flex min-w-full mt-2 py-2 px-4 opacity-75 dark:opacity-75 rounded hover:bg-slate-900 hover:bg-opacity-10 hover:dark:bg-slate-50 hover:dark:bg-opacity-10"
        :class="[props.selected ? 'bg-slate-900 bg-opacity-10 dark:bg-slate-50 dark:bg-opacity-10' : '']"
        @click="handleClick"
    >
        <span class="text-center" style="width: 28pt" v-text="emoji" />
        <span v-text="text" />
    </button>
</template>
