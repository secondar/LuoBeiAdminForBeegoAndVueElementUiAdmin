package admin

import (
	"LuoBeiAdminServeForGolang/extend/jwt/jwt_admin"
	"LuoBeiAdminServeForGolang/extend/utils"
	"LuoBeiAdminServeForGolang/models/models_admin"

	beego "github.com/beego/beego/v2/server/web"
)

type AdminMenuController struct {
	beego.Controller
}

// 添加菜单
func (_this *AdminMenuController) Add() {
	ResultJson := utils.ResultJson{}
	pid, err := _this.GetInt("pid")
	if err != nil || pid < 0 {
		ResultJson.Code = 401
		ResultJson.Msg = "请选择正确的父级菜单"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	types, err := _this.GetInt8("type")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "请选择正确的类型"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	show, err := _this.GetInt8("show")
	if err != nil || show < 0 || show > 1 {
		ResultJson.Code = 401
		ResultJson.Msg = "请选择是否显示"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	link, err := _this.GetInt8("link")
	if err != nil && types == 2 || types == 2 && link < 0 || types == 2 && link > 1 {
		ResultJson.Code = 401
		ResultJson.Msg = "请选择是否为外链"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	sort, err := _this.GetInt("sort")
	if err != nil && types == 3 || sort < 0 && types == 3 {
		ResultJson.Code = 401
		ResultJson.Msg = "请填写正确的排序"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	title := utils.TrimSpace(_this.GetString("title"))
	icon := utils.TrimSpace(_this.GetString("icon"))
	characteristic := utils.TrimSpace(_this.GetString("characteristic"))
	router := utils.TrimSpace(_this.GetString("router"))
	component := utils.TrimSpace(_this.GetString("component"))
	path := utils.TrimSpace(_this.GetString("path"))
	api_path := utils.TrimSpace(_this.GetString("api_path"))
	if title == "" {
		ResultJson.Code = 401
		ResultJson.Msg = "请填写菜单标题"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	if characteristic == "" {
		ResultJson.Code = 401
		ResultJson.Msg = "请填写权限标识"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	if types != 3 && router == "" {
		ResultJson.Code = 401
		ResultJson.Msg = "请填写路由地址"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	if types == 2 && link == 0 {
		if component == "" {
			ResultJson.Code = 401
			ResultJson.Msg = "请填写组件名称"
			_this.Data["json"] = ResultJson
			_ = _this.ServeJSON()
			return
		}
		if path == "" {
			ResultJson.Code = 401
			ResultJson.Msg = "请填写组件路径"
			_this.Data["json"] = ResultJson
			_ = _this.ServeJSON()
			return
		}
	}
	AdminMenuData := models_admin.AdminMenu{}
	AdminMenuModels := models_admin.AdminMenuModel{}
	AdminMenuModels.NewAdminMenuQs()
	AdminMenuData.Pid = pid
	AdminMenuData.Type = types
	AdminMenuData.Show = show
	AdminMenuData.Link = link
	AdminMenuData.Sort = &sort
	AdminMenuData.Icon = &icon
	AdminMenuData.Title = title
	AdminMenuData.ApiPath = &api_path
	AdminMenuData.Characteristic = &characteristic
	AdminMenuData.Router = &router
	AdminMenuData.Component = &component
	AdminMenuData.Path = &path
	_, err = AdminMenuModels.Add(AdminMenuData)
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 501
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 编辑
func (_this *AdminMenuController) Edit() {
	ResultJson := utils.ResultJson{}
	id, err := _this.GetInt("id")
	if err != nil || id <= 0 {
		ResultJson.Code = 401
		ResultJson.Msg = "非法请求"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	pid, err := _this.GetInt("pid")
	if err != nil || pid < 0 {
		ResultJson.Code = 401
		ResultJson.Msg = "请选择正确的父级菜单"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	types, err := _this.GetInt8("type")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "请选择正确的类型"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	show, err := _this.GetInt8("show")
	if err != nil || show < 0 || show > 1 {
		ResultJson.Code = 401
		ResultJson.Msg = "请选择是否显示"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	link, err := _this.GetInt8("link")
	if err != nil && types == 2 || types == 2 && link < 0 || types == 2 && link > 1 {
		ResultJson.Code = 401
		ResultJson.Msg = "请选择是否为外链"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	sort, err := _this.GetInt("sort")
	if err != nil && types == 3 || sort < 0 && types == 3 {
		ResultJson.Code = 401
		ResultJson.Msg = "请填写正确的排序"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	title := utils.TrimSpace(_this.GetString("title"))
	icon := utils.TrimSpace(_this.GetString("icon"))
	characteristic := utils.TrimSpace(_this.GetString("characteristic"))
	router := utils.TrimSpace(_this.GetString("router"))
	component := utils.TrimSpace(_this.GetString("component"))
	path := utils.TrimSpace(_this.GetString("path"))
	api_path := utils.TrimSpace(_this.GetString("api_path"))
	if title == "" {
		ResultJson.Code = 401
		ResultJson.Msg = "请填写菜单标题"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	if characteristic == "" {
		ResultJson.Code = 401
		ResultJson.Msg = "请填写权限标识"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	if types != 3 && router == "" {
		ResultJson.Code = 401
		ResultJson.Msg = "请填写路由地址"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	if types == 2 {
		if component == "" {
			ResultJson.Code = 401
			ResultJson.Msg = "请填写组件名称"
			_this.Data["json"] = ResultJson
			_ = _this.ServeJSON()
			return
		}
		if path == "" {
			ResultJson.Code = 401
			ResultJson.Msg = "请填写组件路径"
			_this.Data["json"] = ResultJson
			_ = _this.ServeJSON()
			return
		}
	}
	AdminMenuData := models_admin.AdminMenu{}
	AdminMenuModels := models_admin.AdminMenuModel{}
	AdminMenuModels.NewAdminMenuQs()
	AdminMenuData.Id = id
	AdminMenuData.Pid = pid
	AdminMenuData.Type = types
	AdminMenuData.Show = show
	AdminMenuData.Link = link
	AdminMenuData.Sort = &sort
	AdminMenuData.Icon = &icon
	AdminMenuData.Title = title
	if api_path != "" {
		AdminMenuData.ApiPath = &api_path
	}
	AdminMenuData.Characteristic = &characteristic
	AdminMenuData.Router = &router
	AdminMenuData.Component = &component
	AdminMenuData.Path = &path
	_, err = AdminMenuModels.Edit(AdminMenuData)
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 501
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 删除
func (_this *AdminMenuController) Delete() {
	ResultJson := utils.ResultJson{}
	id, err := _this.GetInt("id")
	if err != nil || id <= 0 {
		ResultJson.Code = 401
		ResultJson.Msg = "非法请求"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	AdminMenuModels := models_admin.AdminMenuModel{}
	AdminMenuModels.NewAdminMenuQs()
	_, err = AdminMenuModels.Delete(id)
	ResultJson.Code = 200
	ResultJson.Msg = "删除成功"
	if err != nil {
		ResultJson.Code = 501
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 获取列表
func (_this *AdminMenuController) List() {
	AdminMenuModel := models_admin.AdminMenuModel{}
	AdminMenu := AdminMenuModel.GetList()
	ResultJson := utils.ResultJson{}
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	ResultJson.Data = AdminMenu
	_this.Data["json"] = ResultJson
	_this.ServeJSON()
}

// 获取账户拥有的路由
func (_this *AdminMenuController) GetAdminMenuRouter() {
	ResultJson := utils.ResultJson{}
	AdminMenuModel := models_admin.AdminMenuModel{}
	AdminModel := models_admin.AdminModel{}
	AdminInfo, err := AdminModel.CtxTokenGetAdminInfo(_this.Ctx.Input.GetData("admin_token_claims").(*jwt_admin.CustomClaims))
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
	}
	AdminMenu := []models_admin.AdminMenu{}
	if AdminInfo.Role != 1 {
		AdminMenu, err = AdminMenuModel.GetAdminMenuRouter(AdminInfo.Id)
	} else {
		// 如果是超级管理员组就直接放行了
		AdminMenu = AdminMenuModel.GetList()
	}
	if err == nil {
		ResultJson.Code = 200
		ResultJson.Msg = "成功"
		ResultJson.Data = AdminMenu
	} else {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
	}
	ResultJson.Data = AdminMenu
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}
