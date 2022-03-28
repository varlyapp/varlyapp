<script setup lang="ts">
import { inject } from 'vue'
import type { Varly } from '@plugins/varly'

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
  <div class="h-full flex">
    <aside
      class="sticky top-0 h-full p-4 overflow-hidden border-r border-slate-900 border-opacity-10 dark:border-slate-50 dark:border-opacity-10"
    >
      <nav
        class="mt-6 text-left text-base text-slate-900 text-opacity-90 dark:text-slate-50 dark:text-opacity-90"
      >
        <ul>
          <li>
            <router-link
              class="flex mt-2 py-2 px-4 bg-slate-900 bg-opacity-10 dark:bg-slate-50 dark:bg-opacity-10 rounded"
              :to="{ name: 'start' }"
            >
              <span class="text-center" style="width: 24pt">🏙</span>
              <span>Recent Documents</span>
            </router-link>
          </li>
          <li>
            <router-link class="flex mt-2 py-2 px-4" :to="{ name: 'artwork.layers' }">
              <span class="text-center" style="width: 24pt">📝</span>
              <span>Start NFT Project</span>
            </router-link>
          </li>
          <li>
            <button
              type="button"
              class="flex items-center justify-start mt-2 py-2 px-4"
              @click="launch('https://twitter.com/varlyapp')"
            >
              <span class="text-center" style="width: 24pt">🦋</span>
              <span>Connect on Twitter</span>
            </button>
          </li>
        </ul>
      </nav>
    </aside>

    <main class="h-full overflow-auto scrollbar-none">
      <div class="p-8">
        <section class="flex flex-wrap -mx-8">
          <div
            class="flex flex-col items-center w-4/12 lg:w-3/12 mt-4 px-8"
            v-for="(document, key) in documents"
            :key="document.id"
          >
            <div class="w-16 bg-purple-700" style="min-height: 64pt"></div>
            <h2 class="mt-4 py-2 px-4">{{ document.name }}</h2>
          </div>
        </section>
      </div>
    </main>
  </div>
</template>
