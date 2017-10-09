<template>
  <div>
    <canvas id="canvas" class="canvas"></canvas>
    <div class="login-box">
      <div class="login-logo">
      <a href="/"><b>{{$i18n.t('message.garbage')}}</b>{{$i18n.t('message.tms')}}</a>
    </div>
  <!-- /.login-logo -->
  <div class="login-box-body">
    <p class="login-box-msg">{{ $t('message.loginSession') }}</p>
    <Form ref="formLogin" :model="formLogin" :rules="ruleLogin">
      <FormItem prop="user">
          <Input type="text" v-model="formLogin.user" v-bind:placeholder="$i18n.t('message.userNameMail')">
              <Icon type="person" slot="prepend" :size="18"></Icon>
          </Input>
      </FormItem>
      <FormItem prop="password">
          <Input type="password" v-model="formLogin.password" v-bind:placeholder="$t('message.password')">
              <Icon type="locked" slot="prepend" :size="18"></Icon>
          </Input>
      </FormItem>
      <FormItem>
          <Checkbox v-model="formLogin.rememberme">{{ $t('message.rememberme') }}</Checkbox>
          <Button type="primary" @click="handleSubmit('formLogin')">{{ $t('message.login') }}</Button>
          <!-- <a href="javascript:void(0);" @click="handleSubmit('formLogin')" style="float:right;">{{ $t('message.login') }}</a> -->
      </FormItem>
      <FormItem>
        <p style="text-align:center;">- OR -</p>
        <a href="register.html" class="text-center">{{ $t('message.register') }}</a>
        <a href="#" style="float:right;">{{ $t('message.forgotPassword') }}</a>
      </FormItem>
    </Form>
  </div>
  <!-- /.login-box-body -->
</div>
    
  </div>
</template>
<script>
  import drawCanvas from '../utils/line-canvas.js'  // 导入首页Canvas动画
  export default {
    name: 'login',
    components: { },
    data () {
      const validateUsername = (rule, value, callback) => {
        if (value === '') {
          callback(new Error(this.$i18n.t('message.inputUserName')))
        } else {
          callback()
        }
      }
      const validatePass = (rule, value, callback) => {
        if (value === '') {
          callback(new Error(this.$i18n.t('message.inputPass')))
        } else if (value.length < 6) {
          callback(new Error(this.$i18n.t('message.passTooSimple')))
        } else {
          callback()
        }
      }
      return {
        formLogin: {
          user: '',
          password: '',
          rememberme: false
        },
        ruleLogin: {
          user: [
            { validator: validateUsername, trigger: 'blur' }
          ],
          password: [
            { validator: validatePass, trigger: 'blur' }
          ]
        }
      }
    },
    mounted () {
      drawCanvas()
    },
    created () {
      var lang = localStorage.getItem('language')
      if (lang != null) {
        this.$i18n.locale = lang
      }
    },
    methods: {
      handleSelectAll (status) {
        this.$refs.selection.selectAll(status)
      },
      handleSubmit (name) {
        var $this = this
        this.$refs[name].validate((valid) => {
          if (valid) {
            debugger
            this.$api.get('topics', $this.formLogin, r => {
            })
            this.$Message.success(this.$i18n.t('message.validateSuccess'))
          } else {
            this.$Message.error(this.$i18n.t('message.validateFail'))
          }
        })
      },
      handleReset (name) {
        this.$refs[name].resetFields()
      }
    }
  }
</script>

<style>
  @import '../style/login.scss'
</style>
