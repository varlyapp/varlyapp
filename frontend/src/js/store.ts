import { ref, type Ref } from 'vue'
import { defineStore } from 'pinia'

const useStore = defineStore('app', () => {
  const isDark = ref<boolean>(false)

  return { isDark }
})

const useCollectionStore = defineStore('collection', () => {
  const traits = ref<Object[]>([])
  const layers = ref<Object>({})
  const directory = ref<string>('') // where the layers were loaded from

  function reset(this: any) {
    this.traits = []
    this.layers = {}
    this.directory = ''
  }

  return { reset, traits, layers, directory }
}, { persist: true })

export { useStore, useCollectionStore }
