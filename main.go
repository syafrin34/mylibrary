package main

import (
	"fmt"

	"github.com/syafrin34/mylibrary/pointer"
)

func main() {
	x := pointer.Of("Hello")
	fmt.Println(*x)
}
