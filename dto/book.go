package dto

import (
	"mime/multipart"
	"time"
)

type CreateBookForm struct {
	Title      string                `form:"title" binding:"required"`
	Desc       string                `form:"desc" binding:"required"`
	Image      *multipart.FileHeader `form:"image" binding:"required"`
	CategoryID uint                  `form:"categoryId" binding:"required"`
}

type CreateBookResponse struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	CategoryID uint
	AuthorID   uint
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type BookItemResponse struct {
	ID        uint      `json:"id"`
	Title     uint      `json:"title"`
	Desc      uint      `json:"desc"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Category  struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	Author struct {
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	} `json:"author"`
}
