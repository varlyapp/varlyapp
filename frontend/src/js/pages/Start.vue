<script setup lang="ts">
import { inject } from 'vue'
import type { Varly } from '@plugins/varly'

import Icon from '@components/Icon.vue'

const varly = inject<Varly>('varly')!
</script>

<template>
  <div class="h-full flex flex-col items-center justify-center">
    <section class="flex items-center justify-center">
      <Icon
        class="mx-8"
        name="new"
        text="Start NFT Project"
        @click="() => varly.router.push({ name: 'collections.start' })"
      />
      <Icon
        class="mx-8"
        name="import"
        text="Import Document"
        @click="() => varly.runtime.BrowserOpenURL('https://twitter.com/varlyapp')"
      />
    </section>
  </div>
</template>
