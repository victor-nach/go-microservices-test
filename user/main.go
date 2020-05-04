package main

import (
	"fmt"
	"github.com/victor-nach/go-microservices-test/something"
)

func main() {
	fmt.Println("starting user service..")
	fmt.Println("from some package: ", some.SomeVariable)
}