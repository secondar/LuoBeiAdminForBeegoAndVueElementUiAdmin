package admin

import (
	"LuoBeiAdminServeForGolang/extend/utils"
	"LuoBeiAdminServeForGolang/models/models_admin"
	beego "github.com/beego/beego/v2/server/web"
)

type SystemController struct {
	beego.Controller
}

// 获取
func (_this *SystemController) GetSystem() {
	ResultJson := utils.ResultJson{}
	SystemModel := models_admin.SystemModel{}
	SystemModel.NewSystemQs()
	SystemInfo, err := SystemModel.GetSystem()
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
	} else {
		ResultJson.Code = 200
		ResultJson.Msg = "成功"
		ResultJson.Data = SystemInfo
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 保存
func (_this *SystemController) SaveSystem() {
	ResultJson := utils.ResultJson{}
	title := utils.TrimSpace(_this.GetString("title"))
	tail := utils.TrimSpace(_this.GetString("tail"))
	keyword := utils.TrimSpace(_this.GetString("keyword"))
	describe := utils.TrimSpace(_this.GetString("describe"))
	if title == "" {
		ResultJson.Code = 401
		ResultJson.Msg = "请填写网站标题"
	}
	if tail == "" {
		ResultJson.Code = 401
		ResultJson.Msg = "请填写网站全站小尾巴"
	}
	if keyword == "" {
		ResultJson.Code = 401
		ResultJson.Msg = "请填写网站关键词"
	}
	if describe == "" {
		ResultJson.Code = 401
		ResultJson.Msg = "请填写网站描述"
	}
	SystemModel := models_admin.SystemModel{}
	SystemModel.NewSystemQs()
	SystemInfo := models_admin.System{Id: 1, Title: title, Tail: tail, Keyword: keyword, Describe: describe}
	_, err := SystemModel.SaveSystem(SystemInfo)
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
	} else {
		ResultJson.Code = 200
		ResultJson.Msg = "成功"
		ResultJson.Data = SystemInfo
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}
