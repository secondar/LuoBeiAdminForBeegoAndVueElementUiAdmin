package models_admin

import (
	"LuoBeiAdminServeForGolang/extend/lib"
	"errors"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type AdminRole struct {
	Id      int      `orm:"pk;auto;size(11)" json:"id"`
	Title   string   `orm:"size(32)" json:"title"`
	Remarks *string  `orm:"size(255)" json:"remarks"`
	State   int8     `orm:"size(1);default(1)" json:"state"`
	Addtime lib.Time `orm:"auto_now_add" json:"addtime"`
}
type AdminRoleModel struct {
	Orm orm.Ormer
	Qs  orm.QuerySeter
}

// 初始化
func init() {
	table_prefix, err := beego.AppConfig.String("mysql::tableprefix")
	if err != nil {
		logs.Error(err)
	}
	orm.RegisterModelWithPrefix(table_prefix, new(AdminRole))
}

// 获取ORM
func (_this *AdminRoleModel) NewAdminRoleOrm() {
	if _this.Orm == nil {
		_this.Orm = orm.NewOrm()
	}
}

// 获取QS
func (_this *AdminRoleModel) NewAdminRoleQs() {
	if _this.Orm == nil {
		_this.NewAdminRoleOrm()
	}
	if _this.Qs == nil {
		_this.Qs = _this.Orm.QueryTable(new(AdminRole))
	}
}

// 添加角色
func (_this *AdminRoleModel) Add(AdminRoleData AdminRole) (int64, error) {
	if _this.Qs == nil {
		_this.NewAdminRoleQs()
	}
	if AdminRoleData.Title == "" {
		return 0, errors.New("请填写角色名称")
	}
	AdminRoleDataCheck := AdminRole{}
	err := _this.Qs.Filter("title", AdminRoleData.Title).One(&AdminRoleDataCheck)
	if err == nil {
		return 0, errors.New("名称已存在")
	}
	row, err := _this.Orm.Insert(&AdminRoleData)
	if err != nil {
		logs.Error(err)
		return row, errors.New("添加失败，如果您是系统管理员，您可以通过错误日志查看错误信息")
	}
	return row, err
}

// 编辑角色
func (_this *AdminRoleModel) Edit(AdminRoleData AdminRole) (int64, error) {
	if _this.Qs == nil {
		_this.NewAdminRoleQs()
	}
	if AdminRoleData.Title == "" {
		return 0, errors.New("请填写角色名称")
	}
	AdminRoleDataCheck := AdminRole{}
	AdminRoleDataCheck.Id = AdminRoleData.Id
	err := _this.Orm.Read(&AdminRoleDataCheck)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("找不到需要编辑的角色ID")
	}
	AdminRoleData.Id = AdminRoleDataCheck.Id
	row, err := _this.Orm.Update(&AdminRoleData)
	if err != nil {
		logs.Error(err)
		return row, errors.New("修改失败，如果您是系统管理员，您可以通过错误日志查看错误信息")
	}
	return row, err
}

// 获取列表
func (_this *AdminRoleModel) GetList() ([]AdminRole, error) {
	if _this.Qs == nil {
		_this.NewAdminRoleQs()
	}
	RoleList := []AdminRole{}
	_, err := _this.Qs.OrderBy("id").All(&RoleList)
	if err != nil {
		logs.Error(err)
		return RoleList, errors.New("获取失败，如果您是系统管理员，您可以通过错误日志查看详细信息")
	}
	return RoleList, err
}

// 删除
func (_this *AdminRoleModel) Delete(Id int) (int64, error) {
	if _this.Qs == nil {
		_this.NewAdminRoleQs()
	}
	row, err := _this.Qs.Filter("id", Id).Delete()
	if err != nil {
		logs.Error(err)
		return row, errors.New("删除失败，如果您的系统管理员，您可以查看错误日志")
	} else {
		return row, err
	}
}
