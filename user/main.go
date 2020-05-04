package main

import (
	"fmt"
	// "github.com/victor-nach/go-microservices-test/something"
	"github.com/victor-nach/go-microservices-test/fancy/fan"
)

func main() {
	fmt.Println("starting user service..")
	fmt.Println("from some package: ", fan.SomeVariable)
}