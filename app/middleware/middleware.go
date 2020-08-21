package middleware

import (
	"fmt"
	"go-projects-server/pkg/resp"
	"net/http"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//func Auth(r *ghttp.Request) {
//	user := new(service.AdminService).GetLoginAdminInfo(r)
//	if g.IsEmpty(user) {
//		resp.Error(r).SetMsg("无权限，非法请求").Json()
//	}
//	if user.Permission == "*" {
//		r.Middleware.Next()
//		return
//	}
//	pas := router.FindPermission(r.Router.Uri)
//	if g.IsEmpty(user.Permission) || !strings.Contains(user.Permission, pas) {
//		resp.Error(r).SetMsg("无权限，非法请求").Json()
//	}
//	r.Middleware.Next()
//}

func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	var m = g.MapIntStr{
		http.StatusNotFound:            fmt.Sprintf("你迷路了，当前请求: %s", r.URL.String()),
		http.StatusInternalServerError: "哎哟我去，服务器居然开小差了，请再试吧！",
	}
	for k, v := range m {
		if r.Response.Status == k {
			r.Response.ClearBuffer()
			resp.Error(r).SetMsg(v).Json()
		}
	}
}
