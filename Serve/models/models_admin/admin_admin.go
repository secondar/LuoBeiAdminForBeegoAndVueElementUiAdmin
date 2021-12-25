package models_admin

import (
	"LuoBeiAdminServeForGolang/extend/jwt/jwt_admin"
	"LuoBeiAdminServeForGolang/extend/lib"
	"LuoBeiAdminServeForGolang/extend/utils"
	"encoding/json"
	"errors"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	jwtgo "github.com/dgrijalva/jwt-go"
)

type Admin struct {
	Id        int      `orm:"pk;auto;size(11)" json:"id"`
	Account   string   `orm:"type(char);size(32)" json:"account"`
	Password  string   `orm:"type(char);size(32)" json:"-"`
	Interfere string   `orm:"type(char);size(32)" json:"-"`
	Role      int      `orm:"size(11)" json:"role"`
	State     int8     `orm:"default(1);size(1)" json:"state"`
	Addtime   lib.Time `orm:"auto_now_add" json:"addtime"`
	Token     string   `orm:"-" json:"token"`
}
type AdminModel struct {
	Orm orm.Ormer
	Qs  orm.QuerySeter
}

// 初始化
func init() {
	table_prefix, err := beego.AppConfig.String("mysql::tableprefix")
	if err != nil {
		logs.Error(err)
	}
	orm.RegisterModelWithPrefix(table_prefix, new(Admin))
}

// 获取ORM
func (_this *AdminModel) NewAdminOrm() {
	if _this.Orm == nil {
		_this.Orm = orm.NewOrm()
	}
}

// 获取QS
func (_this *AdminModel) NewAdminQs() {
	if _this.Orm == nil {
		_this.NewAdminOrm()
	}
	if _this.Qs == nil {
		_this.Qs = _this.Orm.QueryTable(new(Admin))
	}
}

// 新增用户
func (_this *AdminModel) Add(AdminData Admin) (int64, error) {
	if _this.Qs == nil {
		_this.NewAdminQs()
	}
	AdminDataChek := Admin{}
	err := _this.Qs.Filter("account", AdminData.Account).One(&AdminDataChek)
	if err != orm.ErrNoRows {
		return 0, errors.New("账户已存在")
	}
	_ = AdminDataChek
	row, err := _this.Orm.Insert(&AdminData)
	if err != nil {
		logs.Error(err)
		return row, errors.New("新增用户失败，如果您是系统管理员您可以通过错误日志查看详细错误信息")
	}
	return row, err
}

// 编辑用户
func (_this *AdminModel) Edit(AdminData Admin) (int64, error) {
	if _this.Qs == nil {
		_this.NewAdminQs()
	}
	AdminDataChek := Admin{}
	err := _this.Qs.Filter("id", AdminData.Id).One(&AdminDataChek)
	if err == orm.ErrNoRows {
		return 0, errors.New("账户不存在")
	} else if err != nil {
		logs.Error(err)
		return 0, errors.New("查询用户是否存在时出现错误，如果您是系统管理员您可以通过错误日志查看详细错误信息")
	}
	AdminData.Id = AdminDataChek.Id
	if AdminData.Password == "" {
		AdminData.Password = AdminDataChek.Password
		AdminData.Interfere = AdminDataChek.Interfere
	}
	row, err := _this.Orm.Update(&AdminData)
	if err != nil {
		logs.Error(err)
		return row, errors.New("编辑用户失败，如果您是系统管理员您可以通过错误日志查看详细错误信息")
	}
	return row, err
}

// 删除用户
func (_this *AdminModel) Delete(id int) error {
	if _this.Qs == nil {
		_this.NewAdminQs()
	}
	_, err := _this.Qs.Filter("id", id).Delete()
	if err != nil {
		logs.Error(err)
		return errors.New("删除用户失败，如果您是系统管理员您可以通过错误日志查看详细错误信息")
	}
	return err
}

