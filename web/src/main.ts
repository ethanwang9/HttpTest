import {createApp} from 'vue'
import App from './App.vue'
import router from "./router";

// 导入css
import "normalize.css"

const app = createApp(App)
app.use(router)
app.mount('#app')
