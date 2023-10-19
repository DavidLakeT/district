import { createRouter, createWebHistory } from 'vue-router';
import Home from '@/views/HomeView.vue';
import Products from '@/views/ProductsView.vue';
import ProductPage from '@/views/ProductPageView.vue';
import MyCart from '@/views/CartView.vue'; 
import Register from '@/views/Auth/RegisterView.vue';
import Login from  '@/views/Auth/LoginView.vue';
import UserProfile from '@/views/UserProfileView.vue';



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
    path: '/profile/:userId',
    component: UserProfile
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
