<template>
  <div class="h-full flex flex-col bg-slate-50 dark:bg-slate-900">
    <!-- :style="{ backgroundImage: isOnStartScreen ? 'url(assets/images/varly-background.png)' : '' }" -->
    <header data-wails-drag class="p-2 bg-slate-50 shadow dark:bg-slate-900 dark:shadow">
      <h1 class="text-center text-sm font-semibold">✓ Varly</h1>
    </header>
    <router-view class="h-full flex-1 overflow-auto scrollbar-none"></router-view>
  </div>
</template>

<script setup lang="ts">
import { ref, onBeforeUnmount } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from '@root/store'

const router = useRouter()
const store = useStore()

const isOnStartScreen = ref(false)

onBeforeUnmount(() => store.actions = [])

function setIsOnStartScreen(to) {
  isOnStartScreen.value = to.name === 'start'
}

router.beforeResolve((to) => {
  store.actions = [];

  setIsOnStartScreen(to)
})
</script>
