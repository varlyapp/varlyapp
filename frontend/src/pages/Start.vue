<script setup lang="ts">
import { reactive, onBeforeMount } from 'vue'
import Base64Image from '@/components/Base64Image.vue'
import Sidebar from '@/components/Sidebar.vue'
import { getSettings } from '@/utils/Varly'

const documents = reactive<any>([])

onBeforeMount(async () => {
  const settings = await getSettings()
  documents.value = settings.documents
})
</script>

<template>
  <section class="h-full flex">
    <Sidebar />

    <main class="h-full flex-1 overflow-auto scrollbar-none">
      <section class="grid grid-cols-12 gap-8 px-8 py-16 xl:px-16 xl:py-32">
        <div
          class="col-span-6 md:col-span-4 lg:col-span-3 xl:col-span-2 flex flex-col items-center"
          v-for="(document, i) in documents.value"
          :key="i"
        >
          <div class="w-32 h-32 rounded-full">
            <Base64Image
              class="animate__animated animate__fadeIn w-full h-full m-0 p-0 object-cover border-4 border-opacity-50 shadow-md"
              :path="document.preview"
            />
          </div>
          <h2 class="mt-2 lg:mt-4 py-2 px-4">{{ document.title }}</h2>
        </div>
      </section>
    </main>
  </section>
</template>
