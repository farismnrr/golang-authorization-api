/*
Package helper provides helper functions for managing authorization data and performing various tasks.

Functions:
- AuthorizationData: Retrieves authorization data.
- AddDummyAuthorizationData: Adds dummy authorization data for testing purposes.
- AddAuthorizationData: Adds authorization data for a new user.
- RemoveAuthorizationData: Removes authorization data for a user.
- UpdateAuthorizationData: Updates authorization data for a user.
- HashUsername: Generates a hash for a given username.
- IsUsernameExists: Checks if a username exists in the authorization data.
- ShowSuccessFromCloudflareResponse: Parses Cloudflare response and returns success status.
*/

package helper

import (
	"crypto/md5"
	"encoding/hex"
	"errors"

	"github.com/farismnrr/golang-authorization-api/middleware"
	"github.com/farismnrr/golang-authorization-api/model"
	"github.com/google/uuid"
)

var copyrightUsers []model.Copyright
var isDummyDataAdded bool

func AuthorizationData() ([]model.Copyright, error) {
	return copyrightUsers, nil
}

func AddDummyAuthorizationData() {
	if isDummyDataAdded {
		return
	}

	dummyData := []model.Copyright{
		{Id: uuid.New().String(), Username: "farismnrr", CopyrightAuthorization: "b8e457e85d402a1952046ffd0b4a34eb"},
	}
	copyrightUsers = append(copyrightUsers, dummyData...)

	isDummyDataAdded = true
}

func AddAuthorizationData(username string) ([]model.Copyright, error) {
	hash := HashUsername(username)

	newData := model.Copyright{
		Username:               username,
		CopyrightAuthorization: hash,
	}

	newData.Id = uuid.New().String() // Menghasilkan ID baru
	copyrightUsers = append(copyrightUsers, newData)

	return []model.Copyright{newData}, nil
}

func RemoveAuthorizationData(username string) ([]model.Copyright, error) {
	index := -1
	for i, data := range copyrightUsers {
		if data.Username == username {
			index = i
			break
		}
	}

	deletedData := copyrightUsers[index]
	copyrightUsers = append(copyrightUsers[:index], copyrightUsers[index+1:]...)

	return []model.Copyright{deletedData}, nil
}

func UpdateAuthorizationData(requestData map[string]string) ([]model.Copyright, error) {
	username := requestData["username"]
	newUsername := requestData["newUsername"]

	index := -1
	for i, data := range copyrightUsers {
		if data.Username == username {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, errors.New("username not found")
	}

	hash := HashUsername(newUsername)

	updatedData := model.Copyright{
		Username:               newUsername,
		CopyrightAuthorization: hash,
	}

	updatedData.Id = uuid.New().String()
	copyrightUsers[index] = updatedData

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

func ShowSuccessFromCloudflareResponse(jsonData []byte) bool {
	success, err := middleware.ParseCloudflareResponse(jsonData)
	if err != nil {
		return false
	}

	return success
}
