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
  const sourceDirectory = ref<string>('')
  const outputDirectory = ref<string>('')
  const name = ref<string>('')
  const description = ref<string>('')
  const traits = ref<Object[]>([])
  const layers = ref<Object>({})
  const size = ref<number>(0.0)
  const width = ref<number>(0.0)
  const height = ref<number>(0.0)

  function hydrate(this: any, collection: any) {
    this.reset()
    this.sourceDirectory = collection.sourceDirectory || ''
    this.outputDirectory = collection.outputDirectory || ''
    this.name = collection.name || ''
    this.descrption = collection.descrption || ''
    this.traits = collection.traits || []
    this.layers = collection.layers || {}
    this.width = collection.width || 0.0
    this.height = collection.height || 0.0
    this.size = collection.size || 0
  }

  function prepare(this: any): Object {
    return {
      sourceDirectory: this.sourceDirectory,
      outputDirectory: this.outputDirectory,
      name: this.name,
      description: this.description,
      traits: this.traits,
      layers: this.layers,
      width: parseFloat(this.width),
      height: parseFloat(this.height),
      size: parseInt(this.size, 10),
    }
  }
  function reset(this: any) {
    this.sourceDirectory = ''
    this.outputDirectory = ''
    this.name = ''
    this.description = ''
    this.traits = []
    this.layers = {}
    this.width = 0
    this.height = 0
    this.size = 0
  }

  return { hydrate, prepare, reset, traits, layers, sourceDirectory, outputDirectory, name, description, width, height, size }
}, { persist: true })

export { useStore, useCollectionStore }
