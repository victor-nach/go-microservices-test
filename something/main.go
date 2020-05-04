package main

import (
	"fmt"
	"something/some"
)

func main() {
	fmt.Println("starting something main service..")
	fmt.Println("from some package: ", some.SomeVariable)
}