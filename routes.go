package main

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/BaseMax/real-time-notifications-nats-go/controllers"
	"github.com/BaseMax/real-time-notifications-nats-go/helpers"
	"github.com/BaseMax/real-time-notifications-nats-go/middlewares"
)

func todo(c echo.Context) error { return nil }

func InitRoutes() *echo.Echo {
	e := echo.New()
	g := e.Group("/", echojwt.WithConfig(echojwt.Config{SigningKey: helpers.GetJwtSecret()}))

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)
	g.POST("refresh", controllers.Refresh)
	g.GET("users/:id", controllers.FetchUser, middlewares.IsAdmin)
	g.GET("users", controllers.FetchAllUsers, middlewares.IsAdmin)
	g.DELETE("users/:id", controllers.DeleteUser, middlewares.IsAdmin)
	g.PUT("users:/id", controllers.EditUser, middlewares.IsAdmin)

	g.GET("notifications", todo)

	g.POST("products", todo)
	g.GET("products/:id", todo)
	g.GET("products", todo)
	g.GET("products/:order_id", todo)
	g.PUT("products/:id", todo)
	g.DELETE("products/:id", todo)

	g.GET("activities", todo)
	g.POST("activities/seen", todo)

	g.POST("orders", todo)
	g.GET("orders/:id", todo)
	g.GET("orders", todo)
	g.PUT("orders/:id", todo)
	g.DELETE("orders/:id", todo)
	g.GET("orders/:id/status", todo)
	g.POST("orders/:id/cancel", todo)
	g.GET("orders/first", todo)
	g.POST("orders/first/done", todo)
	g.POST("orders/first/cancel", todo)

	g.POST("refunds/:order_id", todo)
	g.GET("refunds/:id", todo)
	g.GET("refunds", todo)
	g.PUT("refunds/:id", todo)
	g.DELETE("refunds/:id", todo)
	g.GET("refunds/:id/status", todo)
	g.POST("refunds/:id/cancel", todo)
	g.GET("refunds/first", todo)
	g.POST("refunds/first/done", todo)
	g.POST("refunds/first/cancel", todo)

	return e
}
