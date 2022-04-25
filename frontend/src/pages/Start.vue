<script setup lang="ts">
import { nextTick, onActivated, onBeforeMount, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { Collection } from '@/wailsjs/go/models'
import { CogIcon, CollectionIcon, PlayIcon } from '@heroicons/vue/outline'
import Sidebar from '@/components/Sidebar.vue'
import { useStore, useCollectionStore } from '@/store'
import rpc from '@/rpc'

const route = useRoute()
const router = useRouter()
const store = useStore()
const collectionStore = useCollectionStore()

const intl = useI18n({ useScope: 'global' })
const { t } = intl

const projects = ref<Array<Collection>>([])

onBeforeMount(() => rpc.setPageTitle('Welcome'))

onActivated(() => {
  load()
  nextTick(() => {
    window.runtime.EventsOn('shortcut.view.refresh', () => {
      if (route.name === 'start') load()
    })
    window.runtime.EventsOn('shortcut.view.hard-refresh', () => {
      if (route.name === 'start') {
        store.documents = []
        load()
      }
    })
  })
})

async function load() {
  projects.value = []

  store.documents.forEach(async (file) => {
    const contents = await rpc.FileSystemService.ReadFileContents(file)
    const collection = Collection.createFrom(JSON.parse(contents))
    projects.value = [...projects.value, collection]
  })
}

async function loadLayers() {
  const sourceDirectory = await rpc.app.OpenDirectoryDialog(t('open_layers_folder'))

  if (!sourceDirectory) return

  collectionStore.sourceDirectory = sourceDirectory

  const collection: Collection = await rpc.CollectionService.LoadCollectionFromDirectory(sourceDirectory)

  collectionStore.layers = { ...collection.layers }

  const traits: Object[] = []

  for (const trait in collectionStore.layers) {
    if (Object.hasOwnProperty.call(collectionStore.layers, trait)) {
      traits.push({ name: trait, collapsed: false })
    }
  }

  collectionStore.traits = [...traits]

  nextTick(() => {
    router.push({ name: `artwork.layers` })
  })
}

function loadCollection(collection: Collection) {
  collectionStore.hydrate(collection)
  router.push({ name: 'artwork.layers' })
}
</script>

<template>
  <section class="h-full flex">
    <Sidebar :links="[
      { icon: CollectionIcon, text: t('layer_setup'), to: 'artwork.layers', selected: false },
      { icon: CogIcon, text: t('build_settings'), to: 'artwork.build', selected: false },
      { icon: PlayIcon, text: t('run'), to: 'artwork.run', selected: false },
    ]" />

    <main class="h-full flex-1 overflow-auto scrollbar-none">
      <section class="p-16 grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 xl:grid-cols-8 gap-8">
        <div class="col-span-1">
          <div>
            <img
              @click="loadLayers"
              class="animate__animated animate__fadeIn w-full h-full m-0 p-0 object-cover"
              src="../assets/varly-new-document.png"
              alt=""
            >
          </div>
          <h2 class="mt-2 py-2 px-4 text-center leading-tight" v-text="t('start_new_project')" />
        </div>

        <div v-if="collectionStore.layers && Object.keys(collectionStore.layers).length" class="col-span-1">
          <div>
            <img
              @click="() => router.push({ name: 'artwork.layers' })"
              class="animate__animated animate__fadeIn w-full h-full m-0 p-0 object-cover"
              src="../assets/varly-current-document.png"
              alt=""
            >
          </div>
          <h2 class="mt-2 py-2 px-4 text-center leading-tight" v-text="t('open_current_project')" />
        </div>

        <div
          class="col-span-1" v-for="(collection, i) in projects" :key="i"
          :class="[i === 0 ? 'col-start-1' : '']"
          >
          <div>
            <img
              @click="loadCollection(collection)"
              class="animate__animated animate__fadeIn w-full h-full m-0 p-0 object-cover"
              src="../assets/varly-document.png"
              alt=""
            >
          </div>
          <h2 class="mt-2 py-2 px-4 text-center leading-tight">{{ collection.name }}</h2>
        </div>
      </section>
    </main>
  </section>
</template>
