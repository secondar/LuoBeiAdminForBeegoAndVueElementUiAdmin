package admin

import (
	"LuoBeiAdminServeForGolang/extend/utils"
	"LuoBeiAdminServeForGolang/models/models"
	"LuoBeiAdminServeForGolang/models/models_admin"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type AdminAdminController struct {
	beego.Controller
}

// 获取用户列表
func (_this *AdminAdminController) GetList() {
	role, _ := _this.GetInt("role")
	account := utils.TrimSpace(_this.GetString("account"))
	ResultJson := utils.ResultJson{}
	AdminModels := models_admin.AdminModel{}
	AdminModels.NewAdminOrm()
	tablePrefix, err := beego.AppConfig.String("mysql::tableprefix")
	if err != nil {
		logs.Error(err)
	}
	TsqlWhere := ""
	Tsql := fmt.Sprintf("SELECT admin.id,admin.account,admin.addtime,admin.role,admin.state,admin_role.title as rolecn FROM %sadmin admin LEFT JOIN %sadmin_role admin_role ON admin.role = admin_role.id", tablePrefix, tablePrefix)
	where := make(map[string]string)
	if role > 0 {
		if TsqlWhere == "" {
			TsqlWhere = " where admin.role=?"
		} else {
			TsqlWhere += " and admin.role=?"
		}
		where["role"] = fmt.Sprintf("%d", role)
	}
	if account != "" {
		if TsqlWhere == "" {
			TsqlWhere = " where admin.account like ?"
		} else {
			TsqlWhere += " and admin.account like ?"
		}
		where["account"] = "%" + account + "%"
	}
	var RawData orm.RawSeter
	Tsql += TsqlWhere
	if len(where) > 0 {
		wh := models.CreateWhere(where)
		logs.Error(Tsql)
		logs.Error(wh)
		RawData = AdminModels.Orm.Raw(Tsql, wh)
	} else {
		RawData = AdminModels.Orm.Raw(Tsql)
	}
	var maps []orm.Params
	_, err = RawData.Values(&maps)
	if err != nil {
		logs.Error(err)
		ResultJson.Code = 503
		ResultJson.Msg = "获取管理员列表失败，如果您是系统管理员，您可以通过错误日志查看详细信息"
	} else {
		ResultJson.Code = 200
		ResultJson.Msg = "成功"
		ResultJson.Data = maps
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 新增
func (_this *AdminAdminController) Add() {
	ResultJson := utils.ResultJson{}
	account := utils.TrimSpace(_this.GetString("account"))
	password := utils.TrimSpace(_this.GetString("password"))
	role, err := _this.GetInt("role")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "请填选择角色"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	state, err := _this.GetInt8("state")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "请填选择是否启用"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	if account == "" {
		ResultJson.Code = 401
		ResultJson.Msg = "请填用户名"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	if password == "" {
		ResultJson.Code = 401
		ResultJson.Msg = "请填密码"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	AdminModel := models_admin.AdminModel{}
	AdminModel.NewAdminQs()
	AdminData := models_admin.Admin{}
	AdminData.Account = account
	AdminData.Interfere = utils.StrToMd5(utils.GetRandomKey(32))
	AdminData.Password = utils.Password(password, AdminData.Interfere)
	AdminData.Role = role
	AdminData.State = state
	_, err = AdminModel.Add(AdminData)
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 编辑
func (_this *AdminAdminController) Edit() {
	ResultJson := utils.ResultJson{}
	account := utils.TrimSpace(_this.GetString("account"))
	password := utils.TrimSpace(_this.GetString("password"))
	id, err := _this.GetInt("id")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "用户ID获取失败"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	role, err := _this.GetInt("role")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "请选择角色"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	state, err := _this.GetInt8("state")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "请填选择是否启用"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	AdminModel := models_admin.AdminModel{}
	AdminModel.NewAdminQs()
	AdminData := models_admin.Admin{}
	AdminData.Id = id
	AdminData.Account = account
	if password != "" {
		AdminData.Interfere = utils.StrToMd5(utils.GetRandomKey(32))
		AdminData.Password = utils.Password(password, AdminData.Interfere)
	}
	AdminData.Role = role
	AdminData.State = state
	_, err = AdminModel.Edit(AdminData)
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 删除
func (_this *AdminAdminController) Delete() {
	ResultJson := utils.ResultJson{}
	id, err := _this.GetInt("id")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "用户ID获取失败"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	AdminModel := models_admin.AdminModel{}
	AdminModel.NewAdminQs()
	err = AdminModel.Delete(id)
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}
