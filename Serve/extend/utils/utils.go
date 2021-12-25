package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/beego/beego/v2/adapter/cache"
	"github.com/beego/beego/v2/adapter/logs"
	beego "github.com/beego/beego/v2/server/web"
)

//字符串截取 需要截取的自费，截取位置，截取几个字符
func Substr(text string, left int, right int) string {
	nameRune := []rune(text)
	return string(nameRune[left:right]) + "..."
}

//验证邮箱师傅正确
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//加密密码 明文，key
func Password(Plaintext string, encryption string) string {
	left := StrToMd5(Plaintext)
	return StrToMd5(left + encryption)
}

//string to md5
func StrToMd5(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

//获取随机的KEY
func GetRandomKey(l int) string {
	var Str = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM`123456789990-===[];',./~!@#$%^&*()_+～！@#￥%……&×（）——+{}：”》？《"
	bytes := []rune(Str)
	result := []rune{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 初始化缓存
func InitCache() (cache.Cache, error) {
	adapterName, err := beego.AppConfig.String("cache::adaptername")
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	config, err := beego.AppConfig.String("cache::config")
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	Cache, err := cache.NewCache(adapterName, config)
	if err != nil {
		logs.Error(err)
	}
	return Cache, err
}

// 去除字符串所有空格
func TrimSpace(str string) string {
	return strings.Replace(strings.Replace(strings.Replace(strings.Replace(str, " ", "", -1), "\n", "", -1), "	", "", -1), "	", "", -1)
}