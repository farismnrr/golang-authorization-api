package model

// Copyright adalah struktur untuk menyimpan data copyright
type Copyright struct {
	Username               string `json:"username"`
	CopyrightAuthorization string `json:"copyrightAuthorization"`
}

type AuthorizationData struct {
	PrivateKey string `json:"private_key"`
}
