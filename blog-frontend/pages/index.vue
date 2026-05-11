<template>
  <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <!-- Hero Section -->
    <div class="text-center mb-12">
      <h1 class="text-4xl font-bold text-gray-900 mb-4">Welcome to My Blog</h1>
      <p class="text-lg text-gray-600">分享技术、生活与思考</p>
    </div>

    <!-- Categories -->
    <div v-if="categories.length" class="mb-8">
      <div class="flex flex-wrap gap-2 justify-center">
        <button
          v-for="cat in categories"
          :key="cat.id"
          @click="selectedCategory = cat.id"
          :class="[
            'px-4 py-2 rounded-full text-sm font-medium transition-colors',
            selectedCategory === cat.id
              ? 'bg-primary-600 text-white'
              : 'bg-white text-gray-700 hover:bg-gray-100 border'
          ]"
        >
          {{ cat.name }}
        </button>
        <button
          @click="selectedCategory = null"
          :class="[
            'px-4 py-2 rounded-full text-sm font-medium transition-colors',
            selectedCategory === null
              ? 'bg-primary-600 text-white'
              : 'bg-white text-gray-700 hover:bg-gray-100 border'
          ]"
        >
          全部
        </button>
      </div>
    </div>

    <!-- Article List -->
    <div v-if="loading" class="text-center py-12">
      <div class="inline-block animate-spin rounded-full h-8 w-8 border-4 border-gray-300 border-t-primary-600"></div>
      <p class="mt-4 text-gray-600">加载中...</p>
    </div>

    <div v-else-if="articles.length" class="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
      <article
        v-for="article in articles"
        :key="article.id"
        class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow"
      >
        <!-- Cover Image -->
        <div v-if="article.cover_image" class="h-48 overflow-hidden">
          <img
            :src="article.cover_image"
            :alt="article.title"
            class="w-full h-full object-cover hover:scale-105 transition-transform duration-300"
          />
        </div>
        <div v-else class="h-48 bg-gradient-to-br from-primary-400 to-primary-600 flex items-center justify-center">
          <span class="text-white text-2xl font-bold">{{ article.title.charAt(0) }}</span>
        </div>

        <!-- Content -->
        <div class="p-6">
          <div class="flex items-center text-sm text-gray-500 mb-2">
            <span v-if="article.category">{{ article.category.name }}</span>
            <span v-if="article.category && article.publish_time" class="mx-2">·</span>
            <span v-if="article.publish_time">{{ formatDate(article.publish_time) }}</span>
          </div>

          <h2 class="text-xl font-semibold text-gray-900 mb-2 line-clamp-2">
            <NuxtLink :to="`/article/${article.id}`" class="hover:text-primary-600">
              {{ article.title }}
            </NuxtLink>
          </h2>

          <p v-if="article.summary" class="text-gray-600 text-sm mb-4 line-clamp-3">
            {{ article.summary }}
          </p>

          <div class="flex items-center justify-between">
            <div class="flex flex-wrap gap-1">
              <span
                v-for="tag in article.tags?.slice(0, 3)"
                :key="tag.id"
                class="px-2 py-1 bg-gray-100 text-gray-600 text-xs rounded"
              >
                {{ tag.name }}
              </span>
            </div>
            <span class="text-sm text-gray-500">{{ article.view_count }} 阅读</span>
          </div>
        </div>
      </article>
    </div>

    <div v-else class="text-center py-12">
      <p class="text-gray-600">暂无文章</p>
    </div>

    <!-- Pagination -->
    <div v-if="totalPages > 1" class="mt-8 flex justify-center">
      <nav class="flex space-x-2">
        <button
          v-for="page in totalPages"
          :key="page"
          @click="currentPage = page"
          :class="[
            'px-4 py-2 rounded-lg text-sm font-medium transition-colors',
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
</template>

<script setup lang="ts">
definePageMeta({
  layout: 'default',
})

const api = useApi()

const articles = ref<any[]>([])
const categories = ref<any[]>([])
const loading = ref(true)
const currentPage = ref(1)
const totalPages = ref(1)
const selectedCategory = ref<number | null>(null)

const fetchArticles = async () => {
  loading.value = true
  try {
    const params: any = {
      page: currentPage.value,
      page_size: 9,
    }
    if (selectedCategory.value) {
      params.category_id = selectedCategory.value
    }
    const res: any = await api.pub.getArticles(params)
    articles.value = res.list || []
    totalPages.value = Math.ceil((res.total || 0) / 9)
  } catch (error) {
    console.error('Failed to fetch articles:', error)
  } finally {
    loading.value = false
  }
}

const fetchCategories = async () => {
  try {
    const res: any = await api.pub.getCategories()
    categories.value = res || []
  } catch (error) {
    console.error('Failed to fetch categories:', error)
  }
}

const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  })
}

watch([currentPage, selectedCategory], () => {
  fetchArticles()
})

onMounted(() => {
  fetchCategories()
  fetchArticles()
})
</script>
