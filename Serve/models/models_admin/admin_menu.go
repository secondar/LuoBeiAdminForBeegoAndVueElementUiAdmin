package models_admin

import (
	"LuoBeiAdminServeForGolang/extend/lib"
	"errors"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// 注意，这个如果变动的话记得在/extend/jwt/jwt_admin/jwt.go 第五十四行一并更改
type AdminMenu struct {
	Id             int         `orm:"pk;auto;size(11)" json:"id"`
	Pid            int         `orm:"size(11)" json:"pid"`
	Title          string      `orm:"size(255)" json:"title"`
	Type           int8        `orm:"size(1)" json:"type"`
	Icon           *string     `orm:"size(255)" json:"icon"`
	Show           int8        `orm:"size(1)" json:"show"`
	Link           int8        `orm:"size(1)" json:"link"`
	ApiPath        *string     `orm:"size(255)" json:"api_path"`
	Characteristic *string     `orm:"size(255)" json:"characteristic"`
	Router         *string     `orm:"size(255)" json:"router"`
	Sort           *int        `orm:"size(11)" json:"sort"`
	Component      *string     `orm:"size(255)" json:"component"`
	Path           *string     `orm:"size(255)" json:"path"`
	Addtime        lib.Time    `orm:"auto_now_add" json:"addtime"`
	Children       []AdminMenu `orm:"-" json:"children"`
}
type AdminMenuModel struct {
	Orm orm.Ormer
	Qs  orm.QuerySeter
}

// 初始化
func init() {
	table_prefix, err := beego.AppConfig.String("mysql::tableprefix")
	if err != nil {
		logs.Error(err)
	}
	orm.RegisterModelWithPrefix(table_prefix, new(AdminMenu))
}

// 获取ORM
func (_this *AdminMenuModel) NewAdminMenuOrm() {
	if _this.Orm == nil {
		_this.Orm = orm.NewOrm()
	}
}

// 获取QS
func (_this *AdminMenuModel) NewAdminMenuQs() {
	if _this.Orm == nil {
		_this.NewAdminMenuOrm()
	}
	if _this.Qs == nil {
		_this.Qs = _this.Orm.QueryTable(new(AdminMenu))
	}
}

// 添加菜单
func (_this *AdminMenuModel) Add(AdminMenuData AdminMenu) (int64, error) {
	if _this.Qs == nil {
		_this.NewAdminMenuQs()
	}
	AdminMenuDataCheck := AdminMenu{}
	if AdminMenuData.Pid != 0 {
		AdminMenuDataCheck.Id = AdminMenuData.Pid
		err := _this.Orm.Read(&AdminMenuDataCheck)
		if err != nil || AdminMenuDataCheck.Title == "" {
			return 0, errors.New("上级菜单不存在")
		}
	}
	AdminMenuDataCheck = AdminMenu{}
	cond := orm.NewCondition()
	where := cond.AndCond(cond.And("title", AdminMenuData.Title))                  //这个是一定存在的
	where = where.OrCond(cond.And("characteristic", AdminMenuData.Characteristic)) //这个一定存在
	if AdminMenuData.Type != 3 {
		where = where.OrCond(cond.And("router", AdminMenuData.Router)) // 这个不等于按钮时存在
	}
	if AdminMenuData.Type == 2 {
		where = where.OrCond(cond.And("component", AdminMenuData.Component)) //菜单是存在
		where = where.OrCond(cond.And("path", AdminMenuData.Path))           //菜单是存在
	}
	err := _this.Qs.SetCond(where).One(&AdminMenuDataCheck)
	if err == nil {
		if AdminMenuDataCheck.Title == AdminMenuData.Title {
			return 0, errors.New("标题已存在")
		}
		if AdminMenuDataCheck.Characteristic == AdminMenuData.Characteristic {
			return 0, errors.New("权限标识已存在")
		}
		if AdminMenuData.Type != 3 {
			if AdminMenuDataCheck.Router == AdminMenuData.Router {
				return 0, errors.New("路由地址已存在")
			}
		}
		if AdminMenuData.Type == 2 {
			if AdminMenuDataCheck.Component == AdminMenuData.Component {
				return 0, errors.New("组件名称已存在")
			}
			if AdminMenuDataCheck.Path == AdminMenuData.Path {
				return 0, errors.New("组件路径已存在")
			}
		}
	}
	i, err := _this.Orm.Insert(&AdminMenuData)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("添加失败，如果您是系统管理员可以查看错误日志")
	}
	return i, nil
}

// 编辑菜单
func (_this *AdminMenuModel) Edit(AdminMenuData AdminMenu) (int64, error) {
	if _this.Qs == nil {
		_this.NewAdminMenuQs()
	}
	AdminMenuDataCheck := AdminMenu{}
	if AdminMenuData.Pid != 0 {
		AdminMenuDataCheck.Id = AdminMenuData.Pid
		err := _this.Orm.Read(&AdminMenuDataCheck)
		if err != nil || AdminMenuDataCheck.Title == "" {
			return 0, errors.New("上级菜单不存在")
		}
	}
	tmp := AdminMenu{}
	tmp.Id = AdminMenuData.Id
	err := _this.Orm.Read(&tmp)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("找不到需要编辑的菜单ID")
	}
	AdminMenuData.Id = tmp.Id
	i, err := _this.Orm.Update(&AdminMenuData)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("修改失败，如果您是系统管理员可以查看错误日志")
	}
	return i, nil
}

// 删除
func (_this *AdminMenuModel) Delete(Id int) (int64, error) {
	if _this.Qs == nil {
		_this.NewAdminMenuQs()
	}
	cou, err := _this.Qs.Filter("pid", Id).Count()
	if err == nil && cou > 0 {
		return 0, errors.New("删除失败，该菜单下还有子菜单，如果需要删除，请先删除该菜单下的所有子菜单")
	}
	row, err := _this.Qs.Filter("id", Id).Delete()
	if err != nil {
		return row, errors.New("删除失败，如果您的系统管理员，您可以查看错误日志")
	} else {
		return row, err
	}
}

// 获取列表
func (_this *AdminMenuModel) GetList() []AdminMenu {
	AdminMenu := []AdminMenu{}
	_this.NewAdminMenuQs()
	_this.Qs.OrderBy("sort").All(&AdminMenu)
	return _this.ToTree(AdminMenu, 0)
}

// 获取用户权限路由
func (_this *AdminMenuModel) GetAdminMenuRouter(role int) ([]AdminMenu, error) {
	if _this.Orm == nil {
		_this.Orm = orm.NewOrm()
	}
	table_prefix, err := beego.AppConfig.String("mysql::tableprefix")
	AdminMenu := []AdminMenu{}
	if err != nil {
		logs.Error(err)
	}
	Tsql := fmt.Sprintf("SELECT admin_menu.* FROM %sadmin_menu admin_menu RIGHT JOIN %sadmin_router admin_router ON admin_router.menu=admin_menu.id WHERE admin_router.role = ?", table_prefix, table_prefix)
	_, err = _this.Orm.Raw(Tsql, role).QueryRows(&AdminMenu)
	if err != nil {
		return AdminMenu, errors.New("获取用户权限路由时出现错误，如果您是系统管理员，您可以通过错误日志查看详细信息")
	}
	return _this.ToTree(AdminMenu, 0), nil
}

// 转树形菜单
func (_this *AdminMenuModel) ToTree(list []AdminMenu, ParentId int) []AdminMenu {
	tree := []AdminMenu{}
	for _, item := range list {
		if item.Pid == ParentId {
			item.Children = _this.ToTree(list, item.Id)
			tree = append(tree, item)
		}
	}
	return tree
}
