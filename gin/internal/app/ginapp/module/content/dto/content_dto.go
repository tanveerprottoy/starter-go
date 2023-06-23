package dto

type CreateUpdateContentDto struct {
	Name string `json:"name" binding:"required"`
}
