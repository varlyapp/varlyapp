<script setup lang="ts">
import { ref } from 'vue'
import Confetti from 'vue-confetti-explosion'
import { CogIcon, CollectionIcon, PlayIcon } from '@heroicons/vue/solid'
import Progress from '@/components/Progress.vue'
import Sidebar from '@/components/Sidebar.vue'
import FloatingButton from '@/components/FloatingButton.vue'
import { useCollectionStore } from '@/store'
import { useVarly } from '@/Varly'

const varly = useVarly()
const store = useCollectionStore()

const steps = ref(0)
const currentStep = ref(0)
const isWorking = ref(false)
const isDone = ref(false)
const loadingText = ref('Loading')
const preview = ref('')

function queueConfetti() {
    isDone.value = true
    varly.runtime.LogInfo('Setting isDone to true')
    setTimeout(() => {
        isDone.value = false
        varly.runtime.LogInfo('Setting isDone to false')
    }, 5000)
}

function toggleIsWorking() {
    isWorking.value = !isWorking.value
}

async function generateCollection() {
    toggleIsWorking()

    window.runtime.EventsOn('collection.generation.started', (data) => {
        loadingText.value = `Preparing collection of ${data.CollectionSize} items`
    })

    window.runtime.EventsOn('collection.item.generated', async (data) => {
        steps.value = data.CollectionSize
        currentStep.value = data.ItemNumber
    })

    const outputDirectory = await varly.openDirectoryDialog()

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
        Width: parseInt(store.width.toString(), 10),
        height: parseInt(store.height.toString(), 10),
        Size: parseInt(store.size.toString(), 10)
    }

    await varly.app.GenerateNewCollectionFromConfig(config)

    toggleIsWorking()
    queueConfetti()
}
</script>

<template>
    <section class="h-full flex">
        <Sidebar
            :links="[
                { icon: CollectionIcon, text: 'Layer Setup', to: 'artwork.layers', selected: false },
                { icon: CogIcon, text: 'Build Settings', to: 'artwork.build', selected: false },
                { icon: PlayIcon, text: 'Run', to: 'artwork.run', selected: true },
            ]"
        />

        <main class="h-full flex-1 overflow-y-scroll scrollbar-none">
            <div v-if="isWorking" class="h-full flex flex-col flex-1 max-w-4xl mx-auto items-center justify-center p-8">
                <Progress :steps="steps" :current-step="currentStep" loading-text="Preparing..." />
            </div>
            <div v-else class="h-full flex flex-col items-center justify-center p-8">
                <div class="flex flex-col items-center">
                    <div class="max-w-xs mx-auto">
                        <div v-if="preview">
                            <img :src="preview" alt="">
                        </div>
                        <h1
                            v-if="!isDone"
                            class="animate__animated animate__fadeIn text-base text-center"
                        >You are ready to generate your beautiful NFT&nbsp;collection🚀</h1>
                        <h1
                            v-else
                            class="animate__animated animate__fadeIn text-6xl text-center font-bold"
                        >Yay 🎉</h1>
                    </div>
                </div>
            </div>
        </main>

        <Confetti
            v-if="isDone"
            :particle-count="1000"
            :particle-size="8"
            :duration="5000"
            class="absolute w-screen h-screen top-0 right-0 bottom-0 left-0"
        />
        <FloatingButton v-if="!isWorking" text="Let&rsquo;s Go" :to="generateCollection"></FloatingButton>
    </section>
</template>