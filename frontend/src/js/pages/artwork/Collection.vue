<script setup lang="ts">
import { inject } from 'vue'
import type { Varly } from '@root/plugins/varly'
import Sidebar from '@components/Sidebar.vue'

const varly = inject<Varly>('varly')!
</script>

<template>
    <section class="h-full flex">
        <Sidebar
            :links="[
                { emoji: '🦋', text: 'Support on Twitter', to: 'start', selected: false },
                { emoji: '', text: '', to: '', selected: false },
                { emoji: '🏙', text: 'Recent Projects', to: 'start', selected: false },
                { emoji: '✨', text: 'Start NFT Project', to: 'artwork.layers', selected: false },
                { emoji: '', text: '', to: '', selected: false },
                { emoji: '🥞', text: 'Layer Setup', to: 'artwork.layers', selected: false },
                { emoji: '📝', text: 'Collection Details', to: 'artwork.collection', selected: true },
                { emoji: '🧩', text: 'Build Settings', to: 'artwork.build', selected: false },
                { emoji: '🚀', text: 'Run', to: '', selected: false },
            ]"
        />

        <main class="h-full flex-1 overflow-y-scroll scrollbar-none">
            <form
                class="p-8 space-y-8 divide-y divide-slate-100 dark:divide-slate-800 flex flex-col justify-between"
            >
                <div>
                    <h1 class="text-lg leading-6 font-medium">New Collection</h1>
                    <p
                        class="mt-1 text-sm opacity-75"
                    >Please provide details about your new NFT collection</p>
                </div>

                <div class="mt-6 grid grid-cols-1 gap-y-6 gap-x-4 sm:grid-cols-6">
                    <div class="sm:col-span-6">
                        <label
                            for="collection-name"
                            class="block text-sm opacity-75"
                        >Collection Name</label>
                        <div class="mt-1">
                            <input
                                type="text"
                                name="collection-name"
                                id="collection-name"
                                autocomplete="none"
                                class="text-slate-700 dark:bg-slate-100 shadow-sm focus:ring-purple-500 focus:border-purple-500 block w-full sm:text-sm border-slate-300 rounded-sm"
                            />
                        </div>
                    </div>

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

                <div class="pt-5">
                    <div class="flex justify-end">
                        <button
                            type="button"
                            class="bg-white py-2 px-4 border border-slate-200 rounded-sm shadow-sm text-sm font-medium text-slate-700 hover:bg-slate-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500"
                        >Cancel</button>
                        <button
                            @click="() => varly.router.push({ name: 'artwork.build' })"
                            type="submit"
                            class="ml-3 inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-sm text-white bg-purple-600 hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500"
                        >Next</button>
                    </div>
                </div>
            </form>
        </main>
    </section>
</template>
