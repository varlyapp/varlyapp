<template>
    <div class="text-center">
        <button @click="$emit('click')" type="button">
            <img class="h-auto" style="width: 64pt" :src="src" alt />
        </button>
        <h1 v-if="text" v-text="text" class="mt-4 text-sm"></h1>
    </div>
</template>

<script setup lang="ts">
import { useTheme } from '@utils/Theme'

const props = defineProps(['name', 'text', 'click'])
const theme = useTheme()

const src = `assets/icons/${theme.isDark ? 'dark/' : 'light/'}${props.name}.svg`
</script>
