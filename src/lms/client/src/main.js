// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import api from './api/index.js'
import utils from './utils/index.js'
// lang
import VueI18n from 'vue-i18n'
import lang from './config/language.js'
// iview
import iView from 'iview'
import 'iview/dist/styles/iview.css'
Vue.use(VueI18n)
Vue.use(iView, {
  i18n: function (path, options) {
    let value = i18n.t(path, options)
    if (value !== null && value !== undefined) return value
    return ''
  }
})

const i18n = new VueI18n({
  locale: 'en', // 默认值
  messages: lang
})

Vue.prototype.$api = api
Vue.prototype.$utils = utils

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  i18n,
  template: '<App/>',
  components: { App }
})
