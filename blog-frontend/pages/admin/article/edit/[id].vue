<template>
  <AdminArticleForm :article-id="Number(route.params.id)" />
</template>

<script setup lang="ts">
definePageMeta({
  layout: false,
})

const route = useRoute()
</script>
