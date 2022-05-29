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
	// 获取所有用户信息
	userGroup.Get("/all", controllers.UserHandler)
	// 获取用户信息
	userGroup.Get("/:id", controllers.UserDetailHandler)
	// 添加用户
	userGroup.Post("/", controllers.UserAddHandler)
	// 编辑用户
	userGroup.Put("/:id", controllers.UserChangeHandler)
	// 删除用户
	userGroup.Delete("/:id", controllers.UserDeleteHandler)
}
