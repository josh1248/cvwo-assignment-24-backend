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
	users, err := models.FindUserByID(c.Param("id"))
	if err != nil {
		//placeholder error code here, for anything not successful.
		//convert to JSON with gin.H.
		//consider cases: server error, invalid username input.
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, users)
	}

}

// special newUser variable that only has the name field. one-off use here only.
type MakeUser struct {
	Name string `json:"name"`
}

// weigh server-side vs client-side data validation. security vs ease of implementation.
func CreateUser(c *gin.Context) {
	var newUser MakeUser

	//Unmarshals JSON data and reads into struct
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadGateway, "Error when creating user.")
	}

	log.Println(newUser.Name, "user created.")
	err = models.CreateUser(newUser.Name)
	if err != nil {
		//placeholder controls.
		c.JSON(http.StatusBadRequest, "username already used.")
	} else {
		c.JSON(http.StatusOK, "successfully created.")
	}
}
