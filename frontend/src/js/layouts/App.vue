<template>
  <!-- url(assets/images/welcome-illustration.svg) -->
  <div
    class="h-full flex flex-col bg-cover bg-center"
    :style="{ backgroundImage: isOnWelcomeScreen ? '' : '' }"
  >
    <header data-wails-drag>
      <nav class="flex items-center justify-end px-2 py-4">
        <ul class="flex items-center">
          <!-- <li>
            <router-link :to="{ name: 'welcome' }" class="px-2 py-4 opacity-75 font-semibold">Home</router-link>
          </li> -->
          <!-- <li>
            <button @click="open('https://twitter.com/varlyapp')" class="p-2 opacity-50">Support</button>
          </li> -->
        </ul>
      </nav>
    </header>
    <router-view class="flex-1 overflow-auto scrollbar-none"></router-view>
    <footer class="text-center px-8 py-4">
      <p class="text-sm opacity-75">&copy;2022 Varly v1.0.0</p>
    </footer>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const isOnWelcomeScreen = ref(false)

function open(url) {
  runtime.BrowserOpenURL(url)
}

function setIsOnWelcomeScreen(to) {
  isOnWelcomeScreen.value = to.name === 'welcome'
}

router.beforeResolve(setIsOnWelcomeScreen)
</script>
