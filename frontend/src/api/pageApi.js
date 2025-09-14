import axios from 'axios'

//创建axios实例
const api = axios.create({
  baseURL: 'http://localhost:8080/api', //后端API的基础路径
  timeout: 5000 //请求超时时间
})

//定义API函数
export const getPage = (title) => {
  return api.get(`/page/${title}`)
}
export const savePage = (pageData) => {
  return api.post('/page/save', pageData)//传递JSON数据: {title, body}
}

//导出axios实例
export default api