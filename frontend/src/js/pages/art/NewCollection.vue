<template>
    <div
        class="h-full flex flex-col"
        :class="hasLayers ? 'justify-start container mx-auto max-w-4xl' : 'justify-center'"
    >
        <section v-if="hasLayers && !isLoading" class="text-slate-900 dark:text-white">
            <nav>
                <button class="px-2" @click="generateCollection">Generate Collection</button>
                <button class="px-2" @click="startOver">Start Over</button>
                <button class="px-2" @click="isCollapsed = !isCollapsed">Expand/Collapse All</button>
            </nav>

            <!-- @see :force-fallback -->
            <!-- Solves issue where dragging works first but second drag requires two clicks -->
            <!-- https://github.com/SortableJS/Vue.Draggable/issues/954 -->
            <draggable
                class=""
                group="trait"
                v-model="store.traits"
                :force-fallback="true"
                @start="isTraitDragging = true"
                @end="isTraitDragging = false"
                item-key="name"
            >
                <template #item="{ element }">
                    <div>
                        <div class="flex items-center justify-between mt-4 py-2 px-4 bg-white bg-opacity-20">
                            <div class="flex items-center">
                                <button @click="toggleCollapsed(element)" class="pr-4">👆</button>
                                <h1 v-text="element.name" class="text-2xl uppercase"></h1>
                            </div>
                            <h2 class="text-lg font-semibold">{{ weightDistributionTotal(store.layers[element.name]) }}&percnt;</h2>
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
                                    class="flex items-center justify-between py-1 p-4 bg-white bg-opacity-5"
                                >
                                    <p class="font-mono">{{ element.Name }}</p>
                                    <input class="appearance-none border-none text-white w-1/12 bg-white bg-opacity-5 text-center" v-model="element.Weight" type="text" />
                                </div>
                            </template>
                        </draggable>
                    </div>
                </template>
            </draggable>
        </section>

        <section v-if="isLoading && loadingText">
            <h1 class="text-xl">{{ loadingText }}</h1>
        </section>

        <section v-else class="text-center">
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
    element.collapsed = !element.collapsed
}

function startOver() {
    store.reset()
    router.push({ name: 'welcome' })
}

async function loadLayers() {
    store.directory = await dialog.openDirectoryDialog()

    window.runtime.EventsOn('collection.generation.started', (data) => {
        loadingText.value = `Preparing collection of ${data.CollectionSize} items`
        console.log(loadingText.value)
    })

    window.runtime.EventsOn('collection.item.generated', (data) => {
        console.log(data)
        loadingText.value = `Generating collection: ${data.ItemNumber}/${data.CollectionSize}`
        console.log(loadingText.value)
    })

    const config = await app.GenerateCollection(store.directory)

    store.layers = { ...config.Layers }

    const traits: Object[] = []

    for (const trait in store.layers) {
        if (Object.hasOwnProperty.call(store.layers, trait)) {
            traits.push({ name: trait })
        }
    }

    store.traits = [...traits]
}

async function generateCollection() {
    toggleIsLoading()

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
        Dir: store.directory,
        Order: [...store.traits].map((item: any) => item.name),
        Layers: layers,
        Width: 1500,
        height: 1500,
        Size: 10
    }

    await app.GenerateCollectionFromConfig(config)

    toggleIsLoading()
}
</script>
