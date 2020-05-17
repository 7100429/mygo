package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github_file/mygo/CommModel"
)

// Operations about object
type StoreController struct {
	beego.Controller
}

// @Title Get
// @Description find storelist by id
// @Param	id	path 	string	true		"the id you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :id is empty
// @router /:id [get]
func (o *StoreController) Get() {
	id := o.Ctx.Input.Param(":id")
	key := id
	haskey := "system_param:new"
	data := make(map[string]interface{})
	data["tesk_id"] = id
	data["myinfo"] = "测试信息"
	logs.Debug("获取信息")
	str := commModels.Strval(data)
	commModels.RDSTools_HSetEx(key, haskey, str,3600)
	//re := models.RDSTools_HGetAll(haskey)
	//if objectId != "" {
	//	ob, err := models.GetOne(objectId)
	//	if err != nil {
	//		o.Data["json"] = err.Error()
	//	} else {
	//		o.Data["json"] = ob
	//	}
	//}
	//var info []string
	//fmt.Println(models.Strval(re))
	//for _,v := range re.([]interface{}) {
	//	models.GetName(v)
	//}
	o.Data["json"] = str
	o.ServeJSON()
	//o.Ctx.WriteString(models.Strval(re))
}
