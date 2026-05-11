<template>
  <div class="max-w-4xl">
    <form @submit.prevent="handleSubmit">
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- Main Content -->
        <div class="lg:col-span-2 space-y-6">
          <div class="bg-white rounded-lg shadow-md p-6">
            <h2 class="text-lg font-semibold text-gray-900 mb-4">{{ isEdit ? '编辑文章' : '写文章' }}</h2>

            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">标题</label>
                <input
                  v-model="form.title"
                  type="text"
                  required
                  class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
                  placeholder="请输入文章标题"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Slug (可选)</label>
                <input
                  v-model="form.slug"
                  type="text"
                  class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
                  placeholder="URL 友好标识，如 my-article"
                />
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">摘要</label>
                <textarea
                  v-model="form.summary"
                  rows="3"
                  class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
                  placeholder="文章摘要（可选）"
                ></textarea>
              </div>

              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">内容 (Markdown)</label>
                <textarea
                  v-model="form.content_md"
                  rows="20"
                  required
                  class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500 font-mono"
                  placeholder="请输入文章内容，支持 Markdown 语法"
                ></textarea>
              </div>
            </div>
          </div>
        </div>

        <!-- Sidebar -->
        <div class="space-y-6">
          <!-- Publish Settings -->
          <div class="bg-white rounded-lg shadow-md p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">发布设置</h3>

            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">状态</label>
                <select
                  v-model="form.status"
                  class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
                >
                  <option :value="0">草稿</option>
                  <option :value="1">发布</option>
                </select>
              </div>

              <div>
                <label class="flex items-center">
                  <input
                    v-model="form.is_top"
                    type="checkbox"
                    :true-value="1"
                    :false-value="0"
                    class="rounded border-gray-300 text-primary-600 focus:ring-primary-500"
                  />
                  <span class="ml-2 text-sm text-gray-700">置顶文章</span>
                </label>
              </div>
            </div>
          </div>

          <!-- Category -->
          <div class="bg-white rounded-lg shadow-md p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">分类</h3>

            <select
              v-model="form.category_id"
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
            >
              <option :value="null">未分类</option>
              <option
                v-for="category in categories"
                :key="category.id"
                :value="category.id"
              >
                {{ category.name }}
              </option>
            </select>
          </div>

          <!-- Tags -->
          <div class="bg-white rounded-lg shadow-md p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">标签</h3>

            <div class="space-y-2 max-h-48 overflow-y-auto">
              <label
                v-for="tag in tags"
                :key="tag.id"
                class="flex items-center"
              >
                <input
                  v-model="form.tag_ids"
                  type="checkbox"
                  :value="tag.id"
                  class="rounded border-gray-300 text-primary-600 focus:ring-primary-500"
                />
                <span class="ml-2 text-sm text-gray-700">{{ tag.name }}</span>
              </label>
            </div>
          </div>

          <!-- Cover Image -->
          <div class="bg-white rounded-lg shadow-md p-6">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">封面图</h3>

            <div class="space-y-4">
              <input
                v-model="form.cover_image"
                type="text"
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
                placeholder="输入图片 URL"
              />

              <div v-if="form.cover_image" class="aspect-video rounded-lg overflow-hidden bg-gray-100">
                <img
                  :src="form.cover_image"
                  alt="Cover preview"
                  class="w-full h-full object-cover"
                />
              </div>
            </div>
          </div>

          <!-- Actions -->
          <div class="bg-white rounded-lg shadow-md p-6">
            <div class="space-y-3">
              <button
                type="submit"
                :disabled="submitting"
                class="w-full py-3 px-4 bg-primary-600 text-white rounded-lg font-medium hover:bg-primary-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
              >
                <span v-if="submitting">保存中...</span>
                <span v-else>{{ isEdit ? '更新文章' : '保存文章' }}</span>
              </button>

              <NuxtLink
                to="/admin/article/list"
                class="block w-full py-3 px-4 bg-gray-100 text-gray-700 rounded-lg font-medium text-center hover:bg-gray-200 transition-colors"
              >
                取消
              </NuxtLink>
            </div>
          </div>
        </div>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  layout: false,
})

const route = useRoute()
const router = useRouter()
const api = useApi()

const isEdit = computed(() => !!route.params.id)
const articleId = computed(() => Number(route.params.id))

const categories = ref<any[]>([])
const tags = ref<any[]>([])
const submitting = ref(false)

const form = reactive({
  title: '',
  slug: '',
  summary: '',
  content_md: '',
  cover_image: '',
  status: 0,
  is_top: 0,
  category_id: null as number | null,
  tag_ids: [] as number[],
})

const fetchArticle = async () => {
  if (!isEdit.value) return

  try {
    const res: any = await api.admin.getArticles({ page: 1, page_size: 100 })
    const article = res.list?.find((a: any) => a.id === articleId.value)
    if (article) {
      form.title = article.title
      form.slug = article.slug || ''
      form.summary = article.summary || ''
      form.content_md = article.content_md || ''
      form.cover_image = article.cover_image || ''
      form.status = article.status
      form.is_top = article.is_top
      form.category_id = article.category_id
      form.tag_ids = article.tags?.map((t: any) => t.id) || []
    }
  } catch (error) {
    console.error('Failed to fetch article:', error)
  }
}

const fetchCategories = async () => {
  try {
    const res: any = await api.admin.getCategories()
    categories.value = res || []
  } catch (error) {
    console.error('Failed to fetch categories:', error)
  }
}

const fetchTags = async () => {
  try {
    const res: any = await api.admin.getTags()
    tags.value = res || []
  } catch (error) {
    console.error('Failed to fetch tags:', error)
  }
}

const handleSubmit = async () => {
  submitting.value = true
  try {
    if (isEdit.value) {
      await api.admin.updateArticle(articleId.value, { ...form })
    } else {
      await api.admin.createArticle({ ...form })
    }
    router.push('/admin/article/list')
  } catch (error) {
    console.error('Failed to save article:', error)
    alert('保存失败')
  } finally {
    submitting.value = false
  }
}

onMounted(() => {
  fetchCategories()
  fetchTags()
  fetchArticle()
})
</script>
