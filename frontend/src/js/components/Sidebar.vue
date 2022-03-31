<script setup lang="ts">
import type { FunctionalComponent, PropType } from 'vue'
import SidebarButton from '@components/SidebarButton.vue'

type Links = Array<{
    text: string
    icon: FunctionalComponent | any
    to: string | Function
    selected: boolean
}>

const props = defineProps({
    links: Array as PropType<Links>,
})
</script>

<template>
    <aside
        class="h-full overflow-auto scrollbar-none border-r border-slate-900 border-opacity-10 dark:border-slate-50 dark:border-opacity-10"
    >
        <nav
            v-if="props.links && props.links.length"
            class="p-4 xl:p-8 mt-8 xl:mt-16 text-left text-sm text-slate-900 text-opacity-80 dark:text-slate-50 dark:text-opacity-80;"
        >
            <ul>
                <li v-for="(link, i) in props.links" :key="i">
                    <SidebarButton
                        v-if="link.text"
                        :text="link.text"
                        :to="link.to"
                        :selected="link.selected"
                    >
                        <component
                            :is="link.icon"
                            class="w-6 mr-2 fill-fuchsia-700 dark:fill-fuchsia-600"
                        />
                    </SidebarButton>
                    <div
                        v-else
                        class="my-4 bg-slate-900 bg-opacity-10 dark:bg-slate-50 dark:bg-opacity-10"
                        style="height: 1pt"
                    ></div>
                </li>
            </ul>
        </nav>
    </aside>
</template>
