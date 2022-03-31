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
  const name = ref<string>('')
  const description = ref<string>('')
  const size = ref<number>(0)
  const width = ref<number>(0)
  const height = ref<number>(0)

  function reset(this: any) {
    this.traits = []
    this.layers = {}
    this.directory = ''
    this.name = ''
    this.size = 0
    this.width = 0
    this.height = 0
  }

  return { reset, traits, layers, directory, name, description, size, width, height }
}, { persist: true })

export { useStore, useCollectionStore }
