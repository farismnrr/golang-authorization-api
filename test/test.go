package test

import (
	"fmt"
	"log"

	"github.com/farismnrr/golang-authorization-api/handler"
	"github.com/farismnrr/golang-authorization-api/helper"
)

func UnitTest() {
	authType, _, _, err := helper.ReadJsonFile()

	if err != nil {
		fmt.Println("Gagal membaca file JSON:", err)
	}

	if authType != "golang_authorization" {
		log.Fatal("Invalid Authorization type!")
	}

	if !handler.CopyrightHandler() {
		log.Fatal("Unauthorized! Please contact the owner's code")
	}
}
