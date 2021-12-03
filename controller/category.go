package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/babelcoder-courses/bookstore/dto"
	"github.com/babelcoder-courses/bookstore/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Category struct {
	CategoryService *service.Category
}

func (c Category) FindAll(ctx *gin.Context) {
	categories, err := c.CategoryService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var serializedCategories []dto.CategoryItemResponse
	copier.Copy(&serializedCategories, categories)
	ctx.JSON(http.StatusOK, gin.H{"categories": serializedCategories})
}

func (c Category) FindOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := c.CategoryService.FindOne(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var serializedCategory dto.CategoryItemResponse
	copier.Copy(&serializedCategory, category)
	ctx.JSON(http.StatusCreated, gin.H{"category": serializedCategory})
}

func (c Category) Create(ctx *gin.Context) {
	var form dto.CreateCategoryForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := c.CategoryService.Create(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var serializedCategory dto.CategoryItemResponse
	copier.Copy(&serializedCategory, category)
	ctx.JSON(http.StatusCreated, gin.H{"category": serializedCategory})
}

func (c Category) Update(ctx *gin.Context) {
	var form dto.UpdateCategoryForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := c.CategoryService.Update(id, &form)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var serializedCategory dto.CategoryItemResponse
	copier.Copy(&serializedCategory, category)
	ctx.JSON(http.StatusCreated, gin.H{"category": serializedCategory})
}

func (c Category) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.CategoryService.Delete(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
