<script setup lang="ts">
</script>

<template>
    <div class="card">
        <slot/>
    </div>
</template>

<style scoped lang="scss">
    .card{
        padding: 15px;
        color: black;
        width: calc(100% + 4em);
        background: rgba(218, 246, 244, 0.4);
        backdrop-filter: blur(2px);
        border-radius: 15px;
        position: absolute !important;
        top: 100%;
        left: 50%;
        transform: translate(-50%, 15px);
        :deep(.title-box){
            margin-bottom: 15px;
        }
        &::before {
            content: '';
            position: absolute;
            top: -15px;
            left: 0;
            right: 0;
            height: 15px;
        }
    }
    .dark {
        .card{
            color: rgba(255, 255, 255, 0.8);
            background: rgba(2, 12, 12, 0.35);
        }
    }
</style>