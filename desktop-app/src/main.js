import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import VueToast from '@meforma/vue-toaster'
// import 'vue-toast-notification/dist/theme-sugar.css';
import './assets/css/main.css'

createApp(App).use(store).use(router).use(VueToast).mount('#app')
