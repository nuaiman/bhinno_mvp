import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPersist from 'pinia-plugin-persistedstate'
import { globalLoader } from 'vue-global-loader'

import App from './App.vue'
import router from './router'

const app = createApp(App)

// Create Pinia instance
const pinia = createPinia()

// Apply persistence plugin
pinia.use(piniaPersist)

// Register Pinia with the app
app.use(pinia)

// Register router
app.use(router)

// Setup global loader (you can customize colors)
app.use(globalLoader, {
  backgroundColor: '#000',
  foregroundColor: '#fff',
  backgroundOpacity: 0.7,
  backgroundBlur: 4,
  screenReaderMessage: 'Loading, please wait...',
})

app.mount('#app')
