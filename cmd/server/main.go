package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/josh1248/cvwo-assignment-24-backend/internal/database"
	"github.com/josh1248/cvwo-assignment-24-backend/internal/router"
	"github.com/josh1248/cvwo-assignment-24-backend/repotest"
	// dependency check with go mod tidy.
	// _ "github.com/lib/pq"
)

// rmb to capitalize for public access
type message struct {
	Msg string `json:"message"`
}

var m message = message{"hihihi!"}

func throwGarbage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, m)
}

func main() {
	fmt.Println(repotest.Hello())
	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%T", db) //silence compiler warnings

	router.StartRoutes()
}
