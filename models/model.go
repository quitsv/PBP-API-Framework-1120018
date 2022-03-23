package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password int    `json:"password"`
}

type MessageResp struct {
	Message string `json:"message"`
}
