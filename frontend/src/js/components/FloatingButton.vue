<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()
const props = defineProps(['emoji', 'text', 'to', 'path', 'icon'])

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
    <div class="absolute z-auto bottom-12 right-8">
        <button
            type="button"
            class="select-none flex min-w-full mt-2 py-2 px-6 items-center rounded-full uppercase text-slate-50 bg-fuchsia-700 shadow-md shadow-fuchsia-900 hover:bg-opacity-90 font-bold"
            @click="handleClick"
        >
            <span>
                <component :is="props.icon" class="w-6 mr-2 fill-fuchsia-100" />
            </span>
            <span v-html="text" />
        </button>
    </div>
</template>
