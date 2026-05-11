// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },

  modules: [
    '@nuxtjs/tailwindcss',
    '@pinia/nuxt',
  ],

  // SSR/SPA 路由分割
  routeRules: {
    '/admin/**': { ssr: false },   // 管理后台纯 SPA，无需 SEO
    '/**': { ssr: true },          // 博客前台 SSR
  },

  // 运行时配置
  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE || 'http://localhost:8080',
    },
  },

  // Element Plus 配置
  elementPlus: {
    importStyle: 'css',
  },

  // Tailwind CSS 配置
  tailwindcss: {
    configPath: 'tailwind.config.ts',
  },

  // TypeScript 配置
  typescript: {
    strict: true,
  },
})
