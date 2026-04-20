<script setup lang="ts">
import avatar from './assets/Avatar.jpg' //el Ui必须要导入不能在src里写
import { useElementSize } from '@vueuse/core'
import {computed, ref} from "vue";
import { Icon } from '@iconify/vue';
import IconCard from "./Card/IconCard.vue";

const HomeTop = ref(null)
const settingShow = ref(false)
const { height } = useElementSize(HomeTop) //通过变量名自动解析

const avatarSize = computed(() => {
    return height.value * 0.6
})

const IconSize = computed(() => {
    return `${height.value * 0.3}px`;
})

const nameSize = computed(() => {
    return height.value * 0.2
})

</script>

<template>
    <div class="home">
        <div ref="HomeTop" class="home-top">
            <el-avatar style="margin: 0 10px; " :size="avatarSize" :src="avatar"/>
            <div style="color: #377ed3" :style="{ fontSize: nameSize + 'px' }">夜月</div>

            <div class="module">
                <Icon class="setting" icon="mdi:settings-outline" @click="settingShow = !settingShow">
                </Icon>

                <IconCard class="setting-card" v-if="settingShow">
                </IconCard>

            </div>
        </div>

        <div class="home-content">
            <div class="content-title">

            </div>
        </div>
    </div>
</template>


<style>
* {
    margin: 0;
    padding: 0;
}
</style>

<style scoped>
.home {
    width: 100%;
    height: 100%;
    background-image: url("./assets/雪莉1.jpeg");
    background-size: cover;
    background-color: rgba(218, 246, 244, 0.4);
    background-blend-mode: overlay;
}

.home-top {
    display: flex;
    width: 100%;
    height: 7vh;
    align-items: center;
    color: black;
    background: rgba(218, 246, 244, 0.4);
    backdrop-filter: blur(5px);
    .module {
        display: flex;
        align-items: center;
        margin-left: auto;
        margin-right: 10px;
        .setting{
            position: relative;
            font-size: v-bind(IconSize);
            cursor: pointer;
            color: rgb(76, 76, 76);
        }
        .setting-card{
            font-size: 12px;
            position: absolute;
            top: 100%;
            right: 0;
            margin-top: 8px;
        }
    }
}

.home-content{
    height: 93vh;

    align-items: center;
}

.content-title{
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-template-rows: repeat(3, 1fr);
    gap: 15px;
}

@media (prefers-color-scheme: dark) {  /*深色模式适配*/
    .home{
        color: white;
        background-color: rgba(2, 12, 12, 0.6);
    }

    .home-top {
        color: rgba(255, 255, 255, 0.8);
        background: rgba(2, 12, 12, 0.35);
        .module{
            .setting{
                color: rgba(198, 196, 196);
            }
        }
    }

    .content-title {
        color: white;
    }
}
</style>
