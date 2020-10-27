package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
)

func BindAndValid(c *gin.Context, params interface{}) string {
	if err := c.ShouldBind(params); err != nil {
		return err.Error()
	}
	if err := gvalid.CheckStruct(params, nil); err != nil {
		return err.FirstString()
	}
	return "success"
}
