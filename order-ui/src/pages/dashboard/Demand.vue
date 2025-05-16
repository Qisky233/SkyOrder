<template>
  <div class="dashboard">
    <SideBar />
    <div class="main-content">
      <TopNav />
      <div class="content">
        <h2>需求列表</h2>
        <div v-if="loading">加载中...</div>
        <div v-else>
          <ul>
            <li v-for="demand in demands" :key="demand.id">
              <p><strong>标题：</strong>{{ demand.title }}</p>
              <p><strong>描述：</strong>{{ demand.descs }}</p>
              <p><strong>价格：</strong>{{ demand.price }}</p>
              <p><strong>状态：</strong>{{ demand.status }}</p>
              <p><strong>用户 ID：</strong>{{ demand.user_id }}</p>
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

const demands = ref([]);
const loading = ref(true);

const fetchDemands = async () => {
  try {
    const token = localStorage.getItem('token');

    const config = {
      headers: {
        Authorization: `Bearer ${token}`
      }
    };

    const response = await axios.get('/api/v1/demand', config);
    demands.value = response.data.data;
  } catch (error) {
    console.error('获取需求列表失败:', error);
  } finally {
    loading.value = false;
  }
};

fetchDemands();
</script>

<style scoped>
.dashboard {
  display: flex;
  width: 100vw;
  height: 100vh;
}

.main-content {
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;
}

.content {
  flex: 1;
  overflow: auto;
  padding: 20px;
}
</style>