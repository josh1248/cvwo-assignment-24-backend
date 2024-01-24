package router

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/josh1248/cvwo-assignment-24-backend/internal/controllers"
)

func StartRoutes() {
	fmt.Println("Test route now at http://localhost:8080")
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
	}))
	router.GET("api/users", controllers.GetAllUsers)
	router.GET("api/users/:name", controllers.GetUserByName)
	router.POST("api/new/user", controllers.CreateUser)
	router.POST("api/auth/login", controllers.VerifyUser)
	router.Run("localhost:8080")
}
