import Vue from 'vue';
import VueMaterial from 'vue-material';
import VueNativeSock from 'vue-native-websocket';
import 'vue-material/dist/vue-material.min.css';
import 'vue-material/dist/theme/default-dark.css';
import store from './store';
import App from './App.vue';

const l = window.location;

Vue.config.productionTip = false;

Vue.use(VueMaterial);
Vue.use(VueNativeSock, `${((l.protocol === 'https:') ? 'wss://' : 'ws://') + l.hostname + (((l.port !== '80') && (l.port !== '443')) ? `:${l.port}` : '')}/api/go-logsink/ws`, {
  reconnection: true, // (Boolean) whether to reconnect automatically (false)
  reconnectionAttempts: 5, // (Number) number of reconnection attempts before giving up (Infinity),
  reconnectionDelay: 3000, // (Number) how long to initially wait before attempting a new (1000)
  store,
  format: 'json',
});

new Vue({
  store,
  render: (h) => h(App),
}).$mount('#app');
