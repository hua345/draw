<template>
  <el-row>
    <el-col :span="24">
      <el-tabs v-if="isUpload" v-model="dataDetailActive">
        <el-tab-pane
          v-for="(item,tabIndex) in dataDetailList"
          :label="item.companyName"
          :key="tabIndex"
        >
          <el-table :data="item.drawResultBody" stripe style="width: 100%">
            <el-table-column v-for="(val,index) in item.drawResultHeader" :key="index" :label="val">
              <template slot-scope="scope">{{scope.row[index]}}</template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
      </el-tabs>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: "DataDetail",
  data: function() {
    return {
      isUpload: false,
      excelData: null,
      dataDetailList: null,
      dataDetailActive: "0"
    };
  },
  mounted: function() {
    this.getUploadStatus();
  },
  methods: {
    getUploadStatus: function() {
      this.excelData = this.$store.state.excelData;
      if (
        null != this.excelData &&
        undefined != this.excelData &&
        this.excelData.length >= 1
      ) {
        var dataDetailList = [];
        for(var i =0; i <this.excelData.length; i++){
          var company = this.excelData[i]
          var dataTable = this.formatData(company.userList);
          var data = {
            companyName: company.companyName,
            drawResultHeader: dataTable.drawResultHeader,
            drawResultBody: dataTable.drawResultBody
          };
          dataDetailList.push(data);
        };
        this.dataDetailList = dataDetailList;
      }
      this.isUpload = this.$store.state.isUpload;
    },
    formatData: function(drawResult) {
      if (
        null != drawResult &&
        undefined != drawResult &&
        drawResult.length >= 1
      ) {
        var userTypeMap = new Map();
        drawResult.forEach(item => {
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
        return {
          drawResultHeader: drawResultHeader,
          drawResultBody: drawResultBody
        };
      }
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
