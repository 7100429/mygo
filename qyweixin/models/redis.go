package models

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
	"mygo/qyweixin/models/redis"
)
var RDSTools redistools.RedisDataStore
func init() {
	//redis test 库初始化
	RDSTools = redistools.RedisDataStore{
		RedisHost: beego.AppConfig.String("redis_host") + ":" + beego.AppConfig.String("redis_prot"),
		RedisPwd: beego.AppConfig.String("redis_password"),
		RedisDB: beego.AppConfig.String("redis_dbnum"),
		TimeOut: 60,
		RedisPool: nil,
	}

	RDSTools.RedisPool = RDSTools.NewPool()
}

func RDSTools_Get(k string) interface{} {
	r, err := RDSTools.Get(k)
	if err != nil{
		logs.Error(err.Error())
	}
	return  r
}

func RDSTools_Set(k string, v string) {
	err := RDSTools.Set(k, v)
	if err != nil{
		logs.Error(err.Error())
	}
}

func RDSTools_SetEx(k string, v string, ttl int64) {
	err := RDSTools.SetEx(k, v, ttl)
	if err != nil{
		logs.Error(err.Error())
	}
}

func RDSTools_HSet(k, f string, v string) {
	err := RDSTools.HSet(f, k, v)
	if err != nil{
		logs.Error(err.Error())
	}
}

func RDSTools_HSetEx(k, f string, v string, ttl int64) {
	err := RDSTools.HSetEx(f, k, v, ttl)
	if err != nil{
		logs.Error(err.Error())
	}
}

func RDSTools_HGet(k, f string) interface{} {
	r, err := RDSTools.HGet(f, k)
	if err != nil{
		logs.Error(err.Error())
	}
	return  r
}

func RDSTools_HGetAll(f string) interface{} {
	r, err := RDSTools.HGetAll(f)
	if err != nil{
		logs.Error(err.Error())
	}
	return  r
}


