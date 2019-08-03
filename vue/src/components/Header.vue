<template>
<el-menu :default-active="this.$route.path" router class="el-menu-demo" mode="horizontal" @select="handleSelect">
  <el-menu-item index="/data">数据中心</el-menu-item>
  <el-menu-item index="/">抽签中心</el-menu-item>
  <el-menu-item index="/result">抽签结果</el-menu-item>
</el-menu>
</template>

<script>
  export default {
    name: 'Header',
    data() {
      return {
        activeIndex: '/',
      };
    },
    methods: {
      handleSelect(key, keyPath) {
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
