package boot

import (
	_ "go-projects-server/router"

	"github.com/gogf/gf/frame/g"
)

// 用于配置初始化.
func init() {
	g.Server()
}
