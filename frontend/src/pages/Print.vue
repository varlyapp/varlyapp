<script setup lang="ts">
import { nextTick, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useRoute } from 'vue-router'
import { Collection } from '@/wailsjs/go/models'
import { useCollectionStore, useStore } from '@/store'
import Preview from '@/components/Preview.vue'
import FloatingButtonBar from '@/components/FloatingButtonBar.vue'
import rpc from '@/rpc'

const { t } = useI18n()
const route = useRoute()
const store = useStore()
const collectionStore = useCollectionStore()

const gif = ref<string>('')
const isWorking = ref<boolean>(false)
onMounted(() => {
    store.setIsGeneratingCollection(false)

    load()
    nextTick(() => {
        window.runtime.EventsOn('shortcut.view.refresh', () => {
            if (route.name === 'print') load()
        })
        window.runtime.EventsOn('shortcut.view.hard-refresh', () => {
            store.setIsGeneratingCollection(false)
            if (route.name === 'print') load()
        })
    })
})

async function exportGIF() {
    if (isWorking.value) return

    console.log('Reloading Print.vue')

    const layers = { ...collectionStore.layers }

    if (!layers || !Object.keys(layers).length) {
        console.log('Nothing to export GIF')
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

    const collection = Collection.createFrom({
        sourceDirectory: collectionStore.sourceDirectory,
        outputDirectory: collectionStore.outputDirectory,
        traits: [...collectionStore.traits],
        layers: layers,
        width: parseFloat(collectionStore.width.toString()),
        height: parseFloat(collectionStore.height.toString()),
        size: parseInt(collectionStore.size.toString(), 10)
    })

    const frames = collectionStore.gif.frames || 10
    const delay = collectionStore.gif.delay || 33

    try {
        gif.value = await rpc.CollectionService.GenerateCollectionGif(collection, parseInt(frames.toString(), 10), parseInt(delay.toString(), 10))
    } catch (error) {
        console.error(error)
    }
}

async function load() {
    if (isWorking.value || store.isGeneratingCollection) return

    console.log('Reloading Run.vue')
    collectionStore.preview = ''
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
        collectionStore.preview = await rpc.CollectionService.GenerateCollectionPreview(collection)
    } catch (error) {
        console.error(error)
    }
}
</script>

<template>
    <section class="h-full flex">
        <main class="h-full flex-1 overflow-auto scrollbar-none">
            <section class="p-16 grid grid-cols-2 gap-8">
                <form class="col-span-2 max-w-4xl mx-auto flex flex-col px-8 animate__animated animate__fadeIn">
                    <h2 class="py-2 font-bold uppercase">Export Animated GIF</h2>
                    <div class="grid grid-cols-12 gap-8">
                        <div class="col-span-6">
                            <label for="gif-frames" class="block text-sm opacity-75">Number of Frames</label>
                            <div class="mt-1">
                                <input type="number" id="gif-frames" class="field" name="gif-frames" autocomplete="off"
                                    autofocus v-model="collectionStore.gif.frames" />
                            </div>
                        </div>
                        <div class="col-span-6">
                            <label for="gif-delay" class="block text-sm opacity-75">Delay Between Frames</label>
                            <div class="mt-1">
                                <select v-model="collectionStore.gif.delay" id="gif-delay"
                                    class="field mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-fuchsia-500 focus:border-fuchsia-500 sm:text-sm rounded-md">
                                    <option value="25">1/4 Second</option>
                                    <option value="33">1/3 Second</option>
                                    <option value="50">1/2 Second</option>
                                    <option value="100">1 Second</option>
                                    <option value="200">2 Seconds</option>
                                    <option value="300">3 Seconds</option>
                                    <option value="400">4 Seconds</option>
                                    <option value="500">5 Seconds</option>
                                </select>
                            </div>
                        </div>

                        <div class="col-span-6">
                            <div v-if="!isWorking">
                                <div v-if="collectionStore.preview">
                                    <h1 class="font-bold uppercase py-1">Frame Preview</h1>
                                    <Preview :source="collectionStore.preview" />
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
                        </div>

                        <div class="col-span-6">
                            <div v-if="gif">
                                <h1 class="font-bold uppercase py-1">Exported GIF Preview</h1>
                                <Preview :source="gif" />
                            </div>
                        </div>
                    </div>
                </form>
            </section>
        </main>

        <FloatingButtonBar>
            <button type="button"
                class="flex mt-2 py-3 px-4 items-center rounded text-slate-50 bg-slate-700 shadow shadow-slate-800 hover:bg-opacity-90"
                @click="$router.back">
                <span v-text="t('close')" />
            </button>

            <button v-if="collectionStore.layers && Object.keys(collectionStore.layers).length" type="button"
                class="flex mt-2 py-3 px-4 items-center rounded text-slate-50 bg-fuchsia-700 shadow shadow-fuchsia-900 hover:bg-opacity-90"
                :class="[isWorking ? 'opacity-50' : '']" :disabled="isWorking" @click="exportGIF">
                <span v-text="t('export_gif')" />
            </button>
        </FloatingButtonBar>
    </section>
</template>
