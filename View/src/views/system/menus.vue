<template>
  <div class="app-container">
    <el-container>
      <el-header>
        <el-row>
          <el-button
            v-permission="'AddMenu'"
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
          <el-table-column prop="title" label="菜单标题" sortable width="200" />
          <el-table-column prop="name" label="图标" sortable width="160">
            <template slot-scope="scope">
              <i :class="scope.row.icon" />
            </template>
          </el-table-column>
          <el-table-column prop="sort" label="排序" />
          <el-table-column prop="characteristic" label="权限标识" width="155" />
          <el-table-column prop="path" label="组件路径" width="155" />
          <el-table-column prop="link" label="外链">
            <template slot-scope="scope">
              <span>{{ scope.row.link == 1 ? "是" : "否" }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="show" label="可见">
            <template slot-scope="scope">
              <span>{{ scope.row.show == 1 ? "是" : "否" }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="addtime" label="创建日期" width="155" />
          <el-table-column prop="address" label="操作" fixed="right" width="118"
            ><template slot-scope="scope">
              <el-button
                v-permission="'EditMenu'"
                size="mini"
                class="el-icon-edit"
                type="primary"
                @click="handleEdit(scope.$index, scope.row)"
              />
              <el-button
                v-permission="'DeleteMenu'"
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
          <el-form-item label="菜单类型">
            <el-radio-group v-model="formData.type" size="small">
              <el-radio-button label="1">目录</el-radio-button>
              <el-radio-button label="2">菜单</el-radio-button>
              <el-radio-button label="3">按钮</el-radio-button>
            </el-radio-group>
          </el-form-item>

          <el-form-item label="菜单图标">
            <el-select
              v-model="formData.icon"
              filterable
              placeholder="请选择,可搜索"
              style="width: 91%"
            >
              <el-option
                v-for="item in ViewControl.Data.Icon"
                :key="item"
                :label="item"
                :value="item"
              >
                <span style="float: left"><i :class="item" /></span>
                <span style="float: right; color: #8492a6; font-size: 13px">{{
                  item
                }}</span>
              </el-option>
            </el-select>
          </el-form-item>

          <div class="flex">
            <el-form-item label="是否可见">
              <el-radio-group v-model="formData.show" size="small">
                <el-radio-button label="1">是</el-radio-button>
                <el-radio-button label="0">否</el-radio-button>
              </el-radio-group>
            </el-form-item>
            <el-form-item
              v-show="formData.type == 2"
              label="是否外链"
              style="margin-left: 1em"
            >
              <el-radio-group v-model="formData.link" size="small">
                <el-radio-button label="1">是</el-radio-button>
                <el-radio-button label="0">否</el-radio-button>
              </el-radio-group>
            </el-form-item>
          </div>

          <el-form-item
            v-show="formData.link == 0 || formData.type != 2"
            label="接口地址"
          >
            <el-input
              v-model="formData.api_path"
              placeholder="接口地址(可留空)"
            />
          </el-form-item>

          <div class="flex">
            <el-form-item label="菜单标题" style="width: 45%">
              <el-input v-model="formData.title" placeholder="菜单标题" />
            </el-form-item>
            <el-form-item label="权限标识" style="margin-left: 1em; width: 45%">
              <el-input
                v-model="formData.characteristic"
                placeholder="权限标识"
              />
            </el-form-item>
          </div>
          <div v-show="formData.type != 3" class="flex">
            <el-form-item label="路由地址" style="width: 45%">
              <el-input v-model="formData.router" placeholder="路由地址" />
            </el-form-item>
            <el-form-item label="菜单排序" style="margin-left: 1em; width: 45%">
              <el-input v-model="formData.sort" placeholder="菜单排序" />
            </el-form-item>
          </div>

          <div v-show="formData.type == 2 && formData.link == 0" class="flex">
            <el-form-item label="组件名称" style="width: 45%">
              <el-input v-model="formData.component" placeholder="组件名称" />
            </el-form-item>
            <el-form-item label="组件路径" style="margin-left: 1em; width: 45%">
              <el-input v-model="formData.path" placeholder="组件路径" />
            </el-form-item>
          </div>
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
  components: { Treeselect },
  directives: { permission },
  data() {
    return {
      ViewControl: {
        Dialog: {
          Show: false,
          Title: "新增菜单",
        },
        Data: {
          Icon: [
            "el-icon-platform-eleme",
            "el-icon-eleme",
            "el-icon-delete-solid",
            "el-icon-delete",
            "el-icon-s-tools",
            "el-icon-setting",
            "el-icon-user-solid",
            "el-icon-user",
            "el-icon-phone",
            "el-icon-phone-outline",
            "el-icon-more",
            "el-icon-more-outline",
            "el-icon-star-on",
            "el-icon-star-off",
            "el-icon-s-goods",
            "el-icon-goods",
            "el-icon-warning",
            "el-icon-warning-outline",
            "el-icon-question",
            "el-icon-info",
            "el-icon-remove",
            "el-icon-circle-plus",
            "el-icon-success",
            "el-icon-error",
            "el-icon-zoom-in",
            "el-icon-zoom-out",
            "el-icon-remove-outline",
            "el-icon-circle-plus-outline",
            "el-icon-circle-check",
            "el-icon-circle-close",
            "el-icon-s-help",
            "el-icon-help",
            "el-icon-minus",
            "el-icon-plus",
            "el-icon-check",
            "el-icon-close",
            "el-icon-picture",
            "el-icon-picture-outline",
            "el-icon-picture-outline-round",
            "el-icon-upload",
            "el-icon-upload2",
            "el-icon-download",
            "el-icon-camera-solid",
            "el-icon-camera",
            "el-icon-video-camera-solid",
            "el-icon-video-camera",
            "el-icon-message-solid",
            "el-icon-bell",
            "el-icon-s-cooperation",
            "el-icon-s-order",
            "el-icon-s-platform",
            "el-icon-s-fold",
            "el-icon-s-unfold",
            "el-icon-s-operation",
            "el-icon-s-promotion",
            "el-icon-s-home",
            "el-icon-s-release",
            "el-icon-s-ticket",
            "el-icon-s-management",
            "el-icon-s-open",
            "el-icon-s-shop",
            "el-icon-s-marketing",
            "el-icon-s-flag",
            "el-icon-s-comment",
            "el-icon-s-finance",
            "el-icon-s-claim",
            "el-icon-s-custom",
            "el-icon-s-opportunity",
            "el-icon-s-data",
            "el-icon-s-check",
            "el-icon-s-grid",
            "el-icon-menu",
            "el-icon-share",
            "el-icon-d-caret",
            "el-icon-caret-left",
            "el-icon-caret-right",
            "el-icon-caret-bottom",
            "el-icon-caret-top",
            "el-icon-bottom-left",
            "el-icon-bottom-right",
            "el-icon-back",
            "el-icon-right",
            "el-icon-bottom",
            "el-icon-top",
            "el-icon-top-left",
            "el-icon-top-right",
            "el-icon-arrow-left",
            "el-icon-arrow-right",
            "el-icon-arrow-down",
            "el-icon-arrow-up",
            "el-icon-d-arrow-left",
            "el-icon-d-arrow-right",
            "el-icon-video-pause",
            "el-icon-video-play",
            "el-icon-refresh",
            "el-icon-refresh-right",
            "el-icon-refresh-left",
            "el-icon-finished",
            "el-icon-sort",
            "el-icon-sort-up",
            "el-icon-sort-down",
            "el-icon-rank",
            "el-icon-loading",
            "el-icon-view",
            "el-icon-c-scale-to-original",
            "el-icon-date",
            "el-icon-edit",
            "el-icon-edit-outline",
            "el-icon-folder",
            "el-icon-folder-opened",
            "el-icon-folder-add",
            "el-icon-folder-remove",
            "el-icon-folder-delete",
            "el-icon-folder-checked",
            "el-icon-tickets",
            "el-icon-document-remove",
            "el-icon-document-delete",
            "el-icon-document-copy",
            "el-icon-document-checked",
            "el-icon-document",
            "el-icon-document-add",
            "el-icon-printer",
            "el-icon-paperclip",
            "el-icon-takeaway-box",
            "el-icon-search",
            "el-icon-monitor",
            "el-icon-attract",
            "el-icon-mobile",
            "el-icon-scissors",
            "el-icon-umbrella",
            "el-icon-headset",
            "el-icon-brush",
            "el-icon-mouse",
            "el-icon-coordinate",
            "el-icon-magic-stick",
            "el-icon-reading",
            "el-icon-data-line",
            "el-icon-data-board",
            "el-icon-pie-chart",
            "el-icon-data-analysis",
            "el-icon-collection-tag",
            "el-icon-film",
            "el-icon-suitcase",
            "el-icon-suitcase-1",
            "el-icon-receiving",
            "el-icon-collection",
            "el-icon-files",
            "el-icon-notebook-1",
            "el-icon-notebook-2",
            "el-icon-toilet-paper",
            "el-icon-office-building",
            "el-icon-school",
            "el-icon-table-lamp",
            "el-icon-house",
            "el-icon-no-smoking",
            "el-icon-smoking",
            "el-icon-shopping-cart-full",
            "el-icon-shopping-cart-1",
            "el-icon-shopping-cart-2",
            "el-icon-shopping-bag-1",
            "el-icon-shopping-bag-2",
            "el-icon-sold-out",
            "el-icon-sell",
            "el-icon-present",
            "el-icon-box",
            "el-icon-bank-card",
            "el-icon-money",
            "el-icon-coin",
            "el-icon-wallet",
            "el-icon-discount",
            "el-icon-price-tag",
            "el-icon-news",
            "el-icon-guide",
            "el-icon-male",
            "el-icon-female",
            "el-icon-thumb",
            "el-icon-cpu",
            "el-icon-link",
            "el-icon-connection",
            "el-icon-open",
            "el-icon-turn-off",
            "el-icon-set-up",
            "el-icon-chat-round",
            "el-icon-chat-line-round",
            "el-icon-chat-square",
            "el-icon-chat-dot-round",
            "el-icon-chat-dot-square",
            "el-icon-chat-line-square",
            "el-icon-message",
            "el-icon-postcard",
            "el-icon-position",
            "el-icon-turn-off-microphone",
            "el-icon-microphone",
            "el-icon-close-notification",
            "el-icon-bangzhu",
            "el-icon-time",
            "el-icon-odometer",
            "el-icon-crop",
            "el-icon-aim",
            "el-icon-switch-button",
            "el-icon-full-screen",
            "el-icon-copy-document",
            "el-icon-mic",
            "el-icon-stopwatch",
            "el-icon-medal-1",
            "el-icon-medal",
            "el-icon-trophy",
            "el-icon-trophy-1",
            "el-icon-first-aid-kit",
            "el-icon-discover",
            "el-icon-place",
            "el-icon-location",
            "el-icon-location-outline",
            "el-icon-location-information",
            "el-icon-add-location",
            "el-icon-delete-location",
            "el-icon-map-location",
            "el-icon-alarm-clock",
            "el-icon-timer",
            "el-icon-watch-1",
            "el-icon-watch",
            "el-icon-lock",
            "el-icon-unlock",
            "el-icon-key",
            "el-icon-service",
            "el-icon-mobile-phone",
            "el-icon-bicycle",
            "el-icon-truck",
            "el-icon-ship",
            "el-icon-basketball",
            "el-icon-football",
            "el-icon-soccer",
            "el-icon-baseball",
            "el-icon-wind-power",
            "el-icon-light-rain",
            "el-icon-lightning",
            "el-icon-heavy-rain",
            "el-icon-sunrise",
            "el-icon-sunrise-1",
            "el-icon-sunset",
            "el-icon-sunny",
            "el-icon-cloudy",
            "el-icon-partly-cloudy",
            "el-icon-cloudy-and-sunny",
            "el-icon-moon",
            "el-icon-moon-night",
            "el-icon-dish",
            "el-icon-dish-1",
            "el-icon-food",
            "el-icon-chicken",
            "el-icon-fork-spoon",
            "el-icon-knife-fork",
            "el-icon-burger",
            "el-icon-tableware",
            "el-icon-sugar",
            "el-icon-dessert",
            "el-icon-ice-cream",
            "el-icon-hot-water",
            "el-icon-water-cup",
            "el-icon-coffee-cup",
            "el-icon-cold-drink",
            "el-icon-goblet",
            "el-icon-goblet-full",
            "el-icon-goblet-square",
            "el-icon-goblet-square-full",
            "el-icon-refrigerator",
            "el-icon-grape",
            "el-icon-watermelon",
            "el-icon-cherry",
            "el-icon-apple",
            "el-icon-pear",
            "el-icon-orange",
            "el-icon-coffee",
            "el-icon-ice-tea",
            "el-icon-ice-drink",
            "el-icon-milk-tea",
            "el-icon-potato-strips",
            "el-icon-lollipop",
            "el-icon-ice-cream-square",
            "el-icon-ice-cream-round",
          ],
        },
      },
      formData: {
        pid: 0,
        type: "1",
        icon: "",
        show: "1",
        link: "0",
        title: "",
        api_path: "",
        characteristic: "",
        router: "",
        sort: 999,
        component: "",
        path: "",
      },
      MenuList: [],
      TreeSelectOptions: [],
    };
  },
  mounted() {
    this.handleGetMenuList();
  },
  methods: {
    handleAddBtn() {
      this.formData.pid = 0;
      this.formData.type = "1";
      this.formData.icon = "";
      this.formData.show = "1";
      this.formData.link = "0";
      this.formData.title = "";
      this.formData.api_path = "";
      this.formData.characteristic = "";
      this.formData.rolter = "";
      this.formData.sort = 999;
      this.formData.component = "";
      this.formData.path = "";
      this.formData.id = null;
      this.ViewControl.Dialog.Show = true;
    },
    handleDialogHide() {
      this.ViewControl.Dialog.Show = false;
    },
    handleGetMenuList() {
      this.$store
        .dispatch("menus/GetList")
        .then((res) => {
          this.MenuList = res.data;
          this.TreeSelectOptions = [{ id: 0, label: "顶级菜单" }].concat(
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
      let store = "menus/Add";
      if (
        this.formData.id !== undefined &&
        this.formData.id != null &&
        this.formData.id > 0
      ) {
        store = "menus/Edit";
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
          this.handleGetMenuList();
        })
        .catch(() => {});
    },
    handleEdit(index, data) {
      this.formData = {
        pid: data.pid,
        type: data.type,
        icon: data.icon,
        show: data.show,
        link: data.link,
        title: data.title,
        api_path: data.api_path,
        characteristic: data.characteristic,
        router: data.router,
        sort: data.sort,
        component: data.component,
        path: data.path,
      };
      this.formData.id = data.id;
      this.ViewControl.Dialog.Show = true;
    },
    handleDelete(index, data) {
      this.$confirm("此操作将删除该菜单及其子菜单, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          this.$store
            .dispatch("menus/Delete", data.id)
            .then((res) => {
              Message({
                message: res.msg || "Error",
                type: "success",
                duration: 5 * 1000,
              });
              this.handleGetMenuList();
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
