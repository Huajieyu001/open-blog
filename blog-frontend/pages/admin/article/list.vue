<template>
  <div>
    <div class="bg-white rounded-lg shadow-md">
      <!-- Header -->
      <div class="px-6 py-4 border-b flex justify-between items-center">
        <h2 class="text-lg font-semibold text-gray-900">文章列表</h2>
        <NuxtLink
          to="/admin/article/create"
          class="px-4 py-2 bg-primary-600 text-white rounded-lg hover:bg-primary-700 transition-colors"
        >
          写文章
        </NuxtLink>
      </div>

      <!-- Filters -->
      <div class="px-6 py-4 border-b bg-gray-50">
        <div class="flex flex-wrap gap-4">
          <select
            v-model="filters.status"
            class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
          >
            <option value="">全部状态</option>
            <option value="0">草稿</option>
            <option value="1">已发布</option>
          </select>
        </div>
      </div>

      <!-- Table -->
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                标题
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                分类
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                状态
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                阅读量
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                创建时间
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                操作
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-if="loading">
              <td colspan="6" class="px-6 py-4 text-center text-gray-500">
                加载中...
              </td>
            </tr>
            <tr v-else-if="articles.length === 0">
              <td colspan="6" class="px-6 py-4 text-center text-gray-500">
                暂无文章
              </td>
            </tr>
            <tr v-for="article in articles" :key="article.id" class="hover:bg-gray-50">
              <td class="px-6 py-4">
                <div class="flex items-center">
                  <div v-if="article.cover_image" class="h-10 w-10 flex-shrink-0 mr-3">
                    <img
                      :src="article.cover_image"
                      :alt="article.title"
                      class="h-10 w-10 rounded object-cover"
                    />
                  </div>
                  <div>
                    <div class="text-sm font-medium text-gray-900">{{ article.title }}</div>
                    <div v-if="article.is_top" class="text-xs text-orange-600">置顶</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">
                {{ article.category?.name || '-' }}
              </td>
              <td class="px-6 py-4">
                <span
                  :class="[
                    'px-2 py-1 text-xs rounded-full',
                    article.status === 1 ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'
                  ]"
                >
                  {{ article.status === 1 ? '已发布' : '草稿' }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">
                {{ article.view_count || 0 }}
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">
                {{ formatDate(article.created_at) }}
              </td>
              <td class="px-6 py-4 text-right text-sm font-medium space-x-2">
                <NuxtLink
                  :to="`/admin/article/edit/${article.id}`"
                  class="text-primary-600 hover:text-primary-700"
                >
                  编辑
                </NuxtLink>
                <button
                  @click="handleDelete(article)"
                  class="text-red-600 hover:text-red-700"
                >
                  删除
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div v-if="totalPages > 1" class="px-6 py-4 border-t flex justify-center">
        <nav class="flex space-x-2">
          <button
            v-for="page in totalPages"
            :key="page"
            @click="currentPage = page"
            :class="[
              'px-3 py-1 rounded text-sm',
              currentPage === page
                ? 'bg-primary-600 text-white'
                : 'bg-white text-gray-700 hover:bg-gray-100 border'
            ]"
          >
            {{ page }}
          </button>
        </nav>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  layout: false,
})

const api = useApi()

const articles = ref<any[]>([])
const loading = ref(true)
const currentPage = ref(1)
const totalPages = ref(1)
const filters = reactive({
  status: '',
})

const fetchArticles = async () => {
  loading.value = true
  try {
    const params: any = {
      page: currentPage.value,
      page_size: 10,
    }
    if (filters.status !== '') {
      params.status = Number(filters.status)
    }
    const res: any = await api.admin.getArticles(params)
    articles.value = res.list || []
    totalPages.value = Math.ceil((res.total || 0) / 10)
  } catch (error) {
    console.error('Failed to fetch articles:', error)
  } finally {
    loading.value = false
  }
}

const handleDelete = async (article: any) => {
  if (!confirm(`确定要删除文章 "${article.title}" 吗？`)) {
    return
  }

  try {
    await api.admin.deleteArticle(article.id)
    fetchArticles()
  } catch (error) {
    console.error('Failed to delete article:', error)
    alert('删除失败')
  }
}

const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  })
}

watch([currentPage, filters], () => {
  fetchArticles()
})

onMounted(() => {
  fetchArticles()
})
</script>
