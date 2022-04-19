import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createI18n } from 'vue-i18n'
import piniaPersistedState from 'pinia-plugin-persistedstate'

import router from '@/router'
import lang from '@/lang'
import App from '@/layouts/App.vue'

import 'animate.css'
import '@/assets/app.css'

const pinia = createPinia()
pinia.use(piniaPersistedState)

const intl = createI18n(lang)

createApp(App)
  .use(router)
  .use(pinia)
  .use(intl)
  .mount('#app')
