<template>
  <keep-alive>
    <router-view class="relative h-full overflow-hidden"></router-view>
  </keep-alive>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { Collection } from '@/wailsjs/go/models'
import { useCollectionStore } from '@/store'
import rpc from '@/rpc'

const router = useRouter()
const store = useCollectionStore()
const isOnStartScreen = ref(false)

onMounted(() => {
  window.runtime.EventsOn('shortcut.collection.open', async () => {
    const collection = await rpc.CollectionService.LoadCollection()
    store.hydrate(collection)
    router.push({ name: 'artwork.layers' })
  })
  window.runtime.EventsOn('shortcut.collection.save', async () => {
    if (store.layers && Object.keys(store.layers).length) {
      const error = await rpc.CollectionService.SaveCollection(Collection.createFrom(store.prepare()))

      if (error) {
        console.error(error)
      }
    }
  })
})

function setIsOnStartScreen(to: any) {
  isOnStartScreen.value = to.name === 'start'
}

router.beforeResolve((to: any) => {
  setIsOnStartScreen(to)
})
</script>
