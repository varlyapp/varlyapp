import { ref } from 'vue'
import { defineStore } from 'pinia'

interface Action {
  type: string
  title: string
  onClick: Function
}

const useStore = defineStore('app', () => {
  const isDark = ref<boolean>(false)
  const isLoading = ref<boolean>(false)
  const actions = ref<Array<Action>>([])

  return { isDark, isLoading, actions }
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
