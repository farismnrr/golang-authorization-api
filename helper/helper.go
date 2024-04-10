package helper

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/farismnrr/golang-authorization-api/model"
)

func GenerateHash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func ReadJsonFile() (string, string) {
	authData, err := ioutil.ReadFile("Authorization.json")
	if err != nil {
		log.Fatal("Please insert the authorization file!")
	}

	var auth model.AuthorizationData
	json.Unmarshal(authData, &auth)

	return auth.Username, auth.PrivateKey
}
