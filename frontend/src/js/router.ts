import { createRouter, createWebHashHistory } from 'vue-router'

import Start from '@root/pages/Start.vue'
import CollectionStart from '@root/pages/collections/StartCollection.vue'
import CollectionLoadLayers from './pages/collections/LoadLayers.vue'
import CollectionGenerate from './pages/collections/GenerateCollection.vue'

const routes = [
  {
    name: 'start',
    path: '/',
    component: Start,
  },
  {
    name: 'collections.start',
    path: '/collections/start',
    component: CollectionStart,
  },
  {
    name: 'collections.loadLayers',
    path: '/collections/load-layers',
    component: CollectionLoadLayers,
  },
  {
    name: 'collections.generate',
    path: '/collections/generate',
    component: CollectionGenerate,
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
});

export default router;
