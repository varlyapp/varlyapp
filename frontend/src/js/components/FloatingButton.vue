<script setup lang="ts">
import { useRouter } from 'vue-router'
import { load, runtime, navigate } from '@utils/Varly'

load({ router: useRouter() })

const props = defineProps(['emoji', 'text', 'to', 'path'])

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
    <div class="absolute z-auto bottom-8 right-8">
        <button
            type="button"
            class="select-none flex min-w-full mt-2 py-2 px-4 items-center rounded-full uppercase text-slate-50 bg-fuchsia-700 shadow-lg shadow-fuchsia-900 hover:bg-opacity-90 font-bold"
            @click="handleClick"
        >
            <span>
                <slot></slot>
            </span>
            <span v-html="text" />
        </button>
    </div>
</template>
