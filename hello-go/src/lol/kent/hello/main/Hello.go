package main

import (
	"fmt"
)

func main() {
	var result = hello(10, "kent")
	fmt.Println("Hello ", result)
}

func hello(a int, b string) string {
	c := string(a) + b
	return c
}
