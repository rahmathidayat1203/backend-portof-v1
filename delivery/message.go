package delivery

import (
	"net/http"
	"rahmathidayat1203/backend-portofolio-v1/helper"
	"rahmathidayat1203/backend-portofolio-v1/model"
	"rahmathidayat1203/backend-portofolio-v1/request"

	validation "github.com/go-ozzo/ozzo-validation"
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
	ctx := c.Request().Context()

	var req request.MessageRequest

	if err := c.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(c, "Error Binding Struct", err.Error())
	}
	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error Validation", errVal)
	}
	msg, err := p.messageUsecase.StoreMessage(ctx, &req)

	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}

	return helper.ResponseSuccessJson(c, "success", msg)
}
