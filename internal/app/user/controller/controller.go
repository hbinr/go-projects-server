package controller

import (
	"go-projects-server/internal/app/user/service"
	"go-projects-server/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	engine      *gin.Engine
	userService service.IUserService
}

func NewUseController(e *gin.Engine, us service.IUserService) (*UserController, error) {
	user := &UserController{
		engine:      e,
		userService: us,
	}
	g := e.Group("/user")
	g.Use(middleware.JWT) // 设置user私有中间件
	{
		g.POST("/signup", user.SignUp)
		g.GET("/get", user.Get)
		g.GET("/delete", user.Delete)
	}
	return user, nil
}
