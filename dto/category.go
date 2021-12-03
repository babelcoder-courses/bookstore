package dto

import "time"

type CreateCategoryForm struct {
	Name string `json:"name" binding:"required"`
	Desc string `json:"desc" binding:"required"`
}

type UpdateCategoryForm struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type CategoryItemResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
