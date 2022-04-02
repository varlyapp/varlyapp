import { ref } from 'vue'
import { defineStore } from 'pinia'

const useStore = defineStore('app', () => {
  const locale = ref<string>('en')
  const isDark = ref<boolean>(false)

  function setLocale(_locale: string) {
    locale.value = _locale
  }

  return { locale, isDark, setLocale }
}, { persist: true })

const useCollectionStore = defineStore('collection', () => {
  const traits      = ref<Object[]>([])
  const layers      = ref<Object>({})
  const directory   = ref<string>('')    // where the layers were loaded from
  const name        = ref<string>('')
  const description = ref<string>('')
  const size        = ref<number>(0)
  const width       = ref<number>(0)
  const height      = ref<number>(0)

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
