// https://youtu.be/ma7rUS_vW9M?t=1693  time spot
package main

import (
	"go-jwt-demo/controllers"
	"go-jwt-demo/initializers"
	"go-jwt-demo/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
