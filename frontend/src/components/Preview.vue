<script setup lang="ts">
const props = defineProps(['source', 'caption'])
</script>

<template>
    <div>
        <p v-if="caption" class="py-2 text-center text-slate-900 dark:text-slate-100 opacity-50 font-semibold" v-text="caption" />
        <img :src="props.source" class="w-full m-1 animate__animated animate__fadeIn shadow-2xl border-2 border-opacity-50" alt="Preview" />
    </div>
</template>
