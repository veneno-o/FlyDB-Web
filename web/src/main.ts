import './assets/main.css';
import 'element-plus/theme-chalk/dark/css-vars.css';
import 'element-plus/dist/index.css';
// 暗黑模式
import 'element-plus/theme-chalk/dark/css-vars.css';
import * as ElementPlusIconsVue from '@element-plus/icons-vue';
import { createApp } from 'vue';
import { createPinia } from 'pinia';
import ElementPlus from 'element-plus';
import App from './App.vue';
import router from './router';
const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(ElementPlus);
app.mount('#app');
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component);
}
