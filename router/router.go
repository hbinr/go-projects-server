package router

import (
	_ "go-projects-server/app/controller"
	"go-projects-server/app/middleware"
	"go-projects-server/pkg/constant"
	"go-projects-server/pkg/router"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

/*
统一路由注册
*/
func init() {
	bindRouter()
	g.DB().SetDebug(g.Config().GetBool("debug"))
}

func bindRouter() {
	s := g.Server()
	s.SetIndexFolder(true)
	s.SetServerRoot(constant.AppPath + "/public")
	s.Use(middleware.ErrorHandler, middleware.CORS)
	gs := s.Group("/")
	registerRouters(gs, router.GroupList)
}
func registerRouters(gs *ghttp.RouterGroup, GroupList []*router.Group) {
	if len(GroupList) > 0 {
		for _, g := range GroupList {
			gs.Group(g.RelativePath, func(group *ghttp.RouterGroup) {
				group.Middleware(g.Handlers...)
				for _, r := range g.Router {
					switch strings.ToUpper(r.Method) {
					case "ANY":
						group.ALL(r.RelativePath, r.HandlerFunc)
					case "GET":
						group.GET(r.RelativePath, r.HandlerFunc)
					case "POST":
						group.POST(r.RelativePath, r.HandlerFunc)
					case "PUT":
						group.PUT(r.RelativePath, r.HandlerFunc)
					case "HEAD":
						group.HEAD(r.RelativePath, r.HandlerFunc)
					case "PATCH":
						group.PATCH(r.RelativePath, r.HandlerFunc)
					case "DELETE":
						group.DELETE(r.RelativePath, r.HandlerFunc)
					case "OPTIONS":
						group.OPTIONS(r.RelativePath, r.HandlerFunc)
					case "CONNECT":
						group.CONNECT(r.RelativePath, r.HandlerFunc)
					case "TRACE":
						group.TRACE(r.RelativePath, r.HandlerFunc)
					}
				}
			})
		}
	}
}
