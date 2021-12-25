<template>
  <div class="upload-container">
    <el-button
      :style="{ background: color, borderColor: color }"
      icon="el-icon-upload"
      size="mini"
      type="primary"
      @click="dialogVisible = true"
    >
      上传图片
    </el-button>
    <el-dialog :visible.sync="dialogVisible">
      <el-upload
        :multiple="true"
        :file-list="fileList"
        :show-file-list="true"
        :on-remove="handleRemove"
        :on-success="handleSuccess"
        :before-upload="beforeUpload"
        class="editor-slide-upload"
        :action="uploadurl"
        :headers="headers"
        list-type="picture-card"
      >
        <el-button size="small" type="primary"> 点我上传 </el-button>
      </el-upload>
      <el-button @click="dialogVisible = false"> 取消 </el-button>
      <el-button type="primary" @click="handleSubmit"> 完成 </el-button>
    </el-dialog>
  </div>
</template>

<script>
import { GetUpLoadImgUrl } from "@/api/conf";
import { getToken } from "@/utils/auth";
export default {
  name: "EditorSlideUpload",
  props: {
    color: {
      type: String,
      default: "#1890ff",
    },
  },
  data() {
    return {
      dialogVisible: false,
      listObj: [],
      fileList: [],
      uploadurl: GetUpLoadImgUrl(),
      headers: {
        "X-Token": getToken(),
      },
    };
  },
  methods: {
    handleSubmit() {
      this.$emit("successCBK", this.listObj);
      this.fileList = [];
      this.dialogVisible = false;
    },
    handleSuccess(response, file) {
      if (response.code == 200) {
        this.listObj.push(response.msg);
      } else {
        handleRemove(file);
        this.$message(response.msg);
      }
    },
    handleRemove(file) {
      delete this.fileList[file.url];
    },
    beforeUpload(file) {
      const _self = this;
      const _URL = window.URL || window.webkitURL;
      const fileName = file.uid;
      this.listObj[fileName] = {};
      return new Promise((resolve, reject) => {
        const img = new Image();
        img.src = _URL.createObjectURL(file);
        img.onload = function () {
          _self.listObj[fileName] = {
            hasSuccess: false,
            uid: file.uid,
            width: this.width,
            height: this.height,
          };
        };
        resolve(true);
      });
    },
  },
};
</script>

<style lang="scss" scoped>
.editor-slide-upload {
  margin-bottom: 20px;
  ::v-deep .el-upload--picture-card {
    width: 100%;
  }
}
</style>
