import { createApp } from 'vue'
import router from './router/index.ts';
import App from './App.vue'
import Antd from 'ant-design-vue';


const app = createApp(App)
app.use(Antd).use(router).mount('#app')
