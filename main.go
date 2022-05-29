package main

import (
	"fiber-mongo-api/databases"
	"fiber-mongo-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// 获取应用实例
	app := fiber.New()
	// 使用日志中间件
	app.Use(logger.New())
	// 获取数据库连接
	databases.Connect()
	// 获取路由
	routes.User(app)
	// 监听服务
	app.Listen("0.0.0.0:8081")
}
