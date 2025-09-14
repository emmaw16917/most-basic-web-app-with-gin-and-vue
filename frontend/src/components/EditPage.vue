<template>
    <div class="edit">
        <h1>编辑页面：{{ title }}</h1>
        <textarea
        v-model="body"
        placeholder="请输入内容"
        class="body-textarea"
        ></textarea>
        <button @click="saveCurrentPage">保存页面</button>
        <button @click="goBack">返回首页</button>
        <div v-if="error" class="error">{{ error }}</div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getPage, savePage } from '../api/pageApi'

const router = useRouter() //获取路由实例
const route = useRoute() //获取当前路由信息
const title = route.params.title //获取路由参数中的标题
const body = ref('') //创建响应式变量body
const error = ref('') //创建响应式变量error 用于存储错误信息(初始为空，无错误时不显示)

//生命周期钩子函数 在组件挂载后执行
onMounted(async () => {
    try {
        const response = await getPage(title) //调用getPage函数获取页面内容
        body.value = response.data.body //将获取到的内容赋值给body
    } catch (err) {
        if (err.response && err.response.status === 404) {
            body.value = '' //如果页面不存在，body为空
        } else {
            error.value = '加载页面失败'
        }
    }
})
//定义saveCurrentPage函数
const saveCurrentPage = async () => {
    try {
        await savePage({ title, body: body.value }) //调用savePage函数保存页面内容
        alert('页面保存成功')
        router.push(`/view/${title}`) //保存成功后跳转到查看页面
    } catch (err) {
        error.value = '保存页面失败：' + (err.response?.data?.error || err.message)
    }
}
//定义goBack函数
const goBack = () => {
    router.push('/')
}
</script>

<style scoped>
    .edit {
        width: 800px;
        margin: 50px auto;
    }
    .body-textarea {
        width: 100%;
        height: 300px;
        padding: 10px;
        margin: 20px 0;
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