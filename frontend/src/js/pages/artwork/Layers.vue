<script setup lang="ts">
import { ref, computed, onBeforeMount, inject } from 'vue'
import { PlusIcon } from '@heroicons/vue/solid'

import { useRouter } from 'vue-router'
import { useDialog } from '@utils/Dialog'
import { useStore, useCollectionStore } from '@root/store'
import draggable from 'vuedraggable'
import type { Varly } from '@root/plugins/varly'
import Sidebar from '@components/Sidebar.vue'

const varly = inject<Varly>('varly')!

const dialog = useDialog(varly?.app)
const router = useRouter()
const store = useStore()
const collectionStore = useCollectionStore()

let loadingText = ref('')
let isLoading = ref(false)
let hasCompleted = ref(false)

const isCollapsed = ref(false)
const isTraitEnabled = ref(true)
const isTraitDragging = ref(false)
const isLayerDragging = ref(false)
const isLayerEnabled = ref(true)

onBeforeMount(() => {
    // console.log(collectionStore.layers)
})

function cancel() {
    collectionStore.reset()
    varly.router.push({ name: 'start' })
}

const hasLayers = computed(() => {
    return collectionStore.layers && Object.keys(collectionStore.layers).length
})

function weightDistributionTotal(items) {
    let total = 0
    items.forEach((item) => {
        total += parseInt(item.Weight.toString(), 10)
    })

    return total
}

function toggleIsLoading() {
    isLoading.value = !isLoading.value
}

function toggleCollapsed(element) {
    console.log(element.collapsed)
    element.collapsed = !element.collapsed
}

function startOver() {
    collectionStore.reset()
    router.push({ name: 'start' })
}

async function loadLayers() {
    collectionStore.directory = await dialog.openDirectoryDialog()

    const config = await varly.app.ReadLayers(collectionStore.directory)

    collectionStore.layers = { ...config.Layers }

    const traits: Object[] = []

    for (const trait in collectionStore.layers) {
        if (Object.hasOwnProperty.call(collectionStore.layers, trait)) {
            traits.push({ name: trait, collapsed: false })
        }
    }

    console.log(config)
    collectionStore.traits = [...traits]
}

async function saveProgress() {
    window.runtime.EventsOn('collection.generation.saved', (data) => {
        loadingText.value = `Preparing collection of ${data.CollectionSize} items`
        console.log(loadingText.value)
    })

    const outputDirectory = await dialog.openDirectoryDialog()

    console.log(outputDirectory)

    const layers = { ...collectionStore.layers }

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
        Order: [...collectionStore.traits].map((item: any) => item.name),
        Layers: layers,
        Width: 512,
        height: 512,
        Size: 1000
    }

    const file = await varly.app.SaveFileDialog()

    await varly.app.SaveFile(file, JSON.stringify(config))
    console.log('Done')
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

    const outputDirectory = await dialog.openDirectoryDialog()

    console.log(outputDirectory)

    const layers = { ...collectionStore.layers }

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
        Order: [...collectionStore.traits].map((item: any) => item.name),
        Layers: layers,
        Width: 512,
        height: 512,
        Size: 1000
    }

    await varly.app.GenerateCollectionFromConfig(config)

    hasCompleted.value = true
    toggleIsLoading()
}

function run(e: any) {
    console.log('Called run()')
    console.log(e)
}

</script>

