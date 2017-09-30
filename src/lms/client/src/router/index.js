import Vue from 'vue'
import iView from 'iview'
import Router from 'vue-router'
import Home from '@/page/home'
import Index from '@/page/index'
import Content from '@/page/content'
import General from '@/page/general'
import Icon from '@/page/icon'
import Login from '@/page/login'

Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '/login',
      component: Login
    },
    {
      path: '/',
      component: Home,
      children: [{
        path: '/artcles',
        component: Index
      }, {
        path: '/general',
        component: General
      }, {
        path: '/icons',
        component: Icon
      }, {
        path: '/artcle/:id',
        component: Content
      }]
    },
    { path: '*', redirect: { path: '/login' } }
  ]
})

// 加载进度条
iView.LoadingBar.config({
  color: '#1c2438',
  failedColor: 'red',
  height: 3
})

router.afterEach(route => {
  iView.LoadingBar.finish()
})

// 验证 token，存在才跳转
router.beforeEach((to, from, next) => {
  // if (to.meta.requireAuth) {
  //   if (token) {
  //     next()
  //   } else {
  //     next({
  //       path: '/login'
  //       // , query: { redirect: to.fullPath }
  //     })
  //   }
  // } else {
  //   iView.LoadingBar.start()
  //   next()
  // }
  iView.LoadingBar.start()
  next()
})

export default router
