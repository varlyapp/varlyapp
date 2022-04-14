<script setup lang="ts">
import { onMounted, ref } from 'vue'
import Confetti from 'vue-confetti-explosion'
import { CogIcon, CollectionIcon, PlayIcon, FolderDownloadIcon } from '@heroicons/vue/solid'
import Progress from '@/components/Progress.vue'
import Sidebar from '@/components/Sidebar.vue'
import { useCollectionStore } from '@/store'
import Preview from '@/components/Preview.vue'
import rpc from '@/rpc'
import { Collection } from '@/wailsjs/go/models'
import FloatingButtonBar from '@/components/FloatingButtonBar.vue'

const store = useCollectionStore()

const steps = ref(0)
const currentStep = ref(0)
const isWorking = ref(false)
const isDone = ref(false)
const loadingText = ref('Loading')
const preview = ref('')

onMounted(() => {
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

    const collection = Collection.createFrom({
        sourceDirectory: store.sourceDirectory,
        outputDirectory: store.outputDirectory,
        traits: [...store.traits],
        layers: layers,
        width: parseInt(store.width.toString(), 10),
        height: parseInt(store.height.toString(), 10),
        size: parseInt(store.size.toString(), 10)
    })

    rpc.CollectionService.GenerateCollectionPreview(collection)
        .then((base64Image) => {
            preview.value = base64Image
        })
        .catch(console.error)
})

function queueConfetti() {
    isDone.value = true
    window.runtime.LogInfo('Setting isDone to true')
    setTimeout(() => {
        isDone.value = false
        window.runtime.LogInfo('Setting isDone to false')
    }, 5000)
}

function toggleIsWorking() {
    isWorking.value = !isWorking.value
}

async function selectOutputDirectory() {
    const outputDirectory = await rpc.app.OpenDirectoryDialog('Select a folder in which to save generated images')

    if (!outputDirectory || outputDirectory === '') {
        alert(`Folder could not be selected, please try again`)
    }

    store.outputDirectory = outputDirectory
}
async function generateCollection() {
    currentStep.value = 0 // reset each time this method is called

    window.runtime.EventsOn('collection.generation.started', (data) => {
        console.log({ msg: `collection generation started`, data })
        loadingText.value = `Preparing collection of ${data.CollectionSize} items`
    })

    window.runtime.EventsOn('collection.item.generated', async (data) => {
        steps.value = data.CollectionSize
        currentStep.value = data.ItemNumber
        console.log({ msg: `collection item generated`, data })
    })
    window.runtime.EventsOn('debug', async (data) => {
        console.log(data)
    })

    const outputDirectory = store.outputDirectory

    if (!outputDirectory) return

    toggleIsWorking()

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

    const collection = Collection.createFrom({
        name: store.name,
        description: store.description,
        sourceDirectory: store.sourceDirectory,
        outputDirectory: outputDirectory,
        traits: [...store.traits],
        layers: layers,
        width: parseInt(store.width.toString(), 10),
        height: parseInt(store.height.toString(), 10),
        size: parseInt(store.size.toString(), 10)
    })

    // Save to file, if desired
    // const data = JSON.stringify(collection)
    // const saved = await varly.app.SaveFile('collection1.json', data)
    // log(saved.toString())

    await rpc.CollectionService.GenerateCollection(collection)

    toggleIsWorking()
    queueConfetti()
}
</script>

<template>
    <section class="h-full flex">
        <Sidebar :links="[
            { icon: CollectionIcon, text: 'Layer Setup', to: 'artwork.layers', selected: false },
            { icon: CogIcon, text: 'Build Settings', to: 'artwork.build', selected: false },
            { icon: PlayIcon, text: 'Run', to: 'artwork.run', selected: true },
        ]" />

        <main class="h-full flex-1 overflow-y-scroll scrollbar-none">
            <div v-if="isWorking" class="h-full flex flex-col flex-1 max-w-4xl mx-auto items-center justify-center p-8">
                <Progress :steps="steps" :current-step="currentStep" loading-text="Preparing..." />
            </div>
            <div v-else class="h-full flex flex-col items-center justify-center p-8">
                <div class="flex flex-col items-center">
                    <div v-if="preview && !isDone" class="p-16">
                        <p v-if="store.outputDirectory" class="pt-4 text-xs text-center">
                            <strong>＊ Output Folder: </strong>
                            {{ store.outputDirectory }}
                        </p>
                        <Preview :source="preview" caption="" />
                    </div>
                    <div class="max-w-xs mx-auto">
                        <h1 v-if="isDone" class="animate__animated animate__fadeIn text-6xl text-center font-bold">Yay
                            🎉</h1>
                    </div>
                </div>
            </div>
        </main>

        <Confetti
            v-if="isDone"
            :particle-count="200"
            :particle-size="10"
            :duration="5000"
            class="absolute w-screen h-screen top-0 right-0 bottom-0 left-0"
        />

        <FloatingButtonBar>
            <button
                type="button"
                class="select-none flex mt-2 py-2 px-6 items-center rounded text-slate-50 bg-slate-700 shadow-md shadow-slate-800 hover:bg-opacity-90 font-bold"
                @click="selectOutputDirectory"
            >
                <span>
                    <FolderDownloadIcon class="w-6 mr-2 fill-slate-100"/>
                </span>
                <span>Select Output Folder</span>
            </button>

            <button v-if="store.outputDirectory"
                type="button"
                class="select-none flex mt-2 py-2 px-6 items-center rounded text-slate-50 bg-fuchsia-700 shadow-md shadow-fuchsia-900 hover:bg-opacity-90 font-bold"
                @click="generateCollection"
            >
                <span>
                    <PlayIcon class="w-6 mr-2 fill-fuchsia-100"/>
                </span>
                <span>Generate Collection</span>
            </button>
        </FloatingButtonBar>
    </section>
</template>
