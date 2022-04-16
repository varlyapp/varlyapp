<script setup lang="ts">
import { nextTick, onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { Collection } from '@/wailsjs/go/models'
import { CogIcon, CollectionIcon, PlayIcon } from '@heroicons/vue/solid'
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

onMounted(() => {
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
      <section v-if="projects.length" class="p-16 grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 xl:grid-cols-8 gap-8">
        <div class="col-span-1" v-for="(collection, i) in projects" :key="i">
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

      <section v-else class="h-full animate__animated animate__fadeIn">
        <div class="h-full p-4 lg:p-8 flex items-center justify-center text-center">
          <div>
            <svg class="mx-auto h-12 w-12" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
              <path vector-effect="non-scaling-stroke" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M9 13h6m-3-3v6m-9 1V7a2 2 0 012-2h6l2 2h6a2 2 0 012 2v8a2 2 0 01-2 2H5a2 2 0 01-2-2z" />
            </svg>
            <h3 class="mt-2 text-sm font-medium" v-text="t('get_started_by_opening_your_layers_folder')"></h3>
            <div class="mt-8">
              <button type="button"
                class="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-fuchsia-700 hover:bg-fuchsia-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-fuchsia-500"
                @click="loadLayers">
                <CollectionIcon class="-ml-1 mr-2 h-5 w-5" aria-hidden="true" />
                <span v-text="t('open_layers_folder')" />
              </button>
            </div>
          </div>
        </div>
      </section>
    </main>
  </section>
</template>
