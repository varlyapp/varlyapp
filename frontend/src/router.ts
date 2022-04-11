import { createRouter, createWebHashHistory } from 'vue-router'

import Start from './pages/Start.vue'
import ArtworkLayers from './pages/artwork/Layers.vue'
import ArtworkBuild from './pages/artwork/Build.vue'
import ArtworkRun from './pages/artwork/Run.vue'

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
    name: 'artwork.build',
    path: '/artwork/build',
    component: ArtworkBuild,
  },
  {
    name: 'artwork.run',
    path: '/artwork/run',
    component: ArtworkRun,
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
});

export default router;
