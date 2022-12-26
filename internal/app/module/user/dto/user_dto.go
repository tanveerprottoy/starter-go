package dto

import "github.com/go-playground/validator/v10"

type CreateUpdateUserDto struct {
	Name           string     `json:"name", validate:"required"`
	Email          string     `json:"email" validate:"required,email"`
	Age            uint8      `json:"email" validate:"gte=0,lte=130"`
	FavouriteColor string     `json:"email" validate:"iscolor"`   // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*UserAddressDto `validate:"required,dive,required"` // a person can have a home and cottage...
}

// Address houses a users address information
type UserAddressDto struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}
