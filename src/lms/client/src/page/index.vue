<template>
  <div>
    <!-- <i-Switch v-model="loading" @on-change="change" stripe></i-Switch> -->
    <Table border ref="selection" :loading="loading" :columns="columns3" :data="list"></Table>
    <div style="margin: 10px;overflow: hidden">
        <div style="float: right;">
            <Page :total="total" :current="page" @on-change="changePage" @on-page-size-change="changePageSize" show-sizer show-total placement="top"></Page>
        </div>
    </div>
  </div>
</template>

<script>
import LHeader from '../components/header.vue'
import LFooter from '../components/footer.vue'
import WangEditor from '../components/wangeditor.vue'
export default {
  components: { LHeader, LFooter, WangEditor },
  data () {
    return {
      list: [],
      spanLeft: 5,
      spanRight: 19,
      total: 0,
      page: 1,
      limit: 10,
      menuText: '中文',
      loading: false,
      columns3: [
        {
          type: 'selection',
          width: 60,
          align: 'center'
        },
        {
          title: '标题',
          key: 'title',
          render: (h, params) => {
            return h('div', [
              h('a', {
                props: {
                  type: 'primary',
                  size: 'small'
                },
                style: {
                  marginRight: '5px'
                },
                on: {
                  click: () => {
                    this.show(params.row.title)
                    this.$router.push('/artcle/' + params.row.id)
                  }
                }
              }, params.row.title)
            ])
          }
        },
        {
          title: '创建时间',
          key: 'create_at'
        }
      ]
    }
  },
  created () {
    this.loading = true
    this.getData(this.page, this.limit)
  },
  computed: {
    iconSize () {
      return this.spanLeft === 5 ? 14 : 24
    }
  },
  methods: {
    getData (page, limit) {
      this.$api.get('topics', null, r => {
        this.total = r.data.length
      })
      this.$api.get('topics', { page: page, limit: limit }, r => {
        this.list = r.data
        this.loading = false
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
    changePage (page) {
      this.loading = true
      this.page = page
      // 这里直接更改了模拟的数据，真实使用场景应该从服务端获取数据
      this.list = this.getData(this.page, this.limit)
    },
    changePageSize (pageSize) {
      this.loading = true
      this.limit = pageSize
      this.list = this.getData(this.page, this.limit)
    },
    show (index) {
      console.log(index)
    }
  }
}
</script>
<style>
</style>