//登录
func (_this *AdminModel) Login(Account string, password string, ip string) (Admin, error) {
	if _this.Qs == nil {
		_this.NewAdminQs()
	}
	AdminInfo := Admin{}
	err := _this.Qs.Filter("account", Account).One(&AdminInfo)
	if err != nil {
		if err == orm.ErrNoRows {
			return AdminInfo, errors.New("账户不存在")
		} else {
			logs.Error("账户%s在登录时数据库出现错误，错误信息：%s", Account, err.Error())
			return AdminInfo, errors.New("数据库出现错误，如果您的管理员，请前往查看错误日志，如果您不是管理员，您可以尝试再次登录")
		}
	}
	if AdminInfo.Password != utils.Password(password, AdminInfo.Interfere) {
		return AdminInfo, errors.New("密码错误")
	}
	// 签发身份
	AdminJwtKey, err := beego.AppConfig.String("jwt::admin_key")
	if err != nil || AdminJwtKey == "" {
		logs.Error("没有配置admin_jwt_key，将使用默认www.bugquit.com")
		AdminJwtKey = "www.bugquit.com"
	}
	AdminJwt := &jwt_admin.JWT{
		[]byte(AdminJwtKey),
	}
	ExpirationTime := time.Now()
	TimeTmp, err := time.ParseDuration("86400s")
	if err != nil {
		logs.Error("定义的JWT过期时间转换失败，使用默认时间86400s")
		TimeTmp, _ = time.ParseDuration("86400s")
	}
	ExpirationTime = ExpirationTime.Add(TimeTmp)
	_ = TimeTmp
	Claims := jwt_admin.CustomClaims{
		AdminInfo.Id,
		AdminInfo,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix()), // 签名生效时间
			ExpiresAt: ExpirationTime.Unix(),    // 过期时间
			Issuer:    AdminJwtKey,              //签名的发行者
		},
	}
	token, err := AdminJwt.CreateToken(Claims)
	if err != nil {
		logs.Error("签名签发错误:%s", err.Error())
		return AdminInfo, errors.New("签名签发错误!")
	}
	AdminInfo.Token = token
	// 登录日志
	LogData := AdminLog{}
	LogData.Aid = AdminInfo.Id
	LogData.Content = "使用IP" + ip + "在" + time.Now().Format("2006-01-02 15:04:05") + "登录"
	LogData.Type = 1
	AdminLogModel := AdminLogModel{}
	AdminLogModel.NewAdminLogModelOrm()
	_, err = AdminLogModel.Add(LogData)
	if err != nil {
		logs.Error("写入登录日志时出错：%s", err.Error())
	}
	// 记录在线用户
	AdminOnLineData := AdminOnLine{}
	AdminOnLineModel := AdminOnLineModel{}
	AdminOnLineModel.NewAdminOnLineOrm()
	AdminOnLineData.Aid = AdminInfo.Id
	AdminOnLineData.Account = AdminInfo.Account
	AdminOnLineData.Token = AdminInfo.Token
	LibTime := lib.Time{ExpirationTime}
	AdminOnLineData.ExpirationTime = LibTime
	_, err = AdminOnLineModel.Add(AdminOnLineData)
	if err != nil {
		logs.Error("写入在线账户时出错：%s", err.Error())
	}
	return AdminInfo, nil
}

// 通过token获取账户信息
func (_this *AdminModel) CtxTokenGetAdminInfo(claims *jwt_admin.CustomClaims) (Admin, error) {
	AdminInfo := Admin{}
	c, err := json.Marshal(claims.AdminInfo)
	if err != nil {
		logs.Error(err)
		return AdminInfo, errors.New("从token中获取账户信息时失败，如果您是系统管理员，您可以通过错误日志查看详细信息")
	}
	err = json.Unmarshal(c, &AdminInfo)
	if err != nil {
		logs.Error(err)
		return AdminInfo, errors.New("从token中获取账户信息时失败，如果您是系统管理员，您可以通过错误日志查看详细信息")
	}
	return AdminInfo, err
}
