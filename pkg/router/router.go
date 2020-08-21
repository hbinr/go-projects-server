package router

import (
	"go-projects-server/pkg/constant"
	"go-projects-server/pkg/jstr"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
)

const (
	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
	DELETE  = "DELETE"
	CONNECT = "CONNECT"
	TRACE   = "TRACE"
)

var GroupList = make([]*Group, 0)

var PermissionMap = make(map[string]string, 0)
var LogNameMap = make(map[string]string, 0)

//路由信息
type router struct {
	Method       string            //方法名称
	RelativePath string            //url路径
	Permission   string            //权限字符串
	HandlerFunc  ghttp.HandlerFunc //执行函数
}

//路由信息
type Opts struct {
	OpenAuth      bool   //是否开启权限，如果开启并且没有Permission就会默认加入权限，
	Permission    string //权限字符串
	RouterLogName string //路由日志名称
}

//路由组信息
type Group struct {
	ServerName   string              //服务名称
	RelativePath string              //实际使用的路由
	Path         string              //保存的路由(用于其他方法变换路由)
	Handlers     []ghttp.HandlerFunc //中间件
	Router       []*router           //路由信息
}

//默认的几个接口意思
var defaultRouterLogNameMap = g.MapStrStr{
	"save":   "添加",
	"edit":   "修改",
	"remove": "删除",
}

//根据url获取权限字符串
func FindPermission(url string) string {
	return PermissionMap[url]
}
func GetAllPermission() map[string]string {
	return PermissionMap
}
func GetAllLogNameMap() map[string]string {
	return LogNameMap
}
func GetLogNameByName(name string) string {
	return LogNameMap[name]
}

//创建一个路由组
func New(serverName, relativePath string, middleware ...ghttp.HandlerFunc) *Group {
	var rg Group
	relativePath = jstr.CompatiblePrefixStr(relativePath)
	rg.ServerName = serverName
	rg.Router = make([]*router, 0)
	rg.RelativePath = relativePath
	rg.Path = relativePath
	rg.Handlers = middleware
	GroupList = append(GroupList, &rg)
	rg.WithIsApiPrefix(true)
	return &rg
}

//是否加api前缀
func (group *Group) WithIsApiPrefix(flag bool) *Group {
	if flag {
		group.RelativePath = jstr.CompatiblePrefixStr(constant.ApiPrefix + group.RelativePath)
	} else {
		group.RelativePath = group.Path
	}
	return group
}

//自定义api前缀
func (group *Group) WithApiPrefix(str string) *Group {
	if str != "" {
		group.RelativePath = jstr.CompatiblePrefixStr(str + group.Path)
	}
	return group
}

//添加路由信息
func (group *Group) Handle(method string, relativePath string, handler ghttp.HandlerFunc, opts ...Opts) *Group {
	var opt Opts
	if len(opts) > 0 {
		opt = opts[0]
	}
	group.Router = append(group.Router, &router{
		Method:       method,
		RelativePath: relativePath,
		Permission:   opt.Permission,
		HandlerFunc:  handler,
	})
	path := jstr.CompatiblePrefixStr(jstr.CompatibleSuffixStr(constant.ApiPrefix))
	routerName := group.RelativePath + jstr.CompatibleSuffixStr(relativePath)
	hasPermission := len(opt.Permission) > 0
	key := path + routerName
	if opt.RouterLogName == "" {
		urls := strings.Split(routerName, "/")
		last := urls[len(urls)-1]
		if _, ok := defaultRouterLogNameMap[last]; ok {
			opt.RouterLogName = defaultRouterLogNameMap[last]
		}
	} else {
		LogNameMap[key] = opt.RouterLogName
	}
	if opt.OpenAuth || hasPermission {
		PermissionMap[key] = gstr.Replace(gstr.Trim(routerName), "/", ":")
		if hasPermission {
			PermissionMap[key] = opt.Permission
		}
	}
	return group
}

//添加路由信息-ANY
func (group *Group) ANY(relativePath string, handler ghttp.HandlerFunc, opts ...Opts) *Group {
	group.Handle("ANY", relativePath, handler, opts...)
	return group
}

//添加路由信息-GET
func (group *Group) GET(relativePath string, handler ghttp.HandlerFunc, opts ...Opts) *Group {
	group.Handle(GET, relativePath, handler, opts...)
	return group
}

//添加路由信息-POST
func (group *Group) POST(relativePath string, handler ghttp.HandlerFunc, opts ...Opts) *Group {
	group.Handle(POST, relativePath, handler, opts...)
	return group
}

//添加路由信息-OPTIONS
func (group *Group) OPTIONS(relativePath string, handler ghttp.HandlerFunc, opts ...Opts) *Group {
	group.Handle(OPTIONS, relativePath, handler, opts...)
	return group
}

//添加路由信息-PUT
func (group *Group) PUT(relativePath string, handler ghttp.HandlerFunc, opts ...Opts) *Group {
	group.Handle(PUT, relativePath, handler, opts...)
	return group
}

//添加路由信息-PATCH
func (group *Group) PATCH(relativePath string, handler ghttp.HandlerFunc, opts ...Opts) *Group {
	group.Handle(PATCH, relativePath, handler, opts...)
	return group
}

//添加路由信息-HEAD
func (group *Group) HEAD(relativePath string, handler ghttp.HandlerFunc, opts ...Opts) *Group {
	group.Handle(HEAD, relativePath, handler, opts...)
	return group
}

//添加路由信息-DELETE
func (group *Group) DELETE(relativePath string, handler ghttp.HandlerFunc, opts ...Opts) *Group {
	group.Handle(DELETE, relativePath, handler, opts...)
	return group
}

//添加路由信息-CONNECT
func (group *Group) CONNECT(relativePath string, handler ghttp.HandlerFunc, opts ...Opts) *Group {
	group.Handle(CONNECT, relativePath, handler, opts...)
	return group
}

//添加路由信息-TRACE
func (group *Group) TRACE(relativePath string, handler ghttp.HandlerFunc, opts ...Opts) *Group {
	group.Handle(TRACE, relativePath, handler, opts...)
	return group
}
