import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createI18n } from 'vue-i18n'

import piniaPersistedState from 'pinia-plugin-persistedstate'

import { useTheme } from '@utils/Theme'
import router from '@root/router'
import lang from '@root/lang'
import App from '@layouts/App.vue'

import 'animate.css'

const intl = createI18n(lang)

const pinia = createPinia()
pinia.use(piniaPersistedState)

const theme = useTheme()
theme.run()

createApp(App)
  .use(intl)
  .use(router)
  .use(pinia)
  .mount('#app')
