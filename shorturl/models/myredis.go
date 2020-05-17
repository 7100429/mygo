package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/cache/redis"
	"time"
)

var  Cache cache.Cache

func init()  {
	beego.LoadAppConfig("ini", "conf/app2.conf")
	name := beego.AppConfig.String("redis_name")
	password := beego.AppConfig.String("redis_password")
	prefix := beego.AppConfig.String("redis_test")
	host := beego.AppConfig.String("redis_host")
	port := beego.AppConfig.String("redis_port")
	dbnum := beego.AppConfig.String("redis_dbnum")

	config := orm.Params{
		"key":prefix,
		"conn":host + ":" + port,
		"dbNum":dbnum,
		"password":password,
	}

	configStr, err := json.Marshal(config)

	logs.Debug(string(configStr))
	if err != nil {
		logs.Error("redis配置模型转换失败")
		return
	}

	Cache, err = cache.NewCache(name, string(configStr))
	logs.Debug(name, configStr)
	if err != nil {
		logs.Error(err.Error())
		return
	}
}

func Set(key string, val interface{}, ttl time.Duration) string {
	err := Cache.Put(key, val, ttl)
	if err !=nil {
		fmt.Println("插入数据错误：", err)
	}else {
		fmt.Println("数据插入成功")
	}
	return "hahahle"
}

func Get(key string) interface{} {
	var cache interface{}
	cache = Cache.Get(key)
	if cache == nil{
		beego.Info("获取信息失败")
	}
	return cache
}
