<template>
  <div class="dashboard">
    <SideBar />
    <div class="main-content">
      <TopNav />
      <div class="content">
        123
      </div>
    </div>
  </div>
</template>

<script setup>
import TopNav from '@/components/topNav.vue';
import SideBar from '@/components/sideBar.vue';
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