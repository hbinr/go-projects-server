package resp

import (
	"fmt"
	"go-projects-server/pkg/errors"
	"go-projects-server/pkg/router"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gvalid"
)

const (
	SUCCESS = 200
	ERROR   = 1
)

//公用接口返回结构
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//公用列表结构
type ListsData struct {
	Total int         `json:"total"`
	Rows  interface{} `json:"rows"`
}
type apiResp struct {
	Resp
	r *ghttp.Request
}
type LogData struct {
	ID   string
	Code string
}

func NewResp(r *ghttp.Request, code int, msg string, data ...interface{}) *apiResp {
	var d interface{}
	if len(data) > 0 {
		d = data[0]
	}
	return &apiResp{
		Resp: Resp{
			Code: code,
			Msg:  msg,
			Data: d,
		},
		r: r,
	}
}
func Success(r *ghttp.Request, data ...interface{}) *apiResp {
	return NewResp(r, SUCCESS, "操作成功", data...)
}
func Error(r *ghttp.Request) *apiResp {
	return NewResp(r, ERROR, "操作失败")
}

func (apiResp *apiResp) SetCode(code int) *apiResp {
	apiResp.Code = code
	return apiResp
}
func (apiResp *apiResp) SetMsg(msg string) *apiResp {
	apiResp.Msg = msg
	return apiResp
}
func (apiResp *apiResp) SetError(err error) *apiResp {
	switch v := err.(type) {
	case *gvalid.Error:
		apiResp.Msg = v.FirstString()
	case *errors.BusinessError:
		apiResp.SetMsg(v.Error()).SetCode(v.GetCode())
	default:
		apiResp.Msg = err.Error()
	}
	glog.Skip(1).Line(true).Println(err.Error())
	return apiResp
}
func (apiResp *apiResp) SetData(data interface{}) *apiResp {
	apiResp.Data = data
	return apiResp
}

//写操作日志
func (apiResp *apiResp) Log(data interface{}) *apiResp {
	switch v := data.(type) {
	case LogData:
		fmt.Println(v.ID)
		fmt.Println(v.Code)
	}
	g.Dump(router.GetLogNameByName(apiResp.r.Router.Uri))
	return apiResp
}

func (apiResp *apiResp) Json() {
	_ = apiResp.r.Response.WriteJsonExit(apiResp.Resp)
}
