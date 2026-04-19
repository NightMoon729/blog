import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from "element-plus"
import 'element-plus/dist/index.css'

const app = createApp(App) //先创建app实例
app.use(ElementPlus)       //注册element plus
app.mount('#app')