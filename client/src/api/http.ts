import { clearUserData } from '@/constants/helper'
import { message } from 'antd'
import axios, { AxiosInstance } from 'axios'
import { throttle } from 'lodash'
import { stringify } from 'qs'
import { renewToken } from '.'
import { getLocalStorage } from '@/utils/xLocalstorage'

export type ApiResponse<T = null> = {
  data: {
    data: T
    message: string
    status: number
    success: boolean
  }
}

const throttleAlert = (msg: string) => throttle(message.error(msg), 1500, { trailing: false })

const instance: AxiosInstance = axios.create({ timeout: 60000 })

instance.interceptors.response.use(
  async response => {
    if (!response?.data.success && response?.data.status !== 401) {
      throttleAlert(response?.data.message)
    }

    if (response?.data.status === 401) {
      const res = await renewToken({ refresh_token: getLocalStorage('refresh_token') })
      if (res?.data.status === 401) return invalidSession()

      document.cookie = `access_token=${res?.data.data.access_token};path=/;domain=${import.meta.env.VITE_DOMAIN};expires=${res?.data.data.access_token_expires_at}`
      const originalRequest = { ...response.config }
      return instance(originalRequest)
    }

    return response
  },
  error => {
    return Promise.reject(error)
  }
)

instance.interceptors.request.use(
  config => {
    return config
  },
  error => Promise.reject(error)
)

const invalidSession = () => {
  throttleAlert('Session is not valid, need to login again')
  clearUserData()
  window.location.href = '/'

  return Promise.reject(new Error('Session expired'))
}

const upload = async <T>(url: string, data): Promise<ApiResponse<T>> => {
  const form = new FormData()

  form.append('file', data)

  //prettier-ignore
  const response = await instance.post(`${import.meta.env.VITE_DOMAIN}${url}`, form, {
    withCredentials: true
  })

  return response as any
}

const post = async <T>(url: string, data = {}): Promise<ApiResponse<T>> =>
  instance.post(`${import.meta.env.VITE_DOMAIN}${url}`, data, {
    withCredentials: true
  })

const get = async <T>(url: string, data = {}): Promise<ApiResponse<T>> =>
  instance.get(
    `${import.meta.env.VITE_DOMAIN}${url}${stringify(data) ? '?' + stringify(data) : ''}`,
    {
      withCredentials: true
    }
  )

export { post, get, upload }
