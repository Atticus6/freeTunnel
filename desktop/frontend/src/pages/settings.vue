<script setup lang="ts">
import { SettingOutlined } from "@ant-design/icons-vue";
</script>

<template>
  <div class="p-6">
    <div class="flex items-center gap-3 mb-6">
      <div
        class="w-8 h-8 bg-gradient-to-br from-gray-600 to-gray-700 rounded-xl flex items-center justify-center"
      >
        <SettingOutlined class="text-white" />
      </div>
      <h1 class="text-xl font-bold text-gray-800">设置</h1>
    </div>
    <div class="bg-white rounded-xl p-5 border border-gray-200 shadow-sm">
      <a-empty description="暂无设置项" />
    </div>
  </div>
</template>
