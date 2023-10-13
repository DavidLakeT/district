import { createRouter, createWebHistory } from 'vue-router';
import Home from '@/components/Home.vue';
import Products from '@/components/Products.vue';

const routes = [
  { path: '/', component: Home },
  { path: '/products', component: Products },
  // Add more routes as needed
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
