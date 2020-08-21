package base

import (
	"go-projects-server/pkg/resp"
	"strings"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

type Controller struct {
}

func (a *Controller) GetTenantId(r *ghttp.Request) string {
	host := strings.Split(r.GetHost(), ":")
	if len(host) > 0 && (host[0] == "127.0.0.1" || host[0] == "localhost") {
		return "dev"
	}
	ten := strings.Split(host[0], ".")
	if len(ten) > 0 && ten[0] != "" {
		return gconv.String(ten[0])
	}
	resp.Error(r).SetMsg("非法租户").Json()
	return ""
}

//func (a *Controller) GetLoginAdminInfo(r *ghttp.Request) *service.LoginInfo {
//	var userInfo *service.LoginInfo
//	_ = gconv.Struct(service.AdminGfToken.GetTokenData(r).Get("data"), &userInfo)
//	return userInfo
//}
