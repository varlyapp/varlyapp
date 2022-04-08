<template>
  <router-view class="relative h-full overflow-hidden"></router-view>
</template>

<script setup lang="ts">
import {  onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useCollectionStore } from '@/store'
import { runtime } from '@/utils/backend'
import rpc from '@/rpc'

const store = useCollectionStore()
const router = useRouter()
const isOnStartScreen = ref(false)

onMounted(() => {
  console.log(store.$state)
  runtime().EventsOn('shortcut.collection.open', async () => {
    const collection = await rpc.CollectionService.LoadCollection()
    console.log({ collection })
    store.hydrate(collection)
    console.log({ state: { ...store.$state }})
    router.push({name: 'artwork.layers'})
  })
  runtime().EventsOn('shortcut.collection.save', () => {
    console.log('Save collection shortcut called')
  })
})

function setIsOnStartScreen(to: any) {
  isOnStartScreen.value = to.name === 'start'
}

router.beforeResolve((to: any) => {
  setIsOnStartScreen(to)
})
</script>
