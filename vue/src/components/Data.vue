<template>
  <el-row type="flex" justify="center">
    <el-col :xs="24" :sm="18" :md="12" :lg="8">
      <el-tabs v-model="activeName">
        <el-tab-pane label="上传文件" name="upload">
          <el-upload
            class="upload-demo"
            drag
            :action="getUrl"
            multiple
            accept=".xlsx, .xls"
            :on-success="uploadSuccess"
            :on-error="uploadFailed"
          >
            <i class="el-icon-upload"></i>
            <div class="el-upload__text">
              将文件拖到此处，或
              <em>点击上传</em>
            </div>
            <div class="el-upload__tip" slot="tip">只能上传xlsx/xls文件，且不超过8M</div>
          </el-upload>
        </el-tab-pane>
        <el-tab-pane label="抽签设置" name="drawInfo">
          <DrawInfo v-if="isUpload"></DrawInfo>
        </el-tab-pane>
        <el-tab-pane label="查看数据" name="third">
          <DataDetail v-if="isUpload"></DataDetail>
        </el-tab-pane>
      </el-tabs>
    </el-col>
  </el-row>
</template>

<script>
import DrawInfo from "./DrawInfo.vue";
import DataDetail from "./DataDetail.vue";

export default {
  name: "Data",
  components: {
    DrawInfo,
    DataDetail
  },
  data: function() {
    return {
      isUpload: false,
      activeName: "upload",
      getUrl: process.env.VUE_APP_BASE_URL + "/api/v2/upload"
    };
  },
  mounted: function() {
    this.getUploadStatus();
  },
  methods: {
    getUploadStatus: function() {
      this.isUpload = this.$store.state.isUpload;
      if (this.isUpload) {
        this.activeName = "drawInfo";
      } else {
        this.activeName = "upload";
      }
    },
    uploadSuccess: function(response, file, fileList) {
      this.$store.commit("SET_EXCEL_DATA", response.data);
      this.isUpload = true;
      this.$store.commit("SET_IsUpload", true);
      this.$notify({
        title: "抽签设置",
        message: "上传文件成功",
        position: "top-left",
        type: "success"
      });
      this.activeName = "drawInfo";
    },
    uploadFailed: function(err, file, fileList) {
      console.log(err);
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
/* .upload-demo {
  padding: 20px 0px;
  margin: 0 auto;
  width: 300px;
} */
</style>
