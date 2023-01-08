package request

import validation "github.com/go-ozzo/ozzo-validation"

type MessageRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Message   string `json:"message"`
}

func (req MessageRequest) Validate() error {
	return validation.ValidateStruct(&req, validation.Field(&req.FirstName, validation.Required), validation.Field(&req.LastName, validation.Required))
}
