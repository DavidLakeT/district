import { createRouter, createWebHistory } from 'vue-router';
import Home from '@/components/Home.vue';
import Products from '@/components/Products.vue';
import ProductPage from '@/components/ProductPage.vue';


const routes = [
  { path: '/', component: Home },
  { path: '/products', component: Products },
  {
    path: '/products/:id',
    component: ProductPage,
    props: true
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
