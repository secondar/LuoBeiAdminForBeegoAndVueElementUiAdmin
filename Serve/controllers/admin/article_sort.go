package admin

import (
	"LuoBeiAdminServeForGolang/extend/utils"
	"LuoBeiAdminServeForGolang/models/models_admin"

	beego "github.com/beego/beego/v2/server/web"
)

type ArticleSortController struct {
	beego.Controller
}

// 获取列表
func (_this *ArticleSortController) GetList() {
	ResultJson := utils.ResultJson{}
	ArticleSortModel := models_admin.ArticleSortModel{}
	ArticleSortModel.NewArticleSortQs()
	ArticleSortList := ArticleSortModel.GetList()
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	ResultJson.Data = ArticleSortList
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 添加
func (_this *ArticleSortController) Add() {
	ResultJson := utils.ResultJson{}
	pid, err := _this.GetInt("pid")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "请选择上级分类"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	sort, err := _this.GetInt("sort")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "请填写排序值，值越低越靠前"
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
	title := _this.GetString("title")
	ArticleSortModel := models_admin.ArticleSortModel{}
	ArticleSortModel.NewArticleSortQs()
	ArticleSortList := models_admin.ArticleSort{}
	ArticleSortList.Pid = pid
	ArticleSortList.Title = title
	ArticleSortList.State = state
	ArticleSortList.Sort = sort
	_, err = ArticleSortModel.Add(ArticleSortList)
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 修改
func (_this *ArticleSortController) Edit() {
	ResultJson := utils.ResultJson{}
	id, err := _this.GetInt("id")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "获取不到分类ID"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	pid, err := _this.GetInt("pid")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "请选择上级分类"
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
	sort, err := _this.GetInt("sort")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "请填写排序值，值越低越靠前"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	title := _this.GetString("title")
	ArticleSortModel := models_admin.ArticleSortModel{}
	ArticleSortModel.NewArticleSortQs()
	ArticleSortList := models_admin.ArticleSort{}
	ArticleSortList.Id = id
	ArticleSortList.Pid = pid
	ArticleSortList.Title = title
	ArticleSortList.State = state
	ArticleSortList.Sort = sort
	_, err = ArticleSortModel.Edit(ArticleSortList)
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}
//删除
func (_this *ArticleSortController)Delete(){
	ResultJson := utils.ResultJson{}
	id, err := _this.GetInt("id")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "获取不到分类ID"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	ArticleSortModel := models_admin.ArticleSortModel{}
	ArticleSortModel.NewArticleSortQs()
	_,err = ArticleSortModel.Delete(id)
	if err != nil{
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
	}else{
		ResultJson.Code = 200
		ResultJson.Msg = "成功"
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}