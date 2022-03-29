<script setup lang="ts">
import { ref } from 'vue'
import { useDialog } from '@utils/Dialog'
import { useCollectionStore } from '@root/store'
import Sidebar from '@components/Sidebar.vue'
import { app, launchTwitter } from '@utils/Varly'
import { BadgeCheckIcon, CogIcon, CollectionIcon, DocumentAddIcon, DocumentDuplicateIcon, FolderOpenIcon, PlusIcon, PlayIcon } from '@heroicons/vue/solid'

const store = useCollectionStore()
const dialog = useDialog(app)

const steps = ref(0)
const currentStep = ref(0)
const isLoading = ref(false)
const loadingText = ref('Loading')
const preview = ref('')

const isCollapsed = ref(false)
const isTraitEnabled = ref(true)
const isTraitDragging = ref(false)
const isLayerDragging = ref(false)
const isLayerEnabled = ref(true)

const WIDTH = 7200
const HEIGHT = 9200
const SIZE = 100

function toggleIsLoading() {
    isLoading.value = !isLoading.value
}

async function generateCollection() {
    toggleIsLoading()

    window.runtime.EventsOn('collection.generation.started', (data) => {
        loadingText.value = `Preparing collection of ${data.CollectionSize} items`
    })

    window.runtime.EventsOn('collection.item.generated', async (data) => {
        steps.value = data.CollectionSize
        currentStep.value = data.ItemNumber
    })

    const outputDirectory = await dialog.openDirectoryDialog()

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
        Width: WIDTH,
        height: HEIGHT,
        Size: SIZE
    }

    await app.GenerateCollectionFromConfig(config)

    toggleIsLoading()
}
</script>

<template>
    <section class="h-full flex">
        <Sidebar
            :links="[
                { icon: BadgeCheckIcon, text: 'Support on Twitter', to: launchTwitter, selected: false },
                { icon: null, text: '', to: '', selected: false },
                { icon: FolderOpenIcon, text: 'Recent Projects', to: 'start', selected: false },
                { icon: DocumentAddIcon, text: 'Start NFT Project', to: 'artwork.layers', selected: false },
                { icon: null, text: '', to: '', selected: false },
                { icon: DocumentDuplicateIcon, text: 'Layer Setup', to: 'artwork.layers', selected: false },
                { icon: CollectionIcon, text: 'Collection Details', to: 'artwork.collection', selected: false },
                { icon: CogIcon, text: 'Build Settings', to: 'artwork.build', selected: true },
                { icon: PlayIcon, text: 'Run', to: 'artwork.run', selected: false },
            ]"
        />

        <main class="h-full flex-1 overflow-y-scroll scrollbar-none">
            <form class="p-8 xl:p-12 flex flex-col justify-between">
                <div class="mt-4 grid grid-cols-1 gap-y-6 gap-x-4 sm:grid-cols-12">
                    <div class="sm:col-span-6">
                        <label
                            for="collection-name"
                            class="block text-sm opacity-75"
                        >How Many Pieces?</label>
                        <div class="mt-1">
                            <input
                                type="number"
                                autofocus
                                name="collection-name"
                                id="collection-name"
                                autocomplete="none"
                                class="text-slate-700 dark:bg-slate-100 shadow-sm focus:ring-purple-500 focus:border-purple-500 block w-full sm:text-sm border-slate-300 rounded-sm"
                            />
                        </div>
                    </div>
                    <div class="sm:col-span-6"/>

                    <div class="sm:col-span-6">
                        <label
                            for="collection-name"
                            class="block text-sm opacity-75"
                        >Select where to put your generated artwork files</label>
                        <div class="mt-1">
                            <input
                                type="number"
                                name="collection-name"
                                id="collection-name"
                                autocomplete="none"
                                disabled
                                class="text-slate-700 dark:bg-slate-100 shadow-sm focus:ring-purple-500 focus:border-purple-500 block w-full sm:text-sm border-slate-300 rounded-sm"
                            />
                        </div>
                    </div>

                    <div class="sm:col-span-6"/>

                    <div class="sm:col-span-6">
                        <label for="about" class="block text-sm opacity-75">Collection Description</label>
                        <div class="mt-1">
                            <textarea
                                id="about"
                                name="about"
                                rows="4"
                                class="text-slate-700 dark:bg-slate-100 shadow-sm focus:ring-purple-500 focus:border-purple-500 block w-full sm:text-sm border border-slate-300 rounded-sm"
                            />
                        </div>
                    </div>
                </div>
            </form>
        </main>
    </section>
</template>
