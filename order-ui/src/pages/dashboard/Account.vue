<template>
  <div class="dashboard">
    <SideBar />
    <div class="main-content">
      <TopNav />
      <div class="content">
        <h2>管理员信息</h2>
        <div v-if="loading">加载中...</div>
        <div v-else>
          <ul>
            <li v-for="admin in adminProfile" :key="admin.id">
              <p><strong>姓名：</strong>{{ admin.username }}</p>
              <p><strong>电话：</strong>{{ admin.role }}</p>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import TopNav from '@/components/topNav.vue';
import SideBar from '@/components/sideBar.vue';
import axios from 'axios';
import { ref } from 'vue';

const adminProfile = ref(null);
const loading = ref(true);

const fetchAdminProfile = async () => {
  try {
    const token = localStorage.getItem('token');

    const config = {
      headers: {
        Authorization: `Bearer ${token}`
      }
    };
    const response = await axios.get('/api/v1/admin', config);
    adminProfile.value = response.data.data.list;
  } catch (error) {
    console.error('获取管理员信息失败:', error);
  } finally {
    loading.value = false;
  }
};

fetchAdminProfile();
</script>

<style scoped>
.dashboard {
  display: flex;
  width: 100vw; /* 占满整个视口宽度 */
  height: 100vh; /* 占满整个视口高度 */
}

.main-content {
  display: flex;
  flex-direction: column; /* 主内容区域为垂直布局 */
  width: 100%; /* 占满剩余宽度 */
  height: 100%; /* 占满剩余高度 */
}

.content {
  flex: 1; /* 主内容区域占据剩余空间 */
  overflow: auto; /* 如果内容超出，允许滚动 */
  padding: 20px; /* 添加一些内边距 */
}
</style>