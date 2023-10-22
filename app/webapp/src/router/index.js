import { createRouter, createWebHistory } from 'vue-router';
import Home from '@/views/HomeView.vue';
import Products from '@/views/ProductsView.vue';
import ProductPage from '@/views/ProductPageView.vue';
import MyCart from '@/views/CartView.vue'; 
import Register from '@/views/Auth/RegisterView.vue';
import Login from  '@/views/Auth/LoginView.vue';
import UserProfile from '@/views/UserProfileView.vue';
import ProductCreateForm from  '@/views/Admin/ProductCreateView.vue'
import store from '@/store';



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
    component: MyCart,
    meta: { requiresAuth: true },
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
    component: UserProfile,
    meta: { requiresAuth: true },
  },
  {
    path: '/admin/product/create',
    component: ProductCreateForm,
    meta: { requiresAuth: true },
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    if (!store.getters['auth/isAuthenticated']) {
      next('/login');
    } else {
      next();
    }
  } else {
    next();
  }
});



export default router;
