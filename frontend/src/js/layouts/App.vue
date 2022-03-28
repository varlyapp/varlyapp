<template>
  <div class="h-full flex flex-col">
    <!-- :style="{ backgroundImage: isOnStartScreen ? 'url(assets/images/varly-background.png)' : '' }" -->
    <header data-wails-drag class="p-1 border-b border-slate-900 border-opacity-10 dark:border-slate-50 dark:border-opacity-10">
      <h1 class="text-center text-base text-slate-900 text-opacity-80 dark:text-slate-50 dark:text-opacity-80">🦄 Varly</h1>
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
