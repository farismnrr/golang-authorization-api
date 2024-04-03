package helper

import (
	"encoding/json"
	"io/ioutil"

	"github.com/farismnrr/golang-authorization-api/model"
	"github.com/google/uuid"
)

// GetAllCopyrights mengembalikan semua data yang ada pada model Copyright
func AuthorizationData() ([]model.Copyright, error) {
	// Contoh dummy data
	copyrightUsers := []model.Copyright{
		{ID: uuid.New().String(), Username: "farismnrr", CopyrightAuthorization: "b8e457e85d402a1952046ffd0b4a34eb"},
		{ID: uuid.New().String(), Username: "user2", CopyrightAuthorization: "authorization2"},
		// Tambahkan data lain sesuai kebutuhan
	}

	return copyrightUsers, nil
}

func AuthorizationKey() model.AuthorizationKey {
	// Inisialisasi middleware untuk otorisasi
	authData, err := ioutil.ReadFile("Authorization.json")
	if err != nil {
		panic(err) // Gagal membaca file, hentikan program
	}

	var auth model.AuthorizationKey
	err = json.Unmarshal(authData, &auth)
	if err != nil {
		panic(err) // Gagal parse JSON, hentikan program
	}

	return auth
}
