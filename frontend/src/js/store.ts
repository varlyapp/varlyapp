import { ref } from 'vue'
import { defineStore } from 'pinia'

const useStore = defineStore('app', () => {
  const isDark = ref(false)

  return { isDark }
})

const useCollectionStore = defineStore('collection', () => {
  const traits = ref([])
  const layers = ref({})
  const directory = ref('') // where the layers were loaded from

  function reset(this: { traits: Array<Object>, layers: Object, directory: string }) {
    this.traits = []
    this.layers = {}
    this.directory = ''
  }

  return { reset, traits, layers, directory }
}, { persist: true })

export { useStore, useCollectionStore }
