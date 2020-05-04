package main

import (
	"fmt"
	"github.com/victor-nach/go-microservices-test/fancy/fan"
)

func main() {
	fmt.Println("starting something main service..")
	fmt.Println("from some package: ", fan.SomeVariable)
}