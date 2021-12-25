package models_admin

import (
	"LuoBeiAdminServeForGolang/extend/lib"
	"errors"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type ArticleSort struct {
	Id       int           `orm:"pk;auto;size(11)" json:"id"`
	Pid      int           `orm:"size(11)" json:"pid"`
	Title    string        `orm:"size(32)" json:"title"`
	State    int8          `orm:"size(1)" json:"state"`
	Sort     int           `orm:"size(11)" json:"sort"`
	Addtime  lib.Time      `orm:"auto_now_add" json:"addtime"`
	Children []ArticleSort `orm:"-" json:"children"`
}
type ArticleSortModel struct {
	Orm orm.Ormer
	Qs  orm.QuerySeter
}

// 初始化
func init() {
	table_prefix, err := beego.AppConfig.String("mysql::tableprefix")
	if err != nil {
		logs.Error(err)
	}
	orm.RegisterModelWithPrefix(table_prefix, new(ArticleSort))
}

// 获取ORM
func (_this *ArticleSortModel) NewArticleSortOrm() {
	if _this.Orm == nil {
		_this.Orm = orm.NewOrm()
	}
}

// 获取QS
func (_this *ArticleSortModel) NewArticleSortQs() {
	if _this.Orm == nil {
		_this.NewArticleSortOrm()
	}
	if _this.Qs == nil {
		_this.Qs = _this.Orm.QueryTable(new(ArticleSort))
	}
}

// 添加
func (_this *ArticleSortModel) Add(ArticleSortInfo ArticleSort) (int64, error) {
	if _this.Qs == nil {
		_this.NewArticleSortOrm()
	}
	if ArticleSortInfo.Pid != 0 {
		CheckPid, err := _this.Qs.Filter("id", ArticleSortInfo.Pid).Count()
		if err == nil && CheckPid <= 0 {
			return 0, errors.New("父级分类不存在")
		}
	}
	row, err := _this.Orm.Insert(&ArticleSortInfo)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("添加文章分类在写入数据库时出现错误，如果您是系统管理员，您可以通过错误日志查看详细错误信息")
	}
	return row, nil
}

// 编辑
func (_this *ArticleSortModel) Edit(ArticleSortInfo ArticleSort) (int64, error) {
	if _this.Qs == nil {
		_this.NewArticleSortOrm()
	}
	if ArticleSortInfo.Pid != 0 {
		CheckPid, err := _this.Qs.Filter("id", ArticleSortInfo.Pid).Count()
		if err == nil && CheckPid <= 0 {
			return 0, errors.New("父级分类不存在")
		}
	}
	row, err := _this.Orm.Update(&ArticleSortInfo)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("编辑文章分类在写入数据库时出现错误，如果您是系统管理员，您可以通过错误日志查看详细错误信息")
	}
	return row, nil
}

// 删除
func (_this *ArticleSortModel) Delete(Id int) (int64, error) {
	if _this.Qs == nil {
		_this.NewArticleSortOrm()
	}
	cou, err := _this.Qs.Filter("pid", Id).Count()
	if err == nil && cou > 0 {
		return 0, errors.New("删除失败，该分类下还有子分类，如果需要删除，请先删除该分类下的所有子分类")
	}
	row, err := _this.Qs.Filter("id", Id).Delete()
	if err != nil {
		return row, errors.New("删除失败，如果您的系统管理员，您可以查看错误日志")
	} else {
		return row, err
	}
}

// 获取列表
func (_this *ArticleSortModel) GetList() []ArticleSort {
	if _this.Qs == nil {
		_this.NewArticleSortQs()
	}
	ArticleSort := []ArticleSort{}
	_this.Qs.OrderBy("sort").All(&ArticleSort)
	return _this.ToTree(ArticleSort, 0)
}

// 转树形菜单
func (_this *ArticleSortModel) ToTree(list []ArticleSort, ParentId int) []ArticleSort {
	tree := []ArticleSort{}
	for _, item := range list {
		if item.Pid == ParentId {
			item.Children = _this.ToTree(list, item.Id)
			tree = append(tree, item)
		}
	}
	return tree
}
