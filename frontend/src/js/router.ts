import { createRouter, createWebHashHistory } from 'vue-router'

import Start from '@root/pages/Start.vue'
import ArtworkLayers from '@root/pages/artwork/Layers.vue'
import ArtworkCollection from '@root/pages/artwork/Collection.vue'
import ArtworkBuild from '@root/pages/artwork/Build.vue'

const routes = [
  {
    name: 'start',
    path: '/',
    component: Start,
  },
  {
    name: 'artwork.layers',
    path: '/artwork/layers',
    component: ArtworkLayers
  },
  {
    name: 'artwork.collection',
    path: '/artwork/collection',
    component: ArtworkCollection,
  },
  {
    name: 'artwork.build',
    path: '/artwork/build',
    component: ArtworkBuild,
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
});

export default router;
