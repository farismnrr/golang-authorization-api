package model

type ResponseData struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []UserData `json:"data"`
}

type UserData struct {
	ID                     string `json:"id"`
	Username               string `json:"username"`
	CopyrightAuthorization string `json:"copyrightAuthorization"`
}

type AuthorizationData struct {
	Username   string `json:"user_id"`
	PrivateKey string `json:"private_key"`
}
