<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import { reactive, onBeforeMount, onMounted } from 'vue'
import Base64Image from '@components/Base64Image.vue'
import Sidebar from '@components/Sidebar.vue'
import { getSettings, launchTwitter } from '@utils/Varly'
import { BadgeCheckIcon, FolderOpenIcon, DocumentAddIcon } from '@heroicons/vue/solid'

const intl = useI18n({ useScope: 'global' })
const { t } = intl

const documents = reactive<any>([])

onBeforeMount(async () => {
  const settings = await getSettings()
  documents.value = settings.documents
  intl.locale.value = 'es'
})
</script>

<template>
  <section class="h-full flex">
    <Sidebar
      :links="[
        { icon: BadgeCheckIcon, text: t('follow_on_twitter'), to: launchTwitter, selected: false },
        { icon: null, text: '', to: '', selected: false },
        { icon: FolderOpenIcon, text: t('recent_projects'), to: 'start', selected: true },
        { icon: DocumentAddIcon, text: t('start_new_project'), to: 'artwork.layers', selected: false },
      ]"
    />

    <main class="h-full flex-1 overflow-auto scrollbar-none">
      <section class="grid grid-cols-12 gap-8 px-8 py-16 xl:px-16 xl:py-32">
        <div
          class="col-span-6 md:col-span-4 lg:col-span-3 xl:col-span-2 flex flex-col items-center"
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
