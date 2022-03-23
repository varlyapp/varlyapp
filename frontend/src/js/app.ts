import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPersistedState from 'pinia-plugin-persistedstate'

import { useTheme } from '@utils/Theme'
import router from '@root/router'
import App from '@layouts/App.vue'
import varly from '@plugins/varly'

const pinia = createPinia()
pinia.use(piniaPersistedState)

const theme = useTheme()
theme.run()

createApp(App)
  .use(router)
  .use(pinia)
  .use(varly, { router })
  .mount('#app')
