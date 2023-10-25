import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min';
import { readAuthToken, decodeAuthToken } from '@/js/auth.js'; 

const app = createApp(App);
app.use(store);
app.use(router);
const authToken = readAuthToken();
if (authToken) {
    const user = decodeAuthToken(authToken);
    store.commit('auth/setAuthToken', authToken);
    store.commit('auth/setUserId', user.userId);
    store.commit('auth/setUserRole', user.isAdmin);
}

app.mount('#app');