<script setup lang="ts">
import { onBeforeMount, reactive } from 'vue'
import { useI18n } from 'vue-i18n'
import Base64Image from '@/components/Base64Image.vue'
import Sidebar from '@/components/Sidebar.vue'
import { useCollectionStore } from '@/store'
import { CogIcon, CollectionIcon, PlayIcon } from '@heroicons/vue/solid'
import rpc from '@/rpc'

const store = useCollectionStore()
const documents = reactive<any>({})

const intl = useI18n({ useScope: 'global' })
const { t } = intl

onBeforeMount(async () => {
  const settings = await rpc.SettingsService.GetSettings()
  documents.value = settings.documents
})
</script>

<template>
  <section class="h-full flex">
    <Sidebar
      v-if="store && store.layers && Object.keys(store.layers).length"
      :links="[
      { icon: CollectionIcon, text: t('layer_setup'), to: 'artwork.layers', selected: true },
      { icon: CogIcon, text: t('build_settings'), to: 'artwork.build', selected: false },
      { icon: PlayIcon, text: t('run'), to: 'artwork.run', selected: false },
    ]" />
    <Sidebar v-else />

    <main class="h-full flex-1 overflow-auto scrollbar-none">
      <section class="grid grid-cols-12 gap-8 px-8 py-16 xl:px-16 xl:py-32">
        <div class="col-span-6 md:col-span-4 lg:col-span-3 xl:col-span-2 flex flex-col items-center"
          v-for="(document, i) in documents.value" :key="i">
          <div class="w-32 h-32 rounded-full">
            <Base64Image
              class="animate__animated animate__fadeIn w-full h-full m-0 p-0 object-cover border-4 border-opacity-50 shadow-md"
              :path="document.preview" />
          </div>
          <h2 class="mt-2 lg:mt-4 py-2 px-4">{{ document.title }}</h2>
        </div>
      </section>
    </main>
  </section>
</template>
