package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/farismnrr/golang-authorization-api/helper"
	"github.com/farismnrr/golang-authorization-api/model"
)

func GetKeyFromAPI() (string, error) {
	_, authToken := helper.ReadJsonFile()

	req, err := http.NewRequest("GET", "https://ruangguru-exercise.as.r.appspot.com/get-key", nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	var response model.ResponseData
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", fmt.Errorf("failed to decode response body: %v", err)
	}

	return response.Data[0].CopyrightAuthorization, nil
}

func GetDataFromAPI() (*model.ResponseData, error) {
	req, err := http.NewRequest("GET", "https://ruangguru-exercise.as.r.appspot.com/copyright", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	authToken, err := GetKeyFromAPI()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	var response model.ResponseData
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body: %v", err)
	}

	return &response, nil
}

func CopyrightHandler() {
	response, err := GetDataFromAPI()
	if err != nil {
		// Handle error
		log.Fatal("Error:")
	}

	username, _ := helper.ReadJsonFile()

	hashed := helper.GenerateHash(username)

	for _, userData := range response.Data {
		if hashed == userData.CopyrightAuthorization {
			log.Fatal("Copyright Authorized by", userData.Username)
		} else {
			log.Fatal("Unauthorized! Please contact the owner's code")
		}
	}

}
