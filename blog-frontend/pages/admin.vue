<template>
  <NuxtLayout name="admin">
    <NuxtPage />
  </NuxtLayout>
</template>

<script setup lang="ts">
definePageMeta({
  layout: false,
})

const authStore = useAuthStore()
const router = useRouter()

onMounted(() => {
  // 初始化 auth 状态
  authStore.init()

  // 检查是否已登录（登录页面除外）
  if (!authStore.isAuthenticated && router.currentRoute.value.path !== '/admin/login') {
    router.push('/admin/login')
  }
})
</script>
