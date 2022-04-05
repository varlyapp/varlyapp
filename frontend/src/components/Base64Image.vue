<script setup lang="ts">
import { ref, onBeforeMount } from 'vue'
import { api } from '@/utils/backend'

const props = defineProps(['path'])
const image = ref('')

onBeforeMount(async () => {
    image.value = await api().EncodeImage(props.path)
})
</script>

<template>
    <img :src="image" alt="Base64 Encoded Image"/>
</template>
