import { createApp } from 'vue'
import App from './App.vue'

import { createRouter, createWebHistory } from 'vue-router'

const routes = [
    { path: '/dogs', name: 'Dogs', component: () => import('@/components/Dogs.vue') },
    { path: '/dog/:collar', name: 'Dog', component: () => import('@/components/Dog.vue') },
    { path: '/dogs/new', name: 'DogsNew', component: () => import('@/components/DogsNew.vue') },
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

createApp(App).use(router).mount('#app')
