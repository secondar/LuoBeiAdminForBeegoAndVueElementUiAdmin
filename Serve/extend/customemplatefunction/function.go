package customemplatefunction

import (
	"LuoBeiAdminServeForGolang/extend/utils"
	"LuoBeiAdminServeForGolang/models/models_admin"
	"time"

	"github.com/beego/beego/v2/adapter/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func Init() {
	_ = beego.AddFuncMap("GetSystemConf", GetSystemConf)
}
func GetSystemConf() (SystemConfInfo models_admin.System) {
	Cache, err := utils.InitCache()
	if err != nil {
		logs.Error("初始化缓存失败,错误信息：%s", err.Error())
		return models_admin.System{}
	}
	SystemConf := Cache.Get("SystemConf")
	if SystemConf == nil {
		SystemModel := models_admin.SystemModel{}
		SystemModel.NewSystemOrm()
		SystemData, err := SystemModel.GetSystem()
		if err != nil {
			logs.Error("从数据可获取网站配置失败,错误信息：%s", err.Error())
			return models_admin.System{}
		}
		SystemConf = SystemData
		err = Cache.Put("SystemConf", SystemData, 12000*60*time.Second)
		if err != nil {
			logs.Error("写网站配置缓存失败,错误信息：%s", err.Error())
			return models_admin.System{}
		}
	}
	return SystemConf.(models_admin.System)
}
