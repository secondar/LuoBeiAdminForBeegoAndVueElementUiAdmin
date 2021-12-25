package admin

import (
	"LuoBeiAdminServeForGolang/extend/utils"
	"LuoBeiAdminServeForGolang/models/models_admin"
	beego "github.com/beego/beego/v2/server/web"
)

type AuthController struct {
	beego.Controller
}

func (_this *AuthController) Login() {
	Account := _this.GetString("account")
	Password := _this.GetString("password")
	Account = utils.TrimSpace(Account)
	Password = utils.TrimSpace(Password)
	ResultJson := utils.ResultJson{}
	if Account == "" {
		ResultJson.Code = 401
		ResultJson.Msg = "请填写用户名"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	if Password == ""{
		ResultJson.Code = 401
		ResultJson.Msg = "请填写用户密码"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	AdminModel := models_admin.AdminModel{}
	AdminModel.NewAdminQs()
	AdminInfo,err := AdminModel.Login(Account,Password,_this.Ctx.Input.IP())
	if err != nil{
		ResultJson.Code = 501
		ResultJson.Msg = err.Error()
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	ResultJson.Code = 200
	ResultJson.Msg = "登录成功"
	ResultJson.Data = AdminInfo
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}