<template>
    <section class="h-full flex">
        <Sidebar
            :links="[
                { emoji: '🦋', text: 'Support on Twitter', to: 'start', selected: false },
                { emoji: '', text: '', to: '', selected: false },
                { emoji: '🏙', text: 'Recent Projects', to: 'start', selected: false },
                { emoji: '✨', text: 'Start NFT Project', to: 'artwork.layers', selected: false },
                { emoji: '', text: '', to: '', selected: false },
                { emoji: '🥞', text: 'Layer Setup', to: 'artwork.layers', selected: true },
                { emoji: '📝', text: 'Collection Details', to: 'artwork.collection', selected: false },
                { emoji: '🧩', text: 'Build Settings', to: 'artwork.build', selected: false },
                { emoji: '🚀', text: 'Run', to: run, selected: false },
            ]"
        />

        <main class="h-full flex-1 overflow-y-scroll scrollbar-none">
            <section v-if="hasLayers" class>
                <div class="p-8">
                    <!-- @see :force-fallback -->
                    <!-- Solves issue where dragging works first but second drag requires two clicks -->
                    <!-- https://github.com/SortableJS/Vue.Draggable/issues/954 -->
                    <draggable
                        class="rounded bg-slate-50 bg-opacity-20 dark:bg-slate-900 dark:bg-opacity-20 border-dashed border-2 border-slate-900 dark:border-slate-50 border-opacity-20 dark:border-opacity-20"
                        group="trait"
                        v-model="collectionStore.traits"
                        :force-fallback="true"
                        @start="isTraitDragging = true"
                        @end="isTraitDragging = false"
                        item-key="name"
                    >
                        <template #item="{ element }">
                            <div class="mt-2 px-4 py-2">
                                <div class="flex items-center justify-between">
                                    <div class="flex items-center">
                                        <button
                                            class="py-1 px-2 mr-2 bg-slate-700 text-slate-100"
                                            @click="toggleCollapsed(element)"
                                        >
                                            <span v-if="element.collapsed">⇣</span>
                                            <span v-else>⇡</span>
                                        </button>
                                        <h1 v-text="element.name" class="text-base"></h1>
                                    </div>
                                    <div>
                                        <input
                                            class="grow-0 text-right appearance-none bg-transparent border-0"
                                            type="text"
                                            :value="`${weightDistributionTotal(collectionStore.layers[element.name])}`"
                                        />
                                    </div>
                                    <!-- <h2
                                class="text-lg font-semibold"
                                    >{{ weightDistributionTotal(collectionStore.layers[element.name]) }}&percnt;</h2>-->
                                </div>

                                <draggable
                                    :class="element.collapsed || isCollapsed ? 'hidden' : 'block'"
                                    group="layer"
                                    :force-fallback="true"
                                    :list="collectionStore.layers[element.name]"
                                    @start="isLayerDragging = true"
                                    @end="isLayerDragging = false"
                                    item-key="name"
                                >
                                    <template #item="{ element }">
                                        <div
                                            :key="element.Name"
                                            class="min-w-full flex justify-between border-dotted border-t border-slate-900 dark:border-slate-50 border-opacity-20 dark:border-opacity-20"
                                        >
                                            <div
                                                class="whitespace-nowrap py-2 px-4 text-sm font-medium sm:px-6 lg:px-8"
                                                v-text="element.Name"
                                            />
                                            <div
                                                class="whitespace-nowrap py-2 px-4 text-sm"
                                                v-text="element.Weight"
                                            />
                                        </div>
                                    </template>
                                </draggable>
                            </div>
                        </template>
                    </draggable>
                </div>
            </section>

            <section v-else>
                <div class="p-4 lg:p-8 text-center">
                    <svg
                        class="mx-auto h-12 w-12 text-gray-400"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                        aria-hidden="true"
                    >
                        <path
                            vector-effect="non-scaling-stroke"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M9 13h6m-3-3v6m-9 1V7a2 2 0 012-2h6l2 2h6a2 2 0 012 2v8a2 2 0 01-2 2H5a2 2 0 01-2-2z"
                        />
                    </svg>
                    <h3 class="mt-2 text-sm font-medium">No Projects</h3>
                    <p class="mt-8 text-sm text-opacity-50">Get started by creating a new one🎉</p>
                    <div class="mt-6">
                        <button
                            type="button"
                            class="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-purple-600 hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500"
                            @click="loadLayers()"
                        >
                            <PlusIcon class="-ml-1 mr-2 h-5 w-5" aria-hidden="true" />New Project
                        </button>
                    </div>
                </div>
            </section>
        </main>
    </section>
</template>
