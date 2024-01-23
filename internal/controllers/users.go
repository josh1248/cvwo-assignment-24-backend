package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josh1248/cvwo-assignment-24-backend/internal/models"
)

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, models.FindAllUsers())
}
