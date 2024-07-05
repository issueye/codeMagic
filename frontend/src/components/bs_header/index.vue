<template>
  <div
    class="w-full h-[55px] flex justify-between items-center bg-white border border-b-1 border-solid border-[#d9d9d9]"
  >
    <div>
      <span
        class="ml-5 text-[#262626] leading-8 text-base font-bold"
        style="vertical-align: middle"
        >{{ title }}</span
      >
      <el-divider direction="vertical" />
      <span class="text-[#BFBFBF] text-xs">{{ description }}</span>
    </div>
    <div class="mr-2.5">
      <slot name="actions"></slot>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps({
  title: {
    type: String,
    default: "",
  },
  description: {
    type: String,
    default: "",
  },
});
</script>

