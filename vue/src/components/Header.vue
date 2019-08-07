<template>
  <el-row>
    <el-col :span="24">
      <el-row type="flex" justify="center" style="text-align: center;">
        <el-col>
          <h2 class="my-title">{{drawName}}</h2>
        </el-col>
      </el-row>
      <el-row type="flex" justify="center">
        <el-col>
          <el-menu :default-active="this.$route.path" router class="el-menu-demo" mode="horizontal">
            <el-menu-item index="/">抽签中心</el-menu-item>
            <el-menu-item index="/result">抽签结果</el-menu-item>
          </el-menu>
        </el-col>
      </el-row>
    </el-col>
  </el-row>
</template>

<script>
import axios from "axios";

export default {
  name: "Header",
  data() {
    return {
      activeIndex: "/",
      drawName: ""
    };
  },
  mounted: function() {
    this.getDrawName();
  },
  methods: {
    getDrawName() {
      axios
        .get(process.env.VUE_APP_BASE_URL + "/api/v2/drawName")
        .then(response => {
          if (undefined != response.data) {
            this.drawName = response.data.data;
            document.title = this.drawName;
          }
        })
        .catch(function(error) {
          console.log(error);
          document.title = "抽签中心";
        });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.my-title{
  text-shadow: 5px 5px 5px #ca9393;
}
</style>
