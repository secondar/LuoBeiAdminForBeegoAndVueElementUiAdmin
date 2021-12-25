<template>
  <div class="app-container">
    <el-container>
      <el-header
        ><el-row>
          <el-button
            v-permission="'AddRole'"
            type="primary"
            icon="el-icon-plus"
            size="small"
            @click="handleBtnAdd()"
            >新增</el-button
          >
        </el-row></el-header
      >
      <el-main>
        <el-table
          ref="multipleTable"
          :data="RoleList"
          tooltip-effect="dark"
          style="width: 100%"
        >
          <el-table-column label="角色名称" prop="title" />
          <el-table-column prop="state" label="状态">
            <template slot-scope="scope">
              {{ scope.row.state == 1 ? "启用" : "禁用" }}
            </template>
          </el-table-column>
          <el-table-column prop="address" label="备注">
            <template slot-scope="scope">
              {{ scope.row.remarks != null ? scope.row.remarks : "-" }}
            </template>
          </el-table-column>
          <el-table-column prop="addtime" label="添加时间" />
          <el-table-column prop="address" label="操作">
            <template slot-scope="scope">
              <el-button
                v-permission="'EditRole'"
                size="mini"
                class="el-icon-edit"
                type="primary"
                @click="handleBtnEdit(scope.$index, scope.row)"
              />
              <el-button
                v-permission="'SettingRole'"
                size="mini"
                class="el-icon-setting"
                type="success"
                @click="handleBtnSettingRouter(scope.$index, scope.row)"
              />

              <el-button
                v-permission="'DeleteRole'"
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
    <div class="dialog">
      <!-- 角色弹出层开始 -->
      <el-dialog
        :title="ViewControl.Dialog.Role.Title"
        :visible.sync="ViewControl.Dialog.Role.Show"
        width="40%"
      >
        <el-form ref="form" :model="formData" label-width="80px">
          <el-form-item label="角色名称" style="width: 95%">
            <el-input v-model="formData.title" placeholder="角色名称" />
          </el-form-item>
          <el-form-item label="是否启用">
            <el-radio-group v-model="formData.state" size="small">
              <el-radio-button label="1">是</el-radio-button>
              <el-radio-button label="0">否</el-radio-button>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="角色备注">
            <el-input v-model="formData.remarks" type="textarea" />
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
          <el-button @click="ViewControl.Dialog.Role.Show = false"
            >取 消</el-button
          >
          <el-button type="primary" @click="handleSave()">确 定</el-button>
        </span>
      </el-dialog>
      <!-- 角色弹出层结束 -->
      <!-- 权限弹出层开始 -->
      <el-dialog
        title="编辑权限"
        :visible.sync="ViewControl.Dialog.Router.Show"
        width="40%"
      >
        <el-form ref="form" label-width="80px">
          <el-form-item label="角色名称" style="width: 95%">
            {{ ViewControl.Dialog.Router.Role }}
          </el-form-item>
          <el-form-item label="分配权限" style="width: 95%">
            <el-tree
              :data="RouterList"
              show-checkbox
              node-key="id"
              :default-checked-keys="RouterListCheck"
              :props="TreeDefaultProps"
              ref="RouterTree"
            >
            </el-tree>
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
          <el-button @click="ViewControl.Dialog.Router.Show = false"
            >取 消</el-button
          >
          <el-button type="primary" @click="handleSettingRouterSave()"
            >确 定</el-button
          >
        </span>
      </el-dialog>
      <!-- 权限弹出层结束 -->
    </div>
    <!-- 弹出层结束 -->
  </div>
</template>

<script>
import { Message } from "element-ui";
import permission from "@/directive/permission/index.js"; // 权限判断指令
export default {
  name: "Role",
  directives: { permission },
  data() {
    return {
      ViewControl: {
        Dialog: {
          Role: {
            Show: false,
            Title: "新增角色",
          },
          Router: {
            Show: false,
            Role: "超级管理员",
          },
        },
      },
      formData: {
        title: "",
        state: 1,
        remarks: "",
      },
      RoleList: [],
      RouterList: [],
      RouterListCheck: [],
      SettingRouterForm: {
        role: null,
        router: null,
      },
      TreeDefaultProps: {
        children: "children",
        label: "label",
      },
    };
  },
  mounted() {
    this.handleGetRoleList();
  },
  methods: {
    handleBtnAdd() {
      this.formData.title = "";
      this.formData.state = 1;
      this.formData.remarks = "";
      this.ViewControl.Dialog.Role.Title = "新增角色";
      this.ViewControl.Dialog.Role.Show = true;
    },
    handleBtnEdit(index, row) {
      this.formData = {
        id: row.id,
        title: row.title,
        state: row.state,
        remarks: row.remarks,
      };
      this.ViewControl.Dialog.Role.Title = "编辑角色";
      this.ViewControl.Dialog.Role.Show = true;
    },
    handleBtnSettingRouter(index, row) {
      this.SettingRouterForm.role = row.id;
      this.handleGetMenuList();
    },
    handleSettingRouterSave() {
      let RouterIds = this.$refs.RouterTree.getCheckedKeys();
      this.SettingRouterForm.router = RouterIds.join(",");
      this.$store
        .dispatch("router/SettingRouter", this.SettingRouterForm)
        .then((res) => {
          Message({
            message: res.msg || "Error",
            type: "success",
            duration: 5 * 1000,
          });
          this.ViewControl.Dialog.Router.Show = false;
        });
    },
    handleDelete(index, row) {
      this.$confirm(
        "此操作将删除该角色,改角色下的用户将不能正常登陆使用, 是否继续?",
        "提示",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning",
        }
      )
        .then(() => {
          this.$store
            .dispatch("role/Delete", row.id)
            .then((res) => {
              Message({
                message: res.msg || "Error",
                type: "success",
                duration: 5 * 1000,
              });
              this.handleGetRoleList();
            })
            .catch(() => {});
        })
        .catch(() => {});
    },
    handleSave() {
      let store = "role/Add";
      if (this.formData.id !== undefined && this.formData.id > 0) {
        store = "role/Edit";
      }
      this.$store.dispatch(store, this.formData).then((res) => {
        this.ViewControl.Dialog.Role.Show = false;
        Message({
          message: res.msg || "Error",
          type: "success",
          duration: 5 * 1000,
        });
        this.handleGetRoleList();
      });
    },
    handleGetRoleList() {
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
        })
        .catch(() => {});
    },
    handleGetMenuList() {
      this.$store
        .dispatch("menus/GetList")
        .then((res) => {
          this.RouterList = this.handleTreeSelectOptions(res.data);
          this.$store
            .dispatch("router/GetRoleRouter", {
              role: this.SettingRouterForm.role,
            })
            .then((res) => {
              console.log(res);
              if (res.code == 200) {
                this.RouterListCheck = [];
                for (let i in res.data) {
                  this.RouterListCheck.push(res.data[i].menu);
                }
                this.ViewControl.Dialog.Router.Show = true;
              } else {
                Message({
                  message: res.msg || "Error",
                  type: "error",
                  duration: 5 * 1000,
                });
              }
            })
            .catch(() => {});
          //
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
  },
};
</script>
<style type="less">
.el-header {
  height: auto !important;
}
</style>
