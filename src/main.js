import Vue from 'vue';
import App from './App.vue';
import axios from 'axios';

Vue.prototype.$http = axios;

Vue.prototype.serverURI = 'http://localhost:8000';

Vue.config.productionTip = false;

new Vue({
  render: h => h(App)
}).$mount('#app');
