package admin

import (
	"LuoBeiAdminServeForGolang/extend/utils"
	"LuoBeiAdminServeForGolang/models/models_admin"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type AdminRouterController struct {
	beego.Controller
}

// 权限权限
func (_this *AdminRouterController) Setting() {
	ResultJson := utils.ResultJson{}
	role, err := _this.GetInt("role")
	if err != nil || role <= 0 {
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

// 获取角色权限路由
func (_this *AdminRouterController) GetRoleRouter() {
	ResultJson := utils.ResultJson{}
	role, err := _this.GetInt("role")
	logs.Error(role)
	if err != nil || role <= 0 {
		ResultJson.Code = 401
		ResultJson.Msg = "非法请求"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	AdminRouterModel := models_admin.AdminRouterModel{}
	AdminRouterModel.NewAdminRouterQs()
	AdminRouterList, err := AdminRouterModel.GetRoleRouter(role)
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
	} else {
		ResultJson.Code = 200
		ResultJson.Msg = "成功"
		ResultJson.Data = AdminRouterList
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}
