<template>
  <div id="main">
    <div class="layout" :class="{'layout-hide-text': spanLeft < 5}">
        <Row type="flex">
          <div class="layout-header">
            <Col :span="spanLeft" style="height:100%;">
              <div class="layout-header-logo">
                <a href="/" class="wrapper-header-nav-logo router-link-active" style="color:White;">Bogerv<span class="layout-text">Admin</span></a>
              </div>
            </Col>
            <Col :span="spanRight" style="height:100%;">
              <Button type="text" @click="toggleClick" style="margin-top: 5px;">
                <Icon type="navicon" size="30" color="white"></Icon>
              </Button>
              <Dropdown style="float: right;margin-top: 2px;margin-right: 5px;" @on-click="profileOpt">
                  <Button type="primary" style="margin-top: 2px;">
                    <Avatar src="https://i.loli.net/2017/08/21/599a521472424.jpg" />
                  </Button>
                  <DropdownMenu slot="list" style="text-align: center;min-width: 0px">
                    <DropdownItem name="profile">{{ $t("message.profile") }}</DropdownItem>
                    <DropdownItem divided name="logout">{{ $t("message.logout") }}</DropdownItem>
                    <!-- <Button type="text" size="small" @click.native="logout">{{ $t("message.logout") }}</Button> -->
                  </DropdownMenu>
              </Dropdown>
              <Dropdown style="float: right;padding: 10px 10px 10px 0;" @on-click="changeText">
                  <Button type="primary">
                    {{menuText}}
                    <Icon type="arrow-down-b"></Icon>
                  </Button>
                  <DropdownMenu slot="list" style="text-align: center;min-width: 60px">
                    <DropdownItem name="zh">中文</DropdownItem>
                    <!-- <DropdownItem disabled>豆汁儿</DropdownItem> -->
                    <DropdownItem name="en">English</DropdownItem>
                  </DropdownMenu>
              </Dropdown>
            </Col>
          </div>
        </Row>
        <Row type="flex">
            <Col :span="spanLeft" class="layout-menu-left">
              <Menu theme="dark" :open-names="['1']" width="auto" accordion @on-select="menuSelect">
                <Submenu name="1">
                    <template slot="title">
                        <Icon type="ios-paper" :size="iconSize"></Icon>
                        <span class="layout-text">{{ $t("message.content") }}{{ $t("message.manangement") }}</span>
                    </template>
                    <MenuItem name="artcles">{{ $t("message.artcle") }}<span class="layout-text">{{ $t("message.manangement") }}</span></MenuItem>
                    <MenuItem name="general">{{ $t("message.comment") }}<span class="layout-text">{{ $t("message.manangement") }}</span></MenuItem>
                    <MenuItem name="icons">{{ $t("message.icon") }}<span class="layout-text">{{ $t("message.manangement") }}</span></MenuItem>
                    <MenuItem name="1">
                      <Icon type="ios-navigate" :size="iconSize"></Icon>
                      <span class="layout-text">选项 1</span>
                    </MenuItem>
                </Submenu>
                <Submenu name="2">
                    <template slot="title">
                        <Icon type="ios-people" :size="iconSize"></Icon>
                        <span class="layout-text">用户管理</span>
                    </template>
                    <MenuItem name="2-1">新增用户</MenuItem>
                    <MenuItem name="2-2">活跃用户</MenuItem>
                </Submenu>
                <Submenu name="3">
                    <template slot="title">
                      <Icon type="stats-bars" :size="iconSize"></Icon>
                        <span class="layout-text">统计分析</span>
                    </template>
                    <MenuGroup title="使用">
                        <MenuItem name="3-1">新增和启动</MenuItem>
                        <MenuItem name="3-2">活跃分析</MenuItem>
                        <MenuItem name="3-3">时段分析</MenuItem>
                    </MenuGroup>
                    <MenuGroup title="留存">
                        <MenuItem name="3-4">用户留存</MenuItem>
                        <MenuItem name="3-5">流失用户</MenuItem>
                    </MenuGroup>
                </Submenu>
              </Menu>
            </Col>
            <Col :span="spanRight" :offset="spanLeft" style="float:right;">
              <div class="layout-breadcrumb">
              <Breadcrumb>
                  <BreadcrumbItem href="#">首页</BreadcrumbItem>
                  <BreadcrumbItem href="#">应用中心</BreadcrumbItem>
                  <BreadcrumbItem>某应用</BreadcrumbItem>
              </Breadcrumb>
            </div>
            <div class="layout-content">
              <div class="layout-content-main">
                <router-view></router-view>
              </div>
            </div>
            <div class="layout-copy">
                <!-- 2017-2017 &copy TalkingData -->
                2017 &copy Bogerv
            </div>
          </Col>
        </Row>
      </div>
  </div>
</template>
<script>
export default {
  name: 'main',
  components: { },
  data () {
    return {
      list: [],
      spanLeft: 5,
      spanRight: 19,
      menuText: ''
    }
  },
  created () {
    this.getData()
    var lang = localStorage.getItem('language')
    if (lang != null) {
      this.$i18n.locale = lang
    }
    if (this.$i18n.locale === 'en') {
      this.menuText = 'English'
    } else {
      this.menuText = '中文'
    }
  },
  mounted () {
  },
  computed: {
    iconSize () {
      return this.spanLeft === 5 ? 14 : 24
    }
  },
  methods: {
    profileOpt (name) {
      if (name === 'logout') {
        this.logout()
      } else {
      }
    },
    logout () {
      this.$router.push('/login')
    },
    handleSelectAll (status) {
      this.$refs.selection.selectAll(status)
    },
    getData () {
      this.$api.get('topics', null, r => {
        this.list = r.data
      })
    },
    toggleClick () {
      if (this.spanLeft === 5) {
        this.spanLeft = 2
        this.spanRight = 22
      } else {
        this.spanLeft = 5
        this.spanRight = 19
      }
    },
    changeText (name) {
      if (name === 'en') {
        this.menuText = 'English'
      } else {
        this.menuText = '中文'
      }
      this.$i18n.locale = name
      localStorage.setItem('language', name)
      // location.reload()
    },
    menuSelect (name) {
      this.$router.push(name)
    }
  }
}
</script>