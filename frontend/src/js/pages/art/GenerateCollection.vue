<template>
    <div class="h-full flex flex-col container mx-auto max-w-4xl py-2 px-4 justify-center">
        <section
            class="h-full flex items-center justify-center text-center font-mono"
            v-if="isLoading && loadingText"
        >
            <h1 class="text-xl">{{ loadingText }} 🚀</h1>
        </section>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useDialog } from '@utils/Dialog'
import { useCollectionStore } from '@root/store'

const app = window.go.main.App

const router = useRouter()
const store = useCollectionStore()
const dialog = useDialog(app)

let isLoading = ref(false)
let loadingText = ref('')

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

    window.runtime.EventsOn('collection.item.generated', (data) => {
        console.log(data)
        loadingText.value = `Generating collection: ${data.ItemNumber}/${data.CollectionSize}`
        console.log(loadingText.value)
    })

    // const outputDirectory = await dialog.openDirectoryDialog()

    // console.log(outputDirectory)

    // const layers = { ...store.layers }

    // for (const trait in Object.keys(layers)) {
    //     if (layers.hasOwnProperty(trait)) {
    //         layers[trait] = layers[trait].map((layer) => {
    //             return {
    //                 ...layer,
    //                 Weight: parseInt(layer.Weight)
    //             }
    //         })
    //     }
    // }

    // const config = {
    //     Dir: outputDirectory,
    //     Order: [...store.traits].map((item: any) => item.name),
    //     Layers: layers,
    //     Width: 1500,
    //     height: 1500,
    //     Size: 1000
    // }

    await app.GenerateCollectionFromConfig(config)

    toggleIsLoading()
}
</script>
