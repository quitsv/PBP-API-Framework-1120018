package models

type Item struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type MessageResp struct {
	Message string `json:"message"`
}
