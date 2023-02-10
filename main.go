package main

import (
	"net/http"

	app "example.com/todolist/app"
	docs "example.com/todolist/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @BasePath /api
func setupSwagger(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/swagger/index.html")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api"

	api := r.Group("/api")

	setupSwagger(r)

	app.HealthCheckAppRoute(api)
	app.TodoAppRoute(api)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
