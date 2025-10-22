package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/introxx/myhttp/internal/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 🔹 Только редирект с /docs на Swagger (удобно)
	r.GET("/docs", func(c *gin.Context) {
		c.Redirect(301, "/swagger/index.html")
	})

	// 🔹 Пример API-группы
	api := r.Group("/api")
	{
		api.GET("/ping", handlers.PingHandler)
		api.GET("/users", handlers.GetUsers)
		api.POST("/users", handlers.CreateUser)
	}

	return r
}
