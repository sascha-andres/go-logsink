import Vue from 'vue';
import Vuex from 'vuex';
import LogLine from './models/logline';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    socket: {
      isConnected: false,
      reconnectError: false,
    },
    logLines: Array<LogLine>(),
    rowLimit: 0,
    filter: '',
    scrolling: true,
  },
  mutations: {
    SOCKET_ONOPEN(state, event) {
      Vue.prototype.$socket = event.currentTarget;
      state.socket.isConnected = true;
    },
    SOCKET_ONCLOSE(state, event) {
      state.socket.isConnected = false;
    },
    SOCKET_ONERROR(state, event) {
      console.error(state, event);
    },
    // default handler called for all methods
    SOCKET_ONMESSAGE(state, message) {
      state.logLines.push(message);
      if (state.rowLimit !== 0) {
        while (state.logLines.length > state.rowLimit) {
          state.logLines.shift();
        }
      }
      if (state.scrolling) {
        window.scrollTo(0, document.body.scrollHeight);
      }
    },
    // mutations for reconnect methods
    SOCKET_RECONNECT(state, count) {
      console.info(state, count);
    },
    SOCKET_RECONNECT_ERROR(state) {
      state.socket.reconnectError = true;
    },
    addLogLine(state, logModel: LogLine) {
      state.logLines.push(logModel);
    },
    setRowLimit(state, limit: number) {
      state.rowLimit = limit;
    },
    setFilter(state, filter: string) {
      state.filter = filter;
    },
    setScrolling(state, value: boolean) {
      state.scrolling = value;
    },
  },
  actions: {
    addLogLine(context, logModel: LogLine) {
      context.commit('addLogLine', logModel);
    },
    setRowLimit(context, limit: number) {
      context.commit('setRowLimit', limit);
    },
    setScrolling(context, value: boolean) {
      context.commit('setScrolling', value);
    },
    setFilter(context, filter: string) {
      context.commit('setFilter', filter);
    },
  },
  getters: {
    filteredLines: (state) => state.logLines.filter((line) => line.Line.match(state.filter)),
  },

});
