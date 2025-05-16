<template>
  <div class="top-nav">
    <el-button class="toggle-button" type="text" @click="toggleSidebar">
      <el-icon :size="24"><Menu /></el-icon>
    </el-button>
    <div class="top-nav-center"></div>
    <div class="top-nav-right">
      <el-dropdown @command="handleCommand">
        <span class="el-dropdown-link">
          <el-avatar icon="el-icon-user-solid"></el-avatar>
        </span>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item command="profile">账号信息</el-dropdown-item>
            <el-dropdown-item command="logout">退出账号</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { Menu } from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';

const isSidebarVisible = ref(true);

const toggleSidebar = () => {
  isSidebarVisible.value = !isSidebarVisible.value;
};

const handleCommand = (command) => {
  if (command === 'logout') {
    ElMessage.info('退出账号');
    // 清除 token 等操作
  } else if (command === 'profile') {
    ElMessage.info('查看账号信息');
  }
};
</script>

<style scoped>
.top-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 60px;
  background-color: #2d3a4b;
  color: white;
  padding: 0 20px;
}

.toggle-button {
  border: none;
  background-color: transparent;
  color: white;
  font-size: 1.2rem;
}

.top-nav-center {
  flex-grow: 1;
}

.top-nav-right {
  display: flex;
  align-items: center;
}

.el-dropdown-link {
  cursor: pointer;
  color: white;
}

.el-avatar {
  margin-right: 10px;
}
</style>