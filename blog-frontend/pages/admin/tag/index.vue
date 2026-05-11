<template>
  <div>
    <div class="bg-white rounded-lg shadow-md">
      <!-- Header -->
      <div class="px-6 py-4 border-b flex justify-between items-center">
        <h2 class="text-lg font-semibold text-gray-900">标签管理</h2>
        <button
          @click="showModal = true; resetForm()"
          class="px-4 py-2 bg-primary-600 text-white rounded-lg hover:bg-primary-700 transition-colors"
        >
          添加标签
        </button>
      </div>

      <!-- Table -->
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                名称
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Slug
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                操作
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-if="loading">
              <td colspan="3" class="px-6 py-4 text-center text-gray-500">
                加载中...
              </td>
            </tr>
            <tr v-else-if="tags.length === 0">
              <td colspan="3" class="px-6 py-4 text-center text-gray-500">
                暂无标签
              </td>
            </tr>
            <tr v-for="tag in tags" :key="tag.id" class="hover:bg-gray-50">
              <td class="px-6 py-4 text-sm font-medium text-gray-900">
                {{ tag.name }}
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">
                {{ tag.slug }}
              </td>
              <td class="px-6 py-4 text-right text-sm font-medium space-x-2">
                <button
                  @click="editTag(tag)"
                  class="text-primary-600 hover:text-primary-700"
                >
                  编辑
                </button>
                <button
                  @click="handleDelete(tag)"
                  class="text-red-600 hover:text-red-700"
                >
                  删除
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Modal -->
    <div
      v-if="showModal"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
      @click.self="showModal = false"
    >
      <div class="bg-white rounded-lg shadow-xl max-w-md w-full mx-4">
        <div class="px-6 py-4 border-b">
          <h3 class="text-lg font-semibold text-gray-900">
            {{ editingTag ? '编辑标签' : '添加标签' }}
          </h3>
        </div>

        <form @submit.prevent="handleSubmit" class="px-6 py-4">
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">名称</label>
              <input
                v-model="form.name"
                type="text"
                required
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
                placeholder="标签名称"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">Slug</label>
              <input
                v-model="form.slug"
                type="text"
                required
                class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
                placeholder="URL 友好标识"
              />
            </div>
          </div>

          <div class="mt-6 flex justify-end space-x-3">
            <button
              type="button"
              @click="showModal = false"
              class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition-colors"
            >
              取消
            </button>
            <button
              type="submit"
              :disabled="submitting"
              class="px-4 py-2 bg-primary-600 text-white rounded-lg hover:bg-primary-700 disabled:opacity-50 transition-colors"
            >
              {{ submitting ? '保存中...' : '保存' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  layout: false,
})

const api = useApi()

const tags = ref<any[]>([])
const loading = ref(true)
const showModal = ref(false)
const submitting = ref(false)
const editingTag = ref<any>(null)

const form = reactive({
  name: '',
  slug: '',
})

const fetchTags = async () => {
  loading.value = true
  try {
    const res: any = await api.admin.getTags()
    tags.value = res || []
  } catch (error) {
    console.error('Failed to fetch tags:', error)
  } finally {
    loading.value = false
  }
}

const resetForm = () => {
  editingTag.value = null
  form.name = ''
  form.slug = ''
}

const editTag = (tag: any) => {
  editingTag.value = tag
  form.name = tag.name
  form.slug = tag.slug
  showModal.value = true
}

const handleSubmit = async () => {
  submitting.value = true
  try {
    if (editingTag.value) {
      await api.admin.updateTag(editingTag.value.id, { ...form })
    } else {
      await api.admin.createTag({ ...form })
    }
    showModal.value = false
    fetchTags()
  } catch (error) {
    console.error('Failed to save tag:', error)
    alert('保存失败')
  } finally {
    submitting.value = false
  }
}

const handleDelete = async (tag: any) => {
  if (!confirm(`确定要删除标签 "${tag.name}" 吗？`)) {
    return
  }

  try {
    await api.admin.deleteTag(tag.id)
    fetchTags()
  } catch (error) {
    console.error('Failed to delete tag:', error)
    alert('删除失败')
  }
}

onMounted(() => {
  fetchTags()
})
</script>
