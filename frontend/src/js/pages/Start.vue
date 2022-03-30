<script setup lang="ts">
import { reactive, onBeforeMount } from 'vue'
import Base64Image from '@components/Base64Image.vue'
import Sidebar from '@components/Sidebar.vue'
import { panic, showMessageDialog, getSettings, launchTwitter } from '@utils/Varly'
import { BadgeCheckIcon, FolderOpenIcon, DocumentAddIcon } from '@heroicons/vue/solid'

const documents = reactive<any>([])

onBeforeMount(async () => {
  const settings = await getSettings()
  documents.value = settings.documents
})
</script>

<template>
  <section class="h-full flex">
    <Sidebar
      :links="[
        { icon: BadgeCheckIcon, text: 'Support on Twitter', to: launchTwitter, selected: false },
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
          <div class="w-32 h-32 rounded-full">
            <Base64Image
              class="w-full h-full m-0 p-0 object-cover border-4 border-opacity-50 shadow-md"
              :path="document.preview"
            />
          </div>
          <h2 class="mt-2 lg:mt-4 py-2 px-4">{{ document.title }}</h2>
        </div>
      </section>
    </main>
  </section>
</template>
