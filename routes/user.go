package routes

import (
	"fiber-mongo-api/controllers"

	"github.com/gofiber/fiber/v2"
)

// 用户路由
// https://docs.gofiber.io/api/app#group
// https://docs.gofiber.io/api/app#route
func User(app *fiber.App) {
	// 用户分组
	userGroup := app.Group("/user")
	// 用户信息
	userGroup.Get("/all/:pageSize/:page", controllers.UserAllHandler)
	// 用户详情
	userGroup.Get("/detail/:id", controllers.UserDetailHandler)
	// 用户添加
	userGroup.Post("/add", controllers.UserAddHandler)
	// 用户编辑
	userGroup.Put("/change/:id", controllers.UserChangeHandler)
	// 用户删除
	userGroup.Delete("/delete/:id", controllers.UserDeleteHandler)
}
