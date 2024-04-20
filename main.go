package main

import (
	"testing"

	"github.com/farismnrr/golang-authorization-api/handler"
)

func main() {
	testingFunc := &testing.T{}
	handler.UnitTesting(testingFunc)
}
