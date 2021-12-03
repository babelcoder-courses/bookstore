package seed

import (
	"github.com/babelcoder-courses/bookstore/config"
	"github.com/babelcoder-courses/bookstore/model"
	"github.com/bxcodec/faker/v3"
)

func Load() {
	db := config.DB()

	numOfCategories := 10
	categories := make([]model.Category, 0, numOfCategories)
	for i := 1; i <= numOfCategories; i++ {
		category := model.Category{
			Name: faker.Word(),
			Desc: faker.Paragraph(),
		}

		db.Create(&category)
		categories = append(categories, category)
	}
}
