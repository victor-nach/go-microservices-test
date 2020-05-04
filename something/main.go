package main

import (
	"fmt"
	"github.com/victor-nach/go-microservices-test/something/some"
)

func main() {
	fmt.Println("starting something main service..")
	fmt.Println("from some package: ", some.SomeVariable)
}