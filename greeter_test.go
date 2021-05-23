package greeter

import (
	"fmt"
	"testing"
)


func TestGreet(t *testing.T) {
	result := Greet("Shubham.")
	fmt.Println(result)
}