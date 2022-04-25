<script setup lang="ts">
import { nextTick, onActivated, onBeforeMount, ref } from 'vue'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import draggable from 'vuedraggable'
import { Collection } from '@/wailsjs/go/models'
import { CogIcon, CollectionIcon, PlayIcon } from '@heroicons/vue/outline'
import { CollectionIcon as CollectionIconSolid } from '@heroicons/vue/solid'
import Sidebar from '@/components/Sidebar.vue'
import { useCollectionStore } from '@/store'
import rpc from '@/rpc'

const route = useRoute()
const intl = useI18n({ useScope: 'global' })
const { t } = intl

const collectionStore = useCollectionStore()

const isCollapsed = ref(false)
const isTraitDragging = ref(false)
const isLayerDragging = ref(false)

const ROUTE_NAME = 'artwork.layers'

onBeforeMount(() => rpc.setPageTitle("Layer Setup"))

onActivated(() => {
    nextTick(() => {
        window.runtime.EventsOn('shortcut.view.refresh', () => {
            if (route.name === ROUTE_NAME && collectionStore.sourceDirectory.length) loadLayersFromDirectory(collectionStore.sourceDirectory)
        })
        window.runtime.EventsOn('shortcut.view.hard-refresh', () => {
            if (route.name === ROUTE_NAME && collectionStore.sourceDirectory.length) loadLayersFromDirectory(collectionStore.sourceDirectory)
        })
    })
})

function weightDistributionTotal(items: any) {
    let total = 0
    items.forEach((item) => {
        total += parseInt(item.weight.toString(), 10)
    })

    if (isNaN(total)) return 0

    return total
}

function toggleCollapsed(element: any) {
    element.collapsed = !element.collapsed
}

async function loadLayers() {
    const sourceDirectory = await rpc.app.OpenDirectoryDialog('Open Collection Folder')

    if (!sourceDirectory) {
        console.error('source directory is empty')
    }

    collectionStore.sourceDirectory = sourceDirectory

    await loadLayersFromDirectory(sourceDirectory)
}

async function loadLayersFromDirectory(sourceDirectory: string) {
    const collection: Collection = await rpc.CollectionService.LoadCollectionFromDirectory(sourceDirectory)

    collectionStore.layers = { ...collection.layers }

    const traits: Object[] = []

    for (const trait in collectionStore.layers) {
        if (Object.hasOwnProperty.call(collectionStore.layers, trait)) {
            traits.push({ name: trait, collapsed: false })
        }
    }

    collectionStore.traits = [...traits]
}
</script>

<template>
    <section class="h-full flex">
        <Sidebar :links="[
            { icon: CollectionIcon, text: t('layer_setup'), to: 'artwork.layers', selected: true },
            { icon: CogIcon, text: t('build_settings'), to: 'artwork.build', selected: false },
            { icon: PlayIcon, text: t('run'), to: 'artwork.run', selected: false },
        ]" />

        <main class="relative h-full flex-1 overflow-y-scroll scrollbar-none">
            <section v-if="collectionStore && collectionStore.layers && Object.keys(collectionStore.layers).length"
                class="h-full animate__animated animate__fadeIn">
                <div class="p-8 mx:p-12 xl:p-16">
                    <!-- @see :force-fallback -->
                    <!-- Solves issue where dragging works first but second drag requires two clicks -->
                    <!-- https://github.com/SortableJS/Vue.Draggable/issues/954 -->
                    <draggable group="trait" v-model="collectionStore.traits" :force-fallback="true"
                        @start="() => isTraitDragging = true" @end="() => isTraitDragging = false" item-key="name">
                        <template #item="{ element }">
                            <div
                                class="mt-8 lg:mt-12 first:mt-2 border border-slate-900 dark:border-slate-100 border-opacity-20 dark:border-opacity-10">
                                <div
                                    class="flex items-center justify-between bg-slate-900 dark:bg-slate-200 bg-opacity-20 dark:bg-opacity-10">
                                    <div @click="toggleCollapsed(element)" class="flex items-center">
                                        <button
                                            class="py-1 px-2 mr-1 text-lg font-bold text-slate-900 dark:text-slate-100"
                                            @click="toggleCollapsed(element)">
                                            <span v-if="element.collapsed">⇣</span>
                                            <span v-else>⇡</span>
                                        </button>
                                        <h1 v-text="element.name" class="text-base font-bold"></h1>
                                    </div>
                                    <div>
                                        <p
                                            class="grow-0 text-base text-right px-2 font-bold bg-transparent border-0"
                                            v-text="`${weightDistributionTotal(collectionStore.layers[element.name])}%`" />
                                    </div>
                                </div>

                                <div :class="element.collapsed || isCollapsed ? 'hidden' : 'block'" group="layer"
                                    item-key="name">
                                    <div v-for="( collection, j) in collectionStore.layers[element.name]" :key="j"
                                        class="min-w-full flex justify-between border-t border-slate-900 dark:border-slate-50 border-opacity-20 dark:border-opacity-20"
                                        :class="[j % 2 === 0 ? `bg-slate-200 dark:bg-slate-800 bg-opacity-10 dark:bg-opacity-5` : `bg-slate-800 dark:bg-slate-400 bg-opacity-5 dark:bg-opacity-5`]">
                                        <div class="whitespace-nowrap py-2 px-4 text-sm font-medium sm:px-6 lg:px-8"
                                            v-text="collection.name" />
                                        <div>
                                            <input
                                                class="field grow-0 text-right appearance-none bg-transparent border-0"
                                                type="text" v-model="collection.weight" />
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </template>
                    </draggable>
                </div>
            </section>

            <section v-else class="h-full animate__animated animate__fadeIn">
                <div class="h-full p-4 lg:p-8 flex items-center justify-center text-center">
                    <div>
                        <h3 class="mt-2 text-sm font-medium" v-text="t('get_started_by_opening_your_layers_folder')" />
                        <div class="mt-8">
                            <button type="button" @click="loadLayers"
                                class="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-fuchsia-700 hover:bg-fuchsia-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-fuchsia-500">
                                <CollectionIconSolid class="-ml-1 mr-2 h-5 w-5" aria-hidden="true" />
                                <span v-text="t('open_layers_folder')" />
                            </button>
                        </div>
                    </div>
                </div>
            </section>
        </main>
    </section>
</template>
