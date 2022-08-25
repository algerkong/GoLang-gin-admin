// 封装axios
import { useStorage } from '@/hooks'
import axios from 'axios'
import { MessageInfoOptions, MessagePlugin } from 'tdesign-vue-next'

export const axiosInstance = axios.create({
  baseURL: '/api',
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json;charset=UTF-8',
  },
})

axiosInstance.interceptors.request.use(
  (config) => {
    config = {
      ...config,
      headers: {
        Authorization: 'Bearer ' + (useStorage.getStorage('token') || ''),
      },
    }
    return config
  },
  (error: any) => {
    return Promise.reject(error)
  }
)

axiosInstance.interceptors.response.use(
  (response) => {
    return response
  },
  (error: any) => {
    MessagePlugin.error(error.response.data.msg || '请求失败')
    return Promise.reject(error)
  }
)

export default axiosInstance
