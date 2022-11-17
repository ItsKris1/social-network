import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import VueToast from '@meforma/vue-toaster'
// import { VOffline } from 'v-offline';
// import 'vue-toast-notification/dist/theme-sugar.css';
import './assets/css/main.css'

const app = createApp(App).use(store).use(router).use(VueToast)
// app.component('VOffline', VOffline)

app.mount('#app')
