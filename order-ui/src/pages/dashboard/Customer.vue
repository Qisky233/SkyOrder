<template>
  <div class="dashboard">
    <SideBar />
    <div class="main-content">
      <TopNav />
      <div class="content">
        <h2>用户列表</h2>
        <div v-if="loading">加载中...</div>
        <div v-else>
          <ul>
            <li v-for="customer in customers" :key="customer.id">
              <p><strong>姓名：</strong>{{ customer.name }}</p>
              <p><strong>电话：</strong>{{ customer.phone }}</p>
              <p><strong>QQ：</strong>{{ customer.qq }}</p>
              <p><strong>微信：</strong>{{ customer.wx }}</p>
              <p><strong>总金额：</strong>{{ customer.total }}</p>
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

const customers = ref([]);
const loading = ref(true);

const fetchCustomers = async () => {
  try {
    const token = localStorage.getItem('token');

    const config = {
      headers: {
        Authorization: `Bearer ${token}`
      }
    };
    const response = await axios.get('/api/v1/customer', config);
    customers.value = response.data.data;
  } catch (error) {
    console.error('获取用户列表失败:', error);
  } finally {
    loading.value = false;
  }
};

fetchCustomers();
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