<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import { ref, computed, onBeforeMount } from 'vue'
import draggable from 'vuedraggable'
import { useDialog } from '@utils/Dialog'
import { useCollectionStore } from '@root/store'
import Sidebar from '@components/Sidebar.vue'
import { app, launchTwitter } from '@utils/Varly'
import { BadgeCheckIcon, CogIcon, CollectionIcon, DocumentAddIcon, FolderOpenIcon, PlusIcon, PlayIcon } from '@heroicons/vue/solid'
import FloatingButton from '@components/FloatingButton.vue'

const intl = useI18n({ useScope: 'global' })
const { t } = intl

const dialog = useDialog(app)
const collectionStore = useCollectionStore()

let loadingText = ref('')
let isLoading = ref(false)
let hasCompleted = ref(false)

const isCollapsed = ref(false)
const isTraitEnabled = ref(true)
const isTraitDragging = ref(false)
const isLayerDragging = ref(false)
const isLayerEnabled = ref(true)

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
    element.collapsed = !element.collapsed
}

async function saveSettings() {
    const docs = [
        {
            title: 'Doodles',
            path: '/Users/selvinortiz/Desktop/doodles.varly',
            preview: '',
        },
        {
            title: 'VeeFriends',
            path: '/Users/selvinortiz/Desktop/vee-friends.varly',
            preview: '',
        },
        {
            title: 'InvisibleFiends',
            path: '/Users/selvinortiz/Desktop/invisible-friends.varly',
            preview: '',
        },
    ]

    await app.SaveDocuments(docs)
}

async function loadLayers() {
    collectionStore.directory = await dialog.openDirectoryDialog()

    const config: { Layers?} = await app.ReadLayers(collectionStore.directory)

    collectionStore.layers = { ...config.Layers }

    const traits: Object[] = []

    for (const trait in collectionStore.layers) {
        if (Object.hasOwnProperty.call(collectionStore.layers, trait)) {
            traits.push({ name: trait, collapsed: false })
        }
    }

    collectionStore.traits = [...traits]
}

async function saveProgress() {
    window.runtime.EventsOn('collection.generation.saved', (data) => {
        loadingText.value = `Preparing collection of ${data.CollectionSize} items`
    })

    const outputDirectory = await dialog.openDirectoryDialog()

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

    await app.SaveFile(file, JSON.stringify(config))
}
</script>

<template>
    <section class="h-full flex">
        <Sidebar
            :links="[
                { icon: BadgeCheckIcon, text: t('follow_on_twitter'), to: launchTwitter, selected: false },
                { icon: null, text: '', to: '', selected: false },
                { icon: FolderOpenIcon, text: t('recent_projects'), to: 'start', selected: false },
                { icon: DocumentAddIcon, text: t('start_new_project'), to: 'artwork.layers', selected: false },
                { icon: null, text: '', to: '', selected: false },
                { icon: CollectionIcon, text: t('layer_setup'), to: 'artwork.layers', selected: true },
                // { icon: CollectionIcon, text: 'Collection Details', to: 'artwork.collection', selected: false },
                { icon: CogIcon, text: t('build_settings'), to: 'artwork.build', selected: false },
                { icon: PlayIcon, text: t('run'), to: 'artwork.run', selected: false },
            ]"
        />

        <main class="h-full flex-1 overflow-y-scroll scrollbar-none">
            <section v-if="hasLayers" class="h-full animate__animated animate__fadeIn">
                <div class="max-w-4xl mx-auto px-8 py-16 xl:py-32">
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
                            <div class="p-4 lg:p-8">
                                <div
                                    class="flex items-center justify-between bg-slate-900 bg-opacity-20 dark:bg-opacity-20"
                                >
                                    <div class="flex items-center">
                                        <button
                                            class="py-1 px-1 mr-1 text-slate-100"
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
                                    <template #item="{ element, index }">
                                        <div
                                            :key="element.Name"
                                            class="min-w-full flex justify-between border-slate-900 dark:border-slate-50 border-opacity-20 dark:border-opacity-20"
                                            :class="[index % 2 === 0 ? `bg-slate-900 dark:bg-slate-100 bg-opacity-5 dark:bg-opacity-5` : `bg-slate-800 dark:bg-slate-400 bg-opacity-5 dark:bg-opacity-5`]"
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

            <section v-else class="h-full animate__animated animate__fadeIn">
                <div class="h-full p-4 lg:p-8 flex items-center justify-center text-center">
                    <div>
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
                                class="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-fuchsia-600 hover:bg-fuchsia-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-fuchsia-500"
                                @click="loadLayers()"
                            >
                                <PlusIcon class="-ml-1 mr-2 h-5 w-5" aria-hidden="true" />New Project
                            </button>
                        </div>
                    </div>
                </div>
            </section>
        </main>

        <FloatingButton
            v-if="hasLayers"
            text="Next&nbsp;→"
            :to="() => $router.push({ name: 'artwork.build' })"
        ></FloatingButton>
    </section>
</template>
