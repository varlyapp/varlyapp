<script setup lang="ts">
import { nextTick, onBeforeMount, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { CogIcon, CollectionIcon, PlayIcon } from '@heroicons/vue/outline'
import { useCollectionStore } from '@/store'
import Sidebar from '@/components/Sidebar.vue'
import rpc from '@/rpc'

const WIDTH = 1500
const HEIGHT = 1500
const SIZE = 100

const route = useRoute()
const collectionStore = useCollectionStore()

const i18n = useI18n()
const { t } = i18n

onBeforeMount(() => rpc.setPageTitle("Build Settings"))

onMounted(() => {
    load()
    nextTick(() => {
        window.runtime.EventsOn('shortcut.view.refresh', () => {
            if (route.name === 'artwork.build') load()
        })
    })
})

async function load() {
    console.log('Reloading Build.vue')
    collectionStore.size = collectionStore.size || SIZE
    collectionStore.height = collectionStore.height || HEIGHT
    collectionStore.width = collectionStore.width || WIDTH

    if (collectionStore.layers && Object.keys(collectionStore.layers).length) {
        for (const trait in collectionStore.layers) {
            if (collectionStore.layers.hasOwnProperty(trait)) {
                if (collectionStore.layers[trait] && collectionStore.layers[trait].length) {
                    const variant = collectionStore.layers[trait][0]
                    if (variant) {
                        const image = await rpc.app.GetImageStats(variant.path)

                        if (image['Width'] && image['Height']) {
                            collectionStore.width = image['Width']
                            collectionStore.height = image['Height']
                        }
                    }
                }
            }
        }
    }
}
</script>

<template>
    <section class="h-full flex">
        <Sidebar :links="[
            { icon: CollectionIcon, text: t('layer_setup'), to: 'artwork.layers', selected: false },
            { icon: CogIcon, text: t('build_settings'), to: 'artwork.build', selected: true },
            { icon: PlayIcon, text: t('run'), to: 'artwork.run', selected: false },
        ]" />

        <main class="h-full flex-1 overflow-y-scroll scrollbar-none">
            <form class="max-w-4xl mx-auto flex flex-col p-8 lg:p-16 xl:p-24 animate__animated animate__fadeIn">
                <div class="grid grid-cols-12 gap-8">
                    <div class="col-span-12">
                        <label for="collection-name" class="block text-sm opacity-75">Collection Name</label>
                        <div class="mt-1">
                            <input type="text" id="collection-name" class="field" name="collection-name"
                                autocomplete="off" placeholder="e.g. Boss Beauties" autofocus v-model="collectionStore.name" />
                        </div>
                    </div>
                    <div class="col-span-12">
                        <label for="collection-name" class="block text-sm opacity-75">Collection Artist</label>
                        <div class="mt-1">
                            <input type="text" id="collection-artist" class="field" name="collection-artist"
                                autocomplete="off" placeholder="e.g Your name or pseudonym" v-model="collectionStore.artist" />
                        </div>
                    </div>
                    <div class="col-span-12">
                        <label for="collection-description" class="block text-sm opacity-75">Collection
                            Description</label>
                        <div class="mt-1">
                            <textarea rows="6" id="collection-description" class="field scrollbar-none"
                                name="collection-description" autocomplete="off"
                                placeholder="A short description about your collection" v-model="collectionStore.description" />
                        </div>
                    </div>

                    <div class="col-span-12">
                        <label for="collection-name" class="block text-sm opacity-75">Collection Base URI</label>
                        <div class="mt-1">
                            <input type="text" id="collection-base-uri" class="field" name="collection-base-uri"
                                autocomplete="off" spellcheck="false" placeholder="e.g ipfs://your-collection-cid/"
                                v-model="collectionStore.baseUri" />
                        </div>
                    </div>

                    <hr class="col-span-12 opacity-0" />

                    <div class="col-span-6">
                        <label for="collection-width" class="block text-sm opacity-75">Layer Width</label>
                        <div class="mt-1">
                            <input type="text" name="collection-width" id="collection-width" autocomplete="off"
                                class="field" v-model="collectionStore.width" />
                        </div>
                    </div>
                    <div class="col-span-6">
                        <label for="collection-height" class="block text-sm opacity-75">Layer Height</label>
                        <div class="mt-1">
                            <input type="text" name="collection-height" id="collection-height" autocomplete="off"
                                class="field" v-model="collectionStore.height" />
                        </div>
                    </div>

                    <div class="col-span-6">
                        <label for="collection-size" class="block text-sm opacity-75">Collection Size</label>
                        <div class="mt-1">
                            <input type="text" name="collection-size" id="collection-size" autocomplete="off"
                                class="field" v-model="collectionStore.size" />
                        </div>
                    </div>
                    <div class="col-span-6"></div>
                </div>
            </form>
        </main>
    </section>
</template>
