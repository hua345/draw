<template>
  <el-row>
    <el-col :span="24">
      <el-tabs v-if="isUpload" v-model="dataDetailActive">
        <el-tab-pane v-for="(item,index) in excelData" :label="item.companyName" :key="index" >
          <el-table :data="item.userList" stripe style="width: 100%">
            <el-table-column prop="name" label="姓名" width="180"></el-table-column>
            <el-table-column label="人员类型" width="180">
              <template slot-scope="scope">
                <el-tag type="success">{{scope.row.type}}</el-tag>
              </template>
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
      dataDetailActive: "0"
    };
  },
  mounted: function() {
    this.getUploadStatus();
  },
  methods: {
    getUploadStatus: function() {
      this.isUpload = this.$store.state.isUpload;
      this.excelData = this.$store.state.excelData;
      if (this.isUpload) {
      }
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
