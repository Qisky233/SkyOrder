<template>
  <div class="sidebar-container">
    <div class="sidebar-header">
      <h2>接单管理系统</h2>
    </div>
    <div class="menu" v-for="(item, index) in menuItems" :key="index">
        <span class="menu-item">
            <router-link :to="item.link">{{ item.name }}</router-link>
        </span>
    </div>
  </div>
</template>

<script setup>
const menuItems = [
  {
    name: '总览',
    link: '/dashboard/index',
  },
  {
    name: '账户管理', 
    link: '/dashboard/account',
  },
  {
    name: '用户管理', 
    link: '/dashboard/customer',
  },
  {
    name: '需求管理', 
    link: '/dashboard/demend',
  },
];

</script>

<style scoped>
.sidebar-container {
  width: 250px;
  height: calc(100vh - 40px);
  background-color: #2d3a4b;
  color: white;
  padding: 20px;
  box-shadow: 2px 0 4px rgba(0, 0, 0, 0.1);
}

.sidebar-header {
  margin-bottom: 20px;
}

.sidebar-header h2 {
  font-size: 1.2rem;
  margin: 0 0 20px;
  text-align: center;
}

/* 新增样式 */
.menu {
  margin: 0;
  padding: 0;
  list-style: none;
}

.menu-item {
  display: block;
  margin: 10px 0;
  padding: 10px;
  border-radius: 5px;
  transition: background-color 0.3s ease;
}

.menu-item:hover {
  background-color: #3a4961;
}

.menu-item a {
  color: white;
  text-decoration: none;
  font-size: 1rem;
  display: block;
}

.menu-item a:hover {
  text-decoration: underline;
}

.menu-item a.router-link-exact-active {
  font-weight: bold;
  color: #4caf50; /* 激活状态的颜色 */
}

@media (max-width: 768px) {
  .sidebar-container {
    display: none;
  }
}
</style>