<template>
    <div
        class="h-full flex flex-col"
        :class="hasLayers ? 'justify-start container mx-auto' : 'justify-center'"
    >
        <nav>
            <button class="px-2" @click="generateCollection">Generate Collection</button>
            <button class="px-2" @click="startOver">Start Over</button>
            <button class="px-2" @click="isCollapsed = !isCollapsed">Expand/Collapse All</button>
        </nav>

        <section v-if="hasLayers" class="text-slate-900 dark:text-white">
            <draggable
                class="mt-4"
                group="trait"
                :list="store.traits"
                :disabled="!isEnabled"
                @start="isDragging = true"
                @end="isDragging = false"
                item-key="name"
            >
                <template #item="{ element }">
                    <div>
                        <div class="flex items-center justify-between">
                            <div class="flex items-center">
                                <button @click="toggleCollapsed(element)" class="pr-4">👆</button>
                                <h1 v-text="element.name" class="text-2xl uppercase"></h1>
                            </div>
                            <h2>{{ weightDistributionTotal(store.layers[element.name]) }}</h2>
                        </div>
                        <draggable
                            class="mt-4"
                            :class="element.collapsed || isCollapsed ? 'hidden' : 'block'"
                            group="layer"
                            :list="store.layers[element.name]"
                            :disabled="!isEnabled"
                            @start="isDragging = true"
                            @end="isDragging = false"
                            item-key="name"
                        >
                            <template #item="{ element }">
                                <div
                                    class="flex items-center justify-between pt-2 py-1 px-2 border-t"
                                >
                                    <p class="font-mono">{{ element.Name }}</p>
                                    <input class="text-black" v-model="element.Weight" type="text" />
                                </div>
                            </template>
                        </draggable>
                    </div>
                </template>
            </draggable>
        </section>

        <section v-else class="text-center">
            <p
                class="mt-1 max-w-md mx-auto"
            >Select the root folder that contains all other folders with layers for each&nbsp;trait.</p>
            <button
                @click="loadLayers()"
                type="button"
                class="mt-8 py-4 px-8 bg-purple-700 text-white rounded shadow-xl hover:bg-opacity-90 font-semibold"
            >Select Folder</button>
        </section>
    </div>
</template>

<script setup>
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
const isDragging = ref(false)
const isEnabled = ref(true)
const isCollapsed = ref(false)

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
    toggleIsLoading()

    store.directory = await dialog.openDirectoryDialog()

    // runtime.EventsOn('collection.generation.started', (data) => {
    //     loadingText.value = `Preparing collection of ${data.CollectionSize} items`
    //     console.log(loadingText.value)
    // })

    // runtime.EventsOn('collection.item.generated', (data) => {
    //     console.log(data)
    //     loadingText.value = `Generating collection: ${data.ItemNumber}/${data.CollectionSize}`
    //     console.log(loadingText.value)
    // })

    const config = await app.GenerateCollection(store.directory)

    store.layers = { ...config.Layers }

    const traits = []
    for (const trait in store.layers) {
        if (Object.hasOwnProperty.call(store.layers, trait)) {
            traits.push({ name: trait })
        }
    }

    store.traits = [...traits]

    toggleIsLoading()
}

function generateCollection() {
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
        Order: Object.keys(layers),
        Layers: layers,
        Width: 1500,
        height: 1500,
        Size: 10
    }

    app.GenerateCollectionFromConfig(config)
}
</script>
