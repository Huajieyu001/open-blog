# 博客系统设计文档

## 一、技术栈

| 层次 | 技术 | 说明 |
|------|------|------|
| 前端 | Nuxt 3 (Vue 3 + Vite) | 混合渲染：博客页 SSR，后台管理 SPA |
| UI | Tailwind CSS + Element Plus | Tailwind 用于博客前台，Element Plus 用于管理后台 |
| 编辑器 | Markdown-it + Tiptap | Markdown + 富文本双模式切换（CSDN 风格） |
| 后端 | Spring Boot 3.2 | Java 21 |
| ORM | MyBatis-Plus 3.5 | 灵活 SQL 控制 |
| 鉴权 | Spring Security + JWT (jjwt 0.12) | 无状态 token 认证 |
| 数据库 | MySQL 8.0 | InnoDB, utf8mb4 |
| 图片存储 | 阿里云 OSS | 后端统一上传接口 |
| 构建 | Maven (后端) + npm (前端) | |

## 二、项目结构

```
blog/
├── DESIGN.md                          # 设计文档（本文件）
├── blog-server/                       # Spring Boot 后端
│   ├── pom.xml
│   └── src/main/
│       ├── java/com/blog/
│       │   ├── BlogApplication.java   # 入口
│       │   ├── common/                # Result, PageResult, 异常处理
│       │   ├── config/                # Security, CORS, OSS, MyBatisPlus
│       │   ├── controller/
│       │   │   ├── admin/             # 后台 API（需 JWT 鉴权）
│       │   │   └── pub/               # 前台 API（公开）
│       │   ├── service/ + impl/       # 业务逻辑
│       │   ├── mapper/                # MyBatis-Plus Mapper
│       │   ├── entity/                # 数据实体
│       │   ├── dto/                   # 数据传输对象
│       │   └── security/              # JWT Provider + Filter
│       └── resources/
│           ├── application.yml         # 通用配置
│           ├── application-dev.yml     # 开发环境配置
│           └── db/schema.sql           # 建表语句
│
├── blog-frontend/                      # Nuxt 3 前端
│   ├── nuxt.config.ts                  # 路由规则（SSR/SPA 分拆）
│   ├── pages/
│   │   ├── index.vue                   # 博客首页（SSR）
│   │   ├── article/[id].vue            # 文章详情（SSR）
│   │   ├── category/[slug].vue         # 分类列表（SSR）
│   │   ├── admin.vue                   # 后台布局
│   │   └── admin/                      # 后台页面（全部 SPA）
│   │       ├── login.vue
│   │       ├── dashboard.vue
│   │       ├── article/create.vue      # 双模式编辑器
│   │       ├── article/edit/[id].vue
│   │       ├── article/list.vue
│   │       ├── category/index.vue
│   │       └── tag/index.vue
│   ├── components/
│   │   ├── blog/ (Header, Footer, ArticleCard, Pagination)
│   │   ├── admin/ (Sidebar)
│   │   └── common/ (ImageUpload)
│   ├── composables/ (useApi, useAuth, useArticle)
│   └── middleware/auth.ts
```

## 三、数据库设计

### users — 管理员用户
| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT PK | 主键 |
| username | VARCHAR(50) UNIQUE | 用户名 |
| password | VARCHAR(255) | BCrypt 加密密码 |
| nickname | VARCHAR(50) | 昵称 |
| avatar | VARCHAR(500) | 头像 URL |
| create_time | DATETIME | 创建时间 |
| update_time | DATETIME | 更新时间 |

### categories — 分类
| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT PK | 主键 |
| name | VARCHAR(50) | 分类名称 |
| slug | VARCHAR(50) UNIQUE | URL 友好标识 |
| description | VARCHAR(200) | 描述 |
| sort_order | INT DEFAULT 0 | 排序 |
| create_time | DATETIME | 创建时间 |
| update_time | DATETIME | 更新时间 |

### tags — 标签
| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT PK | 主键 |
| name | VARCHAR(50) | 标签名 |
| slug | VARCHAR(50) UNIQUE | URL 友好标识 |
| create_time | DATETIME | 创建时间 |
| update_time | DATETIME | 更新时间 |

### articles — 文章
| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT PK | 主键 |
| title | VARCHAR(200) | 标题 |
| slug | VARCHAR(200) UNIQUE | URL 友好标识（可选） |
| summary | VARCHAR(500) | 摘要 |
| content_md | LONGTEXT | Markdown 原文 |
| content_html | LONGTEXT | 渲染后的 HTML |
| cover_image | VARCHAR(500) | 封面图 URL |
| status | TINYINT DEFAULT 0 | 0=草稿, 1=已发布 |
| is_top | TINYINT DEFAULT 0 | 0=否, 1=置顶 |
| view_count | INT DEFAULT 0 | 阅读量 |
| category_id | BIGINT | 分类 ID |
| author_id | BIGINT | 作者 ID |
| publish_time | DATETIME | 发布时间（从草稿变为发布时设置） |
| create_time | DATETIME | 创建时间 |
| update_time | DATETIME | 更新时间 |

### article_tag — 文章标签关联
| 字段 | 类型 | 说明 |
|------|------|------|
| article_id | BIGINT PK | 文章 ID |
| tag_id | BIGINT PK | 标签 ID |
| create_time | DATETIME | 创建时间 |
| update_time | DATETIME | 更新时间 |

## 四、API 设计

### 公开接口（无需鉴权）
| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/pub/articles | 文章列表（分页、可按分类/标签/关键词筛选） |
| GET | /api/pub/articles/{id} | 文章详情（含完整 HTML） |
| GET | /api/pub/articles/slug/{slug} | 按 slug 获取文章 |
| GET | /api/pub/categories | 分类列表 |

### 管理接口（需 JWT 鉴权）
| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/admin/login | 管理员登录 |
| GET | /api/admin/articles | 文章列表（含草稿） |
| POST | /api/admin/articles | 创建文章 |
| PUT | /api/admin/articles/{id} | 更新文章 |
| DELETE | /api/admin/articles/{id} | 删除文章 |
| GET | /api/admin/categories | 分类列表 |
| POST | /api/admin/categories | 创建分类 |
| PUT | /api/admin/categories/{id} | 更新分类 |
| DELETE | /api/admin/categories/{id} | 删除分类 |
| GET | /api/admin/tags | 标签列表 |
| POST | /api/admin/tags | 创建标签 |
| PUT | /api/admin/tags/{id} | 更新标签 |
| DELETE | /api/admin/tags/{id} | 删除标签 |
| POST | /api/admin/upload | 上传图片到 OSS |

## 五、SSR/SPA 路由分割

```ts
// nuxt.config.ts
routeRules: {
  '/admin/**': { ssr: false },   // 管理后台纯 SPA，无需 SEO
  '/**': { ssr: true },          // 博客前台 SSR
}
```

## 六、鉴权流程

1. 前端 POST `/api/admin/login` 提交用户名密码
2. 后端验证后返回 JWT token（24h 有效）
3. 前端存 token 到 localStorage，后续请求带 `Authorization: Bearer <token>`
4. JwtAuthenticationFilter 拦截请求，解析 token 并注入 SecurityContext
5. `/api/admin/**` 路径默认要求认证，`/api/pub/**` 全部公开

## 七、初始管理员

- 用户名: `admin`
- 密码: `admin123`（BCrypt 加密存储，首次登录后请修改）
