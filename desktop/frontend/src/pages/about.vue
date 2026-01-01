<script setup lang="ts">
import { InfoCircleOutlined } from "@ant-design/icons-vue";
import { h } from "vue";
</script>

<template>
  <div class="p-6">
    <div class="flex items-center gap-3 mb-6">
      <div
        class="w-8 h-8 bg-gradient-to-br from-purple-500 to-pink-500 rounded-xl flex items-center justify-center"
      >
        <InfoCircleOutlined class="text-white" />
      </div>

      <h1 class="text-xl font-bold text-gray-800">关于</h1>
    </div>
    <div class="bg-white rounded-xl p-5 border border-gray-200 shadow-sm">
      <div class="text-center py-8">
        <div
          class="w-16 h-16 bg-gradient-to-br from-purple-500 to-pink-500 rounded-2xl flex items-center justify-center mx-auto mb-4"
        >
          <span class="text-white text-2xl font-bold">F</span>
        </div>
        <h2 class="text-xl font-bold text-gray-800 mb-2">freeTunnel</h2>
        <p class="text-gray-500 mb-4">内网穿透工具</p>
        <a-tag color="purple">v1.1.6</a-tag>
      </div>
    </div>
  </div>
</template>
