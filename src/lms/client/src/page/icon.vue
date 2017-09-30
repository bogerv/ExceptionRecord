<template>
  <div>
    <Icon type="edit"></Icon>
    <Icon type="document"></Icon>
    <Icon type="document-text"></Icon>
    <Icon type="trash-a"></Icon>
    <Icon type="settings"></Icon>
    <Icon type="ion-arrow-up-b"></Icon>
    <Icon type="arrow-up-b"></Icon>
    <br>
    <Collapse v-model="pValue">
        <Panel name="1" class="primary">
          {{ $t("message.alert") }}
          <p slot="content">
            <Button @click="success">显示成功提示</Button>
            <Button @click="warning">显示警告提示</Button>
            <Button @click="error">显示错误提示</Button>
          </p>
        </Panel>
        <Panel name="2">
            斯蒂夫·盖瑞·沃兹尼亚克
            <p slot="content">斯蒂夫·盖瑞·沃兹尼亚克（Stephen Gary Wozniak），美国电脑工程师，曾与史蒂夫·乔布斯合伙创立苹果电脑（今之苹果公司）。斯蒂夫·盖瑞·沃兹尼亚克曾就读于美国科罗拉多大学，后转学入美国著名高等学府加州大学伯克利分校（UC Berkeley）并获得电机工程及计算机（EECS）本科学位（1987年）。</p>
        </Panel>
        <Panel name="3">
            乔纳森·伊夫
            <p slot="content">乔纳森·伊夫是一位工业设计师，现任Apple公司设计师兼资深副总裁，英国爵士。他曾参与设计了iPod，iMac，iPhone，iPad等众多苹果产品。除了乔布斯，他是对苹果那些著名的产品最有影响力的人。</p>
        </Panel>
    </Collapse>
  </div>
</template>

<<script>
export default {
  components: { },
  data () {
    return {
      pValue: '1'
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
    success () {
      this.$Message.success('这是一条成功的提示')
    },
    warning () {
      this.$Message.warning('这是一条警告的提示')
    },
    error () {
      this.$Message.error('对方不想说话，并且向你抛出了一个异常')
    }
  }
}
</script>
<style>
</style>
