<script setup lang="ts">
import { nextTick, onBeforeMount, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { types } from '@/wailsjs/go/models'
import { useStore, useCollectionStore } from '@/store'
import rpc from '@/rpc'

const intl = useI18n({ useScope: 'global' })
const router = useRouter()
const store = useStore()
const collectionStore = useCollectionStore()

onBeforeMount(() => {
  intl.locale.value = store.locale || 'en'
})

onMounted(() => {
  rpc.on('shortcut.collection.open', async () => {
    const collection = await rpc.CollectionService.LoadCollection()
    collectionStore.hydrate(collection)
    router.push({ name: 'artwork.layers' })
  })
  rpc.on('shortcut.collection.save', async () => {
    if (collectionStore.layers && Object.keys(collectionStore.layers).length) {
      const path = await rpc.CollectionService.SaveCollection(types.Collection.createFrom(collectionStore.prepare()))
      if (path && path !== '') store.addDocument(path)
    }
  })
  rpc.on('shortcut.language.english', () => {
    intl.locale.value = 'en'
    store.setLocale('en')
    nextTick(() => window.location.reload())
  })
  rpc.on('shortcut.language.spanish', () => {
    intl.locale.value = 'es'
    store.setLocale('es')
    nextTick(() => window.location.reload())
  })
  rpc.on('shortcut.collection.print', () => {
    nextTick(() => router.push({ name: 'print' }))
  })
})
</script>

<template>
  <router-view class="relative h-full overflow-hidden" v-slot="{ Component }">
    <keep-alive>
      <component :is="Component" />
    </keep-alive>
  </router-view>
</template>
