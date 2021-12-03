package service

import (
	"github.com/babelcoder-courses/bookstore/dto"
	"github.com/babelcoder-courses/bookstore/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Book struct {
	DB          *gorm.DB
	FileService *File
}

func (b Book) FindAll() ([]model.Book, error) {
	var books []model.Book

	err := b.DB.Find(&books).Error

	return books, err
}

func (b Book) Create(form *dto.CreateBookForm) (*model.Book, error) {
	var book model.Book

	image, err := b.FileService.Store(form.Image, "books")
	if err != nil {
		return nil, err
	}

	copier.Copy(&book, form)
	book.Image = image
	err = b.DB.Create(&book).Error

	return &book, err
}
