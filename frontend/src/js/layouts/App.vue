<template>
  <div class="h-full flex flex-col" :class="getBackgroundClasses()">
    <!-- :style="{ backgroundImage: isOnStartScreen ? 'url(assets/images/varly-background.png)' : '' }" -->
    <header data-wails-drag>
      <nav class="flex items-center justify-end px-2 py-3">
        <ul class="flex items-center">
          <!-- <li>
            <router-link :to="{ name: 'start' }" class="px-2 py-4 opacity-75 font-semibold">Home</router-link>
          </li>-->
          <!-- <li>
            <button @click="open('https://twitter.com/varlyapp')" class="p-2 opacity-50">Support</button>
          </li>-->
        </ul>
      </nav>
    </header>
    <router-view class="flex-1 overflow-auto scrollbar-none"></router-view>
    <footer class="text-center px-8 py-4">
      <!-- <slot name="footer">
        <p class="text-sm opacity-75">&copy;2022 Varly v1.0.0</p>
      </slot> -->
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const isOnStartScreen = ref(false)

function setIsOnStartScreen(to) {
  isOnStartScreen.value = to.name === 'start'
}

function getBackgroundClasses() {
  if (isOnStartScreen.value) {
    return 'bg-gradient-to-b from-slate-100 to-slate-200 dark:from-purple-600 dark:to-purple-800'
  }

  return 'bg-gradient-to-b from-slate-100 to-slate-200 dark:from-slate-800 dark:to-slate-900'
}

router.beforeResolve(setIsOnStartScreen)
</script>
