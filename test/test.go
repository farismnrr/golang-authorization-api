package test

import (
	"fmt"
	"log"

	"github.com/farismnrr/golang-authorization-api/handler"
	"github.com/farismnrr/golang-authorization-api/helper"
)

func UnitTest() {
	authType, username, private_key, err := helper.ReadJsonFile()

	if err != nil {
		fmt.Println("Gagal membaca file JSON:", err)
	}

	if authType == "" {
		log.Fatal("Please add type on Authorization file!")
	}

	if username == "" {
		log.Fatal("Please add user_id on Authorization file!")
	}

	if private_key == "" {
		log.Fatal("Please add private_key on Authorization file!")
	}

	if authType != "golang_authorization" {
		log.Fatal("Invalid Authorization type!")
	}

	if !handler.CopyrightHandler() {
		log.Fatal("Unauthorized! Please contact the owner's code")
	}
}
