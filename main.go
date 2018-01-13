// file: main.go

package main

import (
	"fmt"
	"time"

	"blizzard/controllers"
	"blizzard/repositories"
	"blizzard/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

func main() {
	// TODO:总入口处理流程
	app := iris.New()
	// You got full debug messages, useful when using MVC and you want to make
	// sure that your code is aligned with the Iris' MVC Architecture.
	app.Logger().SetLevel("debug")

	// Load the template files.
	tmpl := iris.HTML("./views", ".html").
		Layout("shared/layout.html").
		Reload(true)
	app.RegisterView(tmpl)

	app.StaticWeb("/public", "./public")

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "The page you're looking for doesn't exist"))
		ctx.ViewData("LocalIP", ctx.Values().
			GetStringDefault("localip", "192.168.168.105"))
		ctx.View("shared/error.html")
	})
	// TODO: 创建数据库信息对象类 db
	// TODO: 创建数据层信息类userservice
	repo := repositories.NewRepository(app)

	iris.RegisterOnInterrupt(func() {
		repo.Close()
	})
	userService := services.NewUserService(repo)
	fmt.Println("userService:", userService)
	// 管理员账户 用于添加卡牌
	app.Controller("/admin", new(controllers.AdminController),
		// Add the basic authentication(admin:password) middleware
		// for the /users based requests.
		controllers.AdminBasicAuth,
		// Bind the "userService" to the UserController's Service (interface) field.
		userService,
	)

	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessioncookiename",
		Expires: 24 * time.Hour,
	})
	p := new(controllers.UserController)
	// TODO：吧Service绑定入控制器中
	app.Controller("/user", p, userService, sessManager)

	// Start the web server at localhost:8080
	// http://localhost:8080/hello
	// http://localhost:8080/hello/iris
	// http://localhost:8080/users/1
	app.Run(
		iris.Addr(":8080"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations, // enables faster json serialization and more
	)
}
