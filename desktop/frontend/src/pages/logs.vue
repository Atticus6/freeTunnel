<script setup lang="ts">
import { CheckCircleOutlined } from "@ant-design/icons-vue";
</script>

<template>
  <div class="p-6">
    <div class="flex items-center gap-3 mb-6">
      <div
        class="w-8 h-8 bg-gradient-to-br from-green-500 to-emerald-500 rounded-xl flex items-center justify-center"
      >
        <CheckCircleOutlined class="text-white" />
      </div>
      <h1 class="text-xl font-bold text-gray-800">运行日志</h1>
    </div>
    <div class="bg-white rounded-xl p-5 border border-gray-200 shadow-sm">
      <a-empty description="暂无日志" />
    </div>
  </div>
</template>
