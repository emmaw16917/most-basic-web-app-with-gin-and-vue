<template>
    <div class="home">
        <h1>Wiki 首页</h1>
        <div>
            <input 
            v-model="title"
            placeholder="请输入标题（仅字母数字）"
            class="title-input"
            />
            <button @click="goToView">查看页面</button>
            <button @click="goToEdit">编辑页面</button>
        </div>
    </div>
</template>

<script setup>

import { ref } from 'vue'
import { useRouter } from 'vue-router'

const title = ref('') //创建响应式变量title
const router = useRouter() //获取路由实例

//定义goToView函数
const goToView = () => {
    if (title.value.trim() === '') {
        alert('标题不能为空')
        return
    }
    console.log('跳转到查看页面:', title.value.trim()) // 调试日志
    router.push(`/view/${title.value.trim()}`)
}
//定义goToEdit函数
const goToEdit = () => {
    if (title.value.trim() === '') {
        alert('标题不能为空')
        return
    }
    router.push(`/edit/${title.value.trim()}`)
}

</script>

<style scoped>
    .home {
        text-align: center;
        margin-top: 50px;
    }
    .title-input {
        padding: 8px;
        width: 300px;
        margin-right: 10px;
    }
    button {
        padding: 8px 16px;
        margin-right: 10px;
        cursor: pointer;
    }
</style>