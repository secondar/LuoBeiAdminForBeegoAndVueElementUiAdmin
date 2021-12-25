<template>
  <div class="app-container">
    <el-container>
      <el-header>
        <el-form ref="form" :model="SearchForm" label-width="80px">
          <el-row>
            <el-col :span="6"
              ><el-button
                v-permission="'AddUser'"
                type="primary"
                icon="el-icon-plus"
                size="small"
                @click="handleAddBtn"
                >新增</el-button
              ></el-col
            >
            <el-col :span="6">
              <el-form-item label="用户名">
                <el-input v-model="SearchForm.account"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="角色">
                <el-select v-model="SearchForm.role" placeholder="请选择">
                  <el-option
                    v-for="item in RoleList"
                    :key="item.id"
                    :label="item.title"
                    :value="item.id"
                  >
                  </el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-button
                type="primary"
                icon="el-icon-search"
                size="success"
                @click="handleGetAdminList"
                >搜索</el-button
              >

              <el-button
                type="primary"
                icon="el-icon-refresh-right"
                size="info"
                @click="SearchFormReset"
                >重置</el-button
              >
            </el-col>
          </el-row>
        </el-form>
      </el-header>
      <el-main>
        <el-table
          ref="multipleTable"
          :data="AdminList"
          tooltip-effect="dark"
          style="width: 100%"
        >
          <el-table-column prop="id" label="ID"> </el-table-column>
          <el-table-column prop="account" label="用户名"> </el-table-column>
          <el-table-column prop="rolecn" label="角色"> </el-table-column>
          <el-table-column prop="state" label="状态">
            <template slot-scope="scope">
              <span>{{ scope.row.state == 1 ? "启用" : "禁用" }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="addtime" label="添加时间"> </el-table-column>
          <el-table-column prop="address" label="操作" fixed="right" width="118"
            ><template slot-scope="scope">
              <el-button
                v-permission="'EditUser'"
                size="mini"
                class="el-icon-edit"
                type="primary"
                @click="handleEdit(scope.$index, scope.row)"
              />
              <el-button
                v-permission="'DeleteUser'"
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
      <!-- 新增用户 -->
      <el-dialog
        :title="ViewControl.Dialog.Title"
        width="40%"
        :visible.sync="ViewControl.Dialog.Show"
      >
        <el-form ref="form" :model="formData" label-width="80px">
          <el-form-item label="用户名">
            <el-input v-model="formData.account" placeholder="用户名" />
          </el-form-item>
          <el-form-item label="密码">
            <el-input v-model="formData.password" placeholder="密码" />
          </el-form-item>
          <el-form-item label="是否启用">
            <el-radio-group v-model="formData.state" size="small">
              <el-radio-button label="1">是</el-radio-button>
              <el-radio-button label="0">否</el-radio-button>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="角色">
            <el-select v-model="formData.role" placeholder="请选择">
              <el-option
                v-for="item in RoleList"
                :key="item.id"
                :label="item.title"
                :value="item.id"
              >
              </el-option>
            </el-select>
          </el-form-item>
        </el-form>

        <div slot="footer" class="dialog-footer">
          <el-button @click="handleDialogHide()">取 消</el-button>
          <el-button type="primary" @click="handleSave()">确 定</el-button>
        </div>
      </el-dialog>
      <!-- 新增用户结束 -->
    </div>
    <!-- 弹出层结束 -->
  </div>
</template>

<script>
import { Message } from "element-ui";
import permission from "@/directive/permission/index.js"; // 权限判断指令
export default {
  name: "user",
  directives: { permission },
  data() {
    return {
      ViewControl: {
        Dialog: {
          Show: false,
          Title: "新增用户",
        },
      },
      AdminList: [],
      RoleList: [],
      formData: {
        account: "",
        password: "",
        role: "",
        state: 1,
      },
      SearchForm: {
        account: "",
        role: "",
      },
    };
  },
  mounted() {
    this.handleGetAdminList();
    this.handleGetRoleList(false);
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
      this.handleGetRoleList();
    },
    handleDialogHide() {
      this.ViewControl.Dialog.Show = false;
    },
    handleEdit(index, row) {
      this.formData = {
        id: row.id,
        account: row.account,
        password: "",
        role: row.role,
        state: row.state,
      };
      this.ViewControl.Dialog.Title = "编辑用户";
      this.handleGetRoleList();
    },
    handleDelete(index, row) {
      this.$store.dispatch("user/Delete", row.id).then((res) => {
        if (res.code == 200) {
          Message({
            message: res.msg || "Error",
            type: "success",
            duration: 5 * 1000,
          });
          this.handleDialogHide();
          this.handleGetAdminList();
        }
      });
    },
    handleSave() {
      let store = "user/Edit";
      if (this.formData.id == undefined || this.formData == null) {
        store = "user/Add";
      }
      this.$store.dispatch(store, this.formData).then((res) => {
        if (res.code == 200) {
          Message({
            message: res.msg || "Error",
            type: "success",
            duration: 5 * 1000,
          });
          this.handleDialogHide();
          this.handleGetAdminList();
        }
      });
    },
    handleGetAdminList() {
      this.$store.dispatch("user/GetAdminList", this.SearchForm).then((res) => {
        if (res.code == 200) {
          this.AdminList = res.data;
        }
      });
    },
    handleGetRoleList(ShowDialog = true) {
      this.$store
        .dispatch("role/GetList")
        .then((res) => {
          if (res.code != 200) {
            Message({
              message: res.msg || "Error",
              type: "error",
              duration: 5 * 1000,
            });
            return;
          }
          this.RoleList = res.data;
          if (ShowDialog) this.ViewControl.Dialog.Show = true;
        })
        .catch(() => {});
    },
    SearchFormReset() {
      this.SearchForm = {
        account: "",
        role: "",
      };
      this.handleGetAdminList();
    },
  },
};
</script>
