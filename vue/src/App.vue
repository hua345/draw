<template>
  <el-container >
    <el-header style="height: 120px;">
      <el-row :gutter="20" type="flex" justify="center">
        <el-col :xs="24" :sm="18" :md="12" :lg="8">
          <Header></Header>
        </el-col>
      </el-row>
    </el-header>
    <el-main>
      <el-row :gutter="20" type="flex" justify="center">
        <!-- <el-col :md="6" :lg="8">
          <img :src="imgList[0]" style="width: 100%;" />
        </el-col> -->
        <el-col :md="12" :lg="8" style="background-color:#FFF">
          <router-view></router-view>
        </el-col>
        <!-- <el-col :md="6" :lg="8">
          <img :src="imgList[1]" style="width: 100%;" />
        </el-col> -->
      </el-row>
      <!-- <el-row>
        <el-col :span="24">
          <el-carousel :interval="4000" type="card" height="500px">
            <el-carousel-item v-for="item in imgList" :key="item">
              <img :src="item" />
            </el-carousel-item>
          </el-carousel>
        </el-col>
      </el-row> -->
    </el-main>
  </el-container>
</template>

<script>
import Header from "./components/Header.vue";

export default {
  name: "app",
  components: {
    Header
  },
  mounted() {
  document.querySelector('body').setAttribute('style', "background-repeat: no-repeat;background-size: 100% auto;background-image: url(" + process.env.VUE_APP_BASE_URL + "/public/img/img01.jpg" + ")")
},
  data: function() {
    return {
      imgList: [
        process.env.VUE_APP_BASE_URL + "/public/img/img01.jpg",
        process.env.VUE_APP_BASE_URL + "/public/img/img02.jpg",
        process.env.VUE_APP_BASE_URL + "/public/img/img03.jpg",
        process.env.VUE_APP_BASE_URL + "/public/img/img04.jpg"
      ],  bgImg: {
                backgroundImage: "url(" + process.env.VUE_APP_BASE_URL + "/public/img/img02.jpg" + ")",
                backgroundRepeat: "no-repeat",
                backgroundSize: "100% auto"
          }
    };
  },
  methods: {
    beforeRouteEnter(to, from, next) {
      // 在渲染该组件的对应路由被 confirm 前调用
      // 不！能！获取组件实例 `this`
      // 因为当钩子执行前，组件实例还没被创建
      next(vm => {
        console.log("hello")
        document.querySelector("html").style.cssText = `
        background: url(${vm.imgList[0]}) center no-repeat;
        background-size: cover;
        background-attachment: fixed;
      `;
      });
    }
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
