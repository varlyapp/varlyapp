<script setup lang="ts">
import { nextTick, onBeforeMount, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { types } from '@wails/go/models'
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
    debugger
    const layers = { ...collectionStore.layers }

    if (!layers || !Object.keys(layers).length) {
        console.log('No preview to generate')
        return
    }

    for (const trait in layers) {
        if (layers.hasOwnProperty(trait)) {
            layers[trait] = layers[trait].map((layer) => {
                return {
                    ...layer,
                    weight: parseFloat(layer.weight)
                }
            })
        }
    }

    const collection = types.Collection.createFrom({
        ...collectionStore.prepare(),
        sourceDirectory: collectionStore.sourceDirectory,
        outputDirectory: collectionStore.outputDirectory,
        traits: [...collectionStore.traits],
        layers: layers,
        width: parseFloat(collectionStore.width.toString()),
        height: parseFloat(collectionStore.height.toString()),
        size: parseInt(collectionStore.size.toString(), 10)
    })

    console.log(collection)
    const path = await rpc.CollectionService.SaveCollection(types.Collection.createFrom(collection))
    if (path && path !== '') store.addDocument(path)
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
