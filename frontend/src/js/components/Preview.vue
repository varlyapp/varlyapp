<script setup lang="ts">
const props = defineProps(['source', 'caption'])
</script>

<template>
    <div class="border border-slate-800 dark:border-slate-100 border-opacity-25 dark:border-opacity-25">
        <img v-if="props.source && props.source.length" :src="props.source" class="w-full animate__animated animate__fadeIn" alt="Preview" />
    </div>
</template>
