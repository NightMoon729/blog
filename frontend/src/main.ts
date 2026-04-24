import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from "element-plus"
import 'element-plus/dist/index.css'
import '@/assets/quill.css'
import router from './router'

const app = createApp(App) //先创建app实例
app.use(ElementPlus)       //注册element plus
app.use(router)
app.mount('#app')