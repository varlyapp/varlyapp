<script setup lang="ts">
import { ref, onBeforeMount } from 'vue'

const image = ref('')

const props = defineProps(['path'])

onBeforeMount(async () => {
    image.value = await window.go.main.App.EncodeImage(props.path)
})
</script>

<template>
    <img :src="image" alt="Base64 Encoded Image"/>
</template>
