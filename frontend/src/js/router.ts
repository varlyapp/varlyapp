import { createRouter, createWebHashHistory } from 'vue-router'

import Start from '@root/pages/Start.vue'
import NewCollection from '@pages/art/NewCollection.vue'

const routes = [
  {
    name: 'start',
    path: '/',
    component: Start,
  },
  {
    name: 'art.collection.new',
    path: '/art/collection/new',
    component: NewCollection,
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
});

export default router;
