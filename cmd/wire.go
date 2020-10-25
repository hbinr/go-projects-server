//+build wireinject

package main

import (
	"go-projects-server/internal/app/user"
	"go-projects-server/internal/app/user/controller"
	"go-projects-server/pkg/conf"
	"go-projects-server/pkg/database"

	"github.com/google/wire"
)

var controllerSet = wire.NewSet(
	controller.NewUseController)

// initWebApp 注入函数，自定义的函数直接注入就行，不需要使用wire set
func initWebApp() (*WebApp, error) {
	// 逻辑顺序入参，未用到的依赖不需要注入
	wire.Build(
		conf.Init,                     // 初始化配置，自定义
		InitEngine,                    // 初始化web引擎，自定义
		database.Init,                 // 初始化mysql，自定义
		user.Set,                      // user业务provider，wire生成
		controllerSet,                 // 获取业务controller
		wire.Struct(new(WebApp), "*"), // WebApp provider，wire生成
	)
	// 返回值不用管。直接返回nil就行
	return nil, nil
}
