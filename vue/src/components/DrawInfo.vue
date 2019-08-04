<template>
  <div>
    <el-select
      v-if="isInit && null != companyList"
      v-model="value"
      @change="companyChange"
      placeholder="请选择"
      center="true"
    >
      <el-option v-for="item in companyList" :key="item" :label="item" :value="item"></el-option>
    </el-select>
    <el-row>
      <el-col :span="24">
        <el-table :data="selectedTypeInfoList" stripe style="width: 100%" center="true">
          <el-table-column label="人员类型" width="180">
            <template slot-scope="scope">
              <el-tag type="success">{{scope.row.type}}</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="抽奖人数" width="180">
            <template slot-scope="scope">
              <div>
                <el-input-number
                  v-model="scope.row.lotteryNum"
                  :min="0"
                  :max="scope.row.typeSize"
                  label="描述文字"
                ></el-input-number>
              </div>
            </template>
          </el-table-column>
        </el-table>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24">
        <el-button round @click="goLottery">去抽签</el-button>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "DrawInfo",
  data: function() {
    return {
      isInit: false,
      showSub: false,
      value: "",
      num: 1,
      companyList: null,
      selectedTypeInfoList: null,
      selectedCompanyName: null
    };
  },
  mounted: function() {
    this.getCompanyList();
  },
  methods: {
    getCompanyList: function() {
      if (null != this.$store.state.excelData) {
        var excelData = this.$store.state.excelData;
        this.isInit = true;
        this.companyList = this.$store.state.excelData.map(item => {
          return item.companyName;
        });
        console.log(this.companyList);
      }
    },
    companyChange: function(selectedCompanyName) {
      if (
        null == this.$store.state.excelData ||
        undefined == this.$store.state.excelData
      ) {
        return [];
      }
      this.selectedCompanyName = selectedCompanyName;
      var selectedCompany = this.$store.state.excelData.filter(item => {
        return item.companyName == selectedCompanyName;
      });
      selectedCompany = selectedCompany[0];
      console.log(selectedCompany);
      if (null != selectedCompany && undefined != selectedCompany) {
        var typeMap = new Map();
        selectedCompany.userList.forEach(item => {
          let num = typeMap.get(item.type);
          if (null != num && undefined != num && num >= 1) {
            typeMap.set(item.type, num + 1);
          } else {
            typeMap.set(item.type, 1);
          }
        });
        var typeInfoList = [];
        typeMap.forEach(function(value, key, map) {
          var typeInfo = {
            type: key,
            typeSize: value,
            lotteryNum: 0
          };
          typeInfoList.push(typeInfo);
        });
        this.selectedTypeInfoList = typeInfoList;
        this.showSub = true;
      }
    },
    goLottery() {
      if (
        null != this.selectedCompanyName &&
        null != this.selectedTypeInfoList
      ) {
        this.$store.commit("SET_SelectedCompanyName", this.selectedCompanyName);
        this.$store.commit(
          "SET_SelectedTypeInfoList",
          this.selectedTypeInfoList
        );
        var lotteryFlag = this.selectedTypeInfoList.some(infoValue => {
          return infoValue.lotteryNum >= 1;
        });
        if (lotteryFlag) {
          this.$notify({
            title: "抽签设置",
            message: "去抽签",
            position: "top-left"
          });
          this.$router.push({ path: "/" });
        } else {
          this.$notify({
            title: "抽签设置",
            message: "抽奖人数不能都为零",
            position: "top-left"
          });
        }
      } else {
        this.$notify({
          title: "抽签设置",
          message: "需要先选择单位哦",
          position: "top-left"
        });
      }
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
