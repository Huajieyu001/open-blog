-- 博客系统初始数据脚本 (SQLite)

-- 初始管理员账号
-- 密码: admin123 (BCrypt 加密)
INSERT OR IGNORE INTO users (username, password, nickname, avatar) 
VALUES ('admin', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iAt6Z5EH', '管理员', NULL);

-- 示例分类
INSERT OR IGNORE INTO categories (name, slug, description, sort_order) 
VALUES ('技术笔记', 'tech', '技术学习笔记和心得', 1);

INSERT OR IGNORE INTO categories (name, slug, description, sort_order) 
VALUES ('生活随笔', 'life', '生活感悟和随笔', 2);

INSERT OR IGNORE INTO categories (name, slug, description, sort_order) 
VALUES ('项目实战', 'project', '项目开发实战记录', 3);

-- 示例标签
INSERT OR IGNORE INTO tags (name, slug) VALUES ('Go', 'go');
INSERT OR IGNORE INTO tags (name, slug) VALUES ('Vue', 'vue');
INSERT OR IGNORE INTO tags (name, slug) VALUES ('JavaScript', 'javascript');
INSERT OR IGNORE INTO tags (name, slug) VALUES ('数据库', 'database');
INSERT OR IGNORE INTO tags (name, slug) VALUES ('Linux', 'linux');
INSERT OR IGNORE INTO tags (name, slug) VALUES ('Docker', 'docker');

-- 示例文章
INSERT OR IGNORE INTO articles (title, slug, summary, content_md, content_html, status, category_id, author_id, publish_time) 
VALUES (
    '欢迎使用博客系统',
    'welcome',
    '这是一篇欢迎文章，介绍博客系统的基本功能。',
    '# 欢迎使用博客系统

这是一个基于 Nuxt 3 + Go 开发的博客系统。

## 主要功能

- 支持 Markdown 编辑
- 支持富文本编辑
- 支持图片上传
- 支持分类和标签
- 支持文章置顶
- 支持草稿和发布状态

## 技术栈

- 前端：Nuxt 3 (Vue 3 + Vite)
- 后端：Go + Gin
- 数据库：SQLite
- 图片存储：阿里云 OSS

## 开始使用

1. 访问管理后台 `/admin`
2. 使用默认账号登录：admin / admin123
3. 创建你的第一篇文章',
    '<h1>欢迎使用博客系统</h1>
<p>这是一个基于 Nuxt 3 + Go 开发的博客系统。</p>
<h2>主要功能</h2>
<ul>
<li>支持 Markdown 编辑</li>
<li>支持富文本编辑</li>
<li>支持图片上传</li>
<li>支持分类和标签</li>
<li>支持文章置顶</li>
<li>支持草稿和发布状态</li>
</ul>
<h2>技术栈</h2>
<ul>
<li>前端：Nuxt 3 (Vue 3 + Vite)</li>
<li>后端：Go + Gin</li>
<li>数据库：SQLite</li>
<li>图片存储：阿里云 OSS</li>
</ul>
<h2>开始使用</h2>
<ol>
<li>访问管理后台 <code>/admin</code></li>
<li>使用默认账号登录：admin / admin123</li>
<li>创建你的第一篇文章</li>
</ol>',
    1,  -- 已发布
    1,  -- 技术笔记分类
    1,  -- 管理员
    CURRENT_TIMESTAMP
);
