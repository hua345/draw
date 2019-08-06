<template>
  <el-table :data="drawResultBody" v-if="hasResult" stripe style="width: 100%">
    <el-table-column v-for="(val,index) in drawResultHeader" :key="index" :label="val">
      <template slot-scope="scope">{{scope.row[index]}}</template>
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  name: "DataTable",
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
      if (null != this.drawResult && undefined != this.drawResult) {
        var drawResultTable = this.formatData(this.drawResult);
        this.drawResultHeader = drawResultTable.drawResultHeader;
        this.drawResultBody = drawResultTable.drawResultBody;
        this.hasResult = true;
      }
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
