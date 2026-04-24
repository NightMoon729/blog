import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'welcome',
            component: () => import('@/views/Welcome.vue'),

        },{
            path: '/other/algorithm',
            name: 'algorithm',
            component: () => import('@/views/other/Algorithm.vue')
        }
    ]
})

export default router