package main

import (
	"fmt"
	"github.com/handrixn/example-repo/app"
	"github.com/handrixn/example-repo/entity/domain"
)

func main() {
	app.LoadEnvFile("./.env")

	db := app.NewSQLDatabase()

	category := domain.Category{Name: "Cooking"}

	db.Create(&category)

	fmt.Println(category)
}
