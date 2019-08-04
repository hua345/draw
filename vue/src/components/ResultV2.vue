<template>
  <el-row type="flex" justify="center">
    <el-col :xs="24" :sm="18" :md="12" :lg="8">
      <el-table :data="drawResultBody" v-if="hasResult" stripe style="width: 100%">
        <el-table-column
          v-for="(val,index) in drawResultHeader"
          :key="index"
          :label="val"
        >
          <template slot-scope="scope">
            {{scope.row[index]}}
          </template>
        </el-table-column>
      </el-table>

      <el-row type="flex" justify="center">
        <el-col :xs="24" :sm="18" :md="12" :lg="8">
          <el-button round @click="toExcel" class="exportExcel">导出数据</el-button>
        </el-col>
      </el-row>
    </el-col>
  </el-row>
</template>

<script>
import axios from "axios";
export default {
  name: "Result",
  data: () => ({
    hasResult: false,
    drawResult: null,
    drawResultHeader: null,
    drawResultBody: null,
    selectedCompanyName: null
  }),
  mounted: function() {
    this.getDrawResult();
  },
  methods: {
    getDrawResult: function() {
      this.drawResult = this.$store.state.drawResult;
      this.selectedCompanyName = this.$store.state.selectedCompanyName;
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
        console.log("userTypeMap", userTypeMap)
        console.log("maxTypeName", maxTypeName)
        console.log("drawResultHeader",drawResultHeader)
        userTypeMap.get(maxTypeName).forEach(function(val, index) {
          var dataRow = [];
          drawResultHeader.forEach(function(valueType, typeIndex) {
            var currentList = userTypeMap.get(valueType);
            if (currentList.length >= index + 1) {
              console.log("currentList[index]",currentList, index, currentList[index])
              dataRow.push(currentList[index].name);
            } else {
              dataRow.push("");
            }
          });
          console.log(dataRow)
          drawResultBody.push(dataRow);
        });
        this.drawResultHeader = drawResultHeader;
        this.drawResultBody = drawResultBody;
        console.log(this.drawResultHeader);
        console.log(this.drawResultBody);
        this.hasResult = true;
      }
    },
    getNowFormatDate: function() {
      var date = new Date();
      var seperator1 = "_";
      var year = date.getFullYear();
      var month = date.getMonth() + 1;
      var strDate = date.getDate();
      if (month >= 1 && month <= 9) {
        month = "0" + month;
      }
      if (strDate >= 0 && strDate <= 9) {
        strDate = "0" + strDate;
      }
      var currentdate = year + seperator1 + month + seperator1 + strDate;
      return currentdate;
    },
    toExcel: function() {
      if (null == this.drawResult || null == this.selectedCompanyName) {
        this.$notify({
          title: "导出数据",
          message: "导出的数据为空,请先进行抽签",
          position: "top-left"
        });
        return;
      }
      axios({
        method: "post",
        url: process.env.VUE_APP_BASE_URL + "/api/v2/toExcel", // 请求地址
        responseType: "blob",
        data: {
          userList: this.drawResult,
          selectedCompanyName: this.selectedCompanyName
        }
      }).then(res => {
        var dateStr = this.getNowFormatDate();
        const fileName = this.selectedCompanyName + dateStr + ".xlsx";
        let blob = new Blob([res.data], {
          type:
            "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
        });
        if ("download" in document.createElement("a")) {
          // 非IE下载
          const elink = document.createElement("a");
          elink.download = fileName;
          elink.style.display = "none";
          elink.href = URL.createObjectURL(blob);
          document.body.appendChild(elink);
          elink.click();
          URL.revokeObjectURL(elink.href); // 释放URL 对象
          document.body.removeChild(elink);
        } else {
          // IE10+下载
          navigator.msSaveBlob(blob, fileName);
        }
      });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.exportExcel {
  margin: 30px auto;
}
</style>
