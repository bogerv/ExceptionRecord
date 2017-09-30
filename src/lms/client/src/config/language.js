import localeEN from 'iview/dist/locale/en-US'
import localeCN from 'iview/dist/locale/zh-CN'

export default {
  en: Object.assign({
    message: {
      garbage: 'Garbage',
      tms: 'TMS',
      content: 'Content',
      hello: 'Hello World!',
      artcle: 'Artcle',
      comment: 'Genral',
      manangement: ' Mgt.',
      icon: 'Icon',
      alert: 'Global Alert',
      login: 'Sign In',
      rememberme: 'Remember me',
      loginSession: 'Sign in to start your session',
      userNameMail: 'Username or Email',
      forgotPassword: 'Forgot password ?',
      register: 'Sign Up',
      submit: 'Submit',
      reset: 'Reset',
      password: 'Password',
      confirmPass: 'Confirm Password',
      age: 'Age',
      inputUserName: 'Please input username or email',
      inputPass: 'Please input password',
      inputPassAgain: 'Please input password again',
      passTooSimple: 'Password too simple',
      passNotSame: 'Password not same',
      validateSuccess: 'Validate successful!',
      validateFail: 'Validate failure!',
      ageNotNull: 'Age can\'t be null',
      inputInt: 'Input interger',
      must18age: 'Age must lager than 18'
    }
  }, localeEN),
  zh: Object.assign({
    message: {
      garbage: '垃圾',
      tms: '系统',
      content: '内容',
      hello: '你好，世界！',
      artcle: '文章',
      comment: '评论',
      manangement: '管理',
      icon: '图标',
      alert: '全局提醒',
      login: '登录',
      rememberme: '记住密码',
      loginSession: '登录并开始使用',
      userNameMail: '用户名或邮箱',
      forgotPassword: '忘记密码 ?',
      register: '注册',
      submit: '提交',
      reset: '重置',
      password: '密码',
      confirmPass: '确认密码',
      age: '年龄',
      inputUserName: '请输入用户名或邮箱',
      inputPass: '请输入密码',
      inputPassAgain: '请再次输入密码',
      passTooSimple: '密码超过6位数',
      passNotSame: '两次输入密码不一致!',
      validateSuccess: '验证成功!',
      validateFail: '表单验证失败!',
      ageNotNull: '年龄不能为空',
      inputInt: '请输入数字值',
      must18age: '必须年满18岁'
    }
  }, localeCN)
}
