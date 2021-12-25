<template>
  <div class="app-container">
    <el-container>
      <el-header>
        <el-form ref="form" :model="SearchForm" label-width="80px">
          <el-row>
            <el-col :span="6"></el-col>
            <el-col :span="6">
              <el-form-item label="关键词">
                <el-input v-model="SearchForm.key"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="分类">
                <treeselect
                  v-model="SearchForm.sort"
                  :multiple="false"
                  :options="SortList.tree"
                />
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-button
                style="margin-left: 0.5rem"
                type="primary"
                icon="el-icon-search"
                size="success"
                @click="handleGetList(true)"
                >搜索</el-button
              >
              <el-button
                type="primary"
                icon="el-icon-refresh-right"
                size="info"
                @click="SearchFormReset"
                >重置</el-button
              >
              <el-button
                style="margin-left: 0.5rem"
                type="primary"
                icon="el-icon-refresh-right"
                size="success"
                @click="handleGetList()"
                >刷新</el-button
              >
            </el-col>
          </el-row>
        </el-form>
      </el-header>
      <el-main>
        <el-table
          ref="multipleTable"
          :data="ArticleList.list"
          tooltip-effect="dark"
          style="width: 100%"
        >
          <el-table-column prop="id" label="ID"> </el-table-column>
          <el-table-column prop="title" label="标题"> </el-table-column>
          <el-table-column prop="sort" label="分类">
            <template slot-scope="scope">
              {{ handleGetSortIdResCn(scope.row.sort) }}
            </template>
          </el-table-column>
          <el-table-column prop="hot" label="点击数"> </el-table-column>
          <el-table-column prop="addtime" label="添加时间"> </el-table-column>
          <el-table-column label="操作" fixed="right" width="118"
            ><template slot-scope="scope">
              <el-button
                v-permission="'ArticleDetails'"
                size="mini"
                class="el-icon-edit"
                type="primary"
                @click="handleEdit(scope.$index, scope.row)"
              />
              <el-button
                v-permission="'ArticleDelete'"
                size="mini"
                class="el-icon-delete"
                type="danger"
                @click="handleDelete(scope.$index, scope.row)"
              />
            </template>
          </el-table-column>
        </el-table>
        <pagination
          v-show="ArticleList.total > 0"
          :total="ArticleList.total"
          :page.sync="ArticleList.current"
          :limit.sync="ArticleList.limit"
          @pagination="handleGetList"
        />
      </el-main>
    </el-container>
  </div>
</template>

<script>
import { Message } from "element-ui";
import Treeselect from "@riophae/vue-treeselect";
import "@riophae/vue-treeselect/dist/vue-treeselect.css";
import Pagination from "@/components/Pagination";

import permission from "@/directive/permission/index.js"; // 权限判断指令
export default {
  name: "articlelist",
  components: { Treeselect, Pagination },
  directives: { permission },
  data() {
    return {
      ViewControl: {
        Dialog: {
          Show: false,
          Title: "新增用户",
        },
      },
      ArticleList: {
        current: 1,
        last: 1,
        list: [],
        total: 0,
        limit: 30,
      },
      SortList: {
        tree: [],
        list: [],
      },
      SearchForm: {
        key: "",
        sort: 0,
        limit: 30,
        page: 1,
      },
    };
  },
  mounted() {
    this.handleGetSortList();
  },
  methods: {
    handleAddBtn() {
      this.formData = {
        account: "",
        password: "",
        role: "",
        state: 1,
      };
      this.ViewControl.Dialog.Title = "新增用户";
      this.handleGetSortList();
    },
    handleDialogHide() {
      this.ViewControl.Dialog.Show = false;
    },
    handleEdit(index, row) {
      this.$router.push({
        name: "articleedit",
        params: { id: row.id },
      });
    },
    handleDelete(index, row) {
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
          this.$store.dispatch("article/Delete", row.id).then((res) => {
            if (res.code == 200) {
              Message({
                message: res.msg || "Error",
                type: "success",
                duration: 5 * 1000,
              });
              this.handleGetList();
            }
          });
        })
        .catch(() => {});
    },
    handleGetList(SearchBtn = false) {
      if (!SearchBtn) {
        this.SearchForm.limit = this.ArticleList.limit;
        this.SearchForm.page = this.ArticleList.current;
      }
      this.$store.dispatch("article/GetList", this.SearchForm).then((res) => {
        if (res.code == 200) {
          this.ArticleList = res.data;
        }
      });
    },
    handleGetSortList() {
      this.$store
        .dispatch("article/GetSort")
        .then((res) => {
          if (res.code != 200) {
            Message({
              message: res.msg || "Error",
              type: "error",
              duration: 5 * 1000,
            });
            return;
          }
          this.SortList.list = this.handleGetSortListTreeToArr(res.data);
          this.SortList.tree = [{ id: 0, label: "请选择" }].concat(
            this.handleTreeSelectOptions(res.data)
          );
          this.handleGetList();
        })
        .catch((err) => {});
    },
    handleGetSortListTreeToArr(arrData) {
      let arrRes = [];
      for (let i in arrData) {
        let arrTmp = [].concat(arrData[i]);
        delete arrTmp.children;
        arrRes = arrTmp;
        if (arrData[i].children != null && arrData[i].children.length > 0) {
          arrTmp = this.handleGetSortListTreeToArr(arrData[i].children);
          arrRes = arrRes.concat(arrTmp);
        }
      }
      return arrRes;
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
    handleGetSortIdResCn(id) {
      let SortName = "未知分类";
      for (let i in this.SortList.list) {
        if (id == this.SortList.list[i].id) {
          SortName = this.SortList.list[i].title;
          break;
        }
      }
      return SortName;
    },
    SearchFormReset() {
      this.SearchForm = {
        key: "",
        sort: 0,
        limit: 30,
        page: 1,
      };
      this.handleGetList();
    },
  },
};
</script>
