<template>
  <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <div v-if="loading" class="text-center py-12">
      <div class="inline-block animate-spin rounded-full h-8 w-8 border-4 border-gray-300 border-t-primary-600"></div>
      <p class="mt-4 text-gray-600">加载中...</p>
    </div>

    <article v-else-if="article" class="bg-white rounded-lg shadow-md overflow-hidden">
      <!-- Cover Image -->
      <div v-if="article.cover_image" class="h-64 md:h-96 overflow-hidden">
        <img
          :src="article.cover_image"
          :alt="article.title"
          class="w-full h-full object-cover"
        />
      </div>

      <!-- Content -->
      <div class="p-6 md:p-8">
        <!-- Header -->
        <header class="mb-8">
          <h1 class="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
            {{ article.title }}
          </h1>

          <div class="flex items-center text-sm text-gray-500 space-x-4">
            <span v-if="article.category" class="flex items-center">
              <svg class="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
              </svg>
              {{ article.category.name }}
            </span>

            <span v-if="article.publish_time" class="flex items-center">
              <svg class="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
              </svg>
              {{ formatDate(article.publish_time) }}
            </span>

            <span class="flex items-center">
              <svg class="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
              </svg>
              {{ article.view_count }} 阅读
            </span>
          </div>

          <!-- Tags -->
          <div v-if="article.tags?.length" class="mt-4 flex flex-wrap gap-2">
            <NuxtLink
              v-for="tag in article.tags"
              :key="tag.id"
              :to="`/category?tag=${tag.id}`"
              class="px-3 py-1 bg-gray-100 text-gray-600 text-sm rounded-full hover:bg-primary-100 hover:text-primary-700 transition-colors"
            >
              {{ tag.name }}
            </NuxtLink>
          </div>
        </header>

        <!-- Article Content -->
        <div
          class="prose prose-lg max-w-none"
          v-html="article.content_html"
        ></div>

        <!-- Footer -->
        <footer class="mt-8 pt-8 border-t">
          <div class="flex justify-between items-center">
            <NuxtLink
              to="/"
              class="text-primary-600 hover:text-primary-700 flex items-center"
            >
              <svg class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
              </svg>
              返回首页
            </NuxtLink>

            <div class="text-sm text-gray-500">
              最后更新: {{ formatDate(article.updated_at) }}
            </div>
          </div>
        </footer>
      </div>
    </article>

    <div v-else class="text-center py-12">
      <p class="text-gray-600">文章不存在</p>
      <NuxtLink to="/" class="text-primary-600 hover:text-primary-700 mt-4 inline-block">
        返回首页
      </NuxtLink>
    </div>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  layout: 'default',
})

const route = useRoute()
const api = useApi()

const article = ref<any>(null)
const loading = ref(true)

const fetchArticle = async () => {
  loading.value = true
  try {
    const id = Number(route.params.id)
    const res: any = await api.pub.getArticle(id)
    article.value = res
  } catch (error) {
    console.error('Failed to fetch article:', error)
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
  fetchArticle()
})
</script>
