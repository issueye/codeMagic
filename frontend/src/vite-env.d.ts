/// <reference types="vite/client" />
declare module "*.vue" {
  import { ComponentOptions } from "vue";
  const component: ComponentOptions;
  export default component;
}

declare module 'element-plus/dist/locale/zh-cn.mjs';