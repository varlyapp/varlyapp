import { createRouter, createWebHashHistory } from 'vue-router'

import Welcome from '@pages/Welcome.vue'
import NewCollection from '@pages/art/NewCollection.vue'

const routes = [
  {
    name: 'welcome',
    path: '/',
    component: Welcome,
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
