package repository

import (
	"context"
	"rahmathidayat1203/backend-portofolio-v1/config"
	"rahmathidayat1203/backend-portofolio-v1/model"
)

type messageRepository struct {
	Cfg config.Config
}

func NewMessageRepository(cfg config.Config) model.MessageRepository {
	return &messageRepository{
		Cfg: cfg,
	}
}

func (p *messageRepository) Create(ctx context.Context, msg *model.Message) (*model.Message, error) {
	if err := p.Cfg.Database().WithContext(ctx).Create(&msg).Error; err != nil {
		return nil, err
	}

	return msg, nil
}
