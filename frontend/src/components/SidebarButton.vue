<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()
const props = defineProps(['icon', 'text', 'selected', 'to', 'path'])

async function handleClick(): Promise<void> {
    if (props.to && typeof props.to === 'string' && props.to.length) {
        router.push({ name: props.to })
        return
    }

    try {
        return Promise.resolve(props.to())
    } catch (error) {
       console.error(`unable to handle click: ${error}`)
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
        <span>
            <component
                :is="props.icon"
                class="w-6 mr-2 fill-fuchsia-700 dark:fill-fuchsia-500"
            />
        </span>
        <span v-text="text" />
    </button>
</template>
