import { defineConfig } from "vite";
import UnoCSS from 'unocss/vite'

import path from "path";
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import vueDevTools from 'vite-plugin-vue-devtools'
import VueRouter from 'unplugin-vue-router/vite'


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    VueRouter({
      dts:"./src/types/typed-router.d.ts"
    }),
    vue(),
    vueJsx(),
    vueDevTools(),
    UnoCSS()
  ],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
});
