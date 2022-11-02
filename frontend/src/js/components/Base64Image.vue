<script setup lang="ts">
import rpc from '@/rpc'
import { ref, onBeforeMount } from 'vue'

const props = defineProps(['path'])
const image = ref('')

onBeforeMount(async () => {
    image.value = await rpc.app.EncodeImage(props.path)
})
</script>

<template>
    <img :src="image" alt="Base64 Encoded Image"/>
</template>
