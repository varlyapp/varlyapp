<script setup lang="ts">
import { nextTick, onMounted, ref, reactive, onActivated } from 'vue'
import { useI18n } from 'vue-i18n'
import { useRoute } from 'vue-router'
import Confetti from 'vue-confetti-explosion'
import { CogIcon, CollectionIcon, PlayIcon } from '@heroicons/vue/outline'
import { PlayIcon as SolidPlayIcon, RefreshIcon as SolidRefreshIcon, FolderOpenIcon as SolidFolderOpenIcon } from '@heroicons/vue/solid'
import Progress from '@/components/Progress.vue'
import Sidebar from '@/components/Sidebar.vue'
import { types } from '@/wailsjs/go/models'
import { LogInfo } from '@/wailsjs/runtime/runtime'
import FloatingButtonBar from '@/components/FloatingButtonBar.vue'
import Preview from '@/components/Preview.vue'
import StatusBar from '@/components/StatusBar.vue'
import { useStore, useCollectionStore } from '@/store'
import rpc from '@/rpc'

const { t } = useI18n({ useScope: 'global' })

const route = useRoute()
const store = useStore()
const collectionStore = useCollectionStore()

const steps = ref(0)
const currentStep = ref(0)
const currentImage = ref('')
const isWorking = ref(false)
const isDone = ref(false)

onMounted(() => {
    store.setIsGeneratingCollection(false)

    load()
    nextTick(() => {
        rpc.on('shortcut.view.refresh', () => {
            if (route.name === 'artwork.run') load()
        })
        rpc.on('shortcut.view.hard-refresh', () => {
            store.setIsGeneratingCollection(false)
            if (route.name === 'artwork.run') load()
        })
    })
})

onActivated(() => {
    rpc.setPageTitle("Preview & Generate")
})

function canGenerateCollection(): boolean {
    if (isWorking.value === true) return false

    return (collectionStore.layers !== undefined
    && Object.keys(collectionStore.layers).length > 0
    && collectionStore.outputDirectory !== undefined
    && collectionStore.outputDirectory.length > 0)
}

async function load() {
    rpc.setPageTitle("Preview & Generate")

    if (isWorking.value || store.isGeneratingCollection) return

    console.log('Reloading Run.vue')
    const layers = { ...collectionStore.layers }

    if (!layers || !Object.keys(layers).length) {
        console.log('No preview to generate')
        return
    }

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

    const collection = types.Collection.createFrom({
        sourceDirectory: collectionStore.sourceDirectory,
        outputDirectory: collectionStore.outputDirectory,
        traits: [...collectionStore.traits],
        layers: layers,
        width: parseFloat(collectionStore.width.toString()),
        height: parseFloat(collectionStore.height.toString()),
        size: parseInt(collectionStore.size.toString(), 10)
    })

    collectionStore.preview = ''

    try {
        collectionStore.preview = await rpc.CollectionService.GenerateCollectionPreview(collection)
    } catch (error) {
        console.error(error)
    }
}

