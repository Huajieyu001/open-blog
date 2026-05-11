import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:8080/api',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
})

// 请求拦截器
api.interceptors.request.use(
  (config) => {
    if (import.meta.client) {
      // 登录接口不添加 token
      if (!config.url?.includes('/admin/login')) {
        const token = localStorage.getItem('token')
        if (token) {
          config.headers.Authorization = `Bearer ${token}`
        }
      }
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
api.interceptors.response.use(
  (response) => {
    return response.data
  },
  (error) => {
    if (error.response) {
      const { status, data } = error.response
      if (status === 401) {
        if (import.meta.client) {
          localStorage.removeItem('token')
          localStorage.removeItem('user')
          window.location.href = '/admin/login'
        }
      }
      return Promise.reject(data || error.message)
    }
    return Promise.reject(error.message)
  }
)

export const useApi = () => {
  return {
    // 公开接口
    pub: {
      // 文章列表
      getArticles: (params?: {
        page?: number
        page_size?: number
        category_id?: number
        tag_id?: number
        keyword?: string
      }) => api.get('/pub/articles', { params }),

      // 文章详情
      getArticle: (id: number) => api.get(`/pub/articles/${id}`),

      // 按slug获取文章
      getArticleBySlug: (slug: string) => api.get(`/pub/articles/slug/${slug}`),

      // 分类列表
      getCategories: () => api.get('/pub/categories'),

      // 标签列表
      getTags: () => api.get('/pub/tags'),
    },

    // 管理接口
    admin: {
      // 登录
      login: (data: { username: string; password: string }) =>
        api.post('/admin/login', data),

      // 文章管理
      getArticles: (params?: {
        page?: number
        page_size?: number
        status?: number
      }) => api.get('/admin/articles', { params }),

      createArticle: (data: {
        title: string
        slug?: string
        summary?: string
        content_md: string
        cover_image?: string
        status?: number
        is_top?: number
        category_id?: number
        tag_ids?: number[]
      }) => api.post('/admin/articles', data),

      updateArticle: (id: number, data: {
        title?: string
        slug?: string
        summary?: string
        content_md?: string
        cover_image?: string
        status?: number
        is_top?: number
        category_id?: number
        tag_ids?: number[]
      }) => api.put(`/admin/articles/${id}`, data),

      deleteArticle: (id: number) => api.delete(`/admin/articles/${id}`),

      // 分类管理
      getCategories: () => api.get('/admin/categories'),

      createCategory: (data: {
        name: string
        slug: string
        description?: string
        sort_order?: number
      }) => api.post('/admin/categories', data),

      updateCategory: (id: number, data: {
        name?: string
        slug?: string
        description?: string
        sort_order?: number
      }) => api.put(`/admin/categories/${id}`, data),

      deleteCategory: (id: number) => api.delete(`/admin/categories/${id}`),

      // 标签管理
      getTags: () => api.get('/admin/tags'),

      createTag: (data: {
        name: string
        slug: string
      }) => api.post('/admin/tags', data),

      updateTag: (id: number, data: {
        name?: string
        slug?: string
      }) => api.put(`/admin/tags/${id}`, data),

      deleteTag: (id: number) => api.delete(`/admin/tags/${id}`),

      // 文件上传
      upload: (file: File) => {
        const formData = new FormData()
        formData.append('file', file)
        return api.post('/admin/upload', formData, {
          headers: {
            'Content-Type': 'multipart/form-data',
          },
        })
      },
    },
  }
}
