package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/farismnrr/golang-authorization-api/helper"
	"github.com/farismnrr/golang-authorization-api/model"
)

func GetKeyFromAPI() (*model.ResponseData, error) {
	_, authToken, err := helper.ReadJsonFile()
	if err != nil {
		fmt.Println("Gagal membaca file JSON:", err)
	}

	req, err := http.NewRequest("GET", "https://authorization-api-dot-ruangguru-exercise.as.r.appspot.com/get-key", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response model.ResponseData
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func GetDataFromAPI() (*model.ResponseData, error) {
	responseAPI, err := GetKeyFromAPI()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", "https://authorization-api-dot-ruangguru-exercise.as.r.appspot.com/copyright", nil)
	if err != nil {
		return nil, err
	}

	if len(responseAPI.Data) > 0 {
		lastData := responseAPI.Data[len(responseAPI.Data)-1]
		req.Header.Set("Authorization", "Bearer "+lastData.CopyrightAuthorization)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response model.ResponseData
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func CopyrightHandler() bool {
	responseAPI, err := GetDataFromAPI()
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	responseJWT, err := GetKeyFromAPI()
	if err != nil {
		fmt.Println("Gagal mendapatkan kunci dari API:", err)
		return false
	}

	username, _, err := helper.ReadJsonFile()
	if err != nil {
		fmt.Println("Gagal membaca file JSON:", err)
		return false
	}

	hashed := helper.GenerateHash(username)

	authorized := false
	for _, data := range responseAPI.Data {
		if hashed == data.CopyrightAuthorization {
			fmt.Println("Status:", responseJWT.Status)
			fmt.Println(responseJWT.Message)
			log.Println("Copyright Authorized by", data.Username)
			authorized = true
			break
		} else {
			fmt.Println("Status: 401")
			fmt.Println("message:", username, "is not authorized!")
			return false
		}
	}

	if !authorized {
		fmt.Println("Status:", responseAPI.Status)
		fmt.Println(responseJWT.Message)
		fmt.Println("Message:", responseAPI.Message)
	}

	return authorized
}
