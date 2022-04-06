<script setup lang="ts">
import { log } from '@/utils/backend'
import { LoggerType } from '@/types'

const props = defineProps(['icon', 'text'])
</script>

<template>
    <div class="mt-8 text-center mx-auto">
        <button
            type="button"
            class="select-none text-base py-2 px-6 text-center rounded-full uppercase text-slate-50 bg-fuchsia-700 shadow-lg shadow-fuchsia-900 hover:bg-opacity-90 font-bold"
            @click="$emit('tap', $event)"
        >
            <span class="flex items-center">
                <span>
                    <component
                        :is="props.icon"
                        class="w-6 mr-2 fill-slate-800 dark:fill-slate-100"
                    />
                </span>
                <span v-text="text" />
            </span>
        </button>
    </div>
</template>
