<script setup lang="ts">
import { inject } from 'vue'
import type { Varly } from '@plugins/varly'
import Sidebar from '../components/Sidebar.vue'

const varly = inject<Varly>('varly')!

const documents = [
  { id: 1, name: 'Doodles' },
  { id: 2, name: 'VeeFriends' },
  { id: 3, name: 'InvisibleFriends' },
]

function launch(url: string) {
  return varly.runtime.BrowserOpenURL('https://twitter.com/varlyapp')
}
</script>

<template>
  <section class="h-full flex">
    <Sidebar
      :links="[
        { emoji: '🦋', text: 'Support on Twitter', to: 'start', selected: false },
        { emoji: '', text: '', to: '', selected: false },
        { emoji: '🏙', text: 'Recent Projects', to: 'start', selected: true },
        { emoji: '✨', text: 'Start NFT Project', to: 'artwork.layers', selected: false },
      ]"
    />

    <main class="h-full flex-1 overflow-auto scrollbar-none">
      <section class="flex flex-wrap mt-4 -mx-8">
        <div
          class="flex flex-col items-center w-4/12 lg:w-3/12 mt-4 px-8"
          v-for="(document) in documents"
          :key="document.id"
        >
          <div class="w-16 bg-purple-700" style="min-height: 64pt"></div>
          <h2 class="mt-4 py-2 px-4">{{ document.name }}</h2>
        </div>
      </section>
    </main>
  </section>
</template>
