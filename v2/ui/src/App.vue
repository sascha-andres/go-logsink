<template>
  <md-app md-mode="overlap">
    <md-app-toolbar class="md-primary">
      <md-button class="md-icon-button" @click="menuVisible = !menuVisible">
        <md-icon>menu</md-icon>
      </md-button>
      <img src="/assets/logo.png" style="width: 30px;" alt="Logo" />
      <span class="md-title">go-logsink</span>
    </md-app-toolbar>
    <md-app-drawer :md-active.sync="menuVisible">
      <md-toolbar class="md-transparent" md-elevation="0">Settings</md-toolbar>
      <md-list>
        <md-list-item>
          <md-field>
            <label>Limit row numbers</label>
            <md-input :value="this.$store.rowLimit" @input="updateRowLimit" type="number"></md-input>
          </md-field>
        </md-list-item>
        <md-list-item>
          <md-checkbox v-model="scrolling" @change="updateScrolling">Auto scroll</md-checkbox>
        </md-list-item>
        <md-list-item>
          <md-field>
            <label>Filter output</label>
            <md-input :value="this.$store.filter" @input="updateFilter" v-model="type"></md-input>
          </md-field>
        </md-list-item>
        <md-subheader>Links</md-subheader>
        <md-list-item href="https://go-logsink.livingit.de/" target="_blank">
          Homepage
        </md-list-item>
        <md-list-item href="https://github.com/sascha-andres/go-logsink" target="_blank">
          GitHub
        </md-list-item>
        <md-subheader>Connect</md-subheader>
        <md-list-item href="https://livingit.de" target="_blank">
          My homepage
        </md-list-item>
        <md-list-item href="https://www.xing.com/profile/Sascha_Andres" target="_blank">
          My XING profile
        </md-list-item>
        <md-list-item href="https://www.linkedin.com/in/sascha-andres-b7b91935/" target="_blank">
          My LinkedIn profile
        </md-list-item>
      </md-list>
    </md-app-drawer>
    <md-app-content>
      <pre>
        <div v-for="item in this.$store.getters.filteredLines" :key="item.Key">{{ item.Line }}</div>
      </pre>
    </md-app-content>
  </md-app>
</template>

<style lang="scss" scoped>
  .md-app {
    min-height: 100vh;
    border: 1px solid rgba(#000, .12);
  }

   // Demo purposes only
  .md-drawer {
    width: 230px;
    max-width: calc(100vw - 125px);
  }
</style>

<script>
export default {
  name: 'Reveal',
  data: () => ({
    scrolling: true,
    menuVisible: false,
  }),
  methods: {
    updateRowLimit(limit) { this.$store.commit('setRowLimit', limit); },
    updateFilter(filter) { this.$store.commit('setFilter', filter); },
    updateScrolling() { this.$store.commit('setScrolling', !this.$store.state.scrolling); },
  },
};
</script>
