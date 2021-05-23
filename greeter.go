package greeter

import "fmt"

func Greet(message string) string {
	return fmt.Sprintf("Namaste! %s", message)
}