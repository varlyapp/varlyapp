<script setup lang="ts">
import { computed, type FunctionalComponent, type PropType } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { BrowserOpenURL } from '@wails/runtime'
import { CheckBadgeIcon, DocumentPlusIcon, FolderOpenIcon, HeartIcon } from '@heroicons/vue/24/outline'
import SidebarButton from '@/components/SidebarButton.vue'
import { useCollectionStore } from '@/store'
import rpc from '@/rpc'

const route = useRoute()
const router = useRouter()
const collectionStore = useCollectionStore()

const { t } = useI18n({ useScope: 'global' })

type Links = Array<{
    text: string
    icon: FunctionalComponent | any
    to: string | Function
    selected: boolean
}>

const props = defineProps({
    links: Array as PropType<Links>,
})

const startNewProjectAction = () => {
    return async function () {
        if (!Object.keys(collectionStore.layers).length) {
            return router.push({ name: 'artwork.layers' })
        }

        const response = await rpc.app.MessageDialog({
            Title: t(`confirm_start_new_project_title`),
            Message: t(`confirm_start_new_project_message`),
            Buttons: [`OK`, `Cancel`],
            DefaultButton: `OK`
        })

        if (['ok', 'yes'].includes(response.toLowerCase())) {
            collectionStore.reset()
            router.push({ name: 'artwork.layers' })
        }
    }
}

const presets = {
    links: [
        { icon: null, text: t('support'), to: '', selected: false },
        { icon: HeartIcon, text: t('sponsor_on_github'), to: () => BrowserOpenURL('https://github.com/sponsors/selvindev'), selected: false },
        { icon: CheckBadgeIcon, text: t('follow_on_twitter'), to: () => BrowserOpenURL('https://twitter.com/varlyapp'), selected: false },
        { icon: null, text: t('workspace'), to: '', selected: false },
        { icon: FolderOpenIcon, text: t('recent_projects'), to: 'start', selected: route.name === 'start' },
        { icon: DocumentPlusIcon, text: t('start_new_project'), to: startNewProjectAction(), selected: false },
    ]
}

const buttons = computed(() => {
    if (props.links?.length) {
        let added = false
        presets.links.map((link) => {
            if (link.text === t('project')) {
                added = true
            }
        })

        if (!added) presets.links.push({ icon: null, text: t('project'), to: '', selected: false })
    }

    return [...presets.links, ...(props.links || [])]
})
</script>

<template>
    <aside
        class="h-full overflow-auto scrollbar-none border-r border-slate-900 border-opacity-10 dark:border-slate-50 dark:border-opacity-10">
        <nav v-if="buttons"
            class="p-4 xl:p-8 text-left text-sm text-slate-900 text-opacity-80 dark:text-slate-50 dark:text-opacity-80;">
            <ul>
                <li v-for="(link, i) in buttons" :key="i">
                    <SidebarButton v-if="link.icon" :text="link.text" :to="link.to" :selected="link.selected"
                        :icon="link.icon" />
                    <h2 v-else class="mt-6 px-4 text-xs opacity-60 dark:opacity-40 font-bold">{{ link.text }}</h2>
                </li>
            </ul>
        </nav>
    </aside>
</template>
