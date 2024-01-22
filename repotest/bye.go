package repotest

import (
	"fmt"
)

func Bye(name string) string {
	return fmt.Sprintln("bye,", name)
}
