<template>
  <router-view class="relative h-full overflow-hidden"></router-view>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import { ref, onBeforeUnmount, onBeforeMount } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from '@root/store'

const intl = useI18n({ useScope: 'global' })
const router = useRouter()
const store = useStore()

const isOnStartScreen = ref(false)
onBeforeUnmount(() => store.actions = [])

function setIsOnStartScreen(to) {
  isOnStartScreen.value = to.name === 'start'
}

router.beforeResolve((to) => {
  store.actions = []

  setIsOnStartScreen(to)
})
</script>
