<template>
  <div id="app">
    <router-view></router-view>
    <BackTop></BackTop>
  </div>
</template>

<script>
  export default {
    name: 'app',
    components: { },
    data () {
      return {
      }
    },
    created () {
    },
    mounted () {
    },
    computed: {
    },
    methods: {
    }
  }
</script>

<style scoped>
  @import url("./style/base/_color.scss");
  @import url("./style/style.scss");
</style>
