package main

import(
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/middleware/logger"
)

func main() {
	app := iris.Default()
	app.Logger().SetLevel("debug")
	app.Use(logger.New(), recover.New())
	
	
	tmpl := iris.HTML("why-it-is-greate/web/4- framework/views", ".html").Reload(true)
	app.RegisterView(tmpl)
	
	app.Get("/test", func(ctx iris.Context){
		ctx.ViewData("message", "Hello world !")
		ctx.View("index.html")
	})

	app.Run(iris.Addr(":8080"))
}