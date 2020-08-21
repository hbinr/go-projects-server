package controller

import (
	"go-projects-server/pkg/resp"
	"go-projects-server/pkg/router"

	"github.com/gogf/gf/net/ghttp"
)

type HelloController struct {
}

func init() {
	ctl := new(HelloController)
	router.New("hello", "hello").
		GET("/hi", ctl.SayHi, router.Opts{
			OpenAuth: true,
		})

}
func (h *HelloController) SayHi(r *ghttp.Request) {
	resp.Success(r).Json()
}
