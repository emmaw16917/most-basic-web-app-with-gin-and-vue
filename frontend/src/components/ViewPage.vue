<template>
    <div class="view">
        <h1>{{ page.title }}</h1>
        <div class="body" v-if="page.body">{{ page.body }}</div>
        <div v-if="error" class="error">{{ error }}</div>
        <button @click="goToEdit">编辑页面</button>
        <button @click="goBack">返回首页</button>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getPage } from '../api/pageApi'

const router = useRouter() //获取路由实例
const route = useRoute() //获取当前路由信息
const title = route.params.title //获取路由参数中的标题
const page = ref({ title: '', body: '' }) //创建响应式变量page 用于存储页面数据
const error = ref('') //创建响应式变量error 用于存储错误信息(初始为空，无错误时不显示)

//生命周期钩子函数 在组件挂载后执行
onMounted(async () => {
    try {
        const response = await getPage(title) //调用getPage函数获取页面内容
        page.value = response.data //将获取到的内容赋值给page
    } catch (err) {
        if (err.response && err.response.status === 404) {
            error.value = '页面不存在' //如果页面不存在，显示错误信息
        } else {
            error.value = '加载页面失败'
            console.error(err)
        }
    }
})

const goToEdit = () => {
    router.push(`/edit/${title}`) //跳转到编辑页面
}
const goBack = () => {
    router.push('/') //跳转到首页
}
</script>

<style scoped>
    .view {
        width: 800px;
        margin: 50px auto;
    }
    .body {
        margin: 20px 0;
        white-space: pre-wrap; /* 保留换行和空格 */
        line-height: 1.6;
    }
    .error {
        color: red;
        margin-top: 10px;
    }
    button {
        padding: 8px 16px;
        margin-right: 10px;
        cursor: pointer;
    }
</style>