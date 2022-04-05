<script setup lang="ts">
import { computed, type FunctionalComponent, type PropType } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { BadgeCheckIcon, DocumentAddIcon, FolderOpenIcon, TranslateIcon } from '@heroicons/vue/solid'
import SidebarButton from '@/components/SidebarButton.vue'
import { useStore } from '@/store'
import { useVarly } from '@/varly'

const varly = useVarly()
const store = useStore()
const route = useRoute()
const router = useRouter()

const intl = useI18n({ useScope: 'global' })
const { t } = intl


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
    if (route.name === 'start') {
        return 'artwork.layers'
    }

    return () => varly.confirmStartNewProject(t(`confirm_start_new_project_title`), t(`confirm_start_new_project_message`))
}

const presets = {
    links: [
        { icon: null, text: 'Support', to: '', selected: false },
        { icon: TranslateIcon, text: t('switch_language'), to: () => {
            console.log('Setting language')
            store.setLocale('es')
        }, selected: false },
        { icon: BadgeCheckIcon, text: t('follow_on_twitter'), to: varly.launchTwitter, selected: false },
        { icon: null, text: 'Workspace', to: '', selected: false },
        { icon: FolderOpenIcon, text: t('recent_projects'), to: 'start', selected: route.name === 'start' },
        { icon: DocumentAddIcon, text: t('start_new_project'), to: startNewProjectAction(), selected: false },
    ]
}

const buttons = computed(() => {
    if (props.links?.length) {
        presets.links.push({ icon: null, text: 'Project', to: '', selected: false },)
    }

    return [...presets.links, ...(props.links || [])]
})
</script>

<template>
    <aside
        class="h-full overflow-auto scrollbar-none border-r border-slate-900 border-opacity-10 dark:border-slate-50 dark:border-opacity-10"
    >
        <nav
            v-if="buttons"
            class="p-4 xl:p-8 mt-4 xl:mt-8 text-left text-sm text-slate-900 text-opacity-80 dark:text-slate-50 dark:text-opacity-80;"
        >
            <ul>
                <li v-for="(link, i) in buttons" :key="i">
                    <SidebarButton
                        v-if="link.icon"
                        :text="link.text"
                        :to="link.to"
                        :selected="link.selected"
                        :icon="link.icon"
                    />
                    <h2 v-else class="mt-8 text-xs opacity-60 dark:opacity-40 font-bold uppercase">{{ link.text }}</h2>
                </li>
            </ul>
        </nav>
    </aside>
</template>
