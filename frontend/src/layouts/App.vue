<template>
  <router-view class="relative h-full overflow-hidden"></router-view>
</template>

<script setup lang="ts">
import {  ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const isOnStartScreen = ref(false)

function setIsOnStartScreen(to: any) {
  isOnStartScreen.value = to.name === 'start'
}

router.beforeResolve((to: any) => {
  setIsOnStartScreen(to)
})
</script>