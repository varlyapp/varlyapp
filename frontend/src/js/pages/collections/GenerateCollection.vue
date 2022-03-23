<template>
    <div class="h-full flex flex-col container mx-auto max-w-4xl py-2 px-4 justify-center">
        <button class="text-base" @click="generateCollection">Run</button>
        <img :src="preview" alt="" class="w-full" height="h-auto">
    </div>
</template>

<script setup lang="ts">
import { ref, inject } from 'vue'
import { useRouter } from 'vue-router'
import { useDialog } from '@utils/Dialog'
import { useCollectionStore } from '@root/store'
import type { Varly } from '@root/plugins/varly'

const varly = inject<Varly>('varly')!

const app = window.go.main.App

const router = useRouter()
const store = useCollectionStore()
const dialog = useDialog(app)

let isLoading = ref(false)
let loadingText = ref('Loading')
let preview = ref('')

const isCollapsed = ref(false)
const isTraitEnabled = ref(true)
const isTraitDragging = ref(false)
const isLayerDragging = ref(false)
const isLayerEnabled = ref(true)

function toggleIsLoading() {
    isLoading.value = !isLoading.value
}

async function generateCollection() {
    toggleIsLoading()

    window.runtime.EventsOn('collection.generation.started', (data) => {
        loadingText.value = `Preparing collection of ${data.CollectionSize} items`
        console.log(loadingText.value)
    })

    window.runtime.EventsOn('collection.item.generated', async (data) => {
        preview.value = await varly.app.EncodeImage(data.ImagePath)
        loadingText.value = `Generating collection: ${data.ItemNumber}/${data.CollectionSize}`
        console.log(loadingText.value)
    })

    const outputDirectory = await dialog.openDirectoryDialog()

    console.log(outputDirectory)

    const layers = { ...store.layers }

    for (const trait in Object.keys(layers)) {
        if (layers.hasOwnProperty(trait)) {
            layers[trait] = layers[trait].map((layer) => {
                return {
                    ...layer,
                    Weight: parseInt(layer.Weight)
                }
            })
        }
    }

    const config = {
        Dir: outputDirectory,
        Order: [...store.traits].map((item: any) => item.name),
        Layers: layers,
        Width: 3000,
        height: 3000,
        Size: 100
    }

    await app.GenerateCollectionFromConfig(config)

    toggleIsLoading()
}
</script>
