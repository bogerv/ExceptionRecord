<template>
  <div>
    <WangEditor></WangEditor>
  </div>
</template>

<<script>
import WangEditor from '../components/wangeditor.vue'
export default {
  components: { WangEditor },
  data () {
    return {
      list: []
    }
  },
  created () {
  },
  computed: {
    iconSize () {
      return this.spanLeft === 5 ? 14 : 24
    }
  },
  methods: {
  }
}
</script>
<style>
</style>
