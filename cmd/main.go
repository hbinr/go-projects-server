package main

import (
	"fmt"
	"go-projects-server/internal/app/user/controller"
	"go-projects-server/pkg/conf"
	"go-projects-server/pkg/log"

	"github.com/gin-gonic/gin"
)

func main() {
	webApp, err := initWebApp()
	if err != nil {
		panic(err)
	}
	webApp.Start()
}

// WebApp represent a web application
type WebApp struct {
	*gin.Engine
	config *conf.Config
	user   *controller.UserController
}

// InitEngine 初始化gin
func InitEngine(c *conf.Config) (*gin.Engine, error) {
	if c.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.Use(log.GinLogger(), log.GinRecovery(true)) // 设置公共中间件，可以自定义，举例所以使用了gin自带的
	r.Group("/api")
	return r, nil
}

// Start the web app
func (e *WebApp) Start() {
	e.Run(fmt.Sprintf(":%d", e.config.System.Port))
}
