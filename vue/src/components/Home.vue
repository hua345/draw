<template>
  <el-row type="flex" justify="center">
    <el-col :xs="24" :sm="18" :md="12" :lg="8">
      <el-row type="flex" justify="center">
        <el-col :xs="24" :sm="18" :md="12" :lg="8">
          <el-button round @click="beginLottery" class="beginDraw">开始抽签</el-button>
        </el-col>
      </el-row>

      <el-table :data="drawResult" v-if="null != drawResult" stripe style="width: 100%">
        <el-table-column prop="name" label="姓名" width="180"></el-table-column>
        <el-table-column label="人员类型" width="180">
          <template slot-scope="scope">
            <el-tag type="success">{{scope.row.type}}</el-tag>
          </template>
        </el-table-column>
      </el-table>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: "Home",
  data: () => ({
    show3: true,
    isInit: false,
    drawResult: null,
    userTypeMap: null,
    selectedTypeInfoList: null,
    selectedCompanyName: null
  }),
  mounted: function() {
    this.getCompanyList();
  },
  methods: {
    getCompanyList: function() {
      if (null != this.$store.state.excelData) {
        if (
          null == this.$store.state.selectedCompanyName ||
          null == this.$store.state.selectedTypeInfoList
        ) {
          this.$notify({
            title: "抽签设置",
            message: "请先进行抽签设置",
            position: "top-left"
          });
          this.$router.push({ path: "/data" });
          return;
        }
        var excelData = this.$store.state.excelData;
        this.isInit = true;
        this.selectedTypeInfoList = this.$store.state.selectedTypeInfoList;
        this.selectedCompanyName = this.$store.state.selectedCompanyName;

        var selectedCompany = this.$store.state.excelData.filter(item => {
          return item.companyName == this.selectedCompanyName;
        });
        selectedCompany = selectedCompany[0];

        var userTypeMap = new Map();
        selectedCompany.userList.forEach(item => {
          let userTypeInfo = userTypeMap.get(item.type);
          if (undefined != userTypeInfo && undefined != userTypeInfo.userList) {
            userTypeInfo.userList.push(item);
            userTypeMap.set(item.type, userTypeInfo);
          } else {
            userTypeMap.set(item.type, { userList: [item], lotteryNum: 0 });
          }
        });
        this.selectedTypeInfoList.forEach(item => {
          let userTypeInfo = userTypeMap.get(item.type);
          if (undefined != userTypeInfo) {
            userTypeInfo.lotteryNum = item.lotteryNum;
          }
        });
        this.userTypeMap = userTypeMap;
      } else {
        this.$notify({
          title: "抽签设置",
          message: "请先导入数据",
          position: "top-left",
          type: 'warning'
        });
        this.$router.push({ path: "/data" });
      }
    },
    beginLottery: function() {
      var result = [];
      this.userTypeMap.forEach(function(userTypeInfo, userType, map) {
        if (userTypeInfo.lotteryNum >= 1) {
          var drawUserList = [...userTypeInfo.userList];
          for (var i = 0; i < userTypeInfo.lotteryNum; i++) {
            if (0 == drawUserList.length) {
              break;
            }
            let randowmNum = Math.floor(Math.random() * drawUserList.length);
            let deleteElem = drawUserList.splice(randowmNum, 1);
            result.push(deleteElem[0]);
          }
        }
      });
      this.drawResult = result;
      this.$store.commit("SET_DrawResult", result);
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.beginDraw {
  margin: 30px auto;
}
</style>
