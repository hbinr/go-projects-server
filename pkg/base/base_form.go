package base

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

type Form struct {
	Current  int    `form:"current",json:"current"`   // 当前页码
	Page     int    `form:"page",json:"page"`         // (Current - 1)*PageSize
	PageSize int    `form:"pageSize",json:"pageSize"` // 每页多少条
	Total    int    `form:"total",json:"total"`
	Sort     string `form:"sort",json:"sort"`   // 排序字段
	Order    string `form:"order",json:"order"` // 排序模式 asc 或者 desc
	OrderBy  string `form:"orderBy",json:"orderBy"`
	Params   map[string]string
}

var orderMode = g.MapStrStr{
	"asc":  "asc",
	"desc": "desc",
}

func NewForm(params map[string]interface{}) Form {
	form := Form{}
	form.Params = make(map[string]string, 10)
	for key, value := range params {
		form.Params[key] = gconv.String(value)
	}
	form.Current = 1
	form.PageSize = 30
	if value, ok := params["current"]; ok && gconv.Int(params["current"]) > 0 {
		form.Current = gconv.Int(value)
	}
	// 页数
	if value, ok := params["pageSize"]; ok && gconv.Int(params["pageSize"]) > 0 {
		form.PageSize = gconv.Int(value)
	}
	if form.GetParam("sort") != "" {
		form.Sort = gconv.String(form.GetParam("Sort"))
	}
	if form.GetParam("order") != "" {
		form.Order = orderMode[gconv.String(form.GetParam("order"))]
	}

	if form.Sort != "" {
		form.OrderBy = form.Sort + " " + form.Order
	}

	form.Page = (form.Current - 1) * form.PageSize
	return form
}

func (form *Form) SetParam(key string, value string) *Form {
	form.Params[key] = value
	return form
}
func (form *Form) GetParam(key string) string {
	if value, ok := form.Params[key]; ok {
		return value
	}
	return ""
}

func (form *Form) SetParams(params map[string]string) *Form {
	form.Current = gconv.Int(params["current"])
	form.PageSize = gconv.Int(params["pageSize"])
	form.Sort = gconv.String(params["sort"])
	form.Order = gconv.String(params["order"])
	form.Params = params
	return form
}
