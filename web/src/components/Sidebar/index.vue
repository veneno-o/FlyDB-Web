<template>
  <div class="aslide">
    <div class="item">
    <el-icon class="icon"><Message /></el-icon>
      <a class="say" href="mailto:bytestoragecommunity@gmail.com">意见反馈</a>
    </div>
  </div>
</template>

<script setup lang="ts"></script>

<style scoped>
.aslide {
  position: fixed;
  top: 50%;
  right: 20px;
  transform: translateY(-50%);
  box-shadow: 0 0 8px rgba(0, 0, 0, 0.3);
  display: flex;
  flex-direction: column;
  padding: 20px 10px;
  background-color: var(--bg-color);
  border-radius: var(--border-radius);
}
.item {
  display: flex;
  justify-content: center;
  align-items: center;
}
.icon {
  margin-right: 4px;
}

</style>
