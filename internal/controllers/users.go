package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josh1248/cvwo-assignment-24-backend/internal/entities"
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

func GetUserByName(c *gin.Context) {
	user, err := models.FindUserByName(c.Param("name"))
	log.Println(err)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, err)
	} else if err != nil {
		//placeholder error code here, for anything else not successful.
		//consider cases: server error?
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, user)
	}

}

func VerifyUser(c *gin.Context) {
	var loginUser entities.InputUser

	err := c.ShouldBindJSON(&loginUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, "faulty post data.")
		return
	}

	ok, err := models.AuthenticateUser(loginUser)
	if err == sql.ErrNoRows || !ok {
		c.JSON(http.StatusUnauthorized, "Wrong username or password.")
	} else if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
	} else {
		c.JSON(http.StatusOK, ok)
	}
}

// weigh server-side vs client-side data validation. security vs ease of implementation.
func CreateUser(c *gin.Context) {
	var newUser entities.InputUser

	//Unmarshals JSON data and reads into struct
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}

	err = models.CreateUser(newUser)
	if err != nil {
		//placeholder controls.
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, "successfully created.")
		return
	}
}