async function generateCollection() {
    if (isWorking.value || store.isGeneratingCollection) return

    store.setIsGeneratingCollection(true)
    currentStep.value = 0 // reset each time this method is called

    rpc.on('collection.generation.started', (data) => {
        console.log({ msg: `collection generation started`, data })
    })

    rpc.on('collection.item.generated', async (data) => {
        steps.value = data.CollectionSize
        currentStep.value = data.ItemNumber
        currentImage.value = data.ItemPath
    })
    rpc.on('debug', async (data) => {
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

    const collection = types.Collection.createFrom({
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

function queueConfetti() {
    isDone.value = true
    LogInfo('Setting isDone to true')
    setTimeout(() => {
        isDone.value = false
        LogInfo('Setting isDone to false')
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
</script>

<template>
    <section class="h-full flex">
        <Sidebar :links="[
            { icon: CollectionIcon, text: t('layer_setup'), to: 'artwork.layers', selected: false },
            { icon: CogIcon, text: t('build_settings'), to: 'artwork.build', selected: false },
            { icon: PlayIcon, text: t('run'), to: 'artwork.run', selected: true },
        ]" />

        <article class="h-full flex-1">
            <main class="h-full flex flex-col">
                <div
                    class="max-h-12 bg-opacity-10 px-4 border-b border-slate-900 border-opacity-10 dark:border-slate-50 dark:border-opacity-10">
                    <nav class="h-full flex items-center justify-end">
                        <ul class="flex items-center justify-center my-2">
                            <li>
                                <button type="button"
                                    class="flex items-center justify-center px-3 py-1 bg-slate-100 bg-opacity-50 rounded shadow"
                                    @click="load">
                                    <SolidRefreshIcon class="w-6 fill-fuchsia-700" />
                                </button>
                            </li>
                            <li class="px-1" />
                            <li>
                                <button type="button"
                                    class="flex items-center justify-center px-3 py-1 bg-slate-100 bg-opacity-50 rounded shadow"
                                    @click="selectOutputDirectory">
                                    <SolidFolderOpenIcon class="w-6 fill-fuchsia-700" />
                                </button>
                            </li>
                            <li class="px-1" />
                            <li>
                                <button
                                    type="button"
                                    class="flex items-center justify-center px-3 py-1 bg-slate-100 bg-opacity-50 rounded shadow"
                                    :class="[!canGenerateCollection() ? 'opacity-50' : '']" :disabled="!canGenerateCollection()"
                                    @click="generateCollection">
                                    <SolidPlayIcon class="w-6 fill-fuchsia-700" />
                                </button>
                            </li>
                        </ul>
                    </nav>
                </div>

                <div class="flex items-center justify-center flex-1 flex-shrink-0 overflow-y-scroll scrollbar-none">
                    <div v-if="isWorking"
                        class="flex flex-col flex-1 max-w-4xl mx-auto items-center justify-center px-8 lg:px-16 xl:px-24">
                        <img v-if="currentImage !== ''" :src="currentImage" alt="">
                        <Progress :steps="steps" :current-step="currentStep" loading-text="Preparing..." />
                    </div>
                    <div v-else class="flex flex-col items-center justify-center px-8 lg:px-16 xl:px-24">
                        <div class="flex flex-col items-center">
                            <div v-if="!isWorking && !isDone">
                                <div v-if="collectionStore.preview" class="">
                                    <Preview :source="collectionStore.preview" caption="" />
                                </div>
                                <div v-else class="p-16 text-center">
                                    <svg v-if="collectionStore.layers && Object.keys(collectionStore.layers).length"
                                        class="m-0 p-0 mx-auto" width="38" height="38" viewBox="0 0 38 38"
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
                                                <path d="M36 18c0-9.94-8.06-18-18-18" id="Oval-2" stroke="url(#a)"
                                                    stroke-width="2">
                                                    <animateTransform attributeName="transform" type="rotate"
                                                        from="0 18 18" to="360 18 18" dur="0.9s"
                                                        repeatCount="indefinite" />
                                                </path>
                                                <circle fill="#a21caf" cx="36" cy="18" r="1">
                                                    <animateTransform attributeName="transform" type="rotate"
                                                        from="0 18 18" to="360 18 18" dur="0.9s"
                                                        repeatCount="indefinite" />
                                                </circle>
                                            </g>
                                        </g>
                                    </svg>
                                    <p v-else v-text="t('no_preview_yet')" />
                                </div>
                            </div>
                            <div v-if="isDone" class="max-w-xs mx-auto">
                                <h1 class="animate__animated animate__fadeIn text-6xl text-center font-bold">🎉</h1>
                            </div>
                        </div>
                    </div>
                </div>

                <div
                    class="h-10 bg-opacity-10 px-4 border-t border-slate-900 border-opacity-10 dark:border-slate-50 dark:border-opacity-10">
                    <div class="max-w-full overflow-h-scroll">
                        <p v-if="collectionStore.outputDirectory" class="pt-4 text-xs">
                            <strong>Output Folder</strong> {{ collectionStore.outputDirectory }}
                        </p>
                    </div>
                </div>
            </main>
        </article>

        <Confetti v-if="isDone" @done="() => isDone = false" :particle-count="200" :particle-size="10" :duration="4000"
            class="absolute w-full h-full top-0 right-0 bottom-0 left-0" />

        <!-- <FloatingButtonBar>
            <button type="button"
                class="flex mt-2 py-3 px-4 items-center rounded text-slate-50 bg-slate-700 shadow shadow-slate-800 hover:bg-opacity-90"
                @click="selectOutputDirectory">
                <span v-text="t('select_output_folder')" />
            </button>

            <button v-if="collectionStore.layers && Object.keys(collectionStore.layers).length && collectionStore.outputDirectory" type="button"
                class="flex mt-2 py-3 px-4 items-center rounded text-slate-50 bg-fuchsia-700 shadow shadow-fuchsia-900 hover:bg-opacity-90"
                :class="[isWorking ? 'opacity-50' : '']"
                :disabled="isWorking"
                @click="generateCollection">
                <span v-text="t('generate_collection')"/>
            </button>
        </FloatingButtonBar> -->

        <!-- <StatusBar v-if="collectionStore.outputDirectory">
            <p v-if="collectionStore.outputDirectory" class="pt-4 text-xs">
                🏁 {{ collectionStore.outputDirectory }}
            </p>
        </StatusBar> -->
    </section>
</template>
