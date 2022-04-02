<script setup lang="ts">
import { onMounted } from 'vue'
import { useCollectionStore } from '@root/store'
import Sidebar from '@components/Sidebar.vue'
import { load, stat } from '@utils/Varly'
import { CogIcon, CollectionIcon, PlayIcon } from '@heroicons/vue/solid'
import FloatingButton from '@components/FloatingButton.vue'
import { useRouter } from 'vue-router'

const WIDTH = 1500
const HEIGHT = 1500
const SIZE = 100

const store = useCollectionStore()
load({ collectionStore: store, router: useRouter() })

onMounted(async () => {
    store.size = store.size || SIZE
    store.height = store.height || HEIGHT
    store.width = store.width || WIDTH

    if (store.layers && Object.keys(store.layers).length) {
        for (const trait in Object.keys(store.layers)) {
            if (store.layers.hasOwnProperty(trait)) {
                if (store.layers[trait] && store.layers[trait].length) {
                    const image = store.layers[trait][0]
                    if (image) {
                        const stats = await stat(image)
                        console.log(stats)
                    }
                }
            }
        }
    }
})
</script>

<template>
    <section class="h-full flex">
        <Sidebar
            :links="[
                { icon: CollectionIcon, text: 'Layer Setup', to: 'artwork.layers', selected: false },
                { icon: CogIcon, text: 'Build Settings', to: 'artwork.build', selected: true },
                { icon: PlayIcon, text: 'Run', to: 'artwork.run', selected: false },
            ]"
        />

        <main class="h-full flex-1 overflow-y-scroll scrollbar-none">
            <form class="max-w-4xl mx-auto flex flex-col px-8 py-16 xl:py-32 animate__animated animate__fadeIn">
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
                    <div class="col-span-12">
                        <label
                            for="collection-description"
                            class="block text-sm opacity-75"
                        >Collection Description</label>
                        <div class="mt-1">
                            <textarea
                                rows="4"
                                id="collection-description"
                                class="field"
                                name="collection-description"
                                autocomplete="none"
                                autofocus
                                v-model="store.description"
                            />
                        </div>
                    </div>

                    <hr class="col-span-12 opacity-0" />

                    <div class="col-span-6">
                        <label for="collection-width" class="block text-sm opacity-75">Layer Width</label>
                        <div class="mt-1">
                            <input
                                type="text"
                                autofocus
                                name="collection-width"
                                id="collection-width"
                                autocomplete="none"
                                class="field"
                                v-model="store.width"
                            />
                        </div>
                    </div>
                    <div class="col-span-6">
                        <label for="collection-height" class="block text-sm opacity-75">Layer Height</label>
                        <div class="mt-1">
                            <input
                                type="text"
                                autofocus
                                name="collection-height"
                                id="collection-height"
                                autocomplete="none"
                                class="field"
                                v-model="store.height"
                            />
                        </div>
                    </div>

                    <div class="col-span-6">
                        <label
                            for="collection-size"
                            class="block text-sm opacity-75"
                        >Collection Size</label>
                        <div class="mt-1">
                            <input
                                type="text"
                                autofocus
                                name="collection-size"
                                id="collection-size"
                                autocomplete="none"
                                class="field"
                                v-model="store.size"
                            />
                        </div>
                    </div>
                    <div class="col-span-6" />
                </div>
            </form>
        </main>

        <FloatingButton text="Next&nbsp;→" :to="() => $router.push({ name: 'artwork.run' })"></FloatingButton>
    </section>
</template>
