package usecase

import (
	"context"
	"rahmathidayat1203/backend-portofolio-v1/model"
	"rahmathidayat1203/backend-portofolio-v1/request"

	"github.com/google/uuid"
)

type messageUsecase struct {
	messageRepository model.MessageRepository
}

func NewMessageUsecase(msg model.MessageRepository) model.MessageUsecase {
	return &messageUsecase{
		messageRepository: msg,
	}
}

func (p *messageUsecase) StoreMessage(ctx context.Context, req *request.MessageRequest) (*model.Message, error) {
	newMessage := &model.Message{
		ID:        uuid.New().String(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Message:   req.Message,
	}

	msg, err := p.messageRepository.Create(ctx, newMessage)

	if err != nil {
		return nil, err
	}

	return msg, nil
}
