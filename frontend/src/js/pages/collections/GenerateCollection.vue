<template>
    <div class="h-full flex flex-col container mx-auto max-w-4xl py-2 px-4 justify-center">
        <button
            class="text-base px-4 py-2 bg-purple-600 text-white"
            @click="generateCollection"
        >Generate</button>
        <Progress :steps="steps" :current-step="currentStep" />
        <!-- <img :src="preview" alt="" class="m-0 p-0 max-w-full max-h-full object-cover border-0"> -->
    </div>
</template>

<script setup lang="ts">
import { ref, inject } from 'vue'
import { useDialog } from '@utils/Dialog'
import { useCollectionStore } from '@root/store'
import type { Varly } from '@root/plugins/varly'
import Progress from '../../components/Progress.vue'

const varly = inject<Varly>('varly')!

const app = window.go.main.App

const store = useCollectionStore()
const dialog = useDialog(app)

const steps = ref(0)
const currentStep = ref(0)
const isLoading = ref(false)
const loadingText = ref('Loading')
const preview = ref('')

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
        // preview.value = await varly.app.EncodeImage(data.ImagePath)
        steps.value = data.CollectionSize
        currentStep.value = data.ItemNumber
        console.log({ steps: steps.value, currentStep: currentStep.value })
        // loadingText.value = `Generating collection: ${data.ItemNumber}/${data.CollectionSize}`
        // console.log(loadingText.value)
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
        Size: 1000
    }

    await app.GenerateCollectionFromConfig(config)

    toggleIsLoading()
}
</script>
