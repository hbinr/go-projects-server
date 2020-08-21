package constant

import (
	"os"

	"github.com/gogf/gf/frame/g"
)

var ApiPrefix = g.Config().GetString("api-prefix")
var AppPath, _ = os.Getwd()
