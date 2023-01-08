package model

import (
	"context"
	"rahmathidayat1203/backend-portofolio-v1/request"
)

type Message struct {
	ID        string `gorm:"primaryKey" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Message   string `json:"message"`
}

type MessageRepository interface {
	Create(ctx context.Context, msg *Message) (*Message, error)
}

type MessageUsecase interface {
	StoreMessage(ctx context.Context, req *request.MessageRequest) (*Message, error)
}
