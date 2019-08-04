<template>
  <el-row type="flex" justify="center">
    <el-col :xs="24" :sm="18" :md="12" :lg="8">
      <el-row type="flex" justify="center">
        <el-col :xs="24" :sm="18" :md="12" :lg="8">
          <el-button round @click="beginLottery" class="beginDraw">开始抽签</el-button>
        </el-col>
      </el-row>

      <el-table :data="drawResultBody" v-if="hasResult" stripe style="width: 100%">
        <el-table-column v-for="(val,index) in drawResultHeader" :key="index" :label="val">
          <template slot-scope="scope">{{scope.row[index]}}</template>
        </el-table-column>
      </el-table>
    </el-col>
  </el-row>
</template>

<script>
import DataTable from "./DataTable.vue";

export default {
  name: "Home",
  components: {
    DataTable
  },
  data: () => ({
    show3: true,
    isInit: false,
    drawResult: null,
    userTypeMap: null,
    selectedTypeInfoList: null,
    selectedCompanyName: null,
    drawResultHeader: null,
    drawResultBody: null,
    hasResult: false
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
          type: "warning"
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
      if (
        null != this.drawResult &&
        undefined != this.drawResult &&
        this.drawResult.length >= 1
      ) {
        var userTypeMap = new Map();
        this.drawResult.forEach(item => {
          let userList = userTypeMap.get(item.type);
          if (undefined != userList) {
            userList.push(item);
            userTypeMap.set(item.type, userList);
          } else {
            userTypeMap.set(item.type, [item]);
          }
        });
        var drawResultHeader = [];
        var drawResultBody = [];
        //获取数据类型列表
        userTypeMap.forEach(function(valueUserList, keyUserType) {
          drawResultHeader.push(keyUserType);
        });
        //取长度最长的数据类型
        var maxTypeName = "";
        drawResultHeader.forEach(function(valueType, index) {
          if (index == 0) {
            maxTypeName = valueType;
          } else {
            var userList = userTypeMap.get(valueType);
            var originUserList = userTypeMap.get(maxTypeName);
            if (userList.length > originUserList.length) {
              maxTypeName = valueType;
            }
          }
        });
        //获取表数据
        userTypeMap.get(maxTypeName).forEach(function(val, index) {
          var dataRow = [];
          drawResultHeader.forEach(function(valueType, typeIndex) {
            var currentList = userTypeMap.get(valueType);
            if (currentList.length >= index + 1) {
              dataRow.push(currentList[index].name);
            } else {
              dataRow.push("");
            }
          });
          drawResultBody.push(dataRow);
        });
        this.drawResultHeader = drawResultHeader;
        this.drawResultBody = drawResultBody;
        this.hasResult = true;
      }
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
