package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/josh1248/cvwo-assignment-24-backend/internal/controllers"
)

func StartRoutes() {
	fmt.Println("Test route now at http://localhost:8080")
	router := gin.Default()
	router.GET("/users", controllers.GetAllUsers)
	//router.GET("/users/:name", controllers.GetUserByName)
	router.Run("localhost:8080")
}
