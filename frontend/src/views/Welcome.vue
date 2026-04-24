<script setup lang="ts">
import avatar from '@/assets/Avatar.jpg' //el Ui必须要导入不能在src里写
import { useElementSize } from '@vueuse/core'
import {computed, ref} from "vue";
import { Icon } from '@iconify/vue';
import TitleCard from "../Card/titleCard.vue";
import router from "@/router/index";

const HomeTop = ref(null)
const settingShow = ref(false)
const otherShow = ref(false)
const { height } = useElementSize(HomeTop) //通过变量名自动解析
const isDark = ref(false)

const LAD = computed(() => {
    return isDark.value ? "prime:moon" : "prime:sun"
}) //深色模式和浅色模式

const toggleDark = () => {
    isDark.value = !isDark.value
    // 这一行是核心：给 <html> 标签加/删 'dark' class
    document.documentElement.classList.toggle('dark', isDark.value)
}

const avatarSize = computed(() => {
    return height.value * 0.6
})

const iconSize = computed(() => {
    return `${height.value * 0.3}px`;
})

const textSize = computed(() => {
    return `${height.value * 0.2}px`;
})
</script>

<template>
    <div class="home">
        <div class="home-header">
            <div ref="HomeTop" class="home-top">
                <el-avatar style="margin: 0 10px; " :size="avatarSize" :src="avatar"/>
                <div style="color: #377ed3" :style="{ fontSize: textSize }">夜月</div>

                <div class="main-title">
                    <div class="title-box">
                        <Icon icon="tabler:home"/>
                        <div class="text">主页</div>
                    </div>

                    <div class="title-box">
                        <Icon icon="material-symbols:link"/>
                        <div class="text">链接</div>
                    </div>

                    <div class="title-box">
                        <Icon icon="mdi:user"/>
                        <div class="text">我的</div>
                    </div>

                    <div class="title-box"  @mouseenter="otherShow = true" @mouseleave="otherShow = false">
                        <transition name="el-fade-in">
                            <title-card v-if="otherShow">
                                <el-menu>
                                    <div class="title-box">
                                        <Icon icon="ix:project"/>
                                        <div>项目展示</div>
                                    </div>

                                    <div class="title-box">
                                        <Icon icon="grommet-icons:technology"></Icon>
                                        <div>技术展示</div>
                                    </div>

                                    <div class="title-box"  @click="() => { console.log('点击了'); router.push('/other/algorithm'); }">
                                        <Icon icon="material-symbols:problem-outline"></Icon>
                                        <div>算法题解</div>
                                    </div>

                                    <div>时间线</div>
                                </el-menu>
                            </title-card>
                        </transition>


                        <Icon icon="basil:other-1-outline"/>
                        <div class="text">其他</div>
                        <Icon icon="oui:arrow-down"/>
                    </div>
                </div>

                <div class="module">
                    <Icon class="Component" :icon="LAD" @click="toggleDark"/>
                    <Icon class="Component" icon="material-symbols:palette-outline"/>
                    <Icon class="Component" icon="mdi:settings-outline" @click="settingShow = !settingShow"/>
                </div>
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

<style scoped lang="scss">
.home-header {
    width: 100%;
    height: 70vh;
    background-image: url("../assets/雪莉1.jpeg");
    background-size: cover;
    background-color: rgba(218, 246, 244, 0.3);
    background-blend-mode: overlay;
}

.home-top {
    display: flex;
    width: 100%;
    height: 10vh;
    align-items: center;
    color: black;
    background: rgba(218, 246, 244, 0.4);
    backdrop-filter: blur(5px);

    .main-title { /* 所有标题 */
        display: flex;
        align-items: center;
        gap: 3em;
        margin: 0 auto;
        font-size: v-bind(textSize);

        .title-box { /* 选中第一级盒子 标题列表写在了titleCard里面*/
            display: flex;
            align-items: center;
            gap: 3px;
            position: relative;
            cursor: pointer;
            transition: color 0.3s ease-out;

            &:hover {
                transition: color 0.1s ease-in;
                color: #377ed3;
            }

            Icon {
                font-size: v-bind(iconSize);
            }
        }
    }

    .module { /* 右边的图标模块 */
        display: flex;
        align-items: center;
        margin-right: 10px;
        font-size: v-bind(iconSize);
        gap: 0.7em;
        .Component {
            position: relative;
            cursor: pointer;
            color: rgb(76, 76, 76);
            transition: color 0.3s ease-out;
            &:hover {
                transition: color 0.1s ease-in;
                color: #377ed3;
            }
        }
    }
}

.dark {  /*深色模式适配*/
    .home-header{
        color: white;
        background-color: rgba(2, 12, 12, 0.3);
    }

    .home-top {
        color: rgba(255, 255, 255, 0.8);
        background: rgba(2, 12, 12, 0.35);
        .module{
            .Component{
                color: rgba(198, 196, 196);
            }
        }
    }
}
</style>