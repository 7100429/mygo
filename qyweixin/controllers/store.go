package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/favframework/debug"
	"reflect"
)

// Operations about object
type StoreController struct {
	beego.Controller
}

func mytest(v interface{})  {
	defer func() {
		fmt.Println(reflect.TypeOf(v).String())
		godump.Dump(v)
	}()
}

type (
	Class struct {
		Rc int64 `json:"rc"`
		Error string `json:"error"`
		Type string `json:"type"`
		Job_status string `json:"job_status"`
		Souce_Result ResultInfo `json:"result"`
	}

	ResultInfo struct {
		Total_hits int64 `json:"total_hits"`
		Starttime int64 `json:"starttime"`
		Endtime int64 `json:"endtime"`
		Fields []string `json:"fields"`
		Timeline TimelineInfo `json:"timeline"`
	}

	TimelineInfo struct {
		Interval int64 `json:"interval"`
		Starttime int64 `json:"start_ts"`
		Endtime int64 `json:"end_ts"`
		Rows [] RowsInfo `json:"rows"`
	}

	RowsInfo struct {
		Starttime int64 `json:"start_ts"`
		Endtime int64 `json:"end_ts"`
		Number string `json:"number"`
	}
)
// @Title Get
// @Description find storelist by id
// @Param	id	path 	string	true		"the id you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :id is empty
// @router /:id [get]
func (o *StoreController) Get() {
	//id := o.Ctx.Input.Param(":id")
	//resp,err := http.Get("https://j.autohome.com.cn/pcplatform/common/getBrandInfo")
	//json_str := `{"rc":0,"error":"Success","type":"stats","progress":100,"job_status":"COMPLETED","result":{"total_hits":803254,"starttime":1528434707000,"endtime":1528434767000,"fields":[],"timeline":{"interval":1000,"start_ts":1528434707000,"end_ts":1528434767000,"rows":[{"start_ts":1528434707000,"end_ts":1528434708000,"number":"x12887"},{"start_ts":1528434720000,"end_ts":1528434721000,"number":"x13028"},{"start_ts":1528434721000,"end_ts":1528434722000,"number":"x12975"},{"start_ts":1528434722000,"end_ts":1528434723000,"number":"x12879"},{"start_ts":1528434723000,"end_ts":1528434724000,"number":"x13989"}],"total":803254},"total":8}}`
	//res, err := simplejson.NewJson([]byte(json_str))
	//commModels.CheckErr(err)

	//str_list, err := res.Get("result").Get("timeline").Encode();
	//godump.Dump(str_list)
	//commModels.CheckErr(err)
	////fmt.Println(reflect.TypeOf(str_list).String())
	//b := Josn{}
	//godump.Dump(str_list)
	//err =json.Unmarshal([]byte(str_list), &b)
	//commModels.CheckErr(err)
	//godump.Dump(b)
	//o.Data["json"] = id
	//o.Data["json"] = json_str
	//o.ServeJSON()
	//o.Ctx.WriteString(models.Strval(re))

	str := `{"rc":0,"error":"Success","type":"stats","progress":100,"job_status":"COMPLETED","result":
			{"total_hits":803254,"starttime":1528434707000,"endtime":1528434767000,"fields":[],"timeline":
			{"interval":1000,"start_ts":1528434707000,"end_ts":1528434767000,"rows":
			[{"start_ts":1528434707000,"end_ts":1528434708000,"number":"x12887"},
			{"start_ts":1528434720000,"end_ts":1528434721000,"number":"x13028"},
			{"start_ts":1528434721000,"end_ts":1528434722000,"number":"x12975"},
			{"start_ts":1528434722000,"end_ts":1528434723000,"number":"x12879"},
			{"start_ts":1528434723000,"end_ts":1528434724000,"number":"x13989"}],
			"total":803254},"total":8}}`

	var tstr Class
	json.Unmarshal([]byte(str), &tstr)
	godump.Dump(tstr)
}
