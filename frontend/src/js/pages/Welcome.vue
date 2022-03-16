<template>
  <div class="h-full flex flex-col items-center justify-center">
    <section class="flex items-center justify-center">
      <!-- <div class="m-4">
        <router-link
          class="p-8 bg-purple-700 text-white rounded shadow-xl hover:bg-opacity-90 font-bold"
          :to="{ name: 'art.collection.new' }"
        >Start New Project</router-link>
      </div>
      <div class="m-4">
        <button
          class="p-8 bg-blue-500 text-white rounded shadow-xl hover:bg-opacity-90 font-bold"
          @click="open('https://twitter.com/varlyapp')"
        >Tweet @ Varly.app</button>
      </div>-->
      <Icon
        class="mx-8"
        name="new"
        text="Start Project"
        @click="() => router.push({ name: 'art.collection.new' })"
      />
      <Icon
        class="mx-8"
        name="import"
        text="Import Document"
        @click="() => open('https://twitter.com/varlyapp')"
      />
    </section>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import Icon from '@components/Icon.vue'

const router = useRouter()

function open(url) {
  window.runtime.BrowserOpenURL(url)
}
</script>
