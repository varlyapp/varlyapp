<script setup lang="ts">
import { inject } from 'vue'
import type { Varly } from '@plugins/varly'

import Icon from '@components/Icon.vue'

const varly = inject<Varly>('varly')!

function launch(url: string) {
  return varly.runtime.BrowserOpenURL('https://twitter.com/varlyapp')
}

</script>

<template>
  <div class="h-full flex flex-col items-center justify-center p-8">
    <section class="flex items-center justify-center">
      <Icon
        class="mx-8"
        name="new"
        text="Start NFT Project"
        @click="() => varly.router.push({ name: 'artwork.layers' })"
      />
      <Icon
        class="mx-8"
        name="import"
        text="Import Document"
        @click="launch('https://twitter.com/varlyapp')"
      />
    </section>
  </div>
</template>
