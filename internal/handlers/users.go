package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josh1248/cvwo-assignment-24-backend/internal/entities"
)

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, entities.Users)
}

func GetUserByName(c *gin.Context) {
	name := c.Param("name")
	for i, user := range entities.Users {
		if user.Name == name {
			c.JSON(http.StatusOK, entities.Users[i])
		}
	}
}
