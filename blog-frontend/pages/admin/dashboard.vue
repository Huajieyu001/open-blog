<template>
  <div>
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
      <!-- Stats Cards -->
      <div class="bg-white rounded-lg shadow-md p-6">
        <div class="flex items-center">
          <div class="p-3 rounded-full bg-blue-100 text-blue-600">
            <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">文章总数</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.totalArticles }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-md p-6">
        <div class="flex items-center">
          <div class="p-3 rounded-full bg-green-100 text-green-600">
            <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">已发布</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.publishedArticles }}</p>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-md p-6">
        <div class="flex items-center">
          <div class="p-3 rounded-full bg-yellow-100 text-yellow-600">
            <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
            </svg>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600">总阅读量</p>
            <p class="text-2xl font-semibold text-gray-900">{{ stats.totalViews }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Recent Articles -->
    <div class="bg-white rounded-lg shadow-md">
      <div class="px-6 py-4 border-b">
        <h2 class="text-lg font-semibold text-gray-900">最近文章</h2>
      </div>
      <div class="p-6">
        <div v-if="loading" class="text-center py-4">
          <div class="inline-block animate-spin rounded-full h-6 w-6 border-4 border-gray-300 border-t-primary-600"></div>
        </div>
        <div v-else-if="recentArticles.length" class="space-y-4">
          <div
            v-for="article in recentArticles"
            :key="article.id"
            class="flex items-center justify-between p-4 bg-gray-50 rounded-lg"
          >
            <div>
              <h3 class="font-medium text-gray-900">{{ article.title }}</h3>
              <p class="text-sm text-gray-500">
                {{ article.category?.name || '未分类' }} · {{ formatDate(article.created_at) }}
              </p>
            </div>
            <div class="flex items-center space-x-2">
              <span
                :class="[
                  'px-2 py-1 text-xs rounded-full',
                  article.status === 1 ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'
                ]"
              >
                {{ article.status === 1 ? '已发布' : '草稿' }}
              </span>
              <NuxtLink
                :to="`/admin/article/edit/${article.id}`"
                class="text-primary-600 hover:text-primary-700 text-sm"
              >
                编辑
              </NuxtLink>
            </div>
          </div>
        </div>
        <div v-else class="text-center py-4 text-gray-500">
          暂无文章
        </div>
      </div>
    </div>

    <!-- Quick Actions -->
    <div class="mt-8 grid grid-cols-1 md:grid-cols-2 gap-6">
      <NuxtLink
        to="/admin/article/create"
        class="bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow flex items-center"
      >
        <div class="p-3 rounded-full bg-primary-100 text-primary-600">
          <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
          </svg>
        </div>
        <div class="ml-4">
          <p class="font-medium text-gray-900">写新文章</p>
          <p class="text-sm text-gray-500">创建一篇新的博客文章</p>
        </div>
      </NuxtLink>

      <NuxtLink
        to="/admin/article/list"
        class="bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow flex items-center"
      >
        <div class="p-3 rounded-full bg-blue-100 text-blue-600">
          <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16" />
          </svg>
        </div>
        <div class="ml-4">
          <p class="font-medium text-gray-900">管理文章</p>
          <p class="text-sm text-gray-500">查看和管理所有文章</p>
        </div>
      </NuxtLink>
    </div>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  layout: false,
})

const api = useApi()

const loading = ref(true)
const recentArticles = ref<any[]>([])
const stats = reactive({
  totalArticles: 0,
  publishedArticles: 0,
  totalViews: 0,
})

const fetchData = async () => {
  loading.value = true
  try {
    // 获取文章列表
    const res: any = await api.admin.getArticles({ page: 1, page_size: 5 })
    recentArticles.value = res.list || []
    stats.totalArticles = res.total || 0

    // 计算统计数据
    if (res.list) {
      stats.publishedArticles = res.list.filter((a: any) => a.status === 1).length
      stats.totalViews = res.list.reduce((sum: number, a: any) => sum + (a.view_count || 0), 0)
    }
  } catch (error) {
    console.error('Failed to fetch dashboard data:', error)
  } finally {
    loading.value = false
  }
}

const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  })
}

onMounted(() => {
  fetchData()
})
</script>
