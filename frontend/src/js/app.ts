import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createI18n } from 'vue-i18n'

import piniaPersistedState from 'pinia-plugin-persistedstate'

import { useTheme } from '@utils/Theme'
import router from '@root/router'
import App from '@layouts/App.vue'

const i18n = createI18n()
const pinia = createPinia()
pinia.use(piniaPersistedState)

const theme = useTheme()
theme.run()

createApp(App)
  .use(i18n)
  .use(router)
  .use(pinia)
  .mount('#app')
