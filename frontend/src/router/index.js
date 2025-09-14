import {createRouter, createWebHistory} from 'vue-router'
import HomePage from '../components/HomePage.vue'
import ViewPage from '../components/ViewPage.vue'
import EditPage from '../components/EditPage.vue'

// 路由规则：路径与组件的映射关系
const routes = [
    { path: '/', component: HomePage },
    { path: '/view/:title', component: ViewPage},
    { path: '/edit/:title', component: EditPage}
]
// 创建路由实例
const router = createRouter({
    history: createWebHistory(),
    routes
})
// 导出路由实例
export default router