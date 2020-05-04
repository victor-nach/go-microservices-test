package main

import (
	"fmt"
	"github.com/victor-nach/go-microservices-test/fancy/fan"
)

var Lami = "anagrma"

func main() {
	fmt.Println("starting something main service..")
	fmt.Println("from some package: ", fan.SomeVariable)
}