<script setup lang="ts">
import { BadgeCheckIcon, CogIcon, CollectionIcon, DocumentAddIcon, FolderOpenIcon, PlayIcon } from '@heroicons/vue/solid'
import { useCollectionStore } from '@/store'
import { launchTwitter } from '@/utils/Varly'
import Sidebar from '@/components/Sidebar.vue'

const store = useCollectionStore()
</script>

<template>
    <section class="h-full flex">
        <Sidebar
            :links="[
                { icon: BadgeCheckIcon, text: 'Support on Twitter', to: launchTwitter, selected: false },
                { icon: null, text: '', to: '', selected: false },
                { icon: FolderOpenIcon, text: 'Recent Projects', to: 'start', selected: false },
                { icon: DocumentAddIcon, text: 'Start New Project', to: 'artwork.layers', selected: false },
                { icon: null, text: '', to: '', selected: false },
                { icon: CollectionIcon, text: 'Layer Setup', to: 'artwork.layers', selected: false },
                // { icon: CollectionIcon, text: 'Collection Details', to: 'artwork.collection', selected: true },
                { icon: CogIcon, text: 'Build Settings', to: 'artwork.build', selected: false },
                { icon: PlayIcon, text: 'Run', to: 'artwork.run', selected: false },
            ]"
        />

        <main class="h-full flex-1 overflow-y-scroll scrollbar-none">
            <form class="max-w-4xl mx-auto p-8 lg:p-16 flex flex-col">
                <div class="grid grid-cols-12 gap-8">
                    <div class="col-span-12">
                        <label
                            for="collection-name"
                            class="block text-sm opacity-75"
                        >Collection Name</label>
                        <div class="mt-1">
                            <input
                                type="text"
                                id="collection-name"
                                class="field"
                                name="collection-name"
                                autocomplete="none"
                                autofocus
                                v-model="store.name"
                            />
                        </div>
                    </div>
                </div>
            </form>
        </main>
    </section>
</template>
