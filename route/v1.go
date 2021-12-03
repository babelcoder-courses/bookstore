package route

import (
	"github.com/babelcoder-courses/bookstore/config"
	"github.com/babelcoder-courses/bookstore/controller"
	"github.com/babelcoder-courses/bookstore/service"
	"github.com/gin-gonic/gin"
)

func serveV1(r *gin.Engine) {
	db := config.DB()
	v1 := r.Group("/api/v1")
	fileService := &service.File{}
	categoriesService := &service.Category{DB: db}
	booksService := &service.Book{DB: db, FileService: fileService}

	categoriesGroup := v1.Group("categories")
	categoriesController := controller.Category{CategoryService: categoriesService}
	{
		categoriesGroup.GET("", categoriesController.FindAll)
		categoriesGroup.POST("", categoriesController.Create)
		categoriesGroup.GET("/:id", categoriesController.FindOne)
		categoriesGroup.PATCH("/:id", categoriesController.Update)
		categoriesGroup.DELETE("/:id", categoriesController.Delete)
	}

	booksGroup := v1.Group("books")
	booksController := controller.Book{BookService: booksService}
	{
		booksGroup.GET("", booksController.FindAll)
		booksGroup.POST("", booksController.Create)
	}
}
