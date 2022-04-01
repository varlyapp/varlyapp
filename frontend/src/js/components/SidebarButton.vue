<script setup lang="ts">
import { useRouter } from 'vue-router'
import { load, runtime, navigate } from '@utils/Varly'

load({ router: useRouter() })

const props = defineProps(['text', 'selected', 'to', 'path'])

async function handleClick(e: any): Promise<void> {
    if (props.to && typeof props.to === 'string' && props.to.length) {
        return navigate(props.to)
    }

    try {
        return Promise.resolve(props.to())
    } catch (error) {
        runtime.LogError(`unable to handle click: ${error}`)
    }
}
</script>

<template>
    <button
        type="button"
        class="select-none flex min-w-full mt-2 py-2 px-4 items-center rounded hover:bg-slate-900 hover:bg-opacity-5 hover:dark:bg-slate-50 hover:dark:bg-opacity-5"
        :class="[props.selected ? 'bg-slate-900 bg-opacity-10 dark:bg-slate-50 dark:bg-opacity-5' : '']"
        @click="handleClick"
    >
        <span><slot></slot></span>
        <span v-text="text" />
    </button>
</template>
