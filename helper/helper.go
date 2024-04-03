package helper

import (
	"github.com/farismnrr/golang-authorization-api/model"
	"github.com/google/uuid"
)

// GetAllCopyrights mengembalikan semua data yang ada pada model Copyright
func AuthorizationData() ([]model.Copyright, error) {
	copyrightUsers := []model.Copyright{
		{ID: uuid.New().String(), Username: "farismnrr", CopyrightAuthorization: "b8e457e85d402a1952046ffd0b4a34eb"},
		{ID: uuid.New().String(), Username: "user2", CopyrightAuthorization: "authorization2"},
		// Tambahkan data lain sesuai kebutuhan
	}

	return copyrightUsers, nil
}
