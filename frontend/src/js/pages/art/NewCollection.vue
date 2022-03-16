<template>
    <div
        class="h-full flex flex-col container mx-auto max-w-4xl py-2 px-4"
        :class="hasLayers ? 'justify-start' : 'justify-center'"
    >
        <section v-if="showWorkspace" class="text-slate-900 dark:text-white">
            <nav class="flex items-center justify-between">
                <button class @click="startOver">🡠</button>
                <button
                    class="bg-purple-700 text-white px-4 py-2 rounded shadow"
                    @click="generateCollection"
                >✓&nbsp;Generate Collection</button>
                <!-- <button class="pr-2" @click="isCollapsed = !isCollapsed">Expand/Collapse All</button> -->
            </nav>

            <!-- @see :force-fallback -->
            <!-- Solves issue where dragging works first but second drag requires two clicks -->
            <!-- https://github.com/SortableJS/Vue.Draggable/issues/954 -->
            <draggable
                class
                group="trait"
                v-model="store.traits"
                :force-fallback="true"
                @start="isTraitDragging = true"
                @end="isTraitDragging = false"
                item-key="name"
            >
                <template #item="{ element }">
                    <div>
                        <div
                            class="flex items-center justify-between mt-4 pt-1 border-t border-slate-400"
                        >
                            <div class="flex items-center">
                                <button
                                    class="py-1 px-2 mr-2 bg-slate-900 text-slate-100"
                                    @click="toggleCollapsed(element)"
                                >
                                    <span v-if="element.collapsed">⇣</span>
                                    <span v-else>⇡</span>
                                </button>
                                <h1 v-text="element.name" class="text-2xl uppercase"></h1>
                            </div>
                            <h2
                                class="text-lg font-semibold"
                            >{{ weightDistributionTotal(store.layers[element.name]) }}&percnt;</h2>
                        </div>
                        <draggable
                            :class="element.collapsed || isCollapsed ? 'hidden' : 'block'"
                            group="layer"
                            :force-fallback="true"
                            :list="store.layers[element.name]"
                            @start="isLayerDragging = true"
                            @end="isLayerDragging = false"
                            item-key="name"
                        >
                            <template #item="{ element }">
                                <div
                                    class="flex items-center justify-between mt-1 border-t border-dashed border-slate-400"
                                >
                                    <p class="font-mono">{{ element.Name }}</p>
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

        <section class="h-full flex items-start justify-center text-center" v-if="isLoading && loadingText">
            <h1 class="text-xl">{{ loadingText }}</h1>
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
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useDialog } from '@utils/Dialog'
import { useCollectionStore } from '@root/store'
import draggable from 'vuedraggable'

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

const hasLayers = computed(() => {
    return store.layers && Object.keys(store.layers).length
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
    store.reset()
    router.push({ name: 'welcome' })
}

async function loadLayers() {
    store.directory = await dialog.openDirectoryDialog()

    const config = await app.GenerateCollection(store.directory)

    store.layers = { ...config.Layers }

    const traits: Object[] = []

    for (const trait in store.layers) {
        if (Object.hasOwnProperty.call(store.layers, trait)) {
            traits.push({ name: trait, collapsed: false })
        }
    }

    store.traits = [...traits]
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
        Width: 1500,
        height: 1500,
        Size: 1000
    }

    await app.GenerateCollectionFromConfig(config)

    toggleIsLoading()
}
</script>
