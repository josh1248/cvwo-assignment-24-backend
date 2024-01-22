package main

import (
	"fmt"

	"github.com/josh1248/cvwo-assignment-24-backend/internal/models"
	"github.com/josh1248/cvwo-assignment-24-backend/internal/router"

	// dependency check with go mod tidy.
	"github.com/josh1248/cvwo-assignment-24-backend/repotest"
)

func main() {
	fmt.Println(repotest.Hello())
	models.ConnectToDB()
	router.StartRoutes()
}
