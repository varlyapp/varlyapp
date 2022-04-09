<template>
  <router-view class="relative h-full overflow-hidden"></router-view>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useCollectionStore } from '@/store'
import { runtime } from '@/utils/backend'
import rpc from '@/rpc'
import { Collection } from '@/types'
import { Collection as CollectionModel } from '@/wailsjs/go/models'

const store = useCollectionStore()
const router = useRouter()
const isOnStartScreen = ref(false)

onMounted(() => {
  runtime().EventsOn('shortcut.collection.open', async () => {
    const collection = await rpc.CollectionService.LoadCollection()
    store.hydrate(collection)
    router.push({ name: 'artwork.layers' })
  })
  runtime().EventsOn('shortcut.collection.save', async () => {
    if (store.layers && Object.keys(store.layers).length) {
      // const collection: Collection = {
      //   sourceDirectory: '',
      //   outputDirectory: '',
      //   name: '',
      //   description: '',
      //   traits: [ ...store.traits],
      //   layers: { ...store.layers },
      //   width: parseFloat(store.width.toString()),
      //   height: parseFloat(store.height.toString()),
      //   size: parseInt(store.size.toString(), 10),
      // }

      console.log(store.$state)
      await rpc.CollectionService.SaveCollection(CollectionModel.createFrom(store.prepare()))

      console.log('Saved collection?!')
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
