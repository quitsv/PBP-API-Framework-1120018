package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	UserType int    `json:"user_type"`
}

type MessageResp struct {
	Message string `json:"message"`
}
