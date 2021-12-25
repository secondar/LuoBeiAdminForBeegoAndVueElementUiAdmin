package admin

import (
	"LuoBeiAdminServeForGolang/extend/utils"
	"LuoBeiAdminServeForGolang/models/models_admin"

	beego "github.com/beego/beego/v2/server/web"
)

type AdminRoleController struct {
	beego.Controller
}

// 添加菜单
func (_this *AdminRoleController) Add() {
	ResultJson := utils.ResultJson{}
	title := utils.TrimSpace(_this.GetString("title"))
	remarks := utils.TrimSpace(_this.GetString("remarks"))
	state, err := _this.GetInt8("state")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "请选择是否启用"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	if title == "" {
		ResultJson.Code = 401
		ResultJson.Msg = "请填写角色名称"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	AdminRoleData := models_admin.AdminRole{}
	AdminRoleModel := models_admin.AdminRoleModel{}
	AdminRoleModel.NewAdminRoleQs()
	AdminRoleData.Title = title
	AdminRoleData.State = state
	if remarks != "" {
		AdminRoleData.Remarks = &remarks
	}
	_, err = AdminRoleModel.Add(AdminRoleData)
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 501
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 添加菜单
func (_this *AdminRoleController) Edit() {
	ResultJson := utils.ResultJson{}
	title := utils.TrimSpace(_this.GetString("title"))
	remarks := utils.TrimSpace(_this.GetString("remarks"))
	id, err := _this.GetInt("id")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "非法请求"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	state, err := _this.GetInt8("state")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "请选择是否启用"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	if title == "" {
		ResultJson.Code = 401
		ResultJson.Msg = "请填写角色名称"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	AdminRoleData := models_admin.AdminRole{}
	AdminRoleModel := models_admin.AdminRoleModel{}
	AdminRoleModel.NewAdminRoleQs()
	AdminRoleData.Id = id
	AdminRoleData.Title = title
	AdminRoleData.State = state
	if remarks != "" {
		AdminRoleData.Remarks = &remarks
	}
	_, err = AdminRoleModel.Edit(AdminRoleData)
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 501
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 获取列表
func (_this *AdminRoleController) List() {
	ResultJson := utils.ResultJson{}
	AdminRoleModel := models_admin.AdminRoleModel{}
	AdminRoleList, err := AdminRoleModel.GetList()
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = "失败"
	} else {
		ResultJson.Data = AdminRoleList
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 删除
func (_this *AdminRoleController) Delete() {
	ResultJson := utils.ResultJson{}
	id, err := _this.GetInt("id")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "非法请求"
		_this.Data["josn"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	AdminRoleModel := models_admin.AdminRoleModel{}
	AdminRoleModel.NewAdminRoleQs()
	_, err = AdminRoleModel.Delete(id)
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 权限权限
func (_this *AdminRoleController) SettingRouter() {
	ResultJson := utils.ResultJson{}
	role, err := _this.GetInt("role")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "非法请求"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	router := utils.TrimSpace(_this.GetString("router"))
	if router == "" {
		ResultJson.Code = 402
		ResultJson.Msg = "请至少选择一个权限"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	AdminRouterModel := models_admin.AdminRouterModel{}
	AdminRouterModel.NewAdminRouterQs()
	_, err = AdminRouterModel.SaveRouter(role, router)
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}
