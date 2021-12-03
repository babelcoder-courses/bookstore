package controller

import (
	"net/http"

	"github.com/babelcoder-courses/bookstore/dto"
	"github.com/babelcoder-courses/bookstore/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type Book struct {
	BookService *service.Book
}

func (b Book) FindAll(c *gin.Context) {
	books, err := b.BookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var serializedBooks []dto.BookItemResponse
	copier.Copy(&serializedBooks, books)
	c.JSON(http.StatusOK, gin.H{"books": serializedBooks})
}

func (b Book) Create(c *gin.Context) {
	var form dto.CreateBookForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := b.BookService.Create(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var serializedBook dto.CreateBookResponse
	copier.Copy(&serializedBook, book)
	c.JSON(http.StatusCreated, gin.H{"book": serializedBook})
}
