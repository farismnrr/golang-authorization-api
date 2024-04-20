package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/farismnrr/golang-authorization-api/helper"
	"github.com/farismnrr/golang-authorization-api/model"
)

func GetKeyHandler() (*model.ResponseData, error) {
	_, authToken, err := helper.ReadJsonFile()
	if err != nil {
		fmt.Println("Gagal membaca file JSON:", err)
	}

	req, err := http.NewRequest("GET", "https://authorization-api-v1-dot-farismnrr-gclouds.as.r.appspot.com/get-key", nil)
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

func GetDataHandler() (*model.ResponseData, error) {
	responseAPI, err := GetKeyHandler()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", "https://authorization-api-v1-dot-farismnrr-gclouds.as.r.appspot.com/copyright", nil)
	if err != nil {
		return nil, err
	}

	if len(responseAPI.Data) > 0 {
		for _, data := range responseAPI.Data {
			req.Header.Set("Authorization", "Bearer "+data.CopyrightAuthorization)
		}
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

func PostDataHandler(username string) error {
	// Buat map untuk menampung data username
	payload := map[string]string{"username": username}

	// Konversi payload menjadi JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	responseAPI, err := GetKeyHandler()
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://authorization-api-v1-dot-farismnrr-gclouds.as.r.appspot.com/copyright", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}

	if len(responseAPI.Data) > 0 {
		for _, data := range responseAPI.Data {
			req.Header.Set("Authorization", "Bearer "+data.CopyrightAuthorization)
		}
	}

	// Set content type to JSON
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var response model.ResponseData
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err
	}

	fmt.Println("Response:")
	fmt.Println("Status:", response.Status)
	fmt.Println("Message:", response.Message)
	for _, data := range response.Data {
		fmt.Println("Message:", data.ID)
		fmt.Println("Message:", data.Username)
		fmt.Println("Message:", data.CopyrightAuthorization)
	}

	return nil
}

func UpdateDataHandler(username string, newUsername string) error {
	// Buat map untuk menampung data username
	payload := map[string]string{"username": username, "newUsername": newUsername}

	// Konversi payload menjadi JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	responseAPI, err := GetKeyHandler()
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", "https://authorization-api-v1-dot-farismnrr-gclouds.as.r.appspot.com/copyright", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}

	if len(responseAPI.Data) > 0 {
		for _, data := range responseAPI.Data {
			req.Header.Set("Authorization", "Bearer "+data.CopyrightAuthorization)
		}
	}

	// Set content type to JSON
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var response model.ResponseData
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err
	}

	fmt.Println("Response:")
	fmt.Println("Status:", response.Status)
	fmt.Println("Message:", response.Message)
	for _, data := range response.Data {
		fmt.Println("Message:", data.ID)
		fmt.Println("Message:", data.Username)
		fmt.Println("Message:", data.CopyrightAuthorization)
	}

	return nil
}

func DeleteDataHandler(username string) error {
	// Buat map untuk menampung data username
	payload := map[string]string{"username": username}

	// Konversi payload menjadi JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	responseAPI, err := GetKeyHandler()
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", "https://authorization-api-v1-dot-farismnrr-gclouds.as.r.appspot.com/copyright", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}

	if len(responseAPI.Data) > 0 {
		for _, data := range responseAPI.Data {
			req.Header.Set("Authorization", "Bearer "+data.CopyrightAuthorization)
		}
	}

	// Set content type to JSON
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var response model.ResponseData
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err
	}

	fmt.Println("Response:")
	fmt.Println("Status:", response.Status)
	fmt.Println("Message:", response.Message)
	for _, data := range response.Data {
		fmt.Println("Message:", data.ID)
		fmt.Println("Message:", data.Username)
		fmt.Println("Message:", data.CopyrightAuthorization)
	}

	return nil
}

func CopyrightHandler() bool {
	responseAPI, err := GetDataHandler()
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	responseJWT, err := GetKeyHandler()
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