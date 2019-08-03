import Vue from 'vue'
import App from './App.vue'
import './plugins/element.js'
import router from './router/index';
import store from './store/store'

Vue.config.productionTip = false

console.log(process.env);

new Vue({
  store,
  router,
  render: h => h(App),
}).$mount('#app')
