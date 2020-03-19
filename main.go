package main

import (
	"./application"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	//app.Static("/", "./public")

	//定义API路由
	api := app.Group("/api") // /api
	//版本1
	v1 := api.Group("/v1")                                  // /api/v1
	v1.Post("/auth/login", application.LoginHandle())       // /api/v1/auth/login
	v1.Post("/auth/register", application.RegisterHandle()) // /api/v1/auth/register

	app.Listen(3000)
}
