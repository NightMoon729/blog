<script setup lang="ts">
import avatar from './assets/Avatar.jpg' //el Ui必须要导入不能在src里写
import { useElementSize } from '@vueuse/core'
import {computed, ref} from "vue";

const HomeTop = ref(null)
const { height } = useElementSize(HomeTop) //通过变量名自动解析

const avatarSize = computed(() => {
    return height.value * 0.6
})

const nameSize = computed(() => {
    return height.value * 0.2
})

</script>

<template>
    <div ref="HomeTop" class="home-top">
        <el-avatar style="margin: 0 10px; " :size="avatarSize" :src="avatar"/>
        <div :style="{ fontSize: nameSize + 'px' }">夜月</div>

    </div>
    <div class="home-content">

    </div>
</template>


<style>
* {
    margin: 0;
    padding: 0;
}
</style>

<style scoped>
.home-top {
    display: flex;
    width: 100%;
    height: 7vh;
    align-items: center;
    background: #daf6f4;
}
.home-content{
    height: 93vh;
    background-image: url("./assets/雪莉1.jpeg");
    background-size: cover;
    background-color: rgba(218, 246, 244, 0.4);
    background-blend-mode: overlay;
}

@media (prefers-color-scheme: dark) {
    .home-top {
        color: rgba(255, 255, 255, 0.8);
        background: rgb(2, 12, 12);
    }
    .home-content {
        color: white;
        background-color: rgba(2, 12, 12, 0.6);
    }
}
</style>
