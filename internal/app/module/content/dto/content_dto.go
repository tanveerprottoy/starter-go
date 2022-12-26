package dto

import "github.com/go-playground/validator/v10"

type CreateUpdateContentDto struct {
	Name  string `json:"name", validate:"required"`
}
