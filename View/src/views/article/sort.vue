<template>
  <div class="app-container">
    <el-container>
      <el-header>
        <el-row>
          <el-button
            v-permission="'ArticleSortAdd'"
            type="primary"
            icon="el-icon-plus"
            size="small"
            @click="handleAddBtn"
            >新增</el-button
          >
        </el-row>
      </el-header>
      <el-main>
        <el-table
          ref="multipleTable"
          :data="MenuList"
          style="width: 100%; margin-bottom: 20px"
          row-key="id"
          default-expand-all
          tooltip-effect="dark"
          :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
        >
          <el-table-column prop="title" label="分类标题" sortable width="200" />
          <el-table-column prop="state" label="启用">
            <template slot-scope="scope">
              <span>{{ scope.row.state == 1 ? "是" : "否" }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="sort" label="排序" width="155" />
          <el-table-column prop="addtime" label="创建日期" width="155" />
          <el-table-column prop="address" label="操作" fixed="right" width="118"
            ><template slot-scope="scope">
              <el-button
                v-permission="'ArticleSortEdit'"
                size="mini"
                class="el-icon-edit"
                type="primary"
                @click="handleEdit(scope.$index, scope.row)"
              />
              <el-button
                v-permission="'ArticleSortDelete'"
                size="mini"
                class="el-icon-delete"
                type="danger"
                @click="handleDelete(scope.$index, scope.row)"
              />
            </template>
          </el-table-column>
        </el-table>
      </el-main>
    </el-container>
    <!-- 弹出层开始 -->
    <div id="dialog">
      <!-- 新增菜单 -->
      <el-dialog
        :title="ViewControl.Dialog.Title"
        width="40%"
        :visible.sync="ViewControl.Dialog.Show"
      >
        <el-form ref="form" :model="formData" label-width="80px">
          <el-form-item label="分类标题">
            <el-input v-model="formData.title" placeholder="分类标题" />
          </el-form-item>

          <el-form-item label="是否启用">
            <el-radio-group v-model="formData.state" size="small">
              <el-radio-button label="1">是</el-radio-button>
              <el-radio-button label="0">否</el-radio-button>
            </el-radio-group>
          </el-form-item>

          <el-form-item label="分类排序">
            <el-input v-model="formData.sort" placeholder="分类排序" />
          </el-form-item>

          <el-form-item label="上级目录">
            <treeselect
              v-model="formData.pid"
              :multiple="false"
              :options="TreeSelectOptions"
            />
          </el-form-item>
        </el-form>

        <div slot="footer" class="dialog-footer">
          <el-button @click="handleDialogHide()">取 消</el-button>
          <el-button type="primary" @click="handleSave()">确 定</el-button>
        </div>
      </el-dialog>
      <!-- 新增菜单结束 -->
    </div>
    <!-- 弹出层结束 -->
  </div>
</template>

<script>
import Treeselect from "@riophae/vue-treeselect";
import "@riophae/vue-treeselect/dist/vue-treeselect.css";
import { Message } from "element-ui";
import permission from "@/directive/permission/index.js"; // 权限判断指令
export default {
  name: "Menus",
  directives: { permission },
  components: { Treeselect },
  data() {
    return {
      ViewControl: {
        Dialog: {
          Show: false,
          Title: "新增分类",
        },
      },
      formData: {
        pid: 0,
        title: "",
        state: 1,
        sort: 999,
      },
      MenuList: [],
      TreeSelectOptions: [],
    };
  },
  mounted() {
    this.handleGetSortList();
  },
  methods: {
    handleAddBtn() {
      this.formData.id = null;
      this.formData.pid = 0;
      this.formData.title = "";
      this.formData.state = 1;
      this.formData.sort = 999;
      this.ViewControl.Dialog.Show = true;
    },
    handleDialogHide() {
      this.ViewControl.Dialog.Show = false;
    },
    handleGetSortList() {
      this.$store
        .dispatch("article/GetSort")
        .then((res) => {
          this.MenuList = res.data;
          this.TreeSelectOptions = [{ id: 0, label: "顶级分类" }].concat(
            this.handleTreeSelectOptions(this.MenuList)
          );
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
    handleSave() {
      let store = "article/AddSort";
      if (
        this.formData.id !== undefined &&
        this.formData.id != null &&
        this.formData.id > 0
      ) {
        store = "article/EditSort";
      }
      this.$store
        .dispatch(store, this.formData)
        .then((res) => {
          this.handleDialogHide();
          Message({
            message: res.msg || "Error",
            type: "success",
            duration: 5 * 1000,
          });
          this.handleGetSortList();
        })
        .catch(() => {});
    },
    handleEdit(index, data) {
      this.formData = {
        pid: data.pid,
        state: data.state,
        title: data.title,
        sort: data.sort,
      };
      this.formData.id = data.id;
      this.ViewControl.Dialog.Show = true;
    },
    handleDelete(index, data) {
      this.$confirm(
        "此操作将删除该分类及其分类下的所有文章, 是否继续?",
        "提示",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning",
        }
      )
        .then(() => {
          this.$store
            .dispatch("article/DeleteSort", data.id)
            .then((res) => {
              Message({
                message: res.msg || "Error",
                type: "success",
                duration: 5 * 1000,
              });
              this.handleGetSortList();
            })
            .catch(() => {});
        })
        .catch(() => {});
    },
  },
};
</script>
<style type="less">
.el-header {
  height: auto !important;
}
.flex {
  display: flex;
  justify-content: flex-start;
}
</style>
