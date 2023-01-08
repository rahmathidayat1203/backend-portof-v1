package delivery

import (
	"rahmathidayat1203/backend-portofolio-v1/model"

	"github.com/labstack/echo/v4"
)

type messageDelivery struct {
	messageUsecase model.MessageUsecase
}

type MessageDelivery interface {
	Mount(group *echo.Group)
}

func NewMessageDelivery(messageUsecase model.MessageUsecase) MessageDelivery {
	return &messageDelivery{messageUsecase: messageUsecase}
}

func (p *messageDelivery) Mount(group *echo.Group) {
	group.POST("", p.StoreMessage)
}

func (p *messageDelivery) StoreMessage(c echo.Context) error {
	return nil
}
