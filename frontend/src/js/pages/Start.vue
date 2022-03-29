<script setup lang="ts">
import { reactive, inject, onBeforeMount } from 'vue'
import { BadgeCheckIcon, FolderOpenIcon, DocumentAddIcon } from '@heroicons/vue/solid'
import type { Varly } from '@plugins/varly'
import Sidebar from '@components/Sidebar.vue'
import Base64Image from '@components/Base64Image.vue'

const varly = inject<Varly>('varly')!

const documents = reactive<any>([])

onBeforeMount(async () => {
  const settings = await varly.app.GetSettings()
  documents.value = settings.documents
  console.log(documents.value)
})

function launch(url: string) {
  return varly.runtime.BrowserOpenURL('https://twitter.com/varlyapp')
}
</script>

<template>
  <section class="h-full flex">
    <Sidebar
      :links="[
        { icon: BadgeCheckIcon, text: 'Support on Twitter', to: 'start', selected: false },
        { icon: null, text: '', to: '', selected: false },
        { icon: FolderOpenIcon, text: 'Recent Projects', to: 'start', selected: true },
        { icon: DocumentAddIcon, text: 'Start NFT Project', to: 'artwork.layers', selected: false },
      ]"
    />

    <main class="h-full flex-1 overflow-auto scrollbar-none">
      <section class="flex flex-wrap p-8 -mx-8">
        <div
          class="flex flex-col items-center w-4/12 lg:w-2/12 mt-4 px-8 text-center"
          v-for="(document, i) in documents.value"
          :key="i"
        >
          <div class="w-28 h-28 rounded-full">
            <Base64Image style="padding: 2pt" class="w-full h-full object-cover border-4 border-opacity-20 shadow rounded-full" :path="document.preview" />
          </div>
          <h2
            class="mt-2 lg:mt-4 py-2 px-4"
          >{{ document.title }} {{ i === 0 ? 'hello world hello world' : '' }}</h2>
        </div>
      </section>
    </main>
  </section>
</template>
