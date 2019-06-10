package conf

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris"
)

var Logger *golog.Logger

func LoadLog(app *iris.Application){
	if Logger == nil {
		Logger = app.Logger()
	}
}
