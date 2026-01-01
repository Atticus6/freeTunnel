<script setup lang="ts">
import { computed } from "vue";
import { useRoute, useRouter } from "vue-router";

const route = useRoute();
const router = useRouter();
const currentPath = computed(() => route.path);
const navigateTo = (path: string) => router.push(path);
</script>

<template>
  <div class="h-screen bg-gradient-to-br from-gray-50 to-gray-100/80">
    <!-- 标题栏 -->
    <div
      class="h-10 bg-white/90 backdrop-blur border-b border-gray-200 flex items-center px-20"
      style="--wails-draggable: drag"
    >
      <div class="flex items-center gap-2">
        <div
          class="w-6 h-6 bg-gradient-to-br from-purple-500 to-pink-500 rounded-md grid place-items-center"
        >
          <span class="text-white text-xs font-bold">F</span>
        </div>
        <span class="text-gray-700 font-medium text-sm">freeTunnel</span>
      </div>
    </div>

    <div class="flex h-[calc(100%-40px)]">
      <!-- 侧边栏 -->
      <div
        class="w-14 bg-white border-r border-gray-200 flex flex-col items-center py-4 gap-2"
      >
        <a-tooltip placement="right" title="隧道列表">
          <div
            @click="navigateTo('/')"
            :class="[
              'w-9 h-9 rounded-lg cursor-pointer transition-colors grid place-items-center',
              currentPath === '/' ? 'bg-purple-100' : 'hover:bg-gray-100',
            ]"
          >
            <svg
              :class="[
                'w-5 h-5',
                currentPath === '/' ? 'text-purple-600' : 'text-gray-400',
              ]"
              fill="currentColor"
              viewBox="0 0 20 20"
            >
              <path
                d="M3 4a1 1 0 011-1h12a1 1 0 011 1v2a1 1 0 01-1 1H4a1 1 0 01-1-1V4zM3 10a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H4a1 1 0 01-1-1v-6zM14 9a1 1 0 00-1 1v6a1 1 0 001 1h2a1 1 0 001-1v-6a1 1 0 00-1-1h-2z"
              />
            </svg>
          </div>
        </a-tooltip>

        <a-tooltip placement="right" title="运行日志">
          <div
            @click="navigateTo('/logs')"
            :class="[
              'w-9 h-9 rounded-lg cursor-pointer transition-colors grid place-items-center',
              currentPath === '/logs' ? 'bg-green-100' : 'hover:bg-gray-100',
            ]"
          >
            <svg
              :class="[
                'w-5 h-5',
                currentPath === '/logs' ? 'text-green-600' : 'text-gray-400',
              ]"
              fill="currentColor"
              viewBox="0 0 20 20"
            >
              <path
                fill-rule="evenodd"
                d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                clip-rule="evenodd"
              />
            </svg>
          </div>
        </a-tooltip>

        <a-tooltip placement="right" title="账户信息">
          <div
            @click="navigateTo('/account')"
            :class="[
              'w-9 h-9 rounded-lg cursor-pointer transition-colors grid place-items-center',
              currentPath === '/account' ? 'bg-blue-100' : 'hover:bg-gray-100',
            ]"
          >
            <svg
              :class="[
                'w-5 h-5',
                currentPath === '/account' ? 'text-blue-600' : 'text-gray-400',
              ]"
              fill="currentColor"
              viewBox="0 0 20 20"
            >
              <path
                fill-rule="evenodd"
                d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z"
                clip-rule="evenodd"
              />
            </svg>
          </div>
        </a-tooltip>

        <a-tooltip placement="right" title="设置">
          <div
            @click="navigateTo('/settings')"
            :class="[
              'w-9 h-9 rounded-lg cursor-pointer transition-colors grid place-items-center',
              currentPath === '/settings' ? 'bg-gray-200' : 'hover:bg-gray-100',
            ]"
          >
            <svg
              :class="[
                'w-5 h-5',
                currentPath === '/settings' ? 'text-gray-700' : 'text-gray-400',
              ]"
              fill="currentColor"
              viewBox="0 0 20 20"
            >
              <path
                fill-rule="evenodd"
                d="M11.49 3.17c-.38-1.56-2.6-1.56-2.98 0a1.532 1.532 0 01-2.286.948c-1.372-.836-2.942.734-2.106 2.106.54.886.061 2.042-.947 2.287-1.561.379-1.561 2.6 0 2.978a1.532 1.532 0 01.947 2.287c-.836 1.372.734 2.942 2.106 2.106a1.532 1.532 0 012.287.947c.379 1.561 2.6 1.561 2.978 0a1.533 1.533 0 012.287-.947c1.372.836 2.942-.734 2.106-2.106a1.533 1.533 0 01.947-2.287c1.561-.379 1.561-2.6 0-2.978a1.532 1.532 0 01-.947-2.287c.836-1.372-.734-2.942-2.106-2.106a1.532 1.532 0 01-2.287-.947zM10 13a3 3 0 100-6 3 3 0 000 6z"
                clip-rule="evenodd"
              />
            </svg>
          </div>
        </a-tooltip>

        <a-tooltip placement="right" title="关于">
          <div
            @click="navigateTo('/about')"
            :class="[
              'w-9 h-9 rounded-lg cursor-pointer transition-colors grid place-items-center',
              currentPath === '/about' ? 'bg-purple-100' : 'hover:bg-gray-100',
            ]"
          >
            <svg
              :class="[
                'w-5 h-5',
                currentPath === '/about' ? 'text-purple-600' : 'text-gray-400',
              ]"
              fill="currentColor"
              viewBox="0 0 20 20"
            >
              <path
                fill-rule="evenodd"
                d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
                clip-rule="evenodd"
              />
            </svg>
          </div>
        </a-tooltip>
      </div>

      <!-- 主内容区域 -->
      <div class="flex-1 overflow-auto bg-gray-50">
        <RouterView />
      </div>
    </div>
  </div>
</template>
