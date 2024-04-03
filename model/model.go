package model

// Copyright adalah struktur untuk menyimpan data copyright
type Copyright struct {
	ID                     string
	Username               string `json:"username"`
	CopyrightAuthorization string `json:"copyrightAuthorization"`
}

type AuthorizationKey struct {
	PrivateKey string `json:"private_key"`
}
