<script setup lang="ts">
import { ref, onBeforeMount } from 'vue'

const image = ref('')

const props = defineProps(['path'])

onBeforeMount(async () => {
    image.value = await window.go.main.App.EncodeImage(props.path)
    // let old = 0
    // let path = props.path
    // for (let index = 0; index < 10; index++) {
    //     if (index === 0) {
    //     } else {
    //         path = path.replace(`${old}`, `${index}`)

    //         image.value = await window.go.main.App.EncodeImage(path)

    //         old = index - 1

    //         console.log({ old, index })
    //         console.log({ path, props })
    //     }

    //     await new Promise(resolve => setTimeout(resolve, 2000))
    // }
})
</script>

<template>
    <img :src="image" alt="Base 64 Encoded Image" height="128" />
</template>
