package models_admin

import (
	"errors"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type System struct {
	Id       int    `orm:"pk;auto;size(11)" json:"id"`
	Title    string `orm:"size(255)" json:"title"`
	Tail     string `orm:"size(255)" json:"tail"`
	Keyword  string `orm:"size(255)" json:"keyword"`
	Describe string `orm:"size(255)" json:"describe"`
}
type SystemModel struct {
	Orm orm.Ormer
	Qs  orm.QuerySeter
}

// 初始化
func init() {
	tablePrefix, err := beego.AppConfig.String("mysql::tableprefix")
	if err != nil {
		logs.Error(err)
	}
	orm.RegisterModelWithPrefix(tablePrefix, new(System))
}

// 获取ORM
func (_this *SystemModel) NewSystemOrm() {
	if _this.Orm == nil {
		_this.Orm = orm.NewOrm()
	}
}

// 获取QS
func (_this *SystemModel) NewSystemQs() {
	if _this.Orm == nil {
		_this.NewSystemOrm()
	}
	if _this.Qs == nil {
		_this.Qs = _this.Orm.QueryTable(new(System))
	}
}

// 获取
func (_this *SystemModel) GetSystem() (System, error) {
	if _this.Orm == nil {
		_this.NewSystemOrm()
	}
	SystemInfo := System{Id: 1}
	err := _this.Orm.Read(&SystemInfo)
	if err != nil {
		logs.Error(err)
		return SystemInfo, errors.New("读取系统配置失败，如果您是系统管理员，您可以通过错误日志查看详细信息")
	} else {
		return SystemInfo, err
	}
}

// 修改
func (_this *SystemModel) SaveSystem(SystemInfo System) (int64, error) {
	if _this.Orm == nil {
		_this.NewSystemOrm()
	}
	SystemInfoUpdate := System{Id: 1}
	err := _this.Orm.Read(&SystemInfoUpdate)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("读取系统配置失败，如果您是系统管理员，您可以通过错误日志查看详细信息")
	}
	SystemInfoUpdate.Title = SystemInfo.Title
	SystemInfoUpdate.Tail = SystemInfo.Tail
	SystemInfoUpdate.Keyword = SystemInfo.Keyword
	SystemInfoUpdate.Describe = SystemInfo.Describe
	row, err := _this.Orm.Update(&SystemInfoUpdate)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("保存网站配置时失败，如果您是系统管理员，您可以通过错误日志查看详细信息")
	}
	return row, err
}
