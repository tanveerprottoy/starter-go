package dto

type CreateUpdateUserDto struct {
	Name           string            `json:"name" binding:"required"`
	Email          string            `json:"email" binding:"required,email"`
	Age            uint8             `json:"age" binding:"gte=0,lte=130"`
	FavouriteColor string            `json:"favouriteColor" binding:"iscolor"` // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*UserAddressDto `binding:"required,dive,required"`        // a person can multiple addresses
}

// Address houses a users address information
type UserAddressDto struct {
	Street string `json:"street" binding:"required"`
	City   string `json:"city" binding:"required"`
	Phone  string `json:"phone" binding:"required"`
}
