package helper

import (
	"crypto/md5"
	"encoding/hex"
	"errors"

	"github.com/farismnrr/golang-authorization-api/model"
	"github.com/google/uuid"
)

var copyrightUsers []model.Copyright
var isDummyDataAdded bool

// GetAllCopyrights mengembalikan semua data yang ada pada model Copyright
func AuthorizationData() ([]model.Copyright, error) {
	return copyrightUsers, nil
}

func AddDummyAuthorizationData() {
	// Jika data dummy sudah ditambahkan sebelumnya, maka tidak perlu ditambahkan lagi
	if isDummyDataAdded {
		return
	}

	// Tambahkan data dummy hanya jika belum ditambahkan sebelumnya
	dummyData := []model.Copyright{
		{Id: uuid.New().String(), Username: "farismnrr", CopyrightAuthorization: "b8e457e85d402a1952046ffd0b4a34eb"},
		// Tambahkan data lain sesuai kebutuhan
	}
	copyrightUsers = append(copyrightUsers, dummyData...)

	isDummyDataAdded = true
}

func AddAuthorizationData(username string) ([]model.Copyright, error) {
	// Membuat hash dari Username
	hash := HashUsername(username)

	// Membuat data baru
	newData := model.Copyright{
		Username:               username,
		CopyrightAuthorization: hash,
	}

	// Menambahkan newData ke dalam slice copyrightUsers
	newData.Id = uuid.New().String() // Menghasilkan ID baru
	copyrightUsers = append(copyrightUsers, newData)

	// Mengembalikan data yang sudah ditambahkan
	return []model.Copyright{newData}, nil
}

func RemoveAuthorizationData(username string) ([]model.Copyright, error) {
	// Mencari index data yang akan dihapus
	index := -1
	for i, data := range copyrightUsers {
		if data.Username == username {
			index = i
			break
		}
	}

	// Menghapus data dari slice
	deletedData := copyrightUsers[index]
	copyrightUsers = append(copyrightUsers[:index], copyrightUsers[index+1:]...)

	// Mengembalikan data yang dihapus
	return []model.Copyright{deletedData}, nil
}

// helper.go
func UpdateAuthorizationData(requestData map[string]string) ([]model.Copyright, error) {
	// Membaca username dari requestData
	username := requestData["username"]
	newUsername := requestData["newUsername"]

	// Mengecek apakah username sudah ada
	index := -1
	for i, data := range copyrightUsers {
		if data.Username == username {
			index = i
			break
		}
	}

	// Jika username tidak ditemukan, kembalikan error
	if index == -1 {
		return nil, errors.New("username not found")
	}

	// Membuat hash dari newUsername
	hash := HashUsername(newUsername)

	// Membuat data baru
	updatedData := model.Copyright{
		Username:               newUsername,
		CopyrightAuthorization: hash,
	}

	// Mengupdate data
	updatedData.Id = uuid.New().String()
	copyrightUsers[index] = updatedData

	// Mengembalikan data yang sudah diupdate
	return []model.Copyright{updatedData}, nil
}

func HashUsername(username string) string {
	hasher := md5.New()
	hasher.Write([]byte(username))
	return hex.EncodeToString(hasher.Sum(nil))
}

func IsUsernameExists(username string) bool {
	for _, data := range copyrightUsers {
		if data.Username == username {
			return true
		}
	}
	return false
}
