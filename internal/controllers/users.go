package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josh1248/cvwo-assignment-24-backend/internal/models"
)

func GetAllUsers(c *gin.Context) {
	users, err := models.FindAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func GetUserByID(c *gin.Context) {
	user, err := models.FindUserByID(c.Param("id"))
	if err != nil {
		//placeholder controls here.
		c.JSON(http.StatusBadRequest, "User not found/invalid query.")
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func CreateUser(c *gin.Context) {
	log.Println(c.Param("name"))
	err := models.CreateUser(c.Param("name"))
	if err != nil {
		//placeholder controls.
		//consider cases: server error, invalid username input.
		c.JSON(http.StatusBadRequest, "Error when creating user.")
	} else {
		c.JSON(http.StatusOK, "successfully created.")
	}
}
