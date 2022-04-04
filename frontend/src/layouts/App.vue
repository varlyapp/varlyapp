<template>
  <router-view class="relative h-full overflow-hidden"></router-view>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { load } from '@/utils/Varly'

const intl = useI18n({ useScope: 'global' })
const router = useRouter()
const isOnStartScreen = ref(false)

onMounted(() => load({ intl, router }) )

function setIsOnStartScreen(to: any) {
  isOnStartScreen.value = to.name === 'start'
}

router.beforeResolve((to: any) => {
  setIsOnStartScreen(to)
})
</script>
