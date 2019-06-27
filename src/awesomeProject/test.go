package main

import "github.com/kataras/iris"

func main () {
	app := iris.New()
	htmlEngineiris := iris.HTML("./", ".html")
	app.RegisterView(htmlEngineiris)

	app.Get("/", func(ctx iris.Context){
		ctx.WriteString("hello world! -- from iris")
	})
	app.Get("/hello", func(ctx iris.Context){
		ctx.ViewData("title", "测试页面")
		ctx.ViewData("Content","hello world! -- from iris")
		ctx.View("hello.html")
	})

	app.Run(iris.Addr(":8989"), iris.WithCharset("UTF-8"))
}
