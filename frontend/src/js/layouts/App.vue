<template>
  <router-view class="relative h-full overflow-hidden"></router-view>
</template>

<script setup lang="ts">
import { ref, onBeforeUnmount, onBeforeMount } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from '@root/store'

const router = useRouter()
const store = useStore()

const isOnStartScreen = ref(false)

onBeforeMount(() => {
  window.runtime.EventsOn('shortcut.save', () => {
    // @todo Implement
  })
})

onBeforeUnmount(() => store.actions = [])

function setIsOnStartScreen(to) {
  isOnStartScreen.value = to.name === 'start'
}

router.beforeResolve((to) => {
  store.actions = []

  setIsOnStartScreen(to)
})
</script>
