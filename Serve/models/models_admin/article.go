package models_admin

import (
	"LuoBeiAdminServeForGolang/extend/lib"
	"errors"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type Article struct {
	Id        int      `orm:"pk;auto;size(11)" json:"id"`
	Sort      int      `orm:"size(11)" json:"sort"`
	Title     string   `orm:"size(255)" json:"title"`
	Thumbnail string   `orm:"size(255)" json:"thumbnail"`
	Content   string   `orm:"type(text)" json:"content"`
	Describe  string   `orm:"size(255)" json:"describe"`
	Hot       int      `orm:"size(11)" json:"hot"`
	Addtime   lib.Time `orm:"auto_now_add" json:"addtime"`
}
type ArticleModel struct {
	Orm orm.Ormer
	Qs  orm.QuerySeter
}

// 初始化
func init() {
	tablePrefix, err := beego.AppConfig.String("mysql::tableprefix")
	if err != nil {
		logs.Error(err)
	}
	orm.RegisterModelWithPrefix(tablePrefix, new(Article))
}

// 获取ORM
func (_this *ArticleModel) NewArticleOrm() {
	if _this.Orm == nil {
		_this.Orm = orm.NewOrm()
	}
}

// 获取QS
func (_this *ArticleModel) NewArticleQs() {
	if _this.Orm == nil {
		_this.NewArticleOrm()
	}
	if _this.Qs == nil {
		_this.Qs = _this.Orm.QueryTable(new(Article))
	}
}

// 添加
func (_this *ArticleModel) Add(ArticleInfo Article) (int64, error) {
	if _this.Orm == nil {
		_this.NewArticleOrm()
	}
	if ArticleInfo.Title == "" {
		return 0, errors.New("请填写文章标题")
	}
	if ArticleInfo.Sort <= 0 {
		return 0, errors.New("请选择文章分类")
	}
	if ArticleInfo.Content == "" {
		return 0, errors.New("请选择文章内容")
	}
	row, err := _this.Orm.Insert(&ArticleInfo)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("添加文章时出现错误，如果您是系统管理员，您可以通过错误日志查看详细信息")
	}
	return row, err
}

// 修改
func (_this *ArticleModel) Edit(ArticleInfo Article) (int64, error) {
	if _this.Orm == nil {
		_this.NewArticleOrm()
	}
	ArticleInfoUpdate := Article{}
	ArticleInfoUpdate.Id = ArticleInfo.Id
	err := _this.Orm.Read(&ArticleInfoUpdate)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("查询不到需要修改的文章")
	}
	ArticleInfoUpdate.Sort = ArticleInfo.Sort
	ArticleInfoUpdate.Title = ArticleInfo.Title
	ArticleInfoUpdate.Content = ArticleInfo.Content
	ArticleInfoUpdate.Describe = ArticleInfo.Describe
	row, err := _this.Orm.Update(&ArticleInfoUpdate)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("修改文章时失败，如果您是系统管理员，您可以通过错误日志查看详细信息")
	}
	return row, err
}

// 删除
func (_this *ArticleModel) Delete(id int) (int64, error) {
	if _this.Qs == nil {
		_this.NewArticleQs()
	}
	row, err := _this.Qs.Filter("id", id).Delete()
	if err != nil {
		logs.Error(err)
		return 0, errors.New("删除文章失败，如果您是系统管理员，您可以通过错误日志查看详细信息")
	}
	return row, err
}

// 获取详情
func (_this *ArticleModel) Details(id int) (Article, error) {
	if _this.Orm == nil {
		_this.NewArticleOrm()
	}
	ArticleInfo := Article{}
	ArticleInfo.Id = id
	err := _this.Orm.Read(&ArticleInfo)
	if err != nil {
		if err == orm.ErrNoRows {
			return ArticleInfo, errors.New("找不到文章信息")
		} else {
			logs.Error(err)
			return ArticleInfo, errors.New("获取文章信息失败，如果您是系统管理员，您可以通过错误日志查看错误信息")
		}
	}
	return ArticleInfo, err
}
