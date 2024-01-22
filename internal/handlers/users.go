package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josh1248/cvwo-assignment-24-backend/internal/models"
)

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, models.Users)
}

func GetUserByName(c *gin.Context) {
	name := c.Param("name")
	for i, user := range models.Users {
		if user.Name == name {
			c.JSON(http.StatusOK, models.Users[i])
		}
	}
}
