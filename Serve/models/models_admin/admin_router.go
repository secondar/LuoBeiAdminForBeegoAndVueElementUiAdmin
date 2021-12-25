package models_admin

import (
	"LuoBeiAdminServeForGolang/extend/lib"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type AdminRouter struct {
	Id      int      `orm:"pk;auto;size(11)" json:"id"`
	Role    int      `orm:"size(11)" json:"role"`
	Menu    int      `orm:"size(11)" json:"menu"`
	IsPid   int8     `orm:"size(1)" json:"is_pid"`
	Addtime lib.Time `orm:"auto_now_add" json:"addtime"`
}
type AdminRouterModel struct {
	Orm orm.Ormer
	Qs  orm.QuerySeter
}

// 初始化
func init() {
	table_prefix, err := beego.AppConfig.String("mysql::tableprefix")
	if err != nil {
		logs.Error(err)
	}
	orm.RegisterModelWithPrefix(table_prefix, new(AdminRouter))
}

// 获取ORM
func (_this *AdminRouterModel) NewAdminRouterOrm() {
	if _this.Orm == nil {
		_this.Orm = orm.NewOrm()
	}
}

// 获取QS
func (_this *AdminRouterModel) NewAdminRouterQs() {
	if _this.Orm == nil {
		_this.NewAdminRouterOrm()
	}
	if _this.Qs == nil {
		_this.Qs = _this.Orm.QueryTable(new(AdminRouter))
	}
}

// 设置角色权限路由 role 角色ID router 菜单id，使用英文逗号分隔
func (_this *AdminRouterModel) SaveRouter(role int, router string) (int64, error) {
	if _this.Qs == nil {
		_this.NewAdminRouterQs()
	}
	AdminMenuModel := AdminMenuModel{}
	AdminMenuModel.NewAdminMenuOrm()
	AdminMenuData := AdminMenu{}
	MenuIds := make(map[int]int) //定义一个合集，阅历不够，只能用这个方法,用id作为键值，就避免了重复
	var ids []string
	if strings.Contains(router, ",") {
		// 多个权限
		ids = strings.Split(router, ",")
	} else {
		// 单个权限
		ids = []string{router}
	}
	// 不管3721把旧的权限全部删除
	_, _ = _this.Qs.Filter("role", role).Delete()
	// 开始处理权限上下级关系
	for _, item := range ids {
		// 拆分成数组后循环
		id, err := strconv.Atoi(item)
		tmp_id := id
		if err != nil {
			return 0, errors.New("权限ID非法")
		}
		for {
			// 循环获取他的上级id
			AdminMenuData = AdminMenu{}
			AdminMenuData.Id = id
			err = AdminMenuModel.Orm.Read(&AdminMenuData)
			if err != nil {
				logs.Error(err)
				break
			} else {
				MenuIds[id] = 0
				if tmp_id != id {
					MenuIds[id] = 1
				}
				if AdminMenuData.Pid != 0 {
					// 如果他的父级还有父级，再去找他父级的父级，如此循环
					// MenuIds[AdminMenuData.Pid] = AdminMenuData.Pid
					id = AdminMenuData.Pid
				} else {
					// 否则结束本次循环
					break
				}
			}
		}
	}
	AdminRouterData := AdminRouter{}
	isBeErr := false
	var row int64 = 0
	for id, item := range MenuIds {
		AdminRouterData.Role = role
		AdminRouterData.Menu = id
		if item == 1 {
			AdminRouterData.IsPid = 1
		} else {
			AdminRouterData.IsPid = 0
		}
		i, err := _this.Orm.Insert(&AdminRouterData)
		row += i
		if err != nil {
			logs.Error(err)
			isBeErr = true
		}
		AdminRouterData = AdminRouter{}
	}
	if isBeErr {
		return row, errors.New(fmt.Sprintf("存在错误，应有%d个权限被写入，但实际仅写入了%d个权限，您可以重试，如果您是系统管理员，您可以通过错误日志查看错误", len(MenuIds), row))
	}
	return row, nil
}

// 获取角色权限
func (_this *AdminRouterModel) GetRoleRouter(role int) ([]AdminRouter, error) {
	AdminRouterList := []AdminRouter{}
	if _this.Qs == nil {
		_this.NewAdminRouterQs()
	}
	_, err := _this.Qs.Filter("is_pid", 0).Filter("role", role).All(&AdminRouterList)
	if err != nil {
		logs.Error(err)
		return AdminRouterList, errors.New("获取角色权限失败，如果您是系统管理员，您可以通过错误日志查看错误信息")
	}
	return AdminRouterList, err
}
