<script setup lang="ts">
import { useRouter } from 'vue-router'
import Base64Image from '@root/components/Base64Image.vue'

import Icon from '@components/Icon.vue'

const router = useRouter()

function tapped() {
  window.console.log('Hello')
}

function open(url) {
  window.runtime.BrowserOpenURL(url)
}
</script>

<template>
  <div class="h-full flex flex-col items-center justify-center">
    <section class="flex items-center justify-center">
      <Icon
        class="mx-8"
        name="new"
        text="Start NFT Project"
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
