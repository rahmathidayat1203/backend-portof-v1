package model

type Message struct {
	ID        string `gorm:"primaryKey" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Message   string `json:"message"`
}

type MessageRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Message   string `json:"message"`
}
