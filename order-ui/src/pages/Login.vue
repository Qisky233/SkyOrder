<template>
  <div class="login-container">
    <el-card class="login-card">
      <h2 class="login-title">SkyOrder 登录</h2>
      <el-form :model="loginForm" label-position="top">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="loginForm.username" prefix-icon="el-icon-user"></el-input>
        </el-form-item>
        <el-form-item label="密码">
          <el-input
            v-model="loginForm.password"
            type="password"
            prefix-icon="el-icon-lock"
            show-password
          ></el-input>
        </el-form-item>
        <el-form-item label="验证码">
          <el-input v-model="loginForm.captcha" prefix-icon="el-icon-key"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm">登录</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import axios from 'axios';
import { useStore } from 'vuex';
import { ElMessage } from 'element-plus';
import { useRouter } from 'vue-router';

const route = useRouter();

const store = useStore();
const loginForm = ref({
  username: 'root',
  password: '123456',
  captcha: '123',
});
const submitForm = async () => {
  try {
    const response = await axios.post('/api/v1/auth/login', loginForm.value);
    if (response) {
        store.commit('setToken', response.data.data.token);
        ElMessage.success('登录成功');
        // 登录成功后的跳转逻辑
        route.push('/dashboard/index');
    } else {
        ElMessage.error(response.data.message || '登录失败');
    }
    } catch (error) {
    ElMessage.error('登录失败，请检查网络或联系管理员');
  }
};
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #6a11cb, #2575fc);
}

.login-card {
  width: 400px;
  padding: 2rem;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.login-title {
  text-align: center;
  margin-bottom: 2rem;
  color: #333;
}

.el-form-item {
  margin-bottom: 1.5rem;
}

.el-button {
  width: 100%;
}

/* 媒体查询，调整不同屏幕尺寸下的样式 */
@media (max-width: 768px) {
  .login-card {
    width: 90%;
  }
}
</style>