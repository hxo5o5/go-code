package main

import (
	"fmt"
	"gotest/gotesting/reverse"
)

func main() {
	fmt.Println("hello go!")
	fmt.Println((reverse.Reverse("hello go!")))
}
