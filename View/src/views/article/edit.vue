<template>
  <div class="app-container">
    <el-container>
      <el-header>
        <sticky :z-index="10" class="sticky">
          <el-button type="warning" @click="handleRessetForm"> 重置 </el-button>
          <el-button
            v-permission="'ArticleEdit'"
            style="margin-left: 10px"
            @click="handleSave"
            type="success"
          >
            发布
          </el-button>
        </sticky>
      </el-header>
      <el-main>
        <el-form ref="form" :model="formData" class="form-container">
          <el-form-item label-width="70px" label="标题:">
            <el-input
              v-model="formData.title"
              type="textarea"
              class="article-textarea"
              autosize
              placeholder="请填写标题"
            />
          </el-form-item>

          <el-form-item label-width="70px" label="摘要:">
            <el-input
              v-model="formData.describe"
              type="textarea"
              class="article-textarea"
              autosize
              placeholder="请填写摘要，留空将自动填写"
            />
          </el-form-item>

          <el-form-item label="分类:" label-width="70px">
            <div class="tree flex">
              <treeselect
                v-model="formData.sort"
                :multiple="false"
                :options="TreeSelectOptions"
              />
            </div>
          </el-form-item>

          <el-form-item style="margin-bottom: 40px" label-width="70px">
            <Tinymce
              ref="editor"
              :value="formData.content"
              v-model="formData.content"
              :height="400"
            />
          </el-form-item>
        </el-form>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import Tinymce from "@/components/Tinymce";
import Sticky from "@/components/Sticky"; // 粘性header组件
import Treeselect from "@riophae/vue-treeselect";
import "@riophae/vue-treeselect/dist/vue-treeselect.css";
import { Message } from "element-ui";
import permission from "@/directive/permission/index.js"; // 权限判断指令
export default {
  name: "articlesedit",
  directives: { permission },
  components: {
    Tinymce,
    Treeselect,
    Sticky,
  },
  data() {
    return {
      ViewControl: {},
      TreeSelectOptions: [],
      formData: {
        id: 0,
        sort: 0,
        title: "",
        content: "",
        describe: "",
      },
    };
  },
  created() {
    let id = this.$route.params.id;
    if (id == undefined || id == null || id <= 0) {
      this.$store.dispatch("tagsView/delView", this.$route);
      this.$router.push({ path: "/" });
    }
    this.formData.id = id;
  },
  mounted() {
    this.handleGetSortList();
  },
  methods: {
    handleGetSortList() {
      this.$store
        .dispatch("article/GetSort")
        .then((res) => {
          this.TreeSelectOptions = [{ id: 0, label: "请选择" }].concat(
            this.handleTreeSelectOptions(res.data)
          );
          this.handleGetDetails();
        })
        .catch(() => {});
    },
    handleTreeSelectOptions(tree) {
      const ResData = [];
      for (const i in tree) {
        const item = {};
        item.id = tree[i].id;
        item.label = tree[i].title;
        if (tree[i].children.length > 0) {
          item.children = this.handleTreeSelectOptions(tree[i].children);
        }
        ResData.push(item);
      }
      return ResData;
    },
    handleRessetForm() {
      this.formData = {};
      this.formData = {
        id: null,
        sort: 0,
        title: "",
        content: "",
        describe: "",
      };
    },
    handleSave() {
      this.$store
        .dispatch("article/Edit", this.formData)
        .then((res) => {
          this.$alert(res.msg, "提示", {
            confirmButtonText: "确定",
            callback: (action) => {
              if (res.code == 200) {
                this.$store.dispatch("tagsView/delView", this.$route);
                this.$router.go(-1);
              }
            },
          });
        })
        .catch(() => {});
    },
    handleGetDetails() {
      this.$store
        .dispatch("article/Details", this.formData.id)
        .then((res) => {
          if (res.code != 200) {
            this.$store.dispatch("tagsView/delView", this.$route);
            this.$router.push({ path: "/" });
            return;
          } else {
            this.formData = res.data;
          }
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
