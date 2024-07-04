<template>
<div class="bs-header-box">
    <div>
        <span class="ml-5 text-[#262626] leading-8 text-base font-bold" style="vertical-align: middle">{{ title }}</span>
        <el-divider direction="vertical" />
        <span class="text-[#BFBFBF] text-xs">{{ description }}</span>
    </div>
    <div class="mr-2.5">
        <slot name="actions"></slot>
    </div>
</div>
</template>

<script setup>
import { toRefs } from 'vue';

const props = defineProps({
    title: {
        type: String,
        default: '',
    },
    description: {
        type: String,
        default: '',
    },
})

</script>

<style lang="scss" scoped>
.bs-header-box {
    border-radius: 2px 2px 0 0;
    width: 100%;
    height: 45px;
    display: inline-flex;
    justify-content: space-between;
    align-items: center;
    background-color: #fff;
    border-bottom: 1px solid #D9D9D9;
}
</style>
