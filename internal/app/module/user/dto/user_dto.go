package dto

type CreateUpdateUserDto struct {
	Name           string            `json:"name" validate:"required"`
	Email          string            `json:"email" validate:"required,email"`
	Age            uint8             `json:"age" validate:"gte=0,lte=130"`
	FavouriteColor string            `json:"favouriteColor" validate:"iscolor"` // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*UserAddressDto `validate:"required,dive,required"`        // a person can have a home and cottage...
}

// Address houses a users address information
type UserAddressDto struct {
	Street string `json:"street" validate:"required"`
	City   string `json:"city" validate:"required"`
	Phone  string `json:"phone" validate:"required"`
}
