package api

import (
	"fmt"
	"go-projects-server/internal/app/dto"
	userService "go-projects-server/internal/app/service/user"
	"go-projects-server/pkg/ginx"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	var r *dto.RegisterStruct
	// if err := ginx.BindAndValid(c, r); err != nil {
	// 	ginx.FailWithMessage(err.Error(), c)
	// 	return
	// }
	err := userService.Register(r)
	if err != nil {
		ginx.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		ginx.OkWithMessage("注册成功", c)
	}
}
