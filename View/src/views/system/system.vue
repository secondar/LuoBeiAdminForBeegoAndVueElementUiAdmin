<template>
  <div class="app-container">
    <el-container>
      <el-header>
        <sticky :z-index="10" class="sticky">
          <el-button
            style="margin-left: 10px"
            @click="handleSave"
            type="success"
          >
            保存
          </el-button>
        </sticky>
      </el-header>
      <el-main>
        <el-form ref="form" :model="formData" class="form-container">
          <el-form-item label-width="90px" label="网站标题:">
            <el-input
              v-model="formData.title"
              type="textarea"
              class="article-textarea"
              autosize
              placeholder="请填写标题"
            />
          </el-form-item>

          <el-form-item label-width="90px" label="网站小尾巴:">
            <el-input
              v-model="formData.tail"
              type="textarea"
              class="article-textarea"
              autosize
              placeholder="请填写网站小尾巴"
            />
          </el-form-item>

          <el-form-item label-width="90px" label="网站关键词:">
            <el-input
              v-model="formData.keyword"
              type="textarea"
              class="article-textarea"
              autosize
              placeholder="请填写网站关键词"
            />
          </el-form-item>

          <el-form-item label-width="90px" label="网站描述:">
            <el-input
              v-model="formData.describe"
              type="textarea"
              class="article-textarea"
              autosize
              placeholder="请填写网站描述"
            />
          </el-form-item>
        </el-form>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import { Message } from "element-ui";
import Sticky from "@/components/Sticky"; // 粘性header组件
export default {
  name: "articlesadd",
  components: {
    Sticky,
  },
  data() {
    return {
      formData: {
        id: 1,
        title: "",
        tail: "",
        keyword: "",
        describe: "",
      },
    };
  },
  mounted() {
    this.handleGetSystem();
  },
  methods: {
    handleGetSystem() {
      this.$store
        .dispatch("system/GetSystem", this.formData)
        .then((res) => {
          console.log(res);
          this.formData = res.data;
        })
        .catch(() => {});
    },
    handleSave() {
      this.$store
        .dispatch("system/SaveSystem", this.formData)
        .then((res) => {
          this.$alert(res.msg, "提示", {
            confirmButtonText: "确定",
            callback: (action) => {
              if (res.code == 200) {
              }
            },
          });
        })
        .catch(() => {});
    },
  },
};
</script>
<style type="less">
.sticky {
  background: #fff;
  text-align: right;
}
.form-container {
  margin-right: 3rem;
}
.el-header {
  height: auto !important;
}
.flex {
  display: flex;
  justify-content: flex-start;
}
</style>
