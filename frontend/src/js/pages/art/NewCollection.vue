<template>
    <div
        v-if="isLoading || hasCompleted"
        class="h-full flex items-center justify-center flex-col container mx-auto max-w-4xl py-2 px-4"
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
        class="h-full flex flex-col container mx-auto max-w-4xl py-2 px-4"
        :class="hasLayers ? 'justify-start' : 'justify-center'"
    >
        <section v-if="showWorkspace" class="relative text-slate-900 dark:text-white">
            <!-- <div class="sticky top-0">
                <nav class="flex items-center justify-between">
                    <button
                        class="py-2 px-3 rounded-full text-white bg-purple-700 shadow"
                        @click="startOver"
                    >🡠</button>
                    <div>
                        <button
                            class="mx-1 bg-purple-700 text-white px-4 py-2 rounded shadow"
                            @click="generateCollection"
                        >Generate Collection</button>
                        <button
                            class="mx-1 bg-black text-white px-4 py-2 rounded shadow"
                            @click="saveProgress"
                        >✓&nbsp;Save Progress</button>
                    </div>
                </nav>
            </div>-->

            <h1 class="text-lg opacity-75">Layers</h1>
            <p class="opacity-50">You can arrange your layers from botton to top</p>
            <!-- @see :force-fallback -->
            <!-- Solves issue where dragging works first but second drag requires two clicks -->
            <!-- https://github.com/SortableJS/Vue.Draggable/issues/954 -->
            <draggable
                class
                group="trait"
                v-model="collectionStore.traits"
                :force-fallback="true"
                @start="isTraitDragging = true"
                @end="isTraitDragging = false"
                item-key="name"
            >
                <template #item="{ element }">
                    <div>
                        <div class="flex items-center justify-between mt-4">
                            <div class="flex items-center">
                                <button
                                    class="py-1 px-2 mr-2 bg-slate-900 text-slate-100"
                                    @click="toggleCollapsed(element)"
                                >
                                    <span v-if="element.collapsed">⇣</span>
                                    <span v-else>⇡</span>
                                </button>
                                <h1 v-text="element.name" class="text-base"></h1>
                            </div>
                            <div>
                                <input
                                    class="grow-0 text-right appearance-none bg-transparent text-white border-0"
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
                                        class="appearance-none border-none text-center w-1/12"
                                        v-model="element.Weight"
                                        type="text"
                                    />
                                </div>
                            </template>
                        </draggable>
                    </div>
                </template>
            </draggable>
        </section>

        <section v-if="showOpenFolderPrompt" class="text-center">
            <p
                class="mt-1 max-w-md mx-auto"
            >Select the root folder that contains all other folders with layers for each&nbsp;trait.</p>
            <button
                @click="loadLayers()"
                type="button"
                class="mt-8 py-4 px-8 bg-purple-800 text-white rounded shadow-xl hover:bg-opacity-80 font-semibold"
            >Select Folder</button>
        </section>
    </div>
</template>

<script setup lang="ts">
import { ref, computed, onBeforeMount } from 'vue'
import { useRouter } from 'vue-router'
import { useDialog } from '@utils/Dialog'
import { useStore, useCollectionStore } from '@root/store'
import draggable from 'vuedraggable'

const app = window.go.main.App

const dialog = useDialog(app)
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

    const config = await app.ReadLayers(collectionStore.directory)

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

    const file = await app.SaveFileDialog()

    await app.Save(file, JSON.stringify(config))
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

    await app.GenerateCollectionFromConfig(config)

    hasCompleted.value = true
    toggleIsLoading()
}
</script>
