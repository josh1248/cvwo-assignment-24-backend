package main

import (
	"fmt"
	_ "log"
	_ "net/http"

	h "github.com/josh1248/cvwo-assignment-24-backend/repotest"
	// dependency check with go mod tidy.
	// _ "github.com/lib/pq"
)

func main() {
	fmt.Print(h.Hello())
	fmt.Print(h.Bye("hhh"))
}
