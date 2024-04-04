package model

// Copyright adalah struktur untuk menyimpan data copyright
type Copyright struct {
	Id                     string `json:"id"`
	Username               string `json:"username"`
	NewUsername            string `json:"newUsername,omitempty"`
	CopyrightAuthorization string `json:"copyrightAuthorization"`
}

type AuthorizationKey struct {
	PrivateKey string `json:"private_key"`
}
type ResponseStatus struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    []Copyright `json:"data,omitempty"`
}
