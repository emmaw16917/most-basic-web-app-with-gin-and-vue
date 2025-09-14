import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router/index.js' //导入路由实例

createApp(App)
    .use(router)//使用路由
    .mount('#app')
