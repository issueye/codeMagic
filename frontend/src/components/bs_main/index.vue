<template>
    <div class="bs-main-body">
        <div class="body-box">
            <slot name="body"></slot>
        </div>
        <slot name="footer">
        </slot>
    </div>
</template>

<script setup>

</script>

<style lang="scss" scoped>
.bs-main-body {
    height: calc(100% - 46px);
    background-color: #fff;
    border-radius: 0 0 2px 2px;

    .body-box {
        padding: 10px;
        height: calc(100% - 59px);
    }

    .footer-box {
        padding: 10px;
        border-top: 1px solid #d9d9d9;
    }
}
</style>