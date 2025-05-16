<template>
  <div class="home">
    <section class="hero">
      <h1>欢迎来到 SkyOrder</h1>
      <p>程序员的闲暇时间，大学生的作业助手，小团队的外包利器。</p>
      <el-button class="cta-button" type="primary" @click="goToOrder">立即下单</el-button>
    </section>

    <section class="features">
      <el-row :gutter="20">
        <el-col :span="24" :md="8">
          <div class="feature-item">
            <el-icon :size="40"><Coin /></el-icon>
            <h3>价格透明</h3>
            <p>所有服务明码标价，无隐藏费用。您可以清楚地了解每项服务的具体价格，避免任何不必要的支出。</p>
          </div>
        </el-col>
        <el-col :span="24" :md="8">
          <div class="feature-item">
            <el-icon :size="40"><Timer /></el-icon>
            <h3>高效交付</h3>
            <p>专业团队快速响应，按时交付。我们的团队成员都是经验丰富的专业人士，能够高效地完成您的订单。</p>
          </div>
        </el-col>
        <el-col :span="24" :md="8">
          <div class="feature-item">
            <el-icon :size="40"><Service /></el-icon>
            <h3>优质服务</h3>
            <p>贴心售后，解决您的后顾之忧。我们提供全方位的售后服务，确保您在使用过程中没有任何问题。</p>
          </div>
        </el-col>
      </el-row>
    </section>

    <section class="order-stats">
      <h2>订单统计</h2>
      <div ref="orderChartRef" class="order-chart"></div>
    </section>

    <section class="testimonials">
      <h2>客户评价</h2>
      <el-carousel :autoplay="true" :interval="5000" indicator-position="outside" height="auto">
        <el-carousel-item class="testimonial-item" v-for="testimonial in testimonials" :key="testimonial.id">
          <div>
            <p>{{ testimonial.content }}</p>
            <p class="testimonial-author">{{ testimonial.author }}</p>
          </div>
        </el-carousel-item>
      </el-carousel>
    </section>

    <footer class="footer">
      <p>SkyOrder &copy;{{ currentYear }}</p>
    </footer>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import * as echarts from 'echarts';
import { Coin, Timer, Service } from '@element-plus/icons-vue';

const currentYear = new Date().getFullYear();
const orderChartRef = ref(null);
const testimonials = ref([
  { id: 1, content: '非常专业的团队，解决了我的大问题！他们不仅技术精湛，而且服务态度也非常好。', author: '张三' },
  { id: 2, content: '订单交付速度超快，服务态度也很好！我非常满意，强烈推荐给大家。', author: '李四' },
  { id: 3, content: '性价比超高，推荐给大家！价格合理，服务一流，值得信赖。', author: '王五' },
]);

const goToOrder = () => {
  console.log('立即下单按钮被点击');
  // 可以在这里添加跳转到下单页面的逻辑
};

onMounted(() => {
  const orderChart = echarts.init(orderChartRef.value);
  const option = {
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'shadow',
      },
    },
    legend: {
      data: ['订单数量'],
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true,
    },
    xAxis: [
      {
        type: 'category',
        data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],
      },
    ],
    yAxis: [
      {
        type: 'value',
      },
    ],
    series: [
      {
        name: '订单数量',
        type: 'bar',
        data: [10, 20, 15, 30, 25, 18, 22],
      },
    ],
  };
  orderChart.setOption(option);
});
</script>

<style scoped>
.home {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  min-height: 100vh;
  text-align: center;
  font-family: Arial, sans-serif;
  color: #333;
}

.hero {
  padding: 2rem;
  background: linear-gradient(135deg, #6a11cb, #2575fc);
  color: white;
}

.hero h1 {
  font-size: 2.5rem;
  margin-bottom: 1rem;
}

.hero p {
  font-size: 1.2rem;
  margin-bottom: 2rem;
}

.cta-button {
  padding: 0.8rem 1.5rem;
  font-size: 1rem;
  color: white;
  background-color: #ff7f50;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.cta-button:hover {
  background-color: #ff4500;
}

.features {
  padding: 2rem;
  background-color: #f4f4f4;
}

.feature-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 2rem;
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 1rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.feature-item h3 {
  margin: 1rem 0 0.5rem;
}

.feature-item p {
  font-size: 1rem;
}

.order-stats {
  padding: 2rem;
  background-color: #fff;
}

.order-chart {
  width: 100%;
  height: 300px;
}

.testimonials {
  padding: 1rem;
  background-color: #f4f4f4;
}

.el-carousel .el-carousel__container {
    border-radius: 16px;
}

.testimonial-item {
  text-align: left;
  padding: 1rem;
  background-color: #fff;
  width: calc(100% - 2rem);
  border-radius: 16px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  height: 160px;
}

.testimonial-item p {
  font-size: 1rem;
}

.testimonial-author {
  font-weight: bold;
  margin-top: 0.5rem;
}

.footer {
  padding: 1rem;
  background-color: #f4f4f4;
  font-size: 0.9rem;
}

/* 媒体查询，调整不同屏幕尺寸下的轮播图高度 */
@media (max-width: 768px) {
  .testimonial-item {
    min-height: 120px;
  }
}
</style>