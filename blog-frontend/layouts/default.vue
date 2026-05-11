<template>
  <div class="min-h-screen flex flex-col bg-gray-50">
    <!-- Header -->
    <header class="bg-white shadow-sm border-b">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center h-16">
          <!-- Logo -->
          <div class="flex-shrink-0">
            <NuxtLink to="/" class="text-xl font-bold text-gray-900 hover:text-primary-600">
              My Blog
            </NuxtLink>
          </div>

          <!-- Navigation -->
          <nav class="hidden md:flex space-x-8">
            <NuxtLink
              to="/"
              class="text-gray-600 hover:text-gray-900 px-3 py-2 text-sm font-medium"
              active-class="text-primary-600 border-b-2 border-primary-600"
            >
              首页
            </NuxtLink>
            <NuxtLink
              to="/category"
              class="text-gray-600 hover:text-gray-900 px-3 py-2 text-sm font-medium"
              active-class="text-primary-600 border-b-2 border-primary-600"
            >
              分类
            </NuxtLink>
            <NuxtLink
              to="/about"
              class="text-gray-600 hover:text-gray-900 px-3 py-2 text-sm font-medium"
              active-class="text-primary-600 border-b-2 border-primary-600"
            >
              关于
            </NuxtLink>
          </nav>

          <!-- Mobile menu button -->
          <div class="md:hidden">
            <button
              @click="mobileMenuOpen = !mobileMenuOpen"
              class="text-gray-600 hover:text-gray-900 focus:outline-none"
            >
              <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path
                  v-if="!mobileMenuOpen"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M4 6h16M4 12h16M4 18h16"
                />
                <path
                  v-else
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M6 18L18 6M6 6l12 12"
                />
              </svg>
            </button>
          </div>
        </div>
      </div>

      <!-- Mobile menu -->
      <div v-if="mobileMenuOpen" class="md:hidden">
        <div class="px-2 pt-2 pb-3 space-y-1 sm:px-3">
          <NuxtLink
            to="/"
            class="block px-3 py-2 text-base font-medium text-gray-600 hover:text-gray-900 hover:bg-gray-50"
            @click="mobileMenuOpen = false"
          >
            首页
          </NuxtLink>
          <NuxtLink
            to="/category"
            class="block px-3 py-2 text-base font-medium text-gray-600 hover:text-gray-900 hover:bg-gray-50"
            @click="mobileMenuOpen = false"
          >
            分类
          </NuxtLink>
          <NuxtLink
            to="/about"
            class="block px-3 py-2 text-base font-medium text-gray-600 hover:text-gray-900 hover:bg-gray-50"
            @click="mobileMenuOpen = false"
          >
            关于
          </NuxtLink>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="flex-grow">
      <slot />
    </main>

    <!-- Footer -->
    <footer class="bg-white border-t">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div class="text-center text-gray-500 text-sm">
          <p>&copy; {{ new Date().getFullYear() }} My Blog. All rights reserved.</p>
          <p class="mt-2">
            Powered by
            <a href="https://nuxt.com" class="text-primary-600 hover:text-primary-700" target="_blank" rel="noopener noreferrer">Nuxt 3</a>
            &
            <a href="https://gin-gonic.com" class="text-primary-600 hover:text-primary-700" target="_blank" rel="noopener noreferrer">Gin</a>
          </p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
const mobileMenuOpen = ref(false)
</script>
