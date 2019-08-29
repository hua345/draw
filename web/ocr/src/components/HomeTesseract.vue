<template>
  <el-row>
    <el-col :span="24">
      <el-upload
        class="upload-demo"
        action
        :auto-upload="false"
        :on-change="handleChange"
        :on-preview="handlePreview"
        :file-list="fileList"
        list-type="picture"
      >
        <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
        <el-button style="margin-left: 10px;" size="small" type="success" @click="beginOCR">开始OCR</el-button>
        <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div>
      </el-upload>

      <el-table :data="tableData" style="width: 100%">
        <el-table-column prop="text" label="OCR内容"></el-table-column>
        <el-table-column prop="index" label="OCR行号"></el-table-column>
      </el-table>
    </el-col>
  </el-row>
</template>
<script>
import { TesseractWorker, PSM, OEM } from "tesseract.js";
const worker = new TesseractWorker({
      workerPath: './tesseract/worker.min.js',
      langPath: './lang-data',
      corePath: './tesseract/tesseract-core.wasm.js',
    }
);

export default {
  name: "HomeTesseract",
  data() {
    return {
      fileList: [],
      tableData: []
    };
  },
  methods: {
    beginOCR() {
      console.log(this.fileList);
      // worker
      //   .recognize(img, "eng", {
      //     tessedit_ocr_engine_mode: OEM.LSTM_ONLY,
      //     tessedit_pageseg_mode: PSM.SINGLE_BLOCK
      //   })
      //   .then(result => {
      //     console.log(result);
      //   });
    },
    handleChange(file, fileList) {
      console.log(file.url);
      var img = new Image();
      img.src = file.url;
      console.log(img);
      worker
        .recognize(img, "chi_sim", {
          tessedit_ocr_engine_mode: OEM.LSTM_ONLY,
          tessedit_pageseg_mode: PSM.SINGLE_BLOCK
        })
        .progress(p => {
          console.log("progress", p);
        })
        .then(result => {
          console.log(result);
          var tableData = [];
          result.lines.forEach(function(value, index) {
            tableData.push({
              text: value.text,
              index: index
            });
          });
          this.tableData = tableData;
        });
    },
    handlePreview(file) {
      console.log(file);
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
