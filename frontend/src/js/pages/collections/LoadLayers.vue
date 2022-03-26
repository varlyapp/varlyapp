<script setup lang="ts">
import { ref, computed, onBeforeMount, inject } from 'vue'
import { useRouter } from 'vue-router'
import { useDialog } from '@utils/Dialog'
import { useStore, useCollectionStore } from '@root/store'
import draggable from 'vuedraggable'
import type { Varly } from '@root/plugins/varly'

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
    store.actions = [
        { title: '← Start Over', onClick: startOver, type: 'SECONDARY' },
        { title: '✓ Generate', onClick: generateCollection, type: 'PRIMARY' },
    ]
})

function cancel() {
    collectionStore.reset()
    varly.router.push({ name: 'start' })
}

const hasLayers = computed(() => {
    return collectionStore.layers && Object.keys(collectionStore.layers).length
})

const showWorkspace = computed(() => hasLayers.value && !isLoading.value)
const showOpenFolderPrompt = computed(() => !hasLayers.value && !isLoading.value)

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
</script>

<template>
    <div
        v-if="isLoading || hasCompleted"
        class="h-full flex flex-col justify-center container mx-auto p-8"
    >
        <div v-if="hasCompleted">
            <button class="block text-9xl" @click="() => router.back()">👍</button>
        </div>
        <div v-else>
            <h1 v-if="isLoading && loadingText" class="text-xl">{{ loadingText }} 🚀</h1>
            <div v-else>Loading...</div>
        </div>
    </div>
    <div
        v-else
        class="h-full flex justify-center flex-col container mx-auto"
        :class="hasLayers ? 'justify-start' : 'justify-center'"
    >
        <section v-if="showWorkspace" class="h-full flex flex-col justify-between">
            <div>
                <h1 class="text-lg">Layers</h1>
                <p class="opacity-75">You can arrange your layers from botton to top</p>
                <!-- @see :force-fallback -->
                <!-- Solves issue where dragging works first but second drag requires two clicks -->
                <!-- https://github.com/SortableJS/Vue.Draggable/issues/954 -->
                <draggable
                    class="mt-8"
                    group="trait"
                    v-model="collectionStore.traits"
                    :force-fallback="true"
                    @start="isTraitDragging = true"
                    @end="isTraitDragging = false"
                    item-key="name"
                >
                    <template #item="{ element }">
                        <div
                            class="border-t-2 border-t-slate-300 dark:border-t-slate-700 bg-slate-200 dark:bg-slate-900 mt-2 px-4 py-2"
                        >
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
                                    <div class="flex items-center justify-between mt-1">
                                        <p class="pl-4 font-mono">{{ element.Name }}</p>
                                        <input
                                            class="appearance-none dark:text-slate-800 dark:bg-slate-100 border-none text-center w-1/12"
                                            v-model="element.Weight"
                                            type="text"
                                        />
                                    </div>
                                </template>
                            </draggable>
                        </div>
                    </template>
                </draggable>
            </div>

            <div class="py-8 divide-y divide-slate-100 dark:divide-slate-800">
                <div class="flex justify-end">
                    <button
                        @click="cancel"
                        type="button"
                        class="bg-white py-2 px-4 border border-slate-200 rounded-sm shadow-sm text-sm font-medium text-slate-700 hover:bg-slate-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500"
                    >Cancel</button>
                    <button
                        @click="() => varly.router.push({ name: 'collections.generate' })"
                        type="submit"
                        class="ml-3 inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-sm text-white bg-purple-600 hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                    >Next</button>
                </div>
            </div>
        </section>

        <section v-if="showOpenFolderPrompt" class="text-center">
            <p
                class="mt-1 max-w-md mx-auto text-sm"
            >Select the folder with that has your layers in their own folder.</p>
            <button
                @click="loadLayers()"
                type="button"
                class="mt-6 ml-3 inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-sm text-white bg-purple-600 hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >Select Folder</button>
        </section>
    </div>
</template>
