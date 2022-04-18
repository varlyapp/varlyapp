<script setup lang="ts">
import { nextTick, onBeforeUnmount, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import Confetti from 'vue-confetti-explosion'
import { CogIcon, CollectionIcon, PlayIcon, FolderDownloadIcon } from '@heroicons/vue/solid'
import Progress from '@/components/Progress.vue'
import Sidebar from '@/components/Sidebar.vue'
import { useStore, useCollectionStore } from '@/store'
import { Collection } from '@/wailsjs/go/models'
import FloatingButtonBar from '@/components/FloatingButtonBar.vue'
import Preview from '@/components/Preview.vue'
import StatusBar from '@/components/StatusBar.vue'
import { useRoute } from 'vue-router'
import rpc from '@/rpc'

const { t } = useI18n()
const route = useRoute()
const store = useStore()
const collectionStore = useCollectionStore()

const steps = ref(0)
const currentStep = ref(0)
const isWorking = ref(false)
const isDone = ref(false)
const preview = ref('')

onMounted(() => {
    store.setIsGeneratingCollection(false)

    load()
    nextTick(() => {
        window.runtime.EventsOn('shortcut.view.refresh', () => {
            if (route.name === 'artwork.run') load()
        })
        window.runtime.EventsOn('shortcut.view.hard-refresh', () => {
            store.setIsGeneratingCollection(false)
            if (route.name === 'artwork.run') load()
        })
    })
})

async function load() {
    if (isWorking.value || store.isGeneratingCollection) return

    console.log('Reloading Run.vue')
    preview.value = ''
    const layers = { ...collectionStore.layers }

    for (const trait in layers) {
        if (layers.hasOwnProperty(trait)) {
            layers[trait] = layers[trait].map((layer) => {
                return {
                    ...layer,
                    weight: parseFloat(layer.weight)
                }
            })
        }
    }

    const collection = Collection.createFrom({
        sourceDirectory: collectionStore.sourceDirectory,
        outputDirectory: collectionStore.outputDirectory,
        traits: [...collectionStore.traits],
        layers: layers,
        width: parseFloat(collectionStore.width.toString()),
        height: parseFloat(collectionStore.height.toString()),
        size: parseInt(collectionStore.size.toString(), 10)
    })

    try {
        preview.value = await rpc.CollectionService.GenerateCollectionPreview(collection)
    } catch (error) {
        console.error(error)
    }
}

function queueConfetti() {
    isDone.value = true
    window.runtime.LogInfo('Setting isDone to true')
    setTimeout(() => {
        isDone.value = false
        window.runtime.LogInfo('Setting isDone to false')
    }, 3000)
}

function toggleIsWorking() {
    isWorking.value = !isWorking.value
}

async function selectOutputDirectory() {
    const outputDirectory = await rpc.app.OpenDirectoryDialog('Select a folder in which to save generated images')

    if (!outputDirectory || outputDirectory === '') {
        alert(`Folder could not be selected, please try again`)
    }

    collectionStore.outputDirectory = outputDirectory
}

async function generateCollection() {
    if (isWorking.value || store.isGeneratingCollection) return

    store.setIsGeneratingCollection(true)
    currentStep.value = 0 // reset each time this method is called

    window.runtime.EventsOn('collection.generation.started', (data) => {
        console.log({ msg: `collection generation started`, data })
    })

    window.runtime.EventsOn('collection.item.generated', async (data) => {
        steps.value = data.CollectionSize
        currentStep.value = data.ItemNumber
    })
    window.runtime.EventsOn('debug', async (data) => {
        console.log(data)
    })

    const outputDirectory = collectionStore.outputDirectory

    if (!outputDirectory) return

    toggleIsWorking()

    const layers = { ...collectionStore.layers }

    for (const trait in layers) {
        if (layers.hasOwnProperty(trait)) {
            layers[trait] = layers[trait].map((layer) => {
                return {
                    ...layer,
                    weight: parseFloat(layer.weight)
                }
            })
        }
    }

    const collection = Collection.createFrom({
        name: collectionStore.name,
        description: collectionStore.description,
        sourceDirectory: collectionStore.sourceDirectory,
        outputDirectory: outputDirectory,
        traits: [...collectionStore.traits],
        layers: layers,
        width: parseInt(collectionStore.width.toString(), 10),
        height: parseInt(collectionStore.height.toString(), 10),
        size: parseInt(collectionStore.size.toString(), 10)
    })

    // Save to file, if desired
    // const data = JSON.stringify(collection)
    // const saved = await varly.app.SaveFile('collection1.json', data)
    // log(saved.toString())

    await rpc.CollectionService.GenerateCollection(collection)

    toggleIsWorking()
    queueConfetti()
    store.setIsGeneratingCollection(false)
}
</script>

<template>
    <section class="h-full flex">
        <Sidebar :links="[
            { icon: CollectionIcon, text: t('layer_setup'), to: 'artwork.layers', selected: false },
            { icon: CogIcon, text: t('build_settings'), to: 'artwork.build', selected: false },
            { icon: PlayIcon, text: t('run'), to: 'artwork.run', selected: true },
        ]" />

        <main class="h-full flex-1 overflow-y-scroll scrollbar-none">
            <div v-if="isWorking" class="h-full flex flex-col flex-1 max-w-4xl mx-auto items-center justify-center p-8">
                <Progress :steps="steps" :current-step="currentStep" loading-text="Preparing..." />
            </div>
            <div v-else class="h-full flex flex-col items-center justify-center p-8">
                <div class="flex flex-col items-center">
                    <div v-if="!isWorking && !isDone">
                        <div v-if="preview" class="p-16">
                            <Preview :source="preview" caption="" />
                        </div>
                        <div v-else class="p-16 text-center">
                            <p class="py-4">Generating Preview</p>
                            <svg class="m-0 p-0 mx-auto" width="38" height="38" viewBox="0 0 38 38"
                                xmlns="http://www.w3.org/2000/svg">
                                <defs>
                                    <linearGradient x1="8.042%" y1="0%" x2="65.682%" y2="23.865%" id="a">
                                        <stop stop-color="#a21caf" stop-opacity="0" offset="0%" />
                                        <stop stop-color="#a21caf" stop-opacity=".631" offset="63.146%" />
                                        <stop stop-color="#a21caf" offset="100%" />
                                    </linearGradient>
                                </defs>
                                <g fill="none" fill-rule="evenodd">
                                    <g transform="translate(1 1)">
                                        <path d="M36 18c0-9.94-8.06-18-18-18" id="Oval-2" stroke="url(#a)" stroke-width="2">
                                            <animateTransform attributeName="transform" type="rotate" from="0 18 18"
                                                to="360 18 18" dur="0.9s" repeatCount="indefinite" />
                                        </path>
                                        <circle fill="#a21caf" cx="36" cy="18" r="1">
                                            <animateTransform attributeName="transform" type="rotate" from="0 18 18"
                                                to="360 18 18" dur="0.9s" repeatCount="indefinite" />
                                        </circle>
                                    </g>
                                </g>
                            </svg>
                        </div>
                    </div>
                    <div v-if="isDone" class="max-w-xs mx-auto">
                        <h1 class="animate__animated animate__fadeIn text-6xl text-center font-bold">Yay 🎉</h1>
                    </div>
                </div>
            </div>
        </main>

        <Confetti v-if="isDone" @done="() => isDone = false" :particle-count="200" :particle-size="10" :duration="4000" class="absolute w-full h-full top-0 right-0 bottom-0 left-0" />

        <FloatingButtonBar>
            <button type="button"
                class="flex mt-2 py-2 px-4 items-center rounded text-slate-50 bg-slate-700 shadow-md shadow-slate-800 hover:bg-opacity-90"
                @click="selectOutputDirectory">
                <span>
                    <FolderDownloadIcon class="w-6 mr-2 fill-slate-100" />
                </span>
                <span>Select Output Folder</span>
            </button>

            <button v-if="collectionStore.outputDirectory" type="button"
                class="flex mt-2 py-2 px-4 items-center rounded text-slate-50 bg-fuchsia-700 shadow-md shadow-fuchsia-900 hover:bg-opacity-90"
                :class="[isWorking ? 'opacity-50' : '']"
                :disabled="isWorking"
                @click="generateCollection">
                <span>
                    <PlayIcon class="w-6 mr-2 fill-fuchsia-100" />
                </span>
                <span>Generate Collection</span>
            </button>
        </FloatingButtonBar>

        <StatusBar v-if="collectionStore.outputDirectory">
            <p v-if="collectionStore.outputDirectory" class="pt-4 text-xs">
                <strong>🏁 Output Folder: </strong>
                {{ collectionStore.outputDirectory }}
            </p>
        </StatusBar>
    </section>
</template>
