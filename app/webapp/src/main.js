/*import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

createApp(App).use(router).mount('#app')*/

import { createApp } from 'vue';
import App from './App.vue';
import router from './router'; // Assuming your router setup is in router.js
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min';


const app = createApp(App);
app.use(router);
app.mount('#app');


