<script setup lang="ts">
import { ref, h, onMounted } from "vue";
import {
  PlusOutlined,
  AppstoreOutlined,
  PlayCircleOutlined,
  PauseCircleOutlined,
  DeleteOutlined,
  MoreOutlined,
  EditOutlined,
  CopyOutlined,
} from "@ant-design/icons-vue";
import type { MenuProps } from "ant-design-vue";
import { TunnelService } from "../../bindings/github.com/atticus6/freeTunnel/desktop/services";
import { Tunnel } from "../../bindings/github.com/atticus6/freeTunnel/desktop/models";

const showAddModal = ref(false);
const newTunnel = ref({
  name: "",
  host: "127.0.0.1",
  port: 3000,
});

const addTunnel = async () => {
  const res = await TunnelService.CreateTunnel(
    newTunnel.value.name,
    newTunnel.value.host,
    newTunnel.value.port
  );
  if (res) {
    tunnels.value.push(res);
    showAddModal.value = false;
  }
};

const deleteTunnel = async (id: number) => {
  await TunnelService.DeleteById(id);
  getTunnels();
};

const toggleTunnel = async (id: number) => {
  const tar = tunnels.value.find((item) => item.id === id);
  if (!tar) {
    return;
  }
  if (tar.active) {
    await TunnelService.CloseTunnel(id);
  } else {
    await TunnelService.OpenTunnel(id);
  }
  getTunnels();
};

const getMenuItems = (tunnel: any): MenuProps["items"] => [
  { key: "edit", label: "编辑", icon: () => h(EditOutlined) },
  { key: "copy", label: "复制配置", icon: () => h(CopyOutlined) },
  { type: "divider" },
  {
    key: "delete",
    label: "删除",
    danger: true,
    icon: () => h(DeleteOutlined),
    onClick: () => deleteTunnel(tunnel.id),
  },
];

const tunnels = ref<Tunnel[]>([]);

const getTunnels = () =>
  TunnelService.GetAllTunnels().then((res) => {
    tunnels.value = res;
  });

onMounted(() => {
  getTunnels();
});
</script>

<template>
  <div class="p-6">
    <div class="flex items-center justify-between mb-6">
      <div class="flex items-center gap-3">
        <div
          class="w-8 h-8 bg-gradient-to-br from-purple-500 to-pink-500 rounded-xl flex items-center justify-center shadow-lg shadow-purple-500/20"
        >
          <AppstoreOutlined class="text-white" />
        </div>
        <h1 class="text-xl font-bold text-gray-800">穿透列表</h1>
      </div>
      <a-button
        type="primary"
        @click="showAddModal = true"
        class="!bg-purple-600 !border-purple-600 hover:!bg-purple-700 !rounded-lg !h-9"
      >
        <template #icon><PlusOutlined /></template>
        添加隧道
      </a-button>
    </div>

    <div class="space-y-4">
      <div
        v-for="tunnel in tunnels"
        :key="tunnel.id"
        class="bg-white rounded-xl p-5 border border-gray-200 shadow-sm hover:shadow-md transition-all"
      >
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-4">
            <!-- <a-tag
              color="red"
              class="!rounded-md !px-3 !py-1 !text-sm !font-semibold !m-0"
              >{{ tunnel.protocol }}</a-tag
            > -->
            <span class="text-base font-semibold text-gray-800">{{
              tunnel.name
            }}</span>
            <div class="flex items-center gap-2">
              <span
                :class="[
                  'w-2 h-2 rounded-full',
                  tunnel.active ? 'bg-green-500 animate-pulse' : 'bg-gray-400',
                ]"
              ></span>
              <span
                :class="[
                  'text-sm',
                  tunnel.active ? 'text-green-600' : 'text-gray-500',
                ]"
              >
                {{ tunnel.active ? "运行中" : "已停止" }}
              </span>
            </div>
          </div>
          <div class="flex items-center gap-2">
            <a-button
              :type="tunnel.active ? 'default' : 'primary'"
              :danger="tunnel.active"
              @click="toggleTunnel(tunnel.id)"
              class="!rounded-lg"
            >
              <template #icon>
                <PauseCircleOutlined v-if="tunnel.active" />
                <PlayCircleOutlined v-else />
              </template>
              {{ tunnel.active ? "停止" : "启动" }}
            </a-button>
            <a-dropdown :trigger="['click']">
              <a-button class="!rounded-lg"><MoreOutlined /></a-button>
              <template #overlay>
                <a-menu :items="getMenuItems(tunnel)" />
              </template>
            </a-dropdown>
          </div>
        </div>
        <div class="grid grid-cols-2 gap-4 mt-4">
          <div class="bg-gray-50 rounded-lg p-3 border border-gray-100">
            <div class="text-gray-500 text-xs mb-1">内网地址</div>
            <div class="text-gray-800 font-mono font-medium">
              {{ tunnel.host }}
            </div>
          </div>
          <div class="bg-gray-50 rounded-lg p-3 border border-gray-100">
            <div class="text-gray-500 text-xs mb-1">内网端口</div>
            <div class="text-gray-800 font-mono font-medium">
              {{ tunnel.port }}
            </div>
          </div>
        </div>
      </div>
    </div>

    <a-empty v-if="tunnels.length === 0" class="py-16">
      <template #description
        ><span class="text-gray-500">暂无隧道配置</span></template
      >
      <a-button
        type="primary"
        @click="showAddModal = true"
        class="!bg-purple-600 !border-purple-600"
      >
        <template #icon><PlusOutlined /></template>
        创建第一个隧道
      </a-button>
    </a-empty>
  </div>

  <a-modal
    v-model:open="showAddModal"
    title="添加新隧道"
    :width="480"
    @ok="addTunnel"
    ok-text="添加"
    cancel-text="取消"
  >
    <a-form layout="vertical" class="mt-4">
      <a-form-item label="隧道名称" required>
        <a-input v-model:value="newTunnel.name" placeholder="输入隧道名称" />
      </a-form-item>

      <a-row :gutter="16">
        <a-col :span="12">
          <a-form-item label="本地地址" required>
            <a-input v-model:value="newTunnel.host" placeholder="127.0.0.1" />
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item label="本地端口" required>
            <a-input-number
              v-model:value="newTunnel.port"
              :min="1"
              :max="65535"
              class="!w-full"
              placeholder="3000"
            />
          </a-form-item>
        </a-col>
      </a-row>
    </a-form>
  </a-modal>
</template>
