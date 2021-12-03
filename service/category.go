package service

import (
	"github.com/babelcoder-courses/bookstore/dto"
	"github.com/babelcoder-courses/bookstore/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	DB *gorm.DB
}

func (c Category) FindAll() ([]model.Category, error) {
	var categories []model.Category

	err := c.DB.Find(&categories).Error

	return categories, err
}

func (c Category) FindOne(id int) (*model.Category, error) {
	return c.findById(id)
}

func (c Category) Create(form *dto.CreateCategoryForm) (*model.Category, error) {
	var category model.Category

	copier.Copy(&category, form)
	err := c.DB.Create(&category).Error

	return &category, err
}

func (c Category) Update(id int, form *dto.UpdateCategoryForm) (*model.Category, error) {
	category, err := c.findById(id)
	if err != nil {
		return nil, err
	}

	var newCategory model.Category
	copier.Copy(&newCategory, form)
	err = c.DB.Model(&category).Clauses(clause.Returning{}).Updates(newCategory).Error

	return category, err
}

func (c Category) Delete(id int) error {
	category, err := c.findById(id)
	if err != nil {
		return err
	}

	return c.DB.Unscoped().Delete(&category).Error
}

func (c Category) findById(id int) (*model.Category, error) {
	var category model.Category

	err := c.DB.First(&category, id).Error

	return &category, err
}
