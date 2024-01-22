package dummyfolder1

import (
	"fmt"

	_ "github.com/lib/pq"
)

func Hello() string {
	return fmt.Sprintf("Hello, %s", "world!")
}
