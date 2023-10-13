import { createRouter, createWebHistory } from 'vue-router';
import Home from '@/components/Home.vue';
import Products from '@/components/Products.vue';
import ProductPage from '@/components/ProductPage.vue';
import MyCart from '@/components/MyCart.vue'; 
import Register from '@/components/Register.vue';
import Login from '@/components/Login.vue';
import UserProfile from '@/components/UserProfile.vue';



const routes = [
  { 
    path: '/', 
    component: Home 
  },
  { 
    path: '/products', 
    component: Products 
  },
  {
    path: '/products/:id',
    component: ProductPage,
    props: true
  },
  {
    path: '/mycart',
    component: MyCart
  },
  {
    path: '/register',
    component: Register
  },
  {
    path: '/login',
    component: Login
  },
  {
    path: '/profile/:id',
    component: UserProfile
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
